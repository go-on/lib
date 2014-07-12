/*
Package pseudodb provides a simple way to have a json store inside a single file
for rapid prototyping.

It also offers REST routes for all CRUD operations.
All you need, is
*/

package pseudodb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/go-contrib/uuid"
	"github.com/go-on/router"
	"github.com/go-on/router/route"
)

/*
An App offers the following:

- generator for uuids
- generator for unique names
- map to store all items by uuid
- lookup by id
- store id outside of the data
- restroutes

*/

var _ = uuid.DomainGroup

type Store interface {
	Save(*App) error
	Load(*App) error
}

type FileStore struct {
	filename string
	*sync.Mutex
}

func NewFileStore(filename string) Store {
	return &FileStore{
		filename: filename,
		Mutex:    &sync.Mutex{},
	}
}

func (fs *FileStore) Save(a *App) error {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()
	b, err := json.Marshal(a)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(fs.filename, b, 0644)
}

func (fs *FileStore) Load(a *App) error {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()
	b, err := ioutil.ReadFile(fs.filename)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, a)
}

type Storable interface {
	SetUUID(string)
	UUID() string
}

type App struct {
	*sync.RWMutex `json:"-"`
	typeDeps      []string
	registry      map[string]reflect.Type
	GET           *route.Route `json:"-"`
	PATCH         *route.Route `json:"-"`
	POST          *route.Route `json:"-"`
	DELETE        *route.Route `json:"-"`
	INDEX         *route.Route `json:"-"`
	store         Store        `json:"-"`
	Data          map[string]interface{}
	BeforeDelete  func(interface{}) error `json:"-"`
}

type objectWithType struct {
	Object interface{}
	Type   string
}

func (a *App) MarshalJSON() ([]byte, error) {
	data := map[string]objectWithType{}

	for k, v := range a.Data {
		data[k] = objectWithType{
			Object: v,
			Type:   transformType(reflect.TypeOf(v)),
		}
	}

	return json.Marshal(data)
}

func (a *App) UnmarshalJSON(by []byte) error {
	data := map[string]objectWithType{}
	err := json.Unmarshal(by, &data)

	if err != nil {
		return err
	}

	for _, currentty := range a.typeDeps {
		for k, v := range data {
			ty, found := a.registry[v.Type]
			if !found {
				return fmt.Errorf("can't find type %#v in registry", v.Type)
			}

			if v.Type != currentty {
				continue
			}
			val := reflect.New(ty.Elem()).Interface()
			js, errJs := json.Marshal(v.Object)
			if errJs != nil {
				return errJs
			}

			errJs = json.Unmarshal(js, val)
			if errJs != nil {
				return errJs
			}

			a.Data[k] = val
		}
	}
	return nil
}

// NewApp creates a new app, storing in store (if not nil).
// types are pointers to instances of the types that should be
// stored
// the types must be in the order of their dependencies, i.e.
// non dependant types come first. this is important when reloading
// the data from the store
func NewApp(store Store, types ...Storable) *App {
	a := &App{
		RWMutex:  &sync.RWMutex{},
		registry: map[string]reflect.Type{},
		Data:     map[string]interface{}{},
		store:    store,
	}

	for _, t := range types {
		a.register(t)
	}

	return a
}

func _transformType(ty reflect.Type) string {
	str := ty.String()
	i1 := strings.LastIndex(str, ".")
	fmt.Sprintf("i1: %v\n", i1)
	if i1 == -1 {
		return str
	}
	i2 := strings.LastIndex(str[i1+1:], "*")
	fmt.Sprintf("i2: %v\n", i2)
	if i2 == -1 {
		return str[i1+1:]
	}
	return str[i2+1:]
}

func transformType(ty reflect.Type) string {
	return strings.ToLower(_transformType(ty))
}

func (a *App) register(i interface{}) {
	ty := reflect.TypeOf(i)
	if ty.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("%T is no pointer type", i))
	}

	a.Lock()
	defer a.Unlock()
	a.typeDeps = append(a.typeDeps, transformType(ty))
	a.registry[transformType(ty)] = ty
}

func (a *App) Find(key string) (val interface{}, found bool) {
	a.RLock()
	defer a.RUnlock()
	val, found = a.Data[key]
	return
}

func (a *App) Mount(rt *router.Router, prefix string) {
	a.GET = rt.GETFunc(prefix+"/:ressource/:uuid", a.getHandler)
	a.PATCH = rt.PATCHFunc(prefix+"/:ressource/:uuid", a.patchHandler)
	a.POST = rt.POSTFunc(prefix+"/:ressource/", a.postHandler)
	a.DELETE = rt.DELETEFunc(prefix+"/:ressource/:uuid", a.deleteHandler)
	a.INDEX = rt.GETFunc(prefix+"/:ressource/", a.indexHandler)
}

func (a *App) getHandler(rw http.ResponseWriter, req *http.Request) {
	ressource := router.GetRouteParam(req, "ressource")
	uuid := router.GetRouteParam(req, "uuid")

	if ressource == "" || uuid == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("missing params"))
		return
	}

	ty, ok := a.registry[ressource]

	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unknown ressource"))
		return
	}

	val, found := a.Find(uuid)
	if !found {
		http.NotFound(rw, req)
		return
	}

	if reflect.TypeOf(val) != ty {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("wrong ressource type"))
		return
	}

	b, err := json.Marshal(val)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write(b)

}

func filterByType(ty reflect.Type, m map[string]interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	for k, v := range m {
		if reflect.TypeOf(v) == ty {
			res[k] = v
		}
	}
	return res
}

func (a *App) patchHandler(rw http.ResponseWriter, req *http.Request) {
	ressource := router.GetRouteParam(req, "ressource")
	uuid := router.GetRouteParam(req, "uuid")

	if ressource == "" || uuid == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("missing params"))
		return
	}

	ty, ok := a.registry[ressource]

	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unknown ressource"))
		return
	}

	val, found := a.Find(uuid)
	if !found {
		http.NotFound(rw, req)
		return
	}

	if reflect.TypeOf(val) != ty {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("wrong ressource type"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(body, val)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	a.Update(val.(Storable))

	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save: " + err.Error()))
			return
		}
	}
	rw.Write([]byte("ok"))
}

func (a *App) verifyRegistration(val Storable) {
	_, ok := a.registry[transformType(reflect.TypeOf(val))]
	if !ok {
		panic("not in registry: " + reflect.TypeOf(val).String())
	}
}

func (a *App) Update(val Storable) {
	a.verifyRegistration(val)
	a.Lock()
	defer a.Unlock()
	a.Data[val.UUID()] = val
}

func (a *App) postHandler(rw http.ResponseWriter, req *http.Request) {
	ressource := router.GetRouteParam(req, "ressource")

	if ressource == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("missing params"))
		return
	}

	ty, ok := a.registry[ressource]

	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unknown ressource"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	val := reflect.New(ty.Elem()).Interface()

	err = json.Unmarshal(body, val)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}

	_uuid := a.New(val.(Storable))

	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save: " + err.Error()))
			return
		}
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(_uuid))
}

func (a *App) New(val Storable) (uuid_ string) {
	_, ok := a.registry[transformType(reflect.TypeOf(val))]
	if !ok {
		panic("not in registry: " + reflect.TypeOf(val).String())
	}
	a.Lock()
	defer a.Unlock()
	_uuid := uuid.NewV1().String()
	val.SetUUID(_uuid)
	a.Data[_uuid] = val
	return _uuid
}

func (a *App) deleteHandler(rw http.ResponseWriter, req *http.Request) {
	ressource := router.GetRouteParam(req, "ressource")
	uuid := router.GetRouteParam(req, "uuid")

	if ressource == "" || uuid == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("missing params"))
		return
	}

	ty, ok := a.registry[ressource]

	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("unknown ressource"))
		return
	}

	val, found := a.Find(uuid)
	if !found {
		http.NotFound(rw, req)
		return
	}

	if reflect.TypeOf(val) != ty {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("wrong ressource type"))
		return
	}

	var err error

	if a.BeforeDelete != nil {
		err = a.BeforeDelete(val)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
			return
		}
	}

	a.Delete(val.(Storable))

	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save: " + err.Error()))
			return
		}
	}

	rw.Write([]byte("ok"))
}

func (a *App) Delete(val Storable) {
	a.Lock()
	defer a.Unlock()
	delete(a.Data, val.UUID())
}

// List returns the current list for the given pointer to type instance
func (a *App) List(val Storable) map[string]interface{} {
	a.verifyRegistration(val)
	return a.list(reflect.TypeOf(val))
}
func (a *App) list(ty reflect.Type) map[string]interface{} {
	a.RLock()
	defer a.RUnlock()
	return filterByType(ty, a.Data)
}

func (a *App) indexHandler(rw http.ResponseWriter, req *http.Request) {
	ressource := router.GetRouteParam(req, "ressource")

	if ressource == "" {
		rw.Write([]byte("missing params"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	ty, ok := a.registry[ressource]

	if !ok {
		rw.Write([]byte("unknown ressource"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	list := a.list(ty)

	uuids := []string{}

	for k := range list {
		uuids = append(uuids, k)
	}

	sort.Strings(uuids)

	data := make([]interface{}, len(uuids))

	for i, ui := range uuids {
		data[i] = list[ui]
	}

	b, err := json.Marshal(data)

	if err != nil {
		rw.Write([]byte(err.Error()))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(b)
}

func (a *App) Load() error {
	return a.store.Load(a)
}

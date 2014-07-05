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
	"github.com/go-contrib/uuid"
	"github.com/go-on/router"
	"github.com/go-on/router/route"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
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

// TODO: create MarshalJSON and UnmarshalJSON methods that only save the data and the registry
// TODO: check if the registry is compatible when loading, therefor we need to have the
// registrants on NewApp call
type App struct {
	*sync.RWMutex `json:"-"`
	registry      map[string]reflect.Type
	GET           *route.Route `json:"-"`
	PATCH         *route.Route `json:"-"`
	POST          *route.Route `json:"-"`
	DELETE        *route.Route `json:"-"`
	INDEX         *route.Route `json:"-"`
	store         Store        `json:"-"`
	Data          map[string]interface{}
}

func NewApp(store Store) *App {
	return &App{
		RWMutex:  &sync.RWMutex{},
		registry: map[string]reflect.Type{},
		Data:     map[string]interface{}{},
		store:    store,
	}
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

func (a *App) Register(i interface{}) {
	ty := reflect.TypeOf(i)
	if ty.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("%T is no pointer type", i))
	}

	a.Lock()
	defer a.Unlock()

	a.registry[transformType(ty)] = ty
}

func (a *App) Find(key string) (val interface{}, found bool) {
	a.RLock()
	defer a.RUnlock()
	val, found = a.Data[key]
	return
}

func (a *App) Mount(rt *router.Router, prefix string) {
	a.GET = rt.GETFunc(prefix+"/:ressource/:uuid", a.get)
	a.PATCH = rt.PATCHFunc(prefix+"/:ressource/:uuid", a.patch)
	a.POST = rt.POSTFunc(prefix+"/:ressource/", a.post)
	a.DELETE = rt.DELETEFunc(prefix+"/:ressource/:uuid", a.delete)
	a.INDEX = rt.GETFunc(prefix+"/:ressource/", a.index)
}

func (a *App) get(rw http.ResponseWriter, req *http.Request) {
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

func (a *App) patch(rw http.ResponseWriter, req *http.Request) {
	// println("patch called")
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

	a.Lock()
	defer a.Unlock()
	a.Data[uuid] = val

	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save"))
			return
		}
	}
	rw.Write([]byte("ok"))
}

func (a *App) post(rw http.ResponseWriter, req *http.Request) {
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

	val := reflect.New(ty.Elem()).Interface()

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

	a.Lock()
	defer a.Unlock()
	_uuid := uuid.NewV1().String()
	a.Data[_uuid] = val

	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save"))
			return
		}
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(_uuid))
}

func (a *App) delete(rw http.ResponseWriter, req *http.Request) {
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

	a.Lock()
	defer a.Unlock()
	delete(a.Data, uuid)
	if a.store != nil {

		err := a.store.Save(a)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("could not save"))
			return
		}
	}

	rw.Write([]byte("ok"))
}

func (a *App) index(rw http.ResponseWriter, req *http.Request) {
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

	a.RLock()
	defer a.RUnlock()
	data := filterByType(ty, a.Data)

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

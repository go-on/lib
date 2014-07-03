"use strict";
(function() {

Error.stackTraceLimit = -1;

var $global;
if (typeof window !== "undefined") { /* web page */
	$global = window;
} else if (typeof self !== "undefined") { /* web worker */
	$global = self;
} else if (typeof global !== "undefined") { /* Node.js */
	$global = global;
	$global.require = require;
} else {
	console.log("warning: no global object found")
}

var $idCounter = 0;
var $keys = function(m) { return m ? Object.keys(m) : []; };
var $min = Math.min;
var $parseInt = parseInt;
var $parseFloat = function(f) {
	if (f.constructor === Number) {
		return f;
	}
	return parseFloat(f);
};
var $mod = function(x, y) { return x % y; };
var $toString = String;
var $reflect, $newStringPtr;
var $Array = Array;

var $floatKey = function(f) {
	if (f !== f) {
		$idCounter++;
		return "NaN$" + $idCounter;
	}
	return String(f);
};

var $mapArray = function(array, f) {
	var newArray = new array.constructor(array.length), i;
	for (i = 0; i < array.length; i++) {
		newArray[i] = f(array[i]);
	}
	return newArray;
};

var $newType = function(size, kind, string, name, pkgPath, constructor) {
	var typ;
	switch(kind) {
	case "Bool":
	case "Int":
	case "Int8":
	case "Int16":
	case "Int32":
	case "Uint":
	case "Uint8" :
	case "Uint16":
	case "Uint32":
	case "Uintptr":
	case "String":
	case "UnsafePointer":
		typ = function(v) { this.$val = v; };
		typ.prototype.$key = function() { return string + "$" + this.$val; };
		break;

	case "Float32":
	case "Float64":
		typ = function(v) { this.$val = v; };
		typ.prototype.$key = function() { return string + "$" + $floatKey(this.$val); };
		break;

	case "Int64":
		typ = function(high, low) {
			this.high = (high + Math.floor(Math.ceil(low) / 4294967296)) >> 0;
			this.low = low >>> 0;
			this.$val = this;
		};
		typ.prototype.$key = function() { return string + "$" + this.high + "$" + this.low; };
		break;

	case "Uint64":
		typ = function(high, low) {
			this.high = (high + Math.floor(Math.ceil(low) / 4294967296)) >>> 0;
			this.low = low >>> 0;
			this.$val = this;
		};
		typ.prototype.$key = function() { return string + "$" + this.high + "$" + this.low; };
		break;

	case "Complex64":
	case "Complex128":
		typ = function(real, imag) {
			this.real = real;
			this.imag = imag;
			this.$val = this;
		};
		typ.prototype.$key = function() { return string + "$" + this.real + "$" + this.imag; };
		break;

	case "Array":
		typ = function(v) { this.$val = v; };
		typ.Ptr = $newType(4, "Ptr", "*" + string, "", "", function(array) {
			this.$get = function() { return array; };
			this.$val = array;
		});
		typ.init = function(elem, len) {
			typ.elem = elem;
			typ.len = len;
			typ.prototype.$key = function() {
				return string + "$" + Array.prototype.join.call($mapArray(this.$val, function(e) {
					var key = e.$key ? e.$key() : String(e);
					return key.replace(/\\/g, "\\\\").replace(/\$/g, "\\$");
				}), "$");
			};
			typ.extendReflectType = function(rt) {
				rt.arrayType = new $reflect.arrayType.Ptr(rt, elem.reflectType(), undefined, len);
			};
			typ.Ptr.init(typ);
		};
		break;

	case "Chan":
		typ = function() { this.$val = this; };
		typ.prototype.$key = function() {
			if (this.$id === undefined) {
				$idCounter++;
				this.$id = $idCounter;
			}
			return String(this.$id);
		};
		typ.init = function(elem, sendOnly, recvOnly) {
			typ.nil = new typ();
			typ.extendReflectType = function(rt) {
				rt.chanType = new $reflect.chanType.Ptr(rt, elem.reflectType(), sendOnly ? $reflect.SendDir : (recvOnly ? $reflect.RecvDir : $reflect.BothDir));
			};
		};
		break;

	case "Func":
		typ = function(v) { this.$val = v; };
		typ.init = function(params, results, variadic) {
			typ.params = params;
			typ.results = results;
			typ.variadic = variadic;
			typ.extendReflectType = function(rt) {
				var typeSlice = ($sliceType($ptrType($reflect.rtype.Ptr)));
				rt.funcType = new $reflect.funcType.Ptr(rt, variadic, new typeSlice($mapArray(params, function(p) { return p.reflectType(); })), new typeSlice($mapArray(results, function(p) { return p.reflectType(); })));
			};
		};
		break;

	case "Interface":
		typ = { implementedBy: [] };
		typ.init = function(methods) {
			typ.methods = methods;
			typ.extendReflectType = function(rt) {
				var imethods = $mapArray(methods, function(m) {
					return new $reflect.imethod.Ptr($newStringPtr(m[1]), $newStringPtr(m[2]), $funcType(m[3], m[4], m[5]).reflectType());
				});
				var methodSlice = ($sliceType($ptrType($reflect.imethod.Ptr)));
				rt.interfaceType = new $reflect.interfaceType.Ptr(rt, new methodSlice(imethods));
			};
		};
		break;

	case "Map":
		typ = function(v) { this.$val = v; };
		typ.init = function(key, elem) {
			typ.key = key;
			typ.elem = elem;
			typ.extendReflectType = function(rt) {
				rt.mapType = new $reflect.mapType.Ptr(rt, key.reflectType(), elem.reflectType(), undefined, undefined);
			};
		};
		break;

	case "Ptr":
		typ = constructor || function(getter, setter) {
			this.$get = getter;
			this.$set = setter;
			this.$val = this;
		};
		typ.prototype.$key = function() {
			if (this.$id === undefined) {
				$idCounter++;
				this.$id = $idCounter;
			}
			return String(this.$id);
		};
		typ.init = function(elem) {
			typ.nil = new typ($throwNilPointerError, $throwNilPointerError);
			typ.extendReflectType = function(rt) {
				rt.ptrType = new $reflect.ptrType.Ptr(rt, elem.reflectType());
			};
		};
		break;

	case "Slice":
		var nativeArray;
		typ = function(array) {
			if (array.constructor !== nativeArray) {
				array = new nativeArray(array);
			}
			this.array = array;
			this.offset = 0;
			this.length = array.length;
			this.capacity = array.length;
			this.$val = this;
		};
		typ.make = function(length, capacity, zero) {
			capacity = capacity || length;
			var array = new nativeArray(capacity), i;
			for (i = 0; i < capacity; i++) {
				array[i] = zero();
			}
			var slice = new typ(array);
			slice.length = length;
			return slice;
		};
		typ.init = function(elem) {
			typ.elem = elem;
			nativeArray = $nativeArray(elem.kind);
			typ.nil = new typ([]);
			typ.extendReflectType = function(rt) {
				rt.sliceType = new $reflect.sliceType.Ptr(rt, elem.reflectType());
			};
		};
		break;

	case "Struct":
		typ = function(v) { this.$val = v; };
		typ.Ptr = $newType(4, "Ptr", "*" + string, "", "", constructor);
		typ.Ptr.Struct = typ;
		typ.init = function(fields) {
			var i;
			typ.fields = fields;
			typ.Ptr.init(typ);
			/* nil value */
			typ.Ptr.nil = new constructor();
			for (i = 0; i < fields.length; i++) {
				var field = fields[i];
				Object.defineProperty(typ.Ptr.nil, field[1], { get: $throwNilPointerError, set: $throwNilPointerError });
			}
			/* methods for embedded fields */
			for (i = 0; i < typ.methods.length; i++) {
				var method = typ.methods[i];
				if (method[6] != -1) {
					(function(field, methodName) {
						typ.prototype[methodName] = function() {
							var v = this.$val[field[0]];
							return v[methodName].apply(v, arguments);
						};
					})(fields[method[6]], method[0]);
				}
			}
			for (i = 0; i < typ.Ptr.methods.length; i++) {
				var method = typ.Ptr.methods[i];
				if (method[6] != -1) {
					(function(field, methodName) {
						typ.Ptr.prototype[methodName] = function() {
							var v = this[field[0]];
							if (v.$val === undefined) {
								v = new field[3](v);
							}
							return v[methodName].apply(v, arguments);
						};
					})(fields[method[6]], method[0]);
				}
			}
			/* map key */
			typ.prototype.$key = function() {
				var keys = new Array(fields.length);
				for (i = 0; i < fields.length; i++) {
					var v = this.$val[fields[i][0]];
					var key = v.$key ? v.$key() : String(v);
					keys[i] = key.replace(/\\/g, "\\\\").replace(/\$/g, "\\$");
				}
				return string + "$" + keys.join("$");
			};
			/* reflect type */
			typ.extendReflectType = function(rt) {
				var reflectFields = new Array(fields.length), i;
				for (i = 0; i < fields.length; i++) {
					var field = fields[i];
					reflectFields[i] = new $reflect.structField.Ptr($newStringPtr(field[1]), $newStringPtr(field[2]), field[3].reflectType(), $newStringPtr(field[4]), i);
				}
				rt.structType = new $reflect.structType.Ptr(rt, new ($sliceType($reflect.structField.Ptr))(reflectFields));
			};
		};
		break;

	default:
		throw $panic(new $String("invalid kind: " + kind));
	}

	typ.kind = kind;
	typ.string = string;
	typ.typeName = name;
	typ.pkgPath = pkgPath;
	typ.methods = [];
	var rt = null;
	typ.reflectType = function() {
		if (rt === null) {
			rt = new $reflect.rtype.Ptr(size, 0, 0, 0, 0, $reflect.kinds[kind], undefined, undefined, $newStringPtr(string), undefined, undefined);
			rt.jsType = typ;

			var methods = [];
			if (typ.methods !== undefined) {
				var i;
				for (i = 0; i < typ.methods.length; i++) {
					var m = typ.methods[i];
					methods.push(new $reflect.method.Ptr($newStringPtr(m[1]), $newStringPtr(m[2]), $funcType(m[3], m[4], m[5]).reflectType(), $funcType([typ].concat(m[3]), m[4], m[5]).reflectType(), undefined, undefined));
				}
			}
			if (name !== "" || methods.length !== 0) {
				var methodSlice = ($sliceType($ptrType($reflect.method.Ptr)));
				rt.uncommonType = new $reflect.uncommonType.Ptr($newStringPtr(name), $newStringPtr(pkgPath), new methodSlice(methods));
				rt.uncommonType.jsType = typ;
			}

			if (typ.extendReflectType !== undefined) {
				typ.extendReflectType(rt);
			}
		}
		return rt;
	};
	return typ;
};

var $Bool          = $newType( 1, "Bool",          "bool",           "bool",       "", null);
var $Int           = $newType( 4, "Int",           "int",            "int",        "", null);
var $Int8          = $newType( 1, "Int8",          "int8",           "int8",       "", null);
var $Int16         = $newType( 2, "Int16",         "int16",          "int16",      "", null);
var $Int32         = $newType( 4, "Int32",         "int32",          "int32",      "", null);
var $Int64         = $newType( 8, "Int64",         "int64",          "int64",      "", null);
var $Uint          = $newType( 4, "Uint",          "uint",           "uint",       "", null);
var $Uint8         = $newType( 1, "Uint8",         "uint8",          "uint8",      "", null);
var $Uint16        = $newType( 2, "Uint16",        "uint16",         "uint16",     "", null);
var $Uint32        = $newType( 4, "Uint32",        "uint32",         "uint32",     "", null);
var $Uint64        = $newType( 8, "Uint64",        "uint64",         "uint64",     "", null);
var $Uintptr       = $newType( 4, "Uintptr",       "uintptr",        "uintptr",    "", null);
var $Float32       = $newType( 4, "Float32",       "float32",        "float32",    "", null);
var $Float64       = $newType( 8, "Float64",       "float64",        "float64",    "", null);
var $Complex64     = $newType( 8, "Complex64",     "complex64",      "complex64",  "", null);
var $Complex128    = $newType(16, "Complex128",    "complex128",     "complex128", "", null);
var $String        = $newType( 8, "String",        "string",         "string",     "", null);
var $UnsafePointer = $newType( 4, "UnsafePointer", "unsafe.Pointer", "Pointer",    "", null);

var $nativeArray = function(elemKind) {
	return ({ Int: Int32Array, Int8: Int8Array, Int16: Int16Array, Int32: Int32Array, Uint: Uint32Array, Uint8: Uint8Array, Uint16: Uint16Array, Uint32: Uint32Array, Uintptr: Uint32Array, Float32: Float32Array, Float64: Float64Array })[elemKind] || Array;
};
var $toNativeArray = function(elemKind, array) {
	var nativeArray = $nativeArray(elemKind);
	if (nativeArray === Array) {
		return array;
	}
	return new nativeArray(array);
};
var $makeNativeArray = function(elemKind, length, zero) {
	var array = new ($nativeArray(elemKind))(length), i;
	for (i = 0; i < length; i++) {
		array[i] = zero();
	}
	return array;
};
var $arrayTypes = {};
var $arrayType = function(elem, len) {
	var string = "[" + len + "]" + elem.string;
	var typ = $arrayTypes[string];
	if (typ === undefined) {
		typ = $newType(12, "Array", string, "", "", null);
		typ.init(elem, len);
		$arrayTypes[string] = typ;
	}
	return typ;
};

var $chanType = function(elem, sendOnly, recvOnly) {
	var string = (recvOnly ? "<-" : "") + "chan" + (sendOnly ? "<- " : " ") + elem.string;
	var field = sendOnly ? "SendChan" : (recvOnly ? "RecvChan" : "Chan");
	var typ = elem[field];
	if (typ === undefined) {
		typ = $newType(4, "Chan", string, "", "", null);
		typ.init(elem, sendOnly, recvOnly);
		elem[field] = typ;
	}
	return typ;
};

var $funcSig = function(params, results, variadic) {
	var paramTypes = $mapArray(params, function(p) { return p.string; });
	if (variadic) {
		paramTypes[paramTypes.length - 1] = "..." + paramTypes[paramTypes.length - 1].substr(2);
	}
	var string = "(" + paramTypes.join(", ") + ")";
	if (results.length === 1) {
		string += " " + results[0].string;
	} else if (results.length > 1) {
		string += " (" + $mapArray(results, function(r) { return r.string; }).join(", ") + ")";
	}
	return string;
};

var $funcTypes = {};
var $funcType = function(params, results, variadic) {
	var string = "func" + $funcSig(params, results, variadic);
	var typ = $funcTypes[string];
	if (typ === undefined) {
		typ = $newType(4, "Func", string, "", "", null);
		typ.init(params, results, variadic);
		$funcTypes[string] = typ;
	}
	return typ;
};

var $interfaceTypes = {};
var $interfaceType = function(methods) {
	var string = "interface {}";
	if (methods.length !== 0) {
		string = "interface { " + $mapArray(methods, function(m) {
			return (m[2] !== "" ? m[2] + "." : "") + m[1] + $funcSig(m[3], m[4], m[5]);
		}).join("; ") + " }";
	}
	var typ = $interfaceTypes[string];
	if (typ === undefined) {
		typ = $newType(8, "Interface", string, "", "", null);
		typ.init(methods);
		$interfaceTypes[string] = typ;
	}
	return typ;
};
var $emptyInterface = $interfaceType([]);
var $interfaceNil = { $key: function() { return "nil"; } };
var $error = $newType(8, "Interface", "error", "error", "", null);
$error.init([["Error", "Error", "", [], [$String], false]]);

var $Map = function() {};
(function() {
	var names = Object.getOwnPropertyNames(Object.prototype), i;
	for (i = 0; i < names.length; i++) {
		$Map.prototype[names[i]] = undefined;
	}
})();
var $mapTypes = {};
var $mapType = function(key, elem) {
	var string = "map[" + key.string + "]" + elem.string;
	var typ = $mapTypes[string];
	if (typ === undefined) {
		typ = $newType(4, "Map", string, "", "", null);
		typ.init(key, elem);
		$mapTypes[string] = typ;
	}
	return typ;
};

var $throwNilPointerError = function() { $throwRuntimeError("invalid memory address or nil pointer dereference"); };
var $ptrType = function(elem) {
	var typ = elem.Ptr;
	if (typ === undefined) {
		typ = $newType(4, "Ptr", "*" + elem.string, "", "", null);
		typ.init(elem);
		elem.Ptr = typ;
	}
	return typ;
};

var $sliceType = function(elem) {
	var typ = elem.Slice;
	if (typ === undefined) {
		typ = $newType(12, "Slice", "[]" + elem.string, "", "", null);
		typ.init(elem);
		elem.Slice = typ;
	}
	return typ;
};

var $structTypes = {};
var $structType = function(fields) {
	var string = "struct { " + $mapArray(fields, function(f) {
		return f[1] + " " + f[3].string + (f[4] !== "" ? (" \"" + f[4].replace(/\\/g, "\\\\").replace(/"/g, "\\\"") + "\"") : "");
	}).join("; ") + " }";
	var typ = $structTypes[string];
	if (typ === undefined) {
		typ = $newType(0, "Struct", string, "", "", function() {
			this.$val = this;
			var i;
			for (i = 0; i < fields.length; i++) {
				this[fields[i][0]] = arguments[i];
			}
		});
		/* collect methods for anonymous fields */
		var i, j;
		for (i = 0; i < fields.length; i++) {
			var field = fields[i];
			if (field[1] === "") {
				var methods = field[3].methods;
				for (j = 0; j < methods.length; j++) {
					var m = methods[j].slice(0, 6).concat([i]);
					typ.methods.push(m);
					typ.Ptr.methods.push(m);
				}
				if (field[3].kind === "Struct") {
					var methods = field[3].Ptr.methods;
					for (j = 0; j < methods.length; j++) {
						typ.Ptr.methods.push(methods[j].slice(0, 6).concat([i]));
					}
				}
			}
		}
		typ.init(fields);
		$structTypes[string] = typ;
	}
	return typ;
};

var $stringPtrMap = new $Map();
$newStringPtr = function(str) {
	if (str === undefined || str === "") {
		return $ptrType($String).nil;
	}
	var ptr = $stringPtrMap[str];
	if (ptr === undefined) {
		ptr = new ($ptrType($String))(function() { return str; }, function(v) { str = v; });
		$stringPtrMap[str] = ptr;
	}
	return ptr;
};
var $newDataPointer = function(data, constructor) {
	return new constructor(function() { return data; }, function(v) { data = v; });
};

var $coerceFloat32 = function(f) {
	var math = $packages["math"];
	if (math === undefined) {
		return f;
	}
	return math.Float32frombits(math.Float32bits(f));
};
var $flatten64 = function(x) {
	return x.high * 4294967296 + x.low;
};
var $shiftLeft64 = function(x, y) {
	if (y === 0) {
		return x;
	}
	if (y < 32) {
		return new x.constructor(x.high << y | x.low >>> (32 - y), (x.low << y) >>> 0);
	}
	if (y < 64) {
		return new x.constructor(x.low << (y - 32), 0);
	}
	return new x.constructor(0, 0);
};
var $shiftRightInt64 = function(x, y) {
	if (y === 0) {
		return x;
	}
	if (y < 32) {
		return new x.constructor(x.high >> y, (x.low >>> y | x.high << (32 - y)) >>> 0);
	}
	if (y < 64) {
		return new x.constructor(x.high >> 31, (x.high >> (y - 32)) >>> 0);
	}
	if (x.high < 0) {
		return new x.constructor(-1, 4294967295);
	}
	return new x.constructor(0, 0);
};
var $shiftRightUint64 = function(x, y) {
	if (y === 0) {
		return x;
	}
	if (y < 32) {
		return new x.constructor(x.high >>> y, (x.low >>> y | x.high << (32 - y)) >>> 0);
	}
	if (y < 64) {
		return new x.constructor(0, x.high >>> (y - 32));
	}
	return new x.constructor(0, 0);
};
var $mul64 = function(x, y) {
	var high = 0, low = 0, i;
	if ((y.low & 1) !== 0) {
		high = x.high;
		low = x.low;
	}
	for (i = 1; i < 32; i++) {
		if ((y.low & 1<<i) !== 0) {
			high += x.high << i | x.low >>> (32 - i);
			low += (x.low << i) >>> 0;
		}
	}
	for (i = 0; i < 32; i++) {
		if ((y.high & 1<<i) !== 0) {
			high += x.low << i;
		}
	}
	return new x.constructor(high, low);
};
var $div64 = function(x, y, returnRemainder) {
	if (y.high === 0 && y.low === 0) {
		$throwRuntimeError("integer divide by zero");
	}

	var s = 1;
	var rs = 1;

	var xHigh = x.high;
	var xLow = x.low;
	if (xHigh < 0) {
		s = -1;
		rs = -1;
		xHigh = -xHigh;
		if (xLow !== 0) {
			xHigh--;
			xLow = 4294967296 - xLow;
		}
	}

	var yHigh = y.high;
	var yLow = y.low;
	if (y.high < 0) {
		s *= -1;
		yHigh = -yHigh;
		if (yLow !== 0) {
			yHigh--;
			yLow = 4294967296 - yLow;
		}
	}

	var high = 0, low = 0, n = 0, i;
	while (yHigh < 2147483648 && ((xHigh > yHigh) || (xHigh === yHigh && xLow > yLow))) {
		yHigh = (yHigh << 1 | yLow >>> 31) >>> 0;
		yLow = (yLow << 1) >>> 0;
		n++;
	}
	for (i = 0; i <= n; i++) {
		high = high << 1 | low >>> 31;
		low = (low << 1) >>> 0;
		if ((xHigh > yHigh) || (xHigh === yHigh && xLow >= yLow)) {
			xHigh = xHigh - yHigh;
			xLow = xLow - yLow;
			if (xLow < 0) {
				xHigh--;
				xLow += 4294967296;
			}
			low++;
			if (low === 4294967296) {
				high++;
				low = 0;
			}
		}
		yLow = (yLow >>> 1 | yHigh << (32 - 1)) >>> 0;
		yHigh = yHigh >>> 1;
	}

	if (returnRemainder) {
		return new x.constructor(xHigh * rs, xLow * rs);
	}
	return new x.constructor(high * s, low * s);
};

var $divComplex = function(n, d) {
	var ninf = n.real === 1/0 || n.real === -1/0 || n.imag === 1/0 || n.imag === -1/0;
	var dinf = d.real === 1/0 || d.real === -1/0 || d.imag === 1/0 || d.imag === -1/0;
	var nnan = !ninf && (n.real !== n.real || n.imag !== n.imag);
	var dnan = !dinf && (d.real !== d.real || d.imag !== d.imag);
	if(nnan || dnan) {
		return new n.constructor(0/0, 0/0);
	}
	if (ninf && !dinf) {
		return new n.constructor(1/0, 1/0);
	}
	if (!ninf && dinf) {
		return new n.constructor(0, 0);
	}
	if (d.real === 0 && d.imag === 0) {
		if (n.real === 0 && n.imag === 0) {
			return new n.constructor(0/0, 0/0);
		}
		return new n.constructor(1/0, 1/0);
	}
	var a = Math.abs(d.real);
	var b = Math.abs(d.imag);
	if (a <= b) {
		var ratio = d.real / d.imag;
		var denom = d.real * ratio + d.imag;
		return new n.constructor((n.real * ratio + n.imag) / denom, (n.imag * ratio - n.real) / denom);
	}
	var ratio = d.imag / d.real;
	var denom = d.imag * ratio + d.real;
	return new n.constructor((n.imag * ratio + n.real) / denom, (n.imag - n.real * ratio) / denom);
};

var $subslice = function(slice, low, high, max) {
	if (low < 0 || high < low || max < high || high > slice.capacity || max > slice.capacity) {
		$throwRuntimeError("slice bounds out of range");
	}
	var s = new slice.constructor(slice.array);
	s.offset = slice.offset + low;
	s.length = slice.length - low;
	s.capacity = slice.capacity - low;
	if (high !== undefined) {
		s.length = high - low;
	}
	if (max !== undefined) {
		s.capacity = max - low;
	}
	return s;
};

var $sliceToArray = function(slice) {
	if (slice.length === 0) {
		return [];
	}
	if (slice.array.constructor !== Array) {
		return slice.array.subarray(slice.offset, slice.offset + slice.length);
	}
	return slice.array.slice(slice.offset, slice.offset + slice.length);
};

var $decodeRune = function(str, pos) {
	var c0 = str.charCodeAt(pos);

	if (c0 < 0x80) {
		return [c0, 1];
	}

	if (c0 !== c0 || c0 < 0xC0) {
		return [0xFFFD, 1];
	}

	var c1 = str.charCodeAt(pos + 1);
	if (c1 !== c1 || c1 < 0x80 || 0xC0 <= c1) {
		return [0xFFFD, 1];
	}

	if (c0 < 0xE0) {
		var r = (c0 & 0x1F) << 6 | (c1 & 0x3F);
		if (r <= 0x7F) {
			return [0xFFFD, 1];
		}
		return [r, 2];
	}

	var c2 = str.charCodeAt(pos + 2);
	if (c2 !== c2 || c2 < 0x80 || 0xC0 <= c2) {
		return [0xFFFD, 1];
	}

	if (c0 < 0xF0) {
		var r = (c0 & 0x0F) << 12 | (c1 & 0x3F) << 6 | (c2 & 0x3F);
		if (r <= 0x7FF) {
			return [0xFFFD, 1];
		}
		if (0xD800 <= r && r <= 0xDFFF) {
			return [0xFFFD, 1];
		}
		return [r, 3];
	}

	var c3 = str.charCodeAt(pos + 3);
	if (c3 !== c3 || c3 < 0x80 || 0xC0 <= c3) {
		return [0xFFFD, 1];
	}

	if (c0 < 0xF8) {
		var r = (c0 & 0x07) << 18 | (c1 & 0x3F) << 12 | (c2 & 0x3F) << 6 | (c3 & 0x3F);
		if (r <= 0xFFFF || 0x10FFFF < r) {
			return [0xFFFD, 1];
		}
		return [r, 4];
	}

	return [0xFFFD, 1];
};

var $encodeRune = function(r) {
	if (r < 0 || r > 0x10FFFF || (0xD800 <= r && r <= 0xDFFF)) {
		r = 0xFFFD;
	}
	if (r <= 0x7F) {
		return String.fromCharCode(r);
	}
	if (r <= 0x7FF) {
		return String.fromCharCode(0xC0 | r >> 6, 0x80 | (r & 0x3F));
	}
	if (r <= 0xFFFF) {
		return String.fromCharCode(0xE0 | r >> 12, 0x80 | (r >> 6 & 0x3F), 0x80 | (r & 0x3F));
	}
	return String.fromCharCode(0xF0 | r >> 18, 0x80 | (r >> 12 & 0x3F), 0x80 | (r >> 6 & 0x3F), 0x80 | (r & 0x3F));
};

var $stringToBytes = function(str, terminateWithNull) {
	var array = new Uint8Array(terminateWithNull ? str.length + 1 : str.length), i;
	for (i = 0; i < str.length; i++) {
		array[i] = str.charCodeAt(i);
	}
	if (terminateWithNull) {
		array[str.length] = 0;
	}
	return array;
};

var $bytesToString = function(slice) {
	if (slice.length === 0) {
		return "";
	}
	var str = "", i;
	for (i = 0; i < slice.length; i += 10000) {
		str += String.fromCharCode.apply(null, slice.array.subarray(slice.offset + i, slice.offset + Math.min(slice.length, i + 10000)));
	}
	return str;
};

var $stringToRunes = function(str) {
	var array = new Int32Array(str.length);
	var rune, i, j = 0;
	for (i = 0; i < str.length; i += rune[1], j++) {
		rune = $decodeRune(str, i);
		array[j] = rune[0];
	}
	return array.subarray(0, j);
};

var $runesToString = function(slice) {
	if (slice.length === 0) {
		return "";
	}
	var str = "", i;
	for (i = 0; i < slice.length; i++) {
		str += $encodeRune(slice.array[slice.offset + i]);
	}
	return str;
};

var $needsExternalization = function(t) {
	switch (t.kind) {
		case "Int64":
		case "Uint64":
		case "Array":
		case "Func":
		case "Map":
		case "Slice":
		case "String":
			return true;
		case "Interface":
			return t !== $packages["github.com/gopherjs/gopherjs/js"].Object;
		default:
			return false;
	}
};

var $externalize = function(v, t) {
	switch (t.kind) {
	case "Int64":
	case "Uint64":
		return $flatten64(v);
	case "Array":
		if ($needsExternalization(t.elem)) {
			return $mapArray(v, function(e) { return $externalize(e, t.elem); });
		}
		return v;
	case "Func":
		if (v === $throwNilPointerError) {
			return null;
		}
		var convert = false;
		var i;
		for (i = 0; i < t.params.length; i++) {
			convert = convert || (t.params[i] !== $packages["github.com/gopherjs/gopherjs/js"].Object);
		}
		for (i = 0; i < t.results.length; i++) {
			convert = convert || $needsExternalization(t.results[i]);
		}
		if (!convert) {
			return v;
		}
		return function() {
			var args = [], i;
			for (i = 0; i < t.params.length; i++) {
				if (t.variadic && i === t.params.length - 1) {
					var vt = t.params[i].elem, varargs = [], j;
					for (j = i; j < arguments.length; j++) {
						varargs.push($internalize(arguments[j], vt));
					}
					args.push(new (t.params[i])(varargs));
					break;
				}
				args.push($internalize(arguments[i], t.params[i]));
			}
			var result = v.apply(undefined, args);
			switch (t.results.length) {
			case 0:
				return;
			case 1:
				return $externalize(result, t.results[0]);
			default:
				for (i = 0; i < t.results.length; i++) {
					result[i] = $externalize(result[i], t.results[i]);
				}
				return result;
			}
		};
	case "Interface":
		if (v === null) {
			return null;
		}
		if (t === $packages["github.com/gopherjs/gopherjs/js"].Object || v.constructor.kind === undefined) {
			return v;
		}
		return $externalize(v.$val, v.constructor);
	case "Map":
		var m = {};
		var keys = $keys(v), i;
		for (i = 0; i < keys.length; i++) {
			var entry = v[keys[i]];
			m[$externalize(entry.k, t.key)] = $externalize(entry.v, t.elem);
		}
		return m;
	case "Slice":
		if ($needsExternalization(t.elem)) {
			return $mapArray($sliceToArray(v), function(e) { return $externalize(e, t.elem); });
		}
		return $sliceToArray(v);
	case "String":
		var s = "", r, i, j = 0;
		for (i = 0; i < v.length; i += r[1], j++) {
			r = $decodeRune(v, i);
			s += String.fromCharCode(r[0]);
		}
		return s;
	case "Struct":
		var timePkg = $packages["time"];
		if (timePkg && v.constructor === timePkg.Time.Ptr) {
			var milli = $div64(v.UnixNano(), new $Int64(0, 1000000));
			return new Date($flatten64(milli));
		}
		return v;
	default:
		return v;
	}
};

var $internalize = function(v, t, recv) {
	switch (t.kind) {
	case "Bool":
		return !!v;
	case "Int":
		return parseInt(v);
	case "Int8":
		return parseInt(v) << 24 >> 24;
	case "Int16":
		return parseInt(v) << 16 >> 16;
	case "Int32":
		return parseInt(v) >> 0;
	case "Uint":
		return parseInt(v);
	case "Uint8" :
		return parseInt(v) << 24 >>> 24;
	case "Uint16":
		return parseInt(v) << 16 >>> 16;
	case "Uint32":
	case "Uintptr":
		return parseInt(v) >>> 0;
	case "Int64":
	case "Uint64":
		return new t(0, v);
	case "Float32":
	case "Float64":
		return parseFloat(v);
	case "Array":
		if (v.length !== t.len) {
			$throwRuntimeError("got array with wrong size from JavaScript native");
		}
		return $mapArray(v, function(e) { return $internalize(e, t.elem); });
	case "Func":
		return function() {
			var args = [], i;
			for (i = 0; i < t.params.length; i++) {
				if (t.variadic && i === t.params.length - 1) {
					var vt = t.params[i].elem, varargs = arguments[i], j;
					for (j = 0; j < varargs.length; j++) {
						args.push($externalize(varargs.array[varargs.offset + j], vt));
					}
					break;
				}
				args.push($externalize(arguments[i], t.params[i]));
			}
			var result = v.apply(recv, args);
			switch (t.results.length) {
			case 0:
				return;
			case 1:
				return $internalize(result, t.results[0]);
			default:
				for (i = 0; i < t.results.length; i++) {
					result[i] = $internalize(result[i], t.results[i]);
				}
				return result;
			}
		};
	case "Interface":
		if (v === null || t === $packages["github.com/gopherjs/gopherjs/js"].Object) {
			return v;
		}
		switch (v.constructor) {
		case Int8Array:
			return new ($sliceType($Int8))(v);
		case Int16Array:
			return new ($sliceType($Int16))(v);
		case Int32Array:
			return new ($sliceType($Int))(v);
		case Uint8Array:
			return new ($sliceType($Uint8))(v);
		case Uint16Array:
			return new ($sliceType($Uint16))(v);
		case Uint32Array:
			return new ($sliceType($Uint))(v);
		case Float32Array:
			return new ($sliceType($Float32))(v);
		case Float64Array:
			return new ($sliceType($Float64))(v);
		case Array:
			return $internalize(v, $sliceType($emptyInterface));
		case Boolean:
			return new $Bool(!!v);
		case Date:
			var timePkg = $packages["time"];
			if (timePkg) {
				return new timePkg.Time(timePkg.Unix(new $Int64(0, 0), new $Int64(0, v.getTime() * 1000000)));
			}
		case Function:
			var funcType = $funcType([$sliceType($emptyInterface)], [$packages["github.com/gopherjs/gopherjs/js"].Object], true);
			return new funcType($internalize(v, funcType));
		case Number:
			return new $Float64(parseFloat(v));
		case Object:
			var mapType = $mapType($String, $emptyInterface);
			return new mapType($internalize(v, mapType));
		case String:
			return new $String($internalize(v, $String));
		}
		return v;
	case "Map":
		var m = new $Map();
		var keys = $keys(v), i;
		for (i = 0; i < keys.length; i++) {
			var key = $internalize(keys[i], t.key);
			m[key.$key ? key.$key() : key] = { k: key, v: $internalize(v[keys[i]], t.elem) };
		}
		return m;
	case "Slice":
		return new t($mapArray(v, function(e) { return $internalize(e, t.elem); }));
	case "String":
		v = String(v);
		var s = "", i;
		for (i = 0; i < v.length; i++) {
			s += $encodeRune(v.charCodeAt(i));
		}
		return s;
	default:
		return v;
	}
};

var $copySlice = function(dst, src) {
	var n = Math.min(src.length, dst.length), i;
	if (dst.array.constructor !== Array && n !== 0) {
		dst.array.set(src.array.subarray(src.offset, src.offset + n), dst.offset);
		return n;
	}
	for (i = 0; i < n; i++) {
		dst.array[dst.offset + i] = src.array[src.offset + i];
	}
	return n;
};

var $copyString = function(dst, src) {
	var n = Math.min(src.length, dst.length), i;
	for (i = 0; i < n; i++) {
		dst.array[dst.offset + i] = src.charCodeAt(i);
	}
	return n;
};

var $copyArray = function(dst, src) {
	var i;
	for (i = 0; i < src.length; i++) {
		dst[i] = src[i];
	}
};

var $growSlice = function(slice, length) {
	var newCapacity = Math.max(length, slice.capacity < 1024 ? slice.capacity * 2 : Math.floor(slice.capacity * 5 / 4));

	var newArray;
	if (slice.array.constructor === Array) {
		newArray = slice.array;
		if (slice.offset !== 0 || newArray.length !== slice.offset + slice.capacity) {
			newArray = newArray.slice(slice.offset);
		}
		newArray.length = newCapacity;
	} else {
		newArray = new slice.array.constructor(newCapacity);
		newArray.set(slice.array.subarray(slice.offset));
	}

	var newSlice = new slice.constructor(newArray);
	newSlice.length = slice.length;
	newSlice.capacity = newCapacity;
	return newSlice;
};

var $append = function(slice) {
	if (arguments.length === 1) {
		return slice;
	}

	var newLength = slice.length + arguments.length - 1;
	if (newLength > slice.capacity) {
		slice = $growSlice(slice, newLength);
	}

	var array = slice.array;
	var leftOffset = slice.offset + slice.length - 1, i;
	for (i = 1; i < arguments.length; i++) {
		array[leftOffset + i] = arguments[i];
	}

	var newSlice = new slice.constructor(array);
	newSlice.offset = slice.offset;
	newSlice.length = newLength;
	newSlice.capacity = slice.capacity;
	return newSlice;
};

var $appendSlice = function(slice, toAppend) {
	if (toAppend.length === 0) {
		return slice;
	}

	var newLength = slice.length + toAppend.length;
	if (newLength > slice.capacity) {
		slice = $growSlice(slice, newLength);
	}

	var array = slice.array;
	var leftOffset = slice.offset + slice.length, rightOffset = toAppend.offset, i;
	for (i = 0; i < toAppend.length; i++) {
		array[leftOffset + i] = toAppend.array[rightOffset + i];
	}

	var newSlice = new slice.constructor(array);
	newSlice.offset = slice.offset;
	newSlice.length = newLength;
	newSlice.capacity = slice.capacity;
	return newSlice;
};

var $panic = function(value) {
	var message;
	if (value.constructor === $String) {
		message = value.$val;
	} else if (value.Error !== undefined) {
		message = value.Error();
	} else if (value.String !== undefined) {
		message = value.String();
	} else {
		message = value;
	}
	var err = new Error(message);
	err.$panicValue = value;
	return err;
};
var $notSupported = function(feature) {
	var err = new Error("not supported by GopherJS: " + feature);
	err.$notSupported = feature;
	throw err;
};
var $throwRuntimeError; /* set by package "runtime" */

var $errorStack = [], $jsErr = null;

var $pushErr = function(err) {
	if (err.$panicValue === undefined) {
		if (err.$exit || err.$notSupported) {
			$jsErr = err;
			return;
		}
		err.$panicValue = new $packages["github.com/gopherjs/gopherjs/js"].Error.Ptr(err);
	}
	$errorStack.push({ frame: $getStackDepth(), error: err });
};

var $callDeferred = function(deferred) {
	if ($jsErr !== null) {
		throw $jsErr;
	}
	var i;
	for (i = deferred.length - 1; i >= 0; i--) {
		var call = deferred[i];
		try {
			if (call.recv !== undefined) {
				call.recv[call.method].apply(call.recv, call.args);
				continue;
			}
			call.fun.apply(undefined, call.args);
		} catch (err) {
			$errorStack.push({ frame: $getStackDepth(), error: err });
		}
	}
	var err = $errorStack[$errorStack.length - 1];
	if (err !== undefined && err.frame === $getStackDepth()) {
		$errorStack.pop();
		throw err.error;
	}
};

var $recover = function() {
	var err = $errorStack[$errorStack.length - 1];
	if (err === undefined || err.frame !== $getStackDepth()) {
		return null;
	}
	$errorStack.pop();
	return err.error.$panicValue;
};

var $getStack = function() {
	return (new Error()).stack.split("\n");
};

var $getStackDepth = function() {
	var s = $getStack(), d = 0, i;
	for (i = 0; i < s.length; i++) {
		if (s[i].indexOf("$") === -1) {
			d++;
		}
	}
	return d;
};

var $interfaceIsEqual = function(a, b) {
	if (a === b) {
		return true;
	}
	if (a === null || b === null || a === undefined || b === undefined || a.constructor !== b.constructor) {
		return false;
	}
	switch (a.constructor.kind) {
	case "Float32":
		return $float32IsEqual(a.$val, b.$val);
	case "Complex64":
		return $float32IsEqual(a.$val.real, b.$val.real) && $float32IsEqual(a.$val.imag, b.$val.imag);
	case "Complex128":
		return a.$val.real === b.$val.real && a.$val.imag === b.$val.imag;
	case "Int64":
	case "Uint64":
		return a.$val.high === b.$val.high && a.$val.low === b.$val.low;
	case "Array":
		return $arrayIsEqual(a.$val, b.$val);
	case "Ptr":
		if (a.constructor.Struct) {
			return false;
		}
		return $pointerIsEqual(a, b);
	case "Func":
	case "Map":
	case "Slice":
	case "Struct":
		$throwRuntimeError("comparing uncomparable type " + a.constructor);
	case undefined: /* js.Object */
		return false;
	default:
		return a.$val === b.$val;
	}
};
var $float32IsEqual = function(a, b) {
	if (a === b) {
		return true;
	}
	if (a === 0 || b === 0 || a === 1/0 || b === 1/0 || a === -1/0 || b === -1/0 || a !== a || b !== b) {
		return false;
	}
	var math = $packages["math"];
	return math !== undefined && math.Float32bits(a) === math.Float32bits(b);
};
var $arrayIsEqual = function(a, b) {
	if (a.length != b.length) {
		return false;
	}
	var i;
	for (i = 0; i < a.length; i++) {
		if (a[i] !== b[i]) {
			return false;
		}
	}
	return true;
};
var $sliceIsEqual = function(a, ai, b, bi) {
	return a.array === b.array && a.offset + ai === b.offset + bi;
};
var $pointerIsEqual = function(a, b) {
	if (a === b) {
		return true;
	}
	if (a.$get === $throwNilPointerError || b.$get === $throwNilPointerError) {
		return a.$get === $throwNilPointerError && b.$get === $throwNilPointerError;
	}
	var old = a.$get();
	var dummy = new Object();
	a.$set(dummy);
	var equal = b.$get() === dummy;
	a.$set(old);
	return equal;
};

var $typeAssertionFailed = function(obj, expected) {
	var got = "";
	if (obj !== null) {
		got = obj.constructor.string;
	}
	throw $panic(new $packages["runtime"].TypeAssertionError.Ptr("", got, expected.string, ""));
};

var $now = function() { var msec = (new Date()).getTime(); return [new $Int64(0, Math.floor(msec / 1000)), (msec % 1000) * 1000000]; };

var $packages = {};
$packages["github.com/gopherjs/gopherjs/js"] = (function() {
	var $pkg = {}, Object, Error;
	Object = $pkg.Object = $newType(8, "Interface", "js.Object", "Object", "github.com/gopherjs/gopherjs/js", null);
	Error = $pkg.Error = $newType(0, "Struct", "js.Error", "Error", "github.com/gopherjs/gopherjs/js", function(Object_) {
		this.$val = this;
		this.Object = Object_ !== undefined ? Object_ : null;
	});
	Error.Ptr.prototype.Error = function() {
		var err;
		err = this;
		return "JavaScript error: " + $internalize(err.Object.message, $String);
	};
	Error.prototype.Error = function() { return this.$val.Error(); };
	$pkg.init = function() {
		Object.init([["Bool", "Bool", "", [], [$Bool], false], ["Call", "Call", "", [$String, ($sliceType($emptyInterface))], [Object], true], ["Delete", "Delete", "", [$String], [], false], ["Float", "Float", "", [], [$Float64], false], ["Get", "Get", "", [$String], [Object], false], ["Index", "Index", "", [$Int], [Object], false], ["Int", "Int", "", [], [$Int], false], ["Int64", "Int64", "", [], [$Int64], false], ["Interface", "Interface", "", [], [$emptyInterface], false], ["Invoke", "Invoke", "", [($sliceType($emptyInterface))], [Object], true], ["IsNull", "IsNull", "", [], [$Bool], false], ["IsUndefined", "IsUndefined", "", [], [$Bool], false], ["Length", "Length", "", [], [$Int], false], ["New", "New", "", [($sliceType($emptyInterface))], [Object], true], ["Set", "Set", "", [$String, $emptyInterface], [], false], ["SetIndex", "SetIndex", "", [$Int, $emptyInterface], [], false], ["Str", "Str", "", [], [$String], false], ["Uint64", "Uint64", "", [], [$Uint64], false], ["Unsafe", "Unsafe", "", [], [$Uintptr], false]]);
		Error.methods = [["Bool", "Bool", "", [], [$Bool], false, 0], ["Call", "Call", "", [$String, ($sliceType($emptyInterface))], [Object], true, 0], ["Delete", "Delete", "", [$String], [], false, 0], ["Float", "Float", "", [], [$Float64], false, 0], ["Get", "Get", "", [$String], [Object], false, 0], ["Index", "Index", "", [$Int], [Object], false, 0], ["Int", "Int", "", [], [$Int], false, 0], ["Int64", "Int64", "", [], [$Int64], false, 0], ["Interface", "Interface", "", [], [$emptyInterface], false, 0], ["Invoke", "Invoke", "", [($sliceType($emptyInterface))], [Object], true, 0], ["IsNull", "IsNull", "", [], [$Bool], false, 0], ["IsUndefined", "IsUndefined", "", [], [$Bool], false, 0], ["Length", "Length", "", [], [$Int], false, 0], ["New", "New", "", [($sliceType($emptyInterface))], [Object], true, 0], ["Set", "Set", "", [$String, $emptyInterface], [], false, 0], ["SetIndex", "SetIndex", "", [$Int, $emptyInterface], [], false, 0], ["Str", "Str", "", [], [$String], false, 0], ["Uint64", "Uint64", "", [], [$Uint64], false, 0], ["Unsafe", "Unsafe", "", [], [$Uintptr], false, 0]];
		($ptrType(Error)).methods = [["Bool", "Bool", "", [], [$Bool], false, 0], ["Call", "Call", "", [$String, ($sliceType($emptyInterface))], [Object], true, 0], ["Delete", "Delete", "", [$String], [], false, 0], ["Error", "Error", "", [], [$String], false, -1], ["Float", "Float", "", [], [$Float64], false, 0], ["Get", "Get", "", [$String], [Object], false, 0], ["Index", "Index", "", [$Int], [Object], false, 0], ["Int", "Int", "", [], [$Int], false, 0], ["Int64", "Int64", "", [], [$Int64], false, 0], ["Interface", "Interface", "", [], [$emptyInterface], false, 0], ["Invoke", "Invoke", "", [($sliceType($emptyInterface))], [Object], true, 0], ["IsNull", "IsNull", "", [], [$Bool], false, 0], ["IsUndefined", "IsUndefined", "", [], [$Bool], false, 0], ["Length", "Length", "", [], [$Int], false, 0], ["New", "New", "", [($sliceType($emptyInterface))], [Object], true, 0], ["Set", "Set", "", [$String, $emptyInterface], [], false, 0], ["SetIndex", "SetIndex", "", [$Int, $emptyInterface], [], false, 0], ["Str", "Str", "", [], [$String], false, 0], ["Uint64", "Uint64", "", [], [$Uint64], false, 0], ["Unsafe", "Unsafe", "", [], [$Uintptr], false, 0]];
		Error.init([["Object", "", "", Object, ""]]);
		var e;
		e = new Error.Ptr(null);
	};
	return $pkg;
})();
$packages["runtime"] = (function() {
	var $pkg = {}, js = $packages["github.com/gopherjs/gopherjs/js"], TypeAssertionError, errorString, goexit, sizeof_C_MStats;
	TypeAssertionError = $pkg.TypeAssertionError = $newType(0, "Struct", "runtime.TypeAssertionError", "TypeAssertionError", "runtime", function(interfaceString_, concreteString_, assertedString_, missingMethod_) {
		this.$val = this;
		this.interfaceString = interfaceString_ !== undefined ? interfaceString_ : "";
		this.concreteString = concreteString_ !== undefined ? concreteString_ : "";
		this.assertedString = assertedString_ !== undefined ? assertedString_ : "";
		this.missingMethod = missingMethod_ !== undefined ? missingMethod_ : "";
	});
	errorString = $pkg.errorString = $newType(8, "String", "runtime.errorString", "errorString", "runtime", null);
	TypeAssertionError.Ptr.prototype.RuntimeError = function() {
	};
	TypeAssertionError.prototype.RuntimeError = function() { return this.$val.RuntimeError(); };
	TypeAssertionError.Ptr.prototype.Error = function() {
		var e, inter;
		e = this;
		inter = e.interfaceString;
		if (inter === "") {
			inter = "interface";
		}
		if (e.concreteString === "") {
			return "interface conversion: " + inter + " is nil, not " + e.assertedString;
		}
		if (e.missingMethod === "") {
			return "interface conversion: " + inter + " is " + e.concreteString + ", not " + e.assertedString;
		}
		return "interface conversion: " + e.concreteString + " is not " + e.assertedString + ": missing method " + e.missingMethod;
	};
	TypeAssertionError.prototype.Error = function() { return this.$val.Error(); };
	errorString.prototype.RuntimeError = function() {
		var e;
		e = this.$val;
	};
	$ptrType(errorString).prototype.RuntimeError = function() { return new errorString(this.$get()).RuntimeError(); };
	errorString.prototype.Error = function() {
		var e;
		e = this.$val;
		return "runtime error: " + e;
	};
	$ptrType(errorString).prototype.Error = function() { return new errorString(this.$get()).Error(); };
	$pkg.init = function() {
		($ptrType(TypeAssertionError)).methods = [["Error", "Error", "", [], [$String], false, -1], ["RuntimeError", "RuntimeError", "", [], [], false, -1]];
		TypeAssertionError.init([["interfaceString", "interfaceString", "runtime", $String, ""], ["concreteString", "concreteString", "runtime", $String, ""], ["assertedString", "assertedString", "runtime", $String, ""], ["missingMethod", "missingMethod", "runtime", $String, ""]]);
		errorString.methods = [["Error", "Error", "", [], [$String], false, -1], ["RuntimeError", "RuntimeError", "", [], [], false, -1]];
		($ptrType(errorString)).methods = [["Error", "Error", "", [], [$String], false, -1], ["RuntimeError", "RuntimeError", "", [], [], false, -1]];
		sizeof_C_MStats = 3712;
		goexit = $global.eval($externalize("(function() {\n\tvar err = new Error();\n\terr.$exit = true;\n\tthrow err;\n})", $String));
		var e;
		$throwRuntimeError = $externalize((function(msg) {
			throw $panic(new errorString(msg));
		}), ($funcType([$String], [], false)));
		e = new TypeAssertionError.Ptr("", "", "", "");
		if (!((sizeof_C_MStats === 3712))) {
			console.log(sizeof_C_MStats, 3712);
			throw $panic(new $String("MStats vs MemStatsType size mismatch"));
		}
	};
	return $pkg;
})();
$packages["errors"] = (function() {
	var $pkg = {}, errorString, New;
	errorString = $pkg.errorString = $newType(0, "Struct", "errors.errorString", "errorString", "errors", function(s_) {
		this.$val = this;
		this.s = s_ !== undefined ? s_ : "";
	});
	New = $pkg.New = function(text) {
		return new errorString.Ptr(text);
	};
	errorString.Ptr.prototype.Error = function() {
		var e;
		e = this;
		return e.s;
	};
	errorString.prototype.Error = function() { return this.$val.Error(); };
	$pkg.init = function() {
		($ptrType(errorString)).methods = [["Error", "Error", "", [], [$String], false, -1]];
		errorString.init([["s", "s", "errors", $String, ""]]);
	};
	return $pkg;
})();
$packages["sync/atomic"] = (function() {
	var $pkg = {};
	$pkg.init = function() {
	};
	return $pkg;
})();
$packages["sync"] = (function() {
	var $pkg = {}, atomic = $packages["sync/atomic"], runtime_Syncsemcheck;
	runtime_Syncsemcheck = function(size) {
	};
	$pkg.init = function() {
		var s;
		s = $makeNativeArray("Uintptr", 3, function() { return 0; });
		runtime_Syncsemcheck(12);
	};
	return $pkg;
})();
$packages["io"] = (function() {
	var $pkg = {}, errors = $packages["errors"], sync = $packages["sync"], Reader, Writer, errWhence, errOffset;
	Reader = $pkg.Reader = $newType(8, "Interface", "io.Reader", "Reader", "io", null);
	Writer = $pkg.Writer = $newType(8, "Interface", "io.Writer", "Writer", "io", null);
	$pkg.init = function() {
		Reader.init([["Read", "Read", "", [($sliceType($Uint8))], [$Int, $error], false]]);
		Writer.init([["Write", "Write", "", [($sliceType($Uint8))], [$Int, $error], false]]);
		$pkg.ErrShortWrite = errors.New("short write");
		$pkg.ErrShortBuffer = errors.New("short buffer");
		$pkg.EOF = errors.New("EOF");
		$pkg.ErrUnexpectedEOF = errors.New("unexpected EOF");
		$pkg.ErrNoProgress = errors.New("multiple Read calls return no data or error");
		errWhence = errors.New("Seek: invalid whence");
		errOffset = errors.New("Seek: invalid offset");
		$pkg.ErrClosedPipe = errors.New("io: read/write on closed pipe");
	};
	return $pkg;
})();
$packages["unicode"] = (function() {
	var $pkg = {};
	$pkg.init = function() {
	};
	return $pkg;
})();
$packages["unicode/utf8"] = (function() {
	var $pkg = {}, decodeRuneInternal, DecodeRune, DecodeLastRune, EncodeRune, RuneStart;
	decodeRuneInternal = function(p) {
		var r, size, short$1, n, _tmp, _tmp$1, _tmp$2, c0, _tmp$3, _tmp$4, _tmp$5, _tmp$6, _tmp$7, _tmp$8, _tmp$9, _tmp$10, _tmp$11, c1, _tmp$12, _tmp$13, _tmp$14, _tmp$15, _tmp$16, _tmp$17, _tmp$18, _tmp$19, _tmp$20, _tmp$21, _tmp$22, _tmp$23, c2, _tmp$24, _tmp$25, _tmp$26, _tmp$27, _tmp$28, _tmp$29, _tmp$30, _tmp$31, _tmp$32, _tmp$33, _tmp$34, _tmp$35, _tmp$36, _tmp$37, _tmp$38, c3, _tmp$39, _tmp$40, _tmp$41, _tmp$42, _tmp$43, _tmp$44, _tmp$45, _tmp$46, _tmp$47, _tmp$48, _tmp$49, _tmp$50;
		r = 0;
		size = 0;
		short$1 = false;
		n = p.length;
		if (n < 1) {
			_tmp = 65533; _tmp$1 = 0; _tmp$2 = true; r = _tmp; size = _tmp$1; short$1 = _tmp$2;
			return [r, size, short$1];
		}
		c0 = ((0 < 0 || 0 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 0]);
		if (c0 < 128) {
			_tmp$3 = (c0 >> 0); _tmp$4 = 1; _tmp$5 = false; r = _tmp$3; size = _tmp$4; short$1 = _tmp$5;
			return [r, size, short$1];
		}
		if (c0 < 192) {
			_tmp$6 = 65533; _tmp$7 = 1; _tmp$8 = false; r = _tmp$6; size = _tmp$7; short$1 = _tmp$8;
			return [r, size, short$1];
		}
		if (n < 2) {
			_tmp$9 = 65533; _tmp$10 = 1; _tmp$11 = true; r = _tmp$9; size = _tmp$10; short$1 = _tmp$11;
			return [r, size, short$1];
		}
		c1 = ((1 < 0 || 1 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 1]);
		if (c1 < 128 || 192 <= c1) {
			_tmp$12 = 65533; _tmp$13 = 1; _tmp$14 = false; r = _tmp$12; size = _tmp$13; short$1 = _tmp$14;
			return [r, size, short$1];
		}
		if (c0 < 224) {
			r = ((((c0 & 31) >>> 0) >> 0) << 6 >> 0) | (((c1 & 63) >>> 0) >> 0);
			if (r <= 127) {
				_tmp$15 = 65533; _tmp$16 = 1; _tmp$17 = false; r = _tmp$15; size = _tmp$16; short$1 = _tmp$17;
				return [r, size, short$1];
			}
			_tmp$18 = r; _tmp$19 = 2; _tmp$20 = false; r = _tmp$18; size = _tmp$19; short$1 = _tmp$20;
			return [r, size, short$1];
		}
		if (n < 3) {
			_tmp$21 = 65533; _tmp$22 = 1; _tmp$23 = true; r = _tmp$21; size = _tmp$22; short$1 = _tmp$23;
			return [r, size, short$1];
		}
		c2 = ((2 < 0 || 2 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 2]);
		if (c2 < 128 || 192 <= c2) {
			_tmp$24 = 65533; _tmp$25 = 1; _tmp$26 = false; r = _tmp$24; size = _tmp$25; short$1 = _tmp$26;
			return [r, size, short$1];
		}
		if (c0 < 240) {
			r = (((((c0 & 15) >>> 0) >> 0) << 12 >> 0) | ((((c1 & 63) >>> 0) >> 0) << 6 >> 0)) | (((c2 & 63) >>> 0) >> 0);
			if (r <= 2047) {
				_tmp$27 = 65533; _tmp$28 = 1; _tmp$29 = false; r = _tmp$27; size = _tmp$28; short$1 = _tmp$29;
				return [r, size, short$1];
			}
			if (55296 <= r && r <= 57343) {
				_tmp$30 = 65533; _tmp$31 = 1; _tmp$32 = false; r = _tmp$30; size = _tmp$31; short$1 = _tmp$32;
				return [r, size, short$1];
			}
			_tmp$33 = r; _tmp$34 = 3; _tmp$35 = false; r = _tmp$33; size = _tmp$34; short$1 = _tmp$35;
			return [r, size, short$1];
		}
		if (n < 4) {
			_tmp$36 = 65533; _tmp$37 = 1; _tmp$38 = true; r = _tmp$36; size = _tmp$37; short$1 = _tmp$38;
			return [r, size, short$1];
		}
		c3 = ((3 < 0 || 3 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 3]);
		if (c3 < 128 || 192 <= c3) {
			_tmp$39 = 65533; _tmp$40 = 1; _tmp$41 = false; r = _tmp$39; size = _tmp$40; short$1 = _tmp$41;
			return [r, size, short$1];
		}
		if (c0 < 248) {
			r = ((((((c0 & 7) >>> 0) >> 0) << 18 >> 0) | ((((c1 & 63) >>> 0) >> 0) << 12 >> 0)) | ((((c2 & 63) >>> 0) >> 0) << 6 >> 0)) | (((c3 & 63) >>> 0) >> 0);
			if (r <= 65535 || 1114111 < r) {
				_tmp$42 = 65533; _tmp$43 = 1; _tmp$44 = false; r = _tmp$42; size = _tmp$43; short$1 = _tmp$44;
				return [r, size, short$1];
			}
			_tmp$45 = r; _tmp$46 = 4; _tmp$47 = false; r = _tmp$45; size = _tmp$46; short$1 = _tmp$47;
			return [r, size, short$1];
		}
		_tmp$48 = 65533; _tmp$49 = 1; _tmp$50 = false; r = _tmp$48; size = _tmp$49; short$1 = _tmp$50;
		return [r, size, short$1];
	};
	DecodeRune = $pkg.DecodeRune = function(p) {
		var r, size, _tuple;
		r = 0;
		size = 0;
		_tuple = decodeRuneInternal(p); r = _tuple[0]; size = _tuple[1];
		return [r, size];
	};
	DecodeLastRune = $pkg.DecodeLastRune = function(p) {
		var r, size, end, _tmp, _tmp$1, start, _tmp$2, _tmp$3, lim, _tuple, _tmp$4, _tmp$5, _tmp$6, _tmp$7;
		r = 0;
		size = 0;
		end = p.length;
		if (end === 0) {
			_tmp = 65533; _tmp$1 = 0; r = _tmp; size = _tmp$1;
			return [r, size];
		}
		start = end - 1 >> 0;
		r = (((start < 0 || start >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + start]) >> 0);
		if (r < 128) {
			_tmp$2 = r; _tmp$3 = 1; r = _tmp$2; size = _tmp$3;
			return [r, size];
		}
		lim = end - 4 >> 0;
		if (lim < 0) {
			lim = 0;
		}
		start = start - 1 >> 0;
		while (start >= lim) {
			if (RuneStart(((start < 0 || start >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + start]))) {
				break;
			}
			start = start - 1 >> 0;
		}
		if (start < 0) {
			start = 0;
		}
		_tuple = DecodeRune($subslice(p, start, end)); r = _tuple[0]; size = _tuple[1];
		if (!(((start + size >> 0) === end))) {
			_tmp$4 = 65533; _tmp$5 = 1; r = _tmp$4; size = _tmp$5;
			return [r, size];
		}
		_tmp$6 = r; _tmp$7 = size; r = _tmp$6; size = _tmp$7;
		return [r, size];
	};
	EncodeRune = $pkg.EncodeRune = function(p, r) {
		if ((r >>> 0) <= 127) {
			(0 < 0 || 0 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 0] = (r << 24 >>> 24);
			return 1;
		}
		if ((r >>> 0) <= 2047) {
			(0 < 0 || 0 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 0] = (192 | ((r >> 6 >> 0) << 24 >>> 24)) >>> 0;
			(1 < 0 || 1 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 1] = (128 | (((r << 24 >>> 24) & 63) >>> 0)) >>> 0;
			return 2;
		}
		if ((r >>> 0) > 1114111) {
			r = 65533;
		}
		if (55296 <= r && r <= 57343) {
			r = 65533;
		}
		if ((r >>> 0) <= 65535) {
			(0 < 0 || 0 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 0] = (224 | ((r >> 12 >> 0) << 24 >>> 24)) >>> 0;
			(1 < 0 || 1 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 1] = (128 | ((((r >> 6 >> 0) << 24 >>> 24) & 63) >>> 0)) >>> 0;
			(2 < 0 || 2 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 2] = (128 | (((r << 24 >>> 24) & 63) >>> 0)) >>> 0;
			return 3;
		}
		(0 < 0 || 0 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 0] = (240 | ((r >> 18 >> 0) << 24 >>> 24)) >>> 0;
		(1 < 0 || 1 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 1] = (128 | ((((r >> 12 >> 0) << 24 >>> 24) & 63) >>> 0)) >>> 0;
		(2 < 0 || 2 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 2] = (128 | ((((r >> 6 >> 0) << 24 >>> 24) & 63) >>> 0)) >>> 0;
		(3 < 0 || 3 >= p.length) ? $throwRuntimeError("index out of range") : p.array[p.offset + 3] = (128 | (((r << 24 >>> 24) & 63) >>> 0)) >>> 0;
		return 4;
	};
	RuneStart = $pkg.RuneStart = function(b) {
		return !((((b & 192) >>> 0) === 128));
	};
	$pkg.init = function() {
	};
	return $pkg;
})();
$packages["bytes"] = (function() {
	var $pkg = {}, errors = $packages["errors"], io = $packages["io"], utf8 = $packages["unicode/utf8"], unicode = $packages["unicode"], Buffer, readOp, IndexByte, makeSlice;
	Buffer = $pkg.Buffer = $newType(0, "Struct", "bytes.Buffer", "Buffer", "bytes", function(buf_, off_, runeBytes_, bootstrap_, lastRead_) {
		this.$val = this;
		this.buf = buf_ !== undefined ? buf_ : ($sliceType($Uint8)).nil;
		this.off = off_ !== undefined ? off_ : 0;
		this.runeBytes = runeBytes_ !== undefined ? runeBytes_ : $makeNativeArray("Uint8", 4, function() { return 0; });
		this.bootstrap = bootstrap_ !== undefined ? bootstrap_ : $makeNativeArray("Uint8", 64, function() { return 0; });
		this.lastRead = lastRead_ !== undefined ? lastRead_ : 0;
	});
	readOp = $pkg.readOp = $newType(4, "Int", "bytes.readOp", "readOp", "bytes", null);
	IndexByte = $pkg.IndexByte = function(s, c) {
		var _ref, _i, b, i;
		_ref = s;
		_i = 0;
		while (_i < _ref.length) {
			b = ((_i < 0 || _i >= _ref.length) ? $throwRuntimeError("index out of range") : _ref.array[_ref.offset + _i]);
			i = _i;
			if (b === c) {
				return i;
			}
			_i++;
		}
		return -1;
	};
	Buffer.Ptr.prototype.Bytes = function() {
		var b;
		b = this;
		return $subslice(b.buf, b.off);
	};
	Buffer.prototype.Bytes = function() { return this.$val.Bytes(); };
	Buffer.Ptr.prototype.String = function() {
		var b;
		b = this;
		if (b === ($ptrType(Buffer)).nil) {
			return "<nil>";
		}
		return $bytesToString($subslice(b.buf, b.off));
	};
	Buffer.prototype.String = function() { return this.$val.String(); };
	Buffer.Ptr.prototype.Len = function() {
		var b;
		b = this;
		return b.buf.length - b.off >> 0;
	};
	Buffer.prototype.Len = function() { return this.$val.Len(); };
	Buffer.Ptr.prototype.Truncate = function(n) {
		var b;
		b = this;
		b.lastRead = 0;
		if (n < 0 || n > b.Len()) {
			throw $panic(new $String("bytes.Buffer: truncation out of range"));
		} else if (n === 0) {
			b.off = 0;
		}
		b.buf = $subslice(b.buf, 0, (b.off + n >> 0));
	};
	Buffer.prototype.Truncate = function(n) { return this.$val.Truncate(n); };
	Buffer.Ptr.prototype.Reset = function() {
		var b;
		b = this;
		b.Truncate(0);
	};
	Buffer.prototype.Reset = function() { return this.$val.Reset(); };
	Buffer.Ptr.prototype.grow = function(n) {
		var b, m, buf, _q, x;
		b = this;
		m = b.Len();
		if ((m === 0) && !((b.off === 0))) {
			b.Truncate(0);
		}
		if ((b.buf.length + n >> 0) > b.buf.capacity) {
			buf = ($sliceType($Uint8)).nil;
			if (b.buf === ($sliceType($Uint8)).nil && n <= 64) {
				buf = $subslice(new ($sliceType($Uint8))(b.bootstrap), 0);
			} else if ((m + n >> 0) <= (_q = b.buf.capacity / 2, (_q === _q && _q !== 1/0 && _q !== -1/0) ? _q >> 0 : $throwRuntimeError("integer divide by zero"))) {
				$copySlice(b.buf, $subslice(b.buf, b.off));
				buf = $subslice(b.buf, 0, m);
			} else {
				buf = makeSlice((x = b.buf.capacity, (((2 >>> 16 << 16) * x >> 0) + (2 << 16 >>> 16) * x) >> 0) + n >> 0);
				$copySlice(buf, $subslice(b.buf, b.off));
			}
			b.buf = buf;
			b.off = 0;
		}
		b.buf = $subslice(b.buf, 0, ((b.off + m >> 0) + n >> 0));
		return b.off + m >> 0;
	};
	Buffer.prototype.grow = function(n) { return this.$val.grow(n); };
	Buffer.Ptr.prototype.Grow = function(n) {
		var b, m;
		b = this;
		if (n < 0) {
			throw $panic(new $String("bytes.Buffer.Grow: negative count"));
		}
		m = b.grow(n);
		b.buf = $subslice(b.buf, 0, m);
	};
	Buffer.prototype.Grow = function(n) { return this.$val.Grow(n); };
	Buffer.Ptr.prototype.Write = function(p) {
		var n, err, b, m, _tmp, _tmp$1;
		n = 0;
		err = null;
		b = this;
		b.lastRead = 0;
		m = b.grow(p.length);
		_tmp = $copySlice($subslice(b.buf, m), p); _tmp$1 = null; n = _tmp; err = _tmp$1;
		return [n, err];
	};
	Buffer.prototype.Write = function(p) { return this.$val.Write(p); };
	Buffer.Ptr.prototype.WriteString = function(s) {
		var n, err, b, m, _tmp, _tmp$1;
		n = 0;
		err = null;
		b = this;
		b.lastRead = 0;
		m = b.grow(s.length);
		_tmp = $copyString($subslice(b.buf, m), s); _tmp$1 = null; n = _tmp; err = _tmp$1;
		return [n, err];
	};
	Buffer.prototype.WriteString = function(s) { return this.$val.WriteString(s); };
	Buffer.Ptr.prototype.ReadFrom = function(r) {
		var n, err, b, free, newBuf, x, _tuple, m, e, x$1, _tmp, _tmp$1, _tmp$2, _tmp$3;
		n = new $Int64(0, 0);
		err = null;
		b = this;
		b.lastRead = 0;
		if (b.off >= b.buf.length) {
			b.Truncate(0);
		}
		while (true) {
			free = b.buf.capacity - b.buf.length >> 0;
			if (free < 512) {
				newBuf = b.buf;
				if ((b.off + free >> 0) < 512) {
					newBuf = makeSlice((x = b.buf.capacity, (((2 >>> 16 << 16) * x >> 0) + (2 << 16 >>> 16) * x) >> 0) + 512 >> 0);
				}
				$copySlice(newBuf, $subslice(b.buf, b.off));
				b.buf = $subslice(newBuf, 0, (b.buf.length - b.off >> 0));
				b.off = 0;
			}
			_tuple = r.Read($subslice(b.buf, b.buf.length, b.buf.capacity)); m = _tuple[0]; e = _tuple[1];
			b.buf = $subslice(b.buf, 0, (b.buf.length + m >> 0));
			n = (x$1 = new $Int64(0, m), new $Int64(n.high + x$1.high, n.low + x$1.low));
			if ($interfaceIsEqual(e, io.EOF)) {
				break;
			}
			if (!($interfaceIsEqual(e, null))) {
				_tmp = n; _tmp$1 = e; n = _tmp; err = _tmp$1;
				return [n, err];
			}
		}
		_tmp$2 = n; _tmp$3 = null; n = _tmp$2; err = _tmp$3;
		return [n, err];
	};
	Buffer.prototype.ReadFrom = function(r) { return this.$val.ReadFrom(r); };
	makeSlice = function(n) {
		var $deferred = [];
		try {
			$deferred.push({ fun: (function() {
				if (!($interfaceIsEqual($recover(), null))) {
					throw $panic($pkg.ErrTooLarge);
				}
			}), args: [] });
			return ($sliceType($Uint8)).make(n, 0, function() { return 0; });
		} catch($err) {
			$pushErr($err);
			return ($sliceType($Uint8)).nil;
		} finally {
			$callDeferred($deferred);
		}
	};
	Buffer.Ptr.prototype.WriteTo = function(w) {
		var n, err, b, nBytes, _tuple, m, e, _tmp, _tmp$1, _tmp$2, _tmp$3;
		n = new $Int64(0, 0);
		err = null;
		b = this;
		b.lastRead = 0;
		if (b.off < b.buf.length) {
			nBytes = b.Len();
			_tuple = w.Write($subslice(b.buf, b.off)); m = _tuple[0]; e = _tuple[1];
			if (m > nBytes) {
				throw $panic(new $String("bytes.Buffer.WriteTo: invalid Write count"));
			}
			b.off = b.off + (m) >> 0;
			n = new $Int64(0, m);
			if (!($interfaceIsEqual(e, null))) {
				_tmp = n; _tmp$1 = e; n = _tmp; err = _tmp$1;
				return [n, err];
			}
			if (!((m === nBytes))) {
				_tmp$2 = n; _tmp$3 = io.ErrShortWrite; n = _tmp$2; err = _tmp$3;
				return [n, err];
			}
		}
		b.Truncate(0);
		return [n, err];
	};
	Buffer.prototype.WriteTo = function(w) { return this.$val.WriteTo(w); };
	Buffer.Ptr.prototype.WriteByte = function(c) {
		var b, m, x;
		b = this;
		b.lastRead = 0;
		m = b.grow(1);
		(x = b.buf, (m < 0 || m >= x.length) ? $throwRuntimeError("index out of range") : x.array[x.offset + m] = c);
		return null;
	};
	Buffer.prototype.WriteByte = function(c) { return this.$val.WriteByte(c); };
	Buffer.Ptr.prototype.WriteRune = function(r) {
		var n, err, b, _tmp, _tmp$1, _tmp$2, _tmp$3;
		n = 0;
		err = null;
		b = this;
		if (r < 128) {
			b.WriteByte((r << 24 >>> 24));
			_tmp = 1; _tmp$1 = null; n = _tmp; err = _tmp$1;
			return [n, err];
		}
		n = utf8.EncodeRune($subslice(new ($sliceType($Uint8))(b.runeBytes), 0), r);
		b.Write($subslice(new ($sliceType($Uint8))(b.runeBytes), 0, n));
		_tmp$2 = n; _tmp$3 = null; n = _tmp$2; err = _tmp$3;
		return [n, err];
	};
	Buffer.prototype.WriteRune = function(r) { return this.$val.WriteRune(r); };
	Buffer.Ptr.prototype.Read = function(p) {
		var n, err, b, _tmp, _tmp$1;
		n = 0;
		err = null;
		b = this;
		b.lastRead = 0;
		if (b.off >= b.buf.length) {
			b.Truncate(0);
			if (p.length === 0) {
				return [n, err];
			}
			_tmp = 0; _tmp$1 = io.EOF; n = _tmp; err = _tmp$1;
			return [n, err];
		}
		n = $copySlice(p, $subslice(b.buf, b.off));
		b.off = b.off + (n) >> 0;
		if (n > 0) {
			b.lastRead = 2;
		}
		return [n, err];
	};
	Buffer.prototype.Read = function(p) { return this.$val.Read(p); };
	Buffer.Ptr.prototype.Next = function(n) {
		var b, m, data;
		b = this;
		b.lastRead = 0;
		m = b.Len();
		if (n > m) {
			n = m;
		}
		data = $subslice(b.buf, b.off, (b.off + n >> 0));
		b.off = b.off + (n) >> 0;
		if (n > 0) {
			b.lastRead = 2;
		}
		return data;
	};
	Buffer.prototype.Next = function(n) { return this.$val.Next(n); };
	Buffer.Ptr.prototype.ReadByte = function() {
		var c, err, b, _tmp, _tmp$1, x, x$1, _tmp$2, _tmp$3;
		c = 0;
		err = null;
		b = this;
		b.lastRead = 0;
		if (b.off >= b.buf.length) {
			b.Truncate(0);
			_tmp = 0; _tmp$1 = io.EOF; c = _tmp; err = _tmp$1;
			return [c, err];
		}
		c = (x = b.buf, x$1 = b.off, ((x$1 < 0 || x$1 >= x.length) ? $throwRuntimeError("index out of range") : x.array[x.offset + x$1]));
		b.off = b.off + 1 >> 0;
		b.lastRead = 2;
		_tmp$2 = c; _tmp$3 = null; c = _tmp$2; err = _tmp$3;
		return [c, err];
	};
	Buffer.prototype.ReadByte = function() { return this.$val.ReadByte(); };
	Buffer.Ptr.prototype.ReadRune = function() {
		var r, size, err, b, _tmp, _tmp$1, _tmp$2, x, x$1, c, _tmp$3, _tmp$4, _tmp$5, _tuple, n, _tmp$6, _tmp$7, _tmp$8;
		r = 0;
		size = 0;
		err = null;
		b = this;
		b.lastRead = 0;
		if (b.off >= b.buf.length) {
			b.Truncate(0);
			_tmp = 0; _tmp$1 = 0; _tmp$2 = io.EOF; r = _tmp; size = _tmp$1; err = _tmp$2;
			return [r, size, err];
		}
		b.lastRead = 1;
		c = (x = b.buf, x$1 = b.off, ((x$1 < 0 || x$1 >= x.length) ? $throwRuntimeError("index out of range") : x.array[x.offset + x$1]));
		if (c < 128) {
			b.off = b.off + 1 >> 0;
			_tmp$3 = (c >> 0); _tmp$4 = 1; _tmp$5 = null; r = _tmp$3; size = _tmp$4; err = _tmp$5;
			return [r, size, err];
		}
		_tuple = utf8.DecodeRune($subslice(b.buf, b.off)); r = _tuple[0]; n = _tuple[1];
		b.off = b.off + (n) >> 0;
		_tmp$6 = r; _tmp$7 = n; _tmp$8 = null; r = _tmp$6; size = _tmp$7; err = _tmp$8;
		return [r, size, err];
	};
	Buffer.prototype.ReadRune = function() { return this.$val.ReadRune(); };
	Buffer.Ptr.prototype.UnreadRune = function() {
		var b, _tuple, n;
		b = this;
		if (!((b.lastRead === 1))) {
			return errors.New("bytes.Buffer: UnreadRune: previous operation was not ReadRune");
		}
		b.lastRead = 0;
		if (b.off > 0) {
			_tuple = utf8.DecodeLastRune($subslice(b.buf, 0, b.off)); n = _tuple[1];
			b.off = b.off - (n) >> 0;
		}
		return null;
	};
	Buffer.prototype.UnreadRune = function() { return this.$val.UnreadRune(); };
	Buffer.Ptr.prototype.UnreadByte = function() {
		var b;
		b = this;
		if (!((b.lastRead === 1)) && !((b.lastRead === 2))) {
			return errors.New("bytes.Buffer: UnreadByte: previous operation was not a read");
		}
		b.lastRead = 0;
		if (b.off > 0) {
			b.off = b.off - 1 >> 0;
		}
		return null;
	};
	Buffer.prototype.UnreadByte = function() { return this.$val.UnreadByte(); };
	Buffer.Ptr.prototype.ReadBytes = function(delim) {
		var line, err, b, _tuple, slice;
		line = ($sliceType($Uint8)).nil;
		err = null;
		b = this;
		_tuple = b.readSlice(delim); slice = _tuple[0]; err = _tuple[1];
		line = $appendSlice(line, slice);
		return [line, err];
	};
	Buffer.prototype.ReadBytes = function(delim) { return this.$val.ReadBytes(delim); };
	Buffer.Ptr.prototype.readSlice = function(delim) {
		var line, err, b, i, end, _tmp, _tmp$1;
		line = ($sliceType($Uint8)).nil;
		err = null;
		b = this;
		i = IndexByte($subslice(b.buf, b.off), delim);
		end = (b.off + i >> 0) + 1 >> 0;
		if (i < 0) {
			end = b.buf.length;
			err = io.EOF;
		}
		line = $subslice(b.buf, b.off, end);
		b.off = end;
		b.lastRead = 2;
		_tmp = line; _tmp$1 = err; line = _tmp; err = _tmp$1;
		return [line, err];
	};
	Buffer.prototype.readSlice = function(delim) { return this.$val.readSlice(delim); };
	Buffer.Ptr.prototype.ReadString = function(delim) {
		var line, err, b, _tuple, slice, _tmp, _tmp$1;
		line = "";
		err = null;
		b = this;
		_tuple = b.readSlice(delim); slice = _tuple[0]; err = _tuple[1];
		_tmp = $bytesToString(slice); _tmp$1 = err; line = _tmp; err = _tmp$1;
		return [line, err];
	};
	Buffer.prototype.ReadString = function(delim) { return this.$val.ReadString(delim); };
	$pkg.init = function() {
		($ptrType(Buffer)).methods = [["Bytes", "Bytes", "", [], [($sliceType($Uint8))], false, -1], ["Grow", "Grow", "", [$Int], [], false, -1], ["Len", "Len", "", [], [$Int], false, -1], ["Next", "Next", "", [$Int], [($sliceType($Uint8))], false, -1], ["Read", "Read", "", [($sliceType($Uint8))], [$Int, $error], false, -1], ["ReadByte", "ReadByte", "", [], [$Uint8, $error], false, -1], ["ReadBytes", "ReadBytes", "", [$Uint8], [($sliceType($Uint8)), $error], false, -1], ["ReadFrom", "ReadFrom", "", [io.Reader], [$Int64, $error], false, -1], ["ReadRune", "ReadRune", "", [], [$Int32, $Int, $error], false, -1], ["ReadString", "ReadString", "", [$Uint8], [$String, $error], false, -1], ["Reset", "Reset", "", [], [], false, -1], ["String", "String", "", [], [$String], false, -1], ["Truncate", "Truncate", "", [$Int], [], false, -1], ["UnreadByte", "UnreadByte", "", [], [$error], false, -1], ["UnreadRune", "UnreadRune", "", [], [$error], false, -1], ["Write", "Write", "", [($sliceType($Uint8))], [$Int, $error], false, -1], ["WriteByte", "WriteByte", "", [$Uint8], [$error], false, -1], ["WriteRune", "WriteRune", "", [$Int32], [$Int, $error], false, -1], ["WriteString", "WriteString", "", [$String], [$Int, $error], false, -1], ["WriteTo", "WriteTo", "", [io.Writer], [$Int64, $error], false, -1], ["grow", "grow", "bytes", [$Int], [$Int], false, -1], ["readSlice", "readSlice", "bytes", [$Uint8], [($sliceType($Uint8)), $error], false, -1]];
		Buffer.init([["buf", "buf", "bytes", ($sliceType($Uint8)), ""], ["off", "off", "bytes", $Int, ""], ["runeBytes", "runeBytes", "bytes", ($arrayType($Uint8, 4)), ""], ["bootstrap", "bootstrap", "bytes", ($arrayType($Uint8, 64)), ""], ["lastRead", "lastRead", "bytes", readOp, ""]]);
		$pkg.ErrTooLarge = errors.New("bytes.Buffer: too large");
	};
	return $pkg;
})();
$packages["strings"] = (function() {
	var $pkg = {}, js = $packages["github.com/gopherjs/gopherjs/js"], errors = $packages["errors"], io = $packages["io"], utf8 = $packages["unicode/utf8"], unicode = $packages["unicode"], IndexAny;
	IndexAny = $pkg.IndexAny = function(s, chars) {
		var _ref, _i, _rune, c, i, _ref$1, _i$1, _rune$1, m;
		if (chars.length > 0) {
			_ref = s;
			_i = 0;
			while (_i < _ref.length) {
				_rune = $decodeRune(_ref, _i);
				c = _rune[0];
				i = _i;
				_ref$1 = chars;
				_i$1 = 0;
				while (_i$1 < _ref$1.length) {
					_rune$1 = $decodeRune(_ref$1, _i$1);
					m = _rune$1[0];
					if (c === m) {
						return i;
					}
					_i$1 += _rune$1[1];
				}
				_i += _rune[1];
			}
		}
		return -1;
	};
	$pkg.init = function() {
	};
	return $pkg;
})();
$packages["github.com/go-on/html/types"] = (function() {
	var $pkg = {}, bytes = $packages["bytes"], strings = $packages["strings"], Attribute, escape, EscapeString;
	Attribute = $pkg.Attribute = $newType(0, "Struct", "types.Attribute", "Attribute", "github.com/go-on/html/types", function(Key_, Value_) {
		this.$val = this;
		this.Key = Key_ !== undefined ? Key_ : "";
		this.Value = Value_ !== undefined ? Value_ : "";
	});
	escape = function(w, s) {
		var i, _tuple, err, esc, _ref, _tuple$1, err$1, _tuple$2, err$2;
		i = strings.IndexAny(s, "&'<>\"");
		while (!((i === -1))) {
			_tuple = w.WriteString(s.substring(0, i)); err = _tuple[1];
			if (!($interfaceIsEqual(err, null))) {
				return err;
			}
			esc = "";
			_ref = s.charCodeAt(i);
			if (_ref === 38) {
				esc = "&amp;";
			} else if (_ref === 39) {
				esc = "&#39;";
			} else if (_ref === 60) {
				esc = "&lt;";
			} else if (_ref === 62) {
				esc = "&gt;";
			} else if (_ref === 34) {
				esc = "&#34;";
			} else {
				throw $panic(new $String("unrecognized escape character"));
			}
			s = s.substring((i + 1 >> 0));
			_tuple$1 = w.WriteString(esc); err$1 = _tuple$1[1];
			if (!($interfaceIsEqual(err$1, null))) {
				return err$1;
			}
			i = strings.IndexAny(s, "&'<>\"");
		}
		_tuple$2 = w.WriteString(s); err$2 = _tuple$2[1];
		return err$2;
	};
	EscapeString = $pkg.EscapeString = function(s) {
		var buf;
		if (strings.IndexAny(s, "&'<>\"") === -1) {
			return s;
		}
		buf = new bytes.Buffer.Ptr();
		escape(buf, s);
		return buf.String();
	};
	Attribute.Ptr.prototype.String = function() {
		var _struct, a;
		a = (_struct = this, new Attribute.Ptr(_struct.Key, _struct.Value));
		if (a.Key === "") {
			return " " + a.Value;
		}
		return " " + a.Key + "=\"" + EscapeString(a.Value) + "\"";
	};
	Attribute.prototype.String = function() { return this.$val.String(); };
	$pkg.init = function() {
		Attribute.methods = [["String", "String", "", [], [$String], false, -1]];
		($ptrType(Attribute)).methods = [["String", "String", "", [], [$String], false, -1]];
		Attribute.init([["Key", "Key", "", $String, ""], ["Value", "Value", "", $String, ""]]);
	};
	return $pkg;
})();
$packages["/home/benny/Entwicklung/gopath/src/github.com/go-on/html/examples/gopherjstypestest"] = (function() {
	var $pkg = {}, types = $packages["github.com/go-on/html/types"], main;
	main = $pkg.main = function() {
		console.log(new types.Attribute.Ptr("data", "<hiho>").String());
	};
	$pkg.init = function() {
	};
	return $pkg;
})();
$error.implementedBy = [$packages["errors"].errorString.Ptr, $packages["github.com/gopherjs/gopherjs/js"].Error.Ptr, $packages["runtime"].TypeAssertionError.Ptr, $packages["runtime"].errorString, $ptrType($packages["runtime"].errorString)];
$packages["github.com/gopherjs/gopherjs/js"].Object.implementedBy = [$packages["github.com/gopherjs/gopherjs/js"].Error, $packages["github.com/gopherjs/gopherjs/js"].Error.Ptr];
$packages["io"].Reader.implementedBy = [$packages["bytes"].Buffer.Ptr];
$packages["io"].Writer.implementedBy = [$packages["bytes"].Buffer.Ptr];
$packages["github.com/gopherjs/gopherjs/js"].init();
$packages["runtime"].init();
$packages["errors"].init();
$packages["sync/atomic"].init();
$packages["sync"].init();
$packages["io"].init();
$packages["unicode"].init();
$packages["unicode/utf8"].init();
$packages["bytes"].init();
$packages["strings"].init();
$packages["github.com/go-on/html/types"].init();
$packages["/home/benny/Entwicklung/gopath/src/github.com/go-on/html/examples/gopherjstypestest"].init();
$packages["/home/benny/Entwicklung/gopath/src/github.com/go-on/html/examples/gopherjstypestest"].main();

})();
//# sourceMappingURL=gopherjstypestest.js.map

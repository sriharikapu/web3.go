package literals

type Object struct{ object *Object }

func (o *Object) Get(key string) *Object { return o.object.Get(key) }

func (o *Object) Set(key string, value interface{}) { o.object.Set(key, value) }

func (o *Object) Delete(key string) { o.object.Delete(key) }

func (o *Object) Length() int { return o.object.Length() }

func (o *Object) Index(i int) *Object { return o.object.Index(i) }

func (o *Object) SetIndex(i int, value interface{}) { o.object.SetIndex(i, value) }

func (o *Object) Call(name string, args ...interface{}) *Object { return o.object.Call(name, args...) }

func (o *Object) Invoke(args ...interface{}) *Object { return o.object.Invoke(args...) }

func (o *Object) New(args ...interface{}) *Object { return o.object.New(args...) }

func (o *Object) Bool() bool { return o.object.Bool() }

func (o *Object) String() string { return o.object.String() }

func (o *Object) Int() int { return o.object.Int() }

func (o *Object) Int64() int64 { return o.object.Int64() }

func (o *Object) Uint64() uint64 { return o.object.Uint64() }

func (o *Object) Float() float64 { return o.object.Float() }

func (o *Object) Interface() interface{} { return o.object.Interface() }

func (o *Object) Unsafe() uintptr { return o.object.Unsafe() }

type Error struct {
	*Object
}

func (err *Error) Error() string {
	return "JavaScript error: " + err.Get("message").String()
}

func (err *Error) Stack() string {
	return err.Get("stack").String()
}

var Global *Object

var Module *Object

var Undefined *Object

func Debugger() {}

func InternalObject(i interface{}) *Object {
	return nil
}

func MakeFunc(fn func(this *Object, arguments []*Object) interface{}) *Object {
	return Global.Call("$makeFunc", InternalObject(fn))
}

func Keys(o *Object) []string {
	if o == nil || o == Undefined {
		return nil
	}
	a := Global.Get("Object").Call("keys", o)
	s := make([]string, a.Length())
	for i := 0; i < a.Length(); i++ {
		s[i] = a.Index(i).String()
	}
	return s
}

func MakeWrapper(i interface{}) *Object {
	v := InternalObject(i)
	o := Global.Get("Object").New()
	o.Set("__internal_object__", v)
	methods := v.Get("constructor").Get("methods")
	for i := 0; i < methods.Length(); i++ {
		m := methods.Index(i)
		if m.Get("pkg").String() != "" { 
			continue
		}
		o.Set(m.Get("name").String(), func(args ...*Object) *Object {
			return Global.Call("$externalizeFunction", v.Get(m.Get("prop").String()), m.Get("typ"), true).Call("apply", v, args)
		})
	}
	return o
}

func NewArrayBuffer(b []byte) *Object {
	slice := InternalObject(b)
	offset := slice.Get("$offset").Int()
	length := slice.Get("$length").Int()
	return slice.Get("$array").Get("buffer").Call("slice", offset, offset+length)
}

type M map[string]interface{}

type S []interface{}

func init() {
	e := Error{}
	_ = e
}

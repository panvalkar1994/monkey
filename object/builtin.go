package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (bf *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (bf *Builtin) Inspect() string  { return "builtin function" }

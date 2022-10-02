package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
)

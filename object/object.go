package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Q. 왜 이걸 메서드로 구현하지 않지?
func IsError(object Object) bool {
	if object != nil {
		return object.Type() == ErrorObj
	}
	return false
}

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	StringObj      = "STRING"
)

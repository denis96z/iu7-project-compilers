package exec

type Object interface {
	GetName() string
	GetClass() *Class
	GetValue() interface{}
}

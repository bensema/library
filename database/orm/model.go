package orm

type Model interface {
	FieldsName() []string
	FieldsItem() []interface{}
}

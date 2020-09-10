package orm

type Model interface {
	Table() string         // 表名
	Columns() []string     // 字段名列表
	Fields() []interface{} // 字段列表
}

type BSql struct {
	Query string
	Args  []interface{}
}

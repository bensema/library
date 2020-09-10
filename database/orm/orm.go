package orm

import "fmt"

type Orm struct {
	table    string                                // 表名
	callback func(orm *Orm, m Model) (BSql, error) // 回调

	selects []string // 选定指定字段查询
	limit   int      // 限制
	offset  int      // 游标
	order   []string
	group   []string

	whereCond []string
	whereArgs []interface{}
}

func New() *Orm {
	return &Orm{}
}

func (o *Orm) Table(table string) *Orm {
	o.table = table
	return o
}

func (o *Orm) Create() *Orm {
	o.callback = create
	return o
}

func (o *Orm) Delete() *Orm {
	o.callback = delete
	return o
}

func (o *Orm) Update() *Orm {
	o.callback = update
	return o
}

func (o *Orm) Query() *Orm {
	o.callback = query
	return o
}

func (o *Orm) OrderBy(columns ...string) *Orm {
	o.order = append(o.order, columns...)
	return o
}

func (o *Orm) Limit(n int) *Orm {
	o.limit = n
	return o
}

func (o *Orm) Offset() *Orm {
	return o
}

func (o *Orm) Do(m Model) (bs BSql, err error) {
	bs, err = o.callback(o, m)
	return
}

func Asc(column string) string {
	return fmt.Sprintf(" %s ASC ", column)
}

func Desc(column string) string {
	return fmt.Sprintf(" %s DESC ", column)

}

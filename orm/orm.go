package orm

type Orm struct {
	table    string
	buildSql string
	callback func(orm *Orm) error

	selects []string

	whereCond []string
	whereArgs []interface{}
	args      []interface{}

	model Model
}

func New() *Orm {
	return &Orm{}
}

func (o *Orm) Table(t string) *Orm {
	o.table = t
	return o
}

func (o *Orm) Create() *Orm {
	o.callback = Create
	return o
}

func (o *Orm) Delete() *Orm {
	o.callback = Delete
	return o
}

func (o *Orm) Update() *Orm {
	o.callback = Update
	return o
}

func (o *Orm) Query() *Orm {
	o.callback = Query
	return o
}

func (o *Orm) Order() *Orm {
	return o
}

func (o *Orm) Limit() *Orm {
	return o
}

func (o *Orm) Offset() *Orm {
	return o
}

func (o *Orm) Gen(m Model) (err error) {
	o.model = m
	err = o.callback(o)
	return
}

func (o *Orm) String() string {
	return o.buildSql
}

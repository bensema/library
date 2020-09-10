package orm

func (o *Orm) Select(columns ...string) *Orm {
	o.callback = query
	for _, column := range columns {
		o.selects = append(o.selects, column)
	}
	return o
}

package orm

func (o *Orm) Select(fields ...string) *Orm {
	o.callback = Query
	for _, filed := range fields {
		o.selects = append(o.selects, filed)
	}
	return o
}

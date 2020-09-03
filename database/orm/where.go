package orm

func (o *Orm) Where(query string, args ...interface{}) *Orm {
	o.whereCond = append(o.whereCond, query)
	o.whereArgs = append(o.whereArgs, args...)
	return o
}

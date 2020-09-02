package orm

func CheckTable(o *Orm) error {
	if o.table == "" {
		return ErrTable
	}
	return nil
}

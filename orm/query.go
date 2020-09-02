package orm

import (
	"fmt"
	"strings"
)

func Query(o *Orm) (err error) {
	if err = CheckTable(o); err != nil {
		return
	}

	if len(o.selects) != 0 {
		o.buildSql = fmt.Sprintf("SELECT %s ", strings.Join(o.selects, ","))
	} else {
		o.buildSql = fmt.Sprintf("SELECT %s ", strings.Join(o.model.FieldsName(), ","))
	}
	o.buildSql += fmt.Sprintf(" FROM %s ", o.table)

	if len(o.whereCond) != 0 {
		o.buildSql += fmt.Sprintf(" WHERE %s", strings.Join(o.whereCond, " AND "))
	}

	return err
}

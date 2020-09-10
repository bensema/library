package orm

import (
	"fmt"
	"strings"
)

func query(o *Orm, m Model) (bs BSql, err error) {
	if err = CheckTable(o); err != nil {
		return
	}

	if len(o.selects) != 0 {
		bs.Query = fmt.Sprintf("SELECT %s ", strings.Join(o.selects, ","))
	} else {
		bs.Query = fmt.Sprintf("SELECT %s ", strings.Join(m.Columns(), ","))
	}
	bs.Query += fmt.Sprintf(" FROM %s ", o.table)

	if len(o.whereCond) != 0 {
		bs.Query += fmt.Sprintf(" WHERE %s", strings.Join(o.whereCond, " AND "))
	}

	return
}

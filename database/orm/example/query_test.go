package example

import (
	"fmt"
	"github.com/bensema/library/database/orm"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func formatString(str string) string {
	nStr := strings.ToLower(strings.Replace(str, " ", "", -1))
	return nStr
}

func TestQuery(t *testing.T) {
	assert := assert.New(t)

	user := &User{}
	o := orm.New()

	bs, err := o.Table(user.Table()).Query().Where("id = ?", "1").Do(user)
	if err != nil {
		t.Errorf("Select error:%s", err)
	}
	actualSql := fmt.Sprintf("SELECT %s FROM %s WHERE id=?", strings.Join(user.Columns(), ","), user.Table())
	assert.Equal(formatString(bs.Query), formatString(actualSql), "they should be equal")
}

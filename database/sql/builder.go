package sql

import (
	"github.com/facebook/ent/dialect"
	entSql "github.com/facebook/ent/dialect/sql"
)

func SqlBuilder() *entSql.DialectBuilder {
	return entSql.Dialect(dialect.MySQL)
}

package lunar_test

import (
	"database/sql"
	"testing"

	"github.com/chonglou/destiny/lunar"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestPgSqlQuery(t *testing.T) {
	db, err := sql.Open("postgres", "user=postgres dbname=lunar_test_1 sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	testQuery(t, &lunar.PgSQLQuery{Db: db})
}

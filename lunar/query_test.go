package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestQuery(t *testing.T) {
	db, err := gorm.Open("postgres", "user=postgres dbname=lunar_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	q := lunar.Query{Db: db}
	if e := q.Load("tmp"); e != nil {
		t.Fatal(e)
	}

}

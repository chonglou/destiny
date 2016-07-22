package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestCrawler(t *testing.T) {
	db, err := gorm.Open("postgres", "user=postgres dbname=lunar_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(false)

	var qu lunar.Query
	qu = &lunar.GormQuery{Db: db}

	var c lunar.Crawler
	e := c.Fetch()
	if e == nil {
		e = qu.Prepare()
	}
	if e == nil {
		e = c.Store(qu)
	}
	if e != nil {
		t.Fatal(e)
	}

}

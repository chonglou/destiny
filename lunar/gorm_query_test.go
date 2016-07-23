package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestGormQuery(t *testing.T) {
	db, err := gorm.Open("postgres", "user=postgres dbname=lunar_test_0 sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(false)
	testCrawler(t, &lunar.GormQuery{Db: db})
}

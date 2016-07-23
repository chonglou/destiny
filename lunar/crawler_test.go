package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
)

func testCrawler(t *testing.T, q lunar.Query) {

	var c lunar.Crawler
	e := c.Fetch()
	if e == nil {
		e = q.Prepare()
	}
	if e == nil {
		e = c.Store(q)
	}
	if e != nil {
		t.Fatal(e)
	}

}

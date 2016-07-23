package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
)

func testQuery(t *testing.T, q lunar.Query) {
	testCrawler(t, q)
	if d, e := q.FromSolar(2016, 7, 22); e == nil {
		t.Logf("%+v", d)
	} else {
		t.Fatal(e)
	}
	if ds, e := q.FromLunar(2016, 6, 19); e == nil {
		t.Logf("%+v", ds)
	} else {
		t.Fatal(e)
	}
}

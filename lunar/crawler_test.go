package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
)

func TestCrawler(t *testing.T) {
	c := lunar.Crawler{}
	for i := 1901; i <= 2100; i++ {
		if e := c.Fetch(i); e == nil {
			t.Logf("get data of year %d", i)
		} else {
			t.Fatal(e)
		}
	}
}

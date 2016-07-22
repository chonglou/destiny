package lunar_test

import (
	"testing"

	"github.com/chonglou/destiny/lunar"
)

func TestCrawler(t *testing.T) {
	var c lunar.Crawler
	if e := c.Fetch(); e != nil {
		t.Fatal(e)
	}

}

package destiny_test

import (
	"testing"
	"time"

	"github.com/chonglou/destiny"
	"github.com/chonglou/destiny/dizhi"
	"github.com/chonglou/destiny/tiangan"
)

func TestCalendar(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Date(2016, 7, 22, 1, 0, 0, 0, loc)
	cal := destiny.Calendar{T: now}
	if tg, dz, err := cal.Day(); err == nil {
		t.Logf("日干支： %s%s", tg, dz)
		if tg != tiangan.Yi && dz != dizhi.Wei {
			t.Fatalf("want %s%s, get %s%s", tiangan.Yi, dizhi.Wei, tg, dz)
		}
	} else {
		t.Fatal(err)
	}

}

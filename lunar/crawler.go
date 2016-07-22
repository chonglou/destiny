package lunar

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//Crawler 爬虫 抓取香港天文台数据
type Crawler struct {
	lunar map[int]int
}

//Fetch 获得日历数据
func (p *Crawler) Fetch() error {
	for y := BEGIN; y <= END; y++ {
		fn := filename(y)
		e := p.get(y, fn)
		if e != nil {
			return e
		}
	}
	return nil
}

func (p *Crawler) get(year int, file string) error {
	if _, err := os.Stat(file); err == nil {
		return err
	}

	resp, err := http.Get(fmt.Sprintf("http://data.weather.gov.hk/gts/time/calendar/text/T%dc.txt", year))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	enc := traditionalchinese.Big5
	in := transform.NewReader(resp.Body, enc.NewDecoder())
	_, err = io.Copy(out, in)
	return err
}

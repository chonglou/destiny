package lunar

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//Crawler 爬虫 抓取香港天文台数据
type Crawler struct {
}

//Fetch 获得日历数据
func (p *Crawler) Fetch(year int) error {
	dir := "tmp"
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	fn := path.Join(dir, fmt.Sprintf("%d.txt", year))
	if _, err := os.Stat(fn); err == nil {
		return err
	}

	resp, err := http.Get(fmt.Sprintf("http://data.weather.gov.hk/gts/time/calendar/text/T%dc.txt", year))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer out.Close()
	enc := traditionalchinese.Big5
	in := transform.NewReader(resp.Body, enc.NewDecoder())
	_, err = io.Copy(out, in)
	return err
}

//Store generate data tables
func (p *Crawler) Store(name string) error {
	return nil
}

package lunar

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//Crawler 爬虫 抓取香港天文台数据
type Crawler struct {
}

//Store store to database
func (p *Crawler) Store(q Query) error {
	q.Clear()

	last := Date{
		LYear:  1900,
		LMonth: 11,
	}
	for y := BEGIN; y <= END; y++ {
		if err := p.read(y, &last, q); err != nil {
			return err
		}
	}
	return nil
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

func (p *Crawler) current(day string, last *Date) error {
	//mth := last.LMonth
	switch day {
	case "初二":
		last.LDay = 2
	case "初三":
		last.LDay = 3
	case "初四":
		last.LDay = 4
	case "初五":
		last.LDay = 5
	case "初六":
		last.LDay = 6
	case "初七":
		last.LDay = 7
	case "初八":
		last.LDay = 8
	case "初九":
		last.LDay = 9
	case "初十":
		last.LDay = 10
	case "十一":
		last.LDay = 11
	case "十二":
		last.LDay = 12
	case "十三":
		last.LDay = 13
	case "十四":
		last.LDay = 14
	case "十五":
		last.LDay = 15
	case "十六":
		last.LDay = 16
	case "十七":
		last.LDay = 17
	case "十八":
		last.LDay = 18
	case "十九":
		last.LDay = 19
	case "二十":
		last.LDay = 20
	case "廿一":
		last.LDay = 21
	case "廿二":
		last.LDay = 22
	case "廿三":
		last.LDay = 23
	case "廿四":
		last.LDay = 24
	case "廿五":
		last.LDay = 25
	case "廿六":
		last.LDay = 26
	case "廿七":
		last.LDay = 27
	case "廿八":
		last.LDay = 28
	case "廿九":
		last.LDay = 29
	case "三十":
		last.LDay = 30
	default:
		last.LDay = 1
		last.Leap = strings.HasPrefix(day, "閏")

		switch day {
		case "正月":
			last.LMonth = 1
			last.LYear++
		case "二月":
			last.LMonth = 2
		case "三月":
			last.LMonth = 3
		case "四月":
			last.LMonth = 4
		case "五月":
			last.LMonth = 5
		case "六月":
			last.LMonth = 6
		case "七月":
			last.LMonth = 7
		case "八月":
			last.LMonth = 8
		case "九月":
			last.LMonth = 9
		case "十月":
			last.LMonth = 10
		case "十一月":
			last.LMonth = 11
		case "十二月":
			last.LMonth = 12
		case "閏二月":
			last.LMonth = 2
		case "閏三月":
			last.LMonth = 3
		case "閏四月":
			last.LMonth = 4
		case "閏五月":
			last.LMonth = 5
		case "閏六月":
			last.LMonth = 6
		case "閏七月":
			last.LMonth = 7
		case "閏八月":
			last.LMonth = 8
		case "閏九月":
			last.LMonth = 9
		case "閏十月":
			last.LMonth = 10
		case "閏十一月":
			last.LMonth = 11
		default:
			return fmt.Errorf("unknown day %s", day)
		}
	}
	return nil
}

func (p *Crawler) term(t string) (int, error) {
	switch t {
	case "立春":
		return 1, nil
	case "雨水":
		return 2, nil
	case "驚蟄":
		return 3, nil
	case "春分":
		return 4, nil
	case "清明":
		return 5, nil
	case "穀雨":
		return 6, nil
	case "立夏":
		return 7, nil
	case "小滿":
		return 8, nil
	case "芒種":
		return 9, nil
	case "夏至":
		return 10, nil
	case "小暑":
		return 11, nil
	case "大暑":
		return 12, nil
	case "立秋":
		return 13, nil
	case "處暑":
		return 14, nil
	case "白露":
		return 15, nil
	case "秋分":
		return 16, nil
	case "寒露":
		return 17, nil
	case "霜降":
		return 18, nil
	case "立冬":
		return 19, nil
	case "小雪":
		return 20, nil
	case "大雪":
		return 21, nil
	case "冬至":
		return 22, nil
	case "小寒":
		return 23, nil
	case "大寒":
		return 24, nil
	default:
		return 0, fmt.Errorf("unknown term %s", t)

	}
}

func (p *Crawler) read(year int, last *Date, query Query) error {
	fd, err := os.Open(filename(year))
	if err != nil {
		return err
	}
	defer fd.Close()

	sc := bufio.NewScanner(fd)
	re := regexp.MustCompile(`(?P<s_year>[[:digit:]]{4})年(?P<s_month>[[:digit:]]{1,2})月(?P<s_day>[[:digit:]]{1,2})日[[:space:]]+(?P<l_day>[\p{Han}]{2,4})[[:space:]]+[\p{Han}]{3}[[:space:]]+(?P<term>[\p{Han}]*)[[:space:]]+`)

	for sc.Scan() {
		line := sc.Text()
		args := re.FindAllStringSubmatch(sc.Text(), -1)
		if len(args) == 1 {
			rst := args[0][1:]
			var dt Date
			if dt.SYear, err = strconv.Atoi(rst[0]); err != nil {
				return err
			}
			if dt.SMonth, err = strconv.Atoi(rst[1]); err != nil {
				return err
			}
			if dt.SDay, err = strconv.Atoi(rst[2]); err != nil {
				return err
			}
			if err = p.current(rst[3], last); err != nil {
				return err
			}
			// if dt.Term, err = strconv.Atoi(rst[4]); err != nil {
			// 	return err
			// }
			dt.Term = rst[4]

			dt.LYear = last.LYear
			dt.LMonth = last.LMonth
			dt.LDay = last.LDay
			dt.Leap = last.Leap

			//log.Printf("%q", rst)
			//log.Printf("%+v", dt)

			if err = query.Write(&dt); err != nil {
				log.Printf(line)
				return err
			}

		} else {
			log.Printf("ingnore line '%s'", line)
		}

	}
	return nil
}

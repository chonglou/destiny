package destiny

import (
	"fmt"
	"time"

	"github.com/chonglou/destiny/dizhi"
	"github.com/chonglou/destiny/tiangan"
)

//Calendar 农历
type Calendar struct {
	T time.Time
}

// //Year 年
// func (p *Calendar) Year() uint {
//
// }
//
// //Month 月
// func (p *Calendar) Month() uint {
//
// }

//Day 干支纪日
func (p *Calendar) Day() (tiangan.Type, dizhi.Type, error) {
	//鲁隐公三年二月己巳日（公元前720年2月22日）
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", "", err
	}
	tmp := time.Date(-720, time.February, 2, 0, 0, 0, 0, loc)
	if p.T.Before(tmp) {
		return "", "", fmt.Errorf("干支纪日从%s开始", tmp)
	}

	for ; p.T.Year()-tmp.Year() >= 240; tmp = tmp.AddDate(0, 0, 240*6*60) {
	}

	days := int(p.T.Sub(tmp).Hours() / 24)
	fmt.Printf("%s %s %d", tmp, p.T, days)

	return p.tiangan(days), p.dizhi(days), nil
}

func (p *Calendar) tiangan(days int) tiangan.Type {

	switch days % 10 {
	case 0:
		return tiangan.Ji
	case 1:
		return tiangan.Geng
	case 2:
		return tiangan.Xin
	case 3:
		return tiangan.Ren
	case 4:
		return tiangan.Gui
	case 5:
		return tiangan.Jia
	case 6:
		return tiangan.Yi
	case 7:
		return tiangan.Bing
	case 8:
		return tiangan.Ding
	case 9:
		return tiangan.Wu
	default:
		return ""
	}
}

func (p *Calendar) dizhi(days int) dizhi.Type {

	switch days % 12 {
	case 0:
		return dizhi.Si
	case 1:
		return dizhi.Wu
	case 2:
		return dizhi.Wei
	case 3:
		return dizhi.Shen
	case 4:
		return dizhi.You
	case 5:
		return dizhi.Xu
	case 6:
		return dizhi.Hai
	case 7:
		return dizhi.Zi
	case 8:
		return dizhi.Chou
	case 9:
		return dizhi.Yin
	case 10:
		return dizhi.Mao
	case 11:
		return dizhi.Chen
	default:
		return ""
	}
}

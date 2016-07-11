package tiangan

//Type 天干
type Type uint

const (
	_ Type = iota
	//Jia 甲
	Jia Type = iota
	//Yi 乙
	Yi Type = iota
	//Bing 丙
	Bing Type = iota
	//Ding 丁
	Ding Type = iota
	//Wu 戊
	Wu Type = iota
	//Ji 己
	Ji Type = iota
	//Geng 庚
	Geng Type = iota
	//Xin 辛
	Xin Type = iota
	//Ren 壬
	Ren Type = iota
	//Gui 癸
	Gui Type = iota
)

func (p Type) String() string {
	switch p {
	case Jia:
		return "甲"
	case Yi:
		return "乙"
	case Bing:
		return "丙"
	case Ding:
		return "丁"
	case Wu:
		return "戊"
	case Ji:
		return "己"
	case Geng:
		return "庚"
	case Xin:
		return "辛"
	case Ren:
		return "壬"
	case Gui:
		return "癸"
	}
	return ""
}

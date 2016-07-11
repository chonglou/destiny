package wuxing

//Type 五行
type Type uint

const (
	_ Type = iota
	//Jin 金
	Jin Type = iota
	//Mu 木
	Mu Type = iota
	//Shui 水
	Shui Type = iota
	//Huo 火
	Huo Type = iota
	//Tu 土
	Tu Type = iota
)

func (p Type) String() string {
	switch p {
	case Jin:
		return "金"
	case Mu:
		return "木"
	case Shui:
		return "水"
	case Huo:
		return "火"
	case Tu:
		return "土"
	}
	return ""
}

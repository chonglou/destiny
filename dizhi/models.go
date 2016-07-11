package dizhi

//Type 地支
type Type uint

const (
	_ Type = iota
	//Zi 子
	Zi Type = iota
	//Chou 丑
	Chou Type = iota
	//Yin 寅
	Yin Type = iota
	//Mao 卯
	Mao Type = iota
	//Chen 辰
	Chen Type = iota
	//Si 巳
	Si Type = iota
	//Wu 午
	Wu Type = iota
	//Wei 未
	Wei Type = iota
	//Shen 申
	Shen Type = iota
	//You 酉
	You Type = iota
	//Xu 戌
	Xu Type = iota
	//Hai 亥
	Hai Type = iota
)

func (p Type) String() string {
	switch p {
	case Zi:
		return "子"
	case Chou:
		return "丑"
	case Yin:
		return "寅"
	case Mao:
		return "卯"
	case Chen:
		return "辰"
	case Si:
		return "巳"
	case Wu:
		return "午"
	case Wei:
		return "未"
	case Shen:
		return "申"
	case You:
		return "酉"
	case Xu:
		return "戌"
	case Hai:
		return "亥"
	}
	return ""
}

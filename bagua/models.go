package bagua

//Type 八卦
type Type uint

const (
	_ = iota
	//Qian
	Qian = iota
	Dui  = iota
	Li   = iota
	Zhen = iota
	Xun  = iota
	Kan  = iota
	Gen  = iota
	Kun  = iota
)

func (p Type) String() string {
	switch p {
	case Qian:
		return "乾"
	case Dui:
		return "兌"
	case Li:
		return "離"
	case Zhen:
		return "震"
	case Xun:
		return "巽"
	case Kan:
		return "坎"
	case Gen:
		return "艮"
	case Kun:
		return "坤"
	default:
		return ""

	}
}

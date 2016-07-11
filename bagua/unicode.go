package bagua

//Unicode unicode编码
func (p Type) Unicode() string {
	switch p {
	case Qian:
		return "☰"
	case Dui:
		return "☱"
	case Li:
		return "☲"
	case Zhen:
		return "☳"
	case Xun:
		return "☴"
	case Kan:
		return "☵"
	case Gen:
		return "☶"
	case Kun:
		return "☷"
	default:
		return ""
	}
}

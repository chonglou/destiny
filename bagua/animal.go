package bagua

//Animal 动物
func (p Type) Animal() string {
	switch p {
	case Qian:
		return "马"
	case Dui:
		return "羊"
	case Li:
		return "雉"
	case Zhen:
		return "龙"
	case Xun:
		return "鸡"
	case Kan:
		return "豕"
	case Gen:
		return "狗"
	case Kun:
		return "牛"
	default:
		return ""

	}
}

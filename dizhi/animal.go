package dizhi

//Animal 生肖
func (p Type) Animal() string {
	switch p {
	case Zi:
		return "鼠"
	case Chou:
		return "牛"
	case Yin:
		return "虎"
	case Mao:
		return "兔"
	case Chen:
		return "龙"
	case Si:
		return "蛇"
	case Wu:
		return "马"
	case Wei:
		return "羊"
	case Shen:
		return "猴"
	case You:
		return "鸡"
	case Xu:
		return "狗"
	case Hai:
		return "猪"
	default:
		return ""

	}
}

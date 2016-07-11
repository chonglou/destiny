package dizhi

//Month 农历月份
func (p Type) Month() uint {
	switch p {
	case Zi:
		return 11
	case Chou:
		return 12
	case Yin:
		return 1
	case Mao:
		return 2
	case Chen:
		return 3
	case Si:
		return 4
	case Wu:
		return 5
	case Wei:
		return 6
	case Shen:
		return 7
	case You:
		return 8
	case Xu:
		return 9
	case Hai:
		return 10
	default:
		return 0

	}
}

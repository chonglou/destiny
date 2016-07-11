package dizhi

//Hours 时辰
func (p Type) Hours() (uint, uint) {
	switch p {
	case Zi:
		return 23, 1
	case Chou:
		return 1, 3
	case Yin:
		return 3, 5
	case Mao:
		return 5, 7
	case Chen:
		return 7, 9
	case Si:
		return 9, 11
	case Wu:
		return 11, 13
	case Wei:
		return 13, 15
	case Shen:
		return 15, 17
	case You:
		return 17, 19
	case Xu:
		return 19, 21
	case Hai:
		return 21, 23
	default:
		return 0, 0

	}
}

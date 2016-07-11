package dizhi

import "github.com/chonglou/destiny/wuxing"

//WuXing 五行
func (p Type) WuXing() wuxing.Type {
	switch p {
	case Zi:
		return wuxing.Shui
	case Chou:
		return wuxing.Tu
	case Yin:
		return wuxing.Mu
	case Mao:
		return wuxing.Mu
	case Chen:
		return wuxing.Tu
	case Si:
		return wuxing.Huo
	case Wu:
		return wuxing.Huo
	case Wei:
		return wuxing.Tu
	case Shen:
		return wuxing.Jin
	case You:
		return wuxing.Jin
	case Xu:
		return wuxing.Tu
	case Hai:
		return wuxing.Shui
	default:
		return ""
	}
}

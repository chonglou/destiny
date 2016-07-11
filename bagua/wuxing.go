package bagua

import "github.com/chonglou/destiny/wuxing"

//WuXing 五行
func (p Type) WuXing() wuxing.Type {
	switch p {
	case Qian:
		return wuxing.Jin
	case Dui:
		return wuxing.Jin
	case Li:
		return wuxing.Huo
	case Zhen:
		return wuxing.Mu
	case Xun:
		return wuxing.Mu
	case Kan:
		return wuxing.Shui
	case Gen:
		return wuxing.Tu
	case Kun:
		return wuxing.Tu
	default:
		return ""
	}
}

package tiangan

import "github.com/chonglou/destiny/wuxing"

//WuXing 五行
func (p Type) WuXing() wuxing.Type {
	switch p {
	case Jia:
		return wuxing.Mu
	case Yi:
		return wuxing.Mu
	case Bing:
		return wuxing.Huo
	case Ding:
		return wuxing.Huo
	case Wu:
		return wuxing.Tu
	case Ji:
		return wuxing.Tu
	case Geng:
		return wuxing.Jin
	case Xin:
		return wuxing.Jin
	case Ren:
		return wuxing.Shui
	case Gui:
		return wuxing.Shui
	default:
		return ""
	}
}

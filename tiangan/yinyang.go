package tiangan

import "github.com/chonglou/destiny/yao"

//YinYang 阴阳
func (p Type) YinYang() yao.Type {
	switch p {
	case Jia:
		return yao.Yang
	case Yi:
		return yao.Yin
	case Bing:
		return yao.Yang
	case Ding:
		return yao.Yin
	case Wu:
		return yao.Yang
	case Ji:
		return yao.Yin
	case Geng:
		return yao.Yang
	case Xin:
		return yao.Yin
	case Ren:
		return yao.Yang
	case Gui:
		return yao.Yin
	default:
		return ""
	}
}

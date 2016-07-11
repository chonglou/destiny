package dizhi

import (
	"github.com/chonglou/destiny/yinyang"
)

//YinYang 阴阳
func (p Type) YinYang() yinyang.Type {
	switch p {
	case Zi:
		return yinyang.Yang
	case Chou:
		return yinyang.Yin
	case Yin:
		return yinyang.Yang
	case Mao:
		return yinyang.Yin
	case Chen:
		return yinyang.Yang
	case Si:
		return yinyang.Yin
	case Wu:
		return yinyang.Yang
	case Wei:
		return yinyang.Yin
	case Shen:
		return yinyang.Yang
	case You:
		return yinyang.Yin
	case Xu:
		return yinyang.Yang
	case Hai:
		return yinyang.Yin
	default:
		return ""
	}
}

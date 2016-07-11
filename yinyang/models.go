package yinyang

//Type 阴阳
type Type uint

const (
	//Yin 阴
	Yin Type = iota
	//Yang 阳
	Yang Type = iota
)

func (p Type) String() string {
	switch p {
	case Yin:
		return "阴"
	case Yang:
		return "阳"
	}
	return ""
}

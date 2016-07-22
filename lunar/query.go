package lunar

//Query query tools
type Query interface {
	Prepare() error
	Write(*Date) error
	Clear() error
}

package lunar

//Query query tools
type Query interface {
	Prepare() error
	Write(*Date) error
	Clear() error
	FromSolar(y, m, d int) (Date, error)
	FromLunar(y, m, d int) ([]Date, error)
}

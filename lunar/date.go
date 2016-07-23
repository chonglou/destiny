package lunar

const (
	//BEGIN 开始年份
	BEGIN = 1901
	//END 结束年份
	END = 2100
)

//Date 中国农历-公历数据
type Date struct {
	ID uint `gorm:"primary_key"`

	SYear  int `gorm:"type:smallint;not null"`
	SMonth int `gorm:"type:smallint;not null"`
	SDay   int `gorm:"type:smallint;not null"`
	LYear  int `gorm:"type:smallint;not null"`
	LMonth int `gorm:"type:smallint;not null"`
	LDay   int `gorm:"type:smallint;not null"`

	Term string `gorm:"type:varchar(12);index"`
	//Term int  `gorm:"not null"`
	Leap bool `gorm:"not null"`
}

//TableName set Date's table name
func (Date) TableName() string {
	return "lunar_items"
}

package lunar

import (
	"fmt"
	"os"
	"path"
	"time"
)

const (
	//BEGIN 开始年份
	BEGIN = 1901
	//END 结束年份
	END = 2100
)

func filename(year int) string {
	dir := path.Join(os.TempDir(), "lunar")
	os.MkdirAll(dir, 0700)
	return path.Join(dir, fmt.Sprintf("%d.txt", year))
}

//Date 中国农历-公历数据
type Date struct {
	ID uint `gorm:"primary_key"`

	SYear  int `gorm:"not null"`
	SMonth int `gorm:"not null"`
	SDay   int `gorm:"not null"`
	LYear  int `gorm:"not null"`
	LMonth int `gorm:"not null"`
	LDay   int `gorm:"not null"`

	Term string `gorm:"type:varchar(16);index"`
	Leap bool   `gorm:"not null"`

	CreatedAt time.Time
}

//TableName set Date's table name
func (Date) TableName() string {
	return "lunar_dates"
}

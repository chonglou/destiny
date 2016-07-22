package lunar

import "github.com/jinzhu/gorm"

//GormQuery 查询工具类
type GormQuery struct {
	Db *gorm.DB
}

//Clear clear table items
func (p *GormQuery) Clear() error {
	return p.Db.Delete(&Date{}).Error
}

//Prepare create table and indexes
func (p *GormQuery) Prepare() error {
	p.Db.AutoMigrate(&Date{})
	p.Db.Model(&Date{}).AddUniqueIndex("idx_lunar_dates_s", "s_year", "s_month", "s_day")
	p.Db.Model(&Date{}).AddUniqueIndex("idx_lunar_dates_l", "l_year", "l_month", "l_day", "leap")
	return nil
}

//Write create a new record
func (p *GormQuery) Write(item *Date) error {
	return p.Db.Create(item).Error
}

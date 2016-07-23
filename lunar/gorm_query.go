package lunar

import "github.com/jinzhu/gorm"

//GormQuery query by gorm
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
	p.Db.Model(&Date{}).AddUniqueIndex("idx_lunar_items_s", "s_year", "s_month", "s_day")
	p.Db.Model(&Date{}).AddUniqueIndex("idx_lunar_items_l", "l_year", "l_month", "l_day", "leap")
	return nil
}

//Write create a new record
func (p *GormQuery) Write(item *Date) error {
	return p.Db.Create(item).Error
}

//FromSolar 从公历查询
func (p *GormQuery) FromSolar(y, m, d int) (Date, error) {
	var item Date
	err := p.Db.
		Where("s_year = ? AND s_month = ? AND s_day = ?", y, m, d).
		First(&item).Error
	return item, err
}

//FromLunar 从农历查询
func (p *GormQuery) FromLunar(y, m, d int) ([]Date, error) {
	var items []Date
	err := p.Db.
		Where("l_year = ? AND l_month = ? AND l_day = ?", y, m, d).
		Order("ID ASC").
		Find(&items).Error
	return items, err
}

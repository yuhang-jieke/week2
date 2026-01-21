package model

import "gorm.io/gorm"

type Goodss struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(30);comment:名称"`
	Price float64 `gorm:"type:decimal(10,2);comment:价格"`
	Num   int     `gorm:"type:int(11);comment:数量"`
	Title string  `gorm:"type:varchar(30);comment:标题"`
}

func (g *Goodss) Create(db *gorm.DB) error {
	return db.Create(&g).Error
}

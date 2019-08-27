package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Dao dao
type Dao struct {
	db *gorm.DB
}

// New new dao
func New() (dao *Dao) {
	orm, err := gorm.Open("mysql", "root:123456ab@/pp-stand?charset=utf8&parseTime=True&loc=Local")
	orm.SingularTable(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	dao = &Dao{
		db: orm,
	}

	// dao.db.LogMode(true)

	return
}

// Close Close dao
func (dao *Dao) Close() {
	dao.db.Close()
}

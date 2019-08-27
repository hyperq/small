package dao

import (
	"fmt"
	"gin-t/api/dept/model"
)

//QueryDeptByID 根据id获取数据
func (dao *Dao) QueryDeptByID(id interface{}) (s *model.Ee_dept, err error) {
	s = &model.Ee_dept{}
	err = dao.db.Where("id=?", id).First(s).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*s)
	err = dao.db.Model(s).Related(s.Member).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

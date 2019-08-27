package dao

import (
	"fmt"
	"gin-t/api/member/model"
)

//QueryMmeberByID 根据id获取数据
func (dao *Dao) QueryMmeberByID(id interface{}) (member *model.Member, err error) {
	member = &model.Member{}
	err = dao.db.Where("id=?", id).First(member).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

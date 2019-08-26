package dao

import (
	"fmt"
	"small/srv/user/model"
)

//QueryUserByID 根据id获取数据
func (dao *Dao) QueryUserByID(id int64) (s *model.User, err error) {
	s = &model.User{}
	err = dao.db.Where("id=?", id).First(s).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//QueryUserList 获取用户数据
func (dao *Dao) QueryUserList(page, promotionID int64, ob string) (s []model.User, err error) {
	if promotionID != 0 {
		err = dao.db.Where("promotion_id=?", promotionID).Order(ob).Offset(page * 10).Limit(10).Find(s).Error
	} else {
		err = dao.db.Order(ob).Offset(page * 10).Limit(10).Find(&s).Error
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

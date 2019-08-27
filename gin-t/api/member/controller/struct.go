package member

import "gin-t/api/member/dao"

//Member 用户
type Member struct {
	d *dao.Dao
}

//New New Member
func New() (m *Member) {
	return &Member{
		d: dao.New(),
	}
}

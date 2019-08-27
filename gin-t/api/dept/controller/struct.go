package dept

import "gin-t/api/dept/dao"

//Dept 用户
type Dept struct {
	d *dao.Dao
}

//New New Dept
func New() (m *Dept) {
	return &Dept{
		d: dao.New(),
	}
}

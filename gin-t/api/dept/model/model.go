package model

import (
	"time"
)

// Ee_dept 部门struct
type Ee_dept struct {
	ID       int64     `gorm:"primary_key" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	UID      int64     `gorm:"column:uid" json:"uid"`
	GID      int64     `gorm:"column:gid" json:"gid"`
	DeadLine time.Time `gorm:"column:deadline" json:"deadline"`
	Addtime  time.Time `gorm:"column:addtime" json:"addtime"`
	Status   int64     `gorm:"column:status" json:"status"`
}

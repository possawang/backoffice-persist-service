package entity

import (
	"github.com/possawang/go-persist-lib-common/softdelete"
)

type User struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	Username string `gorm:"size:20;unique"`
	Password string `gorm:"type:text"`
	Role     Role   `gorm:"references:id;column:role_id"`
	BranchId uint64 `gorm:""`
}

type Allowed struct {
	softdelete.SoftDeleteModel
	Role     Role `gorm:"references:id;column:role_id"`
	Endpoint string
	Method   string
}

type Role struct {
	softdelete.SoftDeleteModel
	Id   uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	Name string `gorm:"column:name;size:20"`
}

package models

import "github.com/saitofun/items/pkg/depends"

//go:generate tools gen model2 User --database DB --with-comments
// User 用户登陆信息
// @def primary ID
// @def unique_index I_username Username
// @def unique_index I_user_id  UserID

type User struct {
	PrimaryID
	UserRef
	UserBase
	OperationTimes
}

type UserRef struct {
	UserID depends.SFID `db:"f_user_id" json:"userID"` // 用户编号(UUID)
}

type UserBase struct {
	Username string `db:"f_username" json:"username"`
	Password string `db:"f_password" json:"password"`
}

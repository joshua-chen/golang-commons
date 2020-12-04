/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 16:26:34
 */
package models

import (
	"time"

)

type User struct {
	Id         int64     `xorm:"pk autoincr bigint notnull" json:"id"  form:"id"`
	Username   string    `json:"username" form:"username"`
	Password   string    `xorm:"notnull" json:"password" form:"password"`
	Token      string    `json:"-"`
	Enabled    int       `xorm:"notnull tinyint(1)" json:"enabled" form:"enabled"`
	Appid      string    `xorm:"notnull" json:"appid" form:"appid"`
	Nickname    string  `xorm:"notnull" json:"nickname" form:"nickname"`
	Phone      string    `xorm:"null" json:"phone" form:"phone"`
	Mobile      string    `xorm:"null" json:"mobile" form:"mobile"`
	QQ      string    `xorm:"null" json:"qq" form:"qq"`
	Email      string    `xorm:"null" json:"email" form:"email"`
	Userface   string    `xorm:"null" json:"userface" form:"userface"`
	CreateAt time.Time  `xorm:"null" json:"create_at" form:"create_at"`
	UpdateAt time.Time  `xorm:"null" json:"update_at" form:"update_at"`
	Roles  []string   `xorm:"-" json:"roles" `
}

 
//
// Table
//
func(m *User) TableName() string {
	return "sys_user"
}
 
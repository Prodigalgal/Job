package beans

import "time"

type User struct {
	Id         int       `xorm:"not null autoincr pk 'id'" json:"-"`
	Username   string    `xorm:"not null varchar(255) 'username'" json:"username"`
	Password   string    `xorm:"not null varchar(255) 'password'" json:"password"`
	CreateTime time.Time `xorm:"not null created 'create_time'" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated 'update_time'" json:"update_time"`
}

type SecurityUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

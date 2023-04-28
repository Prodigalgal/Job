package beans

import "time"

type Classify struct {
	Id         int       `xorm:"not null autoincr pk 'id'" json:"id"`
	Author     string    `xorm:"varchar(255) 'author'" json:"author"`
	UpdateTime time.Time `xorm:"not null updated 'update_time'" json:"update_time"`
	CreateTime time.Time `xorm:"created 'creat_time'" json:"creat_time"`
	Name       string    `xorm:"not null varchar(255) 'name'" json:"name"`
}

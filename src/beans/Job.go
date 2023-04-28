package beans

import "time"

type Job struct {
	Id         int       `xorm:"not null autoincr pk 'id'" json:"id"`
	Name       string    `xorm:"not null varchar(255) 'name'" json:"name"`
	CreateTime time.Time `xorm:"not null created 'create_time'" json:"create_time"`
	UpdateTime time.Time `xorm:"not null updated 'update_time'" json:"update_time"`
	Path       string    `xorm:"varchar(255) 'path'" json:"path"`
	Author     string    `xorm:"varchar(255) 'author'" json:"author"`
	Dangerous  int       `xorm:"not null 'dangerous'" json:"dangerous"`
}

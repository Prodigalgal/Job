package beans

import "time"

type JobClassifyUnion struct {
	Id         int       `xorm:"not null autoincr pk 'id'" json:"id"`
	JobId      int       `xorm:"not null 'job_id'" json:"job_id"`
	ClassifyId int       `xorm:"not null 'classify_id'" json:"classify_id"`
	CreateTime time.Time `xorm:"not null created 'create_time'" json:"create_time"`
	UpdateTime time.Time `xorm:"updated 'update_time'" json:"update_time"`
}

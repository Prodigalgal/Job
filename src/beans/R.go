package beans

import (
	"github.com/kataras/iris/v12"
	"time"
)

type R struct {
	Data  map[string]interface{} `json:"data,omitempty"`
	Error error                  `json:"error,omitempty"`
	Code  int                    `json:"code"`
}

type RBuilder struct {
	r *R
}

func NewRBuilder() *RBuilder {
	return &RBuilder{r: &R{
		Data: map[string]interface{}{
			"date": time.Now(),
			"code": iris.StatusAccepted,
		},
	}}
}

func (b *RBuilder) SetMessage(message string) *RBuilder {
	b.r.Data["message"] = message
	return b
}

func (b *RBuilder) SetData(data map[string]interface{}) *RBuilder {
	for k, v := range data {
		b.r.Data[k] = v
	}
	return b
}

func (b *RBuilder) SetDate(date time.Time) *RBuilder {
	if !date.IsZero() {
		b.r.Data["date"] = date
	} else {
		b.r.Data["date"] = time.Now()
	}
	return b
}

func (b *RBuilder) SetError(err error) *RBuilder {
	b.r.Error = err
	return b
}

func (b *RBuilder) SetCode(code int) *RBuilder {
	b.r.Code = code
	return b
}

func (b *RBuilder) Build() *R {
	return b.r
}

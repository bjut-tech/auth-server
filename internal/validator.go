package internal

import (
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

type customValidator struct{}

func (cv *customValidator) Validate(i interface{}) error {
	v := validate.New(i)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func init() {
	zhcn.RegisterGlobal()
}

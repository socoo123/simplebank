package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}

	return false
}

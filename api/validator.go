package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool { 
	// interface that contains all information and helper functions to validate a field.
	// reflection value, we have to call .Interface() then we try to convert this value to a string.
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
// Package validation contains the support for validating related struct values
// after decoding web requests.
package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

func init() {

	// Instantiate a validator.
	validate = validator.New()

	// Use JSON tag names for errors instead of Go struct names.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Check validates the provided request struct against it's declared tags.
func Check(val interface{}) error {
	if err := validate.Struct(val); err != nil {
		return err
	}

	return nil
}

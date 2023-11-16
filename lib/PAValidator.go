package lib

import (
	"PAValidator/lib/errors"
	"PAValidator/lib/formats"
	"PAValidator/lib/types"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func Validate(schemaPath string, data interface{}) *errors.Error {
	c := jsonschema.NewCompiler()
	c.AssertFormat = true

	// add format checkers
	c.Formats["type-format"] = formats.TypeFormat

	// create schema
	schema, err := c.Compile(schemaPath)
	if err != nil {
		return errors.NewError(types.ERROR_TYPE_LEXICAL, "Schema file not in json format")
	}

	// validate input file with the schema
	if err := schema.Validate(data); err != nil {
		if verr, ok := err.(*jsonschema.ValidationError); ok {
			var msg string = "\n"
			for i := 0; i < len(verr.BasicOutput().Errors); i++ {
				msg = msg + verr.BasicOutput().Errors[i].KeywordLocation + " " + verr.BasicOutput().Errors[i].Error + "\n"
			}
			return errors.NewError(types.ERROR_TYPE_SYNTAX, msg)
		}
	}

	return nil
}
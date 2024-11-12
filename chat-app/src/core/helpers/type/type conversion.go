package type_conversion

import "strconv"

type ITypeConversion interface {
	// ParseStringToInt parses a string to an integer.
	ParseStringToInt(str string) int
}

// TypeConversion is a struct that implements the ITypeConversion interface.
type TypeConversion struct{}

// NewTypeConversion creates a new TypeConversion.
func NewTypeConversion() ITypeConversion {
	return &TypeConversion{}
}

// ParseStringToInt parses a string to an integer.
func (tc *TypeConversion) ParseStringToInt(str string) int {
	parsedInt, err := strconv.Atoi(str)
	if err != nil {
		panic("Error parsing MAIL_PORT")
	}
	return parsedInt
}

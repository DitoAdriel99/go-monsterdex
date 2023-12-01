package entity

// CustomError is a custom error type that implements the error interface.
type CustomError string

// Error returns the error message for the CustomError.
func (e CustomError) Error() string {
	return string(e)
}

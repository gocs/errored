package errored

// New is a type to handle being a constant
type New string

// Error this implements error interface
func (c New) Error() string {
	return string(c)
}

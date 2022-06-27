package core_utils

// ConstError this is a replacement for errors.New that may be used as a const instead of a variable
type ConstError string

func (e ConstError) Error() string { return string(e) }

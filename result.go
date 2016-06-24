package pulley

// Result represents the result of a command executed via SSH
type Result struct {
	Output []byte
	err    error
}

// String is a convenience method for printing and for getting the stringified output
func (r *Result) String() string {
	return string(r.Output)
}

// Err will return the error that occurred
func (r *Result) Err() error {
	return r.err
}

// Success returns a boolean indicating sucess or failure. true == success
func (r *Result) Success() bool {
	if r.err != nil {
		return false
	}

	return true
}

// Failure returns a boolean indicating sucess or failure. true == failed
func (r *Result) Failure() bool {
	if r.err != nil {
		return true
	}

	return false
}

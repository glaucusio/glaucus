package command

type Error interface {
	error
	ExitCode() int
}

type ExitError struct {
	Err  error
	Code int
}

func (e *ExitError) ExitCode() int {
	switch {
	case e == nil:
		return 0
	case e.Code == 0:
		return -1
	default:
		return e.Code
	}
}

func ExitCode(err error) int {
	switch e := err.(type) {
	case nil:
		return 0
	case Error:
		return e.ExitCode()
	default:
		return -1
	}
}

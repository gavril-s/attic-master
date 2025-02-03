package api

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	FileNotFound      Error = "FileNotFound"
	NotEnoughSpace    Error = "NotEnoughSpace"
	UnexpectedFailure Error = "UnexpectedFailure"
)

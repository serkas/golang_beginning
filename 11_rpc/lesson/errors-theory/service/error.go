package service

var (
	ErrAlreadyExist = LogicError("already exist")
)

type LogicError string

func (e LogicError) Error() string {
	return string(e)
}

type NotFoundError struct {
	msg string
	ID  int64
}

func (e *NotFoundError) Error() string {
	return e.msg
}

func NewNotFoundError(id int64) *NotFoundError {
	return &NotFoundError{
		msg: "entity not found",
		ID:  id,
	}
}

package service

func Process(input string) error {
	switch input {
	case "not_found":
		return NewNotFoundError(123)
	case "conflict":
		return ErrAlreadyExist
	default:
		return nil
	}
}

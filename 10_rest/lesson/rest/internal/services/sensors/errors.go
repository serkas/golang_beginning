package sensors

var ErrEntityNotFound = DataError("entity_not_found")

type DataError string

func (e DataError) Error() string {
	return string(e)
}

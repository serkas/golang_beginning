package entities

import "errors"

var (
	ErrNotFound = errors.New("entity_not_found")
	ErrConflict = errors.New("entity_id_conflict")
)

package models

type Like struct {
	UserID int
	Type   LikeType
}

type LikeType string

const (
	LikeTypeThumbUp LikeType = "thumb_up"
)

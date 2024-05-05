package events

type ArticleLikeDTO struct {
	ArticleID int
	AuthorID  int
	UserID    int
	LikeType  string
}

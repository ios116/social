package entities

import (
     "context"
     "time"
)

type Post struct {
     ID int64
     Content string
     UserID  int64
     Created time.Time
}

type PostQuery struct {
     LastID int64
     Limit  int64
     UserID int64
}

type Poster interface {
     AddPost(ctx context.Context, post *Post) (int64, error)
     DeletePost(ctx context.Context, id int64) error
     Update(ctx context.Context, post *Post) error
     SelectPosts(ctx context.Context, query PostQuery) ([]*Post, error)
}



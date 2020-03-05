package entities

import (
     "context"
)

type Post struct {
     _msgpack struct{} `msgpack:",asArray"`
     ID int64
     UserID  int64
     Content string
     Created string
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



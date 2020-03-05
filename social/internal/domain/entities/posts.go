package entities

import (
     "context"
)

type Post struct {
     _msgpack struct{} `msgpack:",asArray"`
     ID uint64
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
     AddPost(ctx context.Context, post *Post) (uint64, error)
     DeletePost(ctx context.Context, id uint64) error
     Update(ctx context.Context, post *Post) error
     SelectPosts(ctx context.Context, query PostQuery) ([]*Post, error)
}



package storage

import (
	"context"
	"fmt"
	"github.com/tarantool/go-tarantool"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
)
// AddPost add post to tarantool

func (p *Storage) AddPost(ctx context.Context, post *entities.Post) (uint64, error) {
    resp, err := p.TNT.Insert("posts",[]interface{}{nil,post.UserID, post.Content, post.Created})
    if err != nil {
    	return 0, err
	}
    id, ok := resp.Tuples()[0][0].(uint64)
    if !ok {
    	return 0, exceptions.NotInt64
	}
	return id, nil
}

// DeletePost delete post from tarantool
func (p *Storage) DeletePost(ctx context.Context, id uint64) error {
    resp, err := p.TNT.Delete("posts","primary",[]interface{}{id})
    if err !=nil {
    	return err
	}
	fmt.Println(resp.Tuples())
	return nil
}

// Update update post in tarantool by id
func (p *Storage) Update(ctx context.Context, post *entities.Post) error {
    resp, err := p.TNT.Replace("posts",[]interface{}{post.ID,post.UserID, post.Content, post.Created})
    if err != nil {
    	return err
	}
	fmt.Println(resp.Tuples())
    return nil
}

// SelectPosts select post from tarantool by user id and limit with id offset
func (p *Storage) SelectPosts(ctx context.Context, query *entities.PostQuery) ([]*entities.Post, error) {
	var posts []*entities.Post
    err :=p.TNT.SelectTyped("posts","user_id",query.Offset,query.Limit, tarantool.IterGe,[]interface{}{query.UserID},&posts)
    if err != nil {
    	return nil, err
	}
	return posts, err
}

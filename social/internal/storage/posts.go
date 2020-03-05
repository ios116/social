package storage

import (
	"context"
	"social/internal/domain/entities"
	"social/internal/domain/exceptions"
)

func (p *Storage) AddPost(ctx context.Context, post *entities.Post) (uint64, error) {

    resp, err := p.Tar.Insert("posts",[]interface{}{nil,post.UserID, post.Content, post.Created})
    if err != nil {
    	return 0, err
	}
    id, ok := resp.Tuples()[0][0].(uint64)
    if !ok {
    	return 0, exceptions.NotInt64
	}
	return id, nil

}
func (p *Storage) DeletePost(ctx context.Context, id int64) error {

	panic("not implementation")
}
func (p *Storage) Update(ctx context.Context, post *entities.Post) error {

	panic("not implementation")
}
func (p *Storage) SelectPosts(ctx context.Context, query entities.PostQuery) ([]*entities.Post, error) {

	panic("not implementation")
}

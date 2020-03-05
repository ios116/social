package storage

import (
	"context"
	"fmt"
	"social/internal/domain/entities"
)

func (p *Storage) AddPost(ctx context.Context, post *entities.Post) (int64, error) {
	//space := p.Tar.Schema.Spaces["posts"]
    resp, err := p.Tar.Insert("posts",post)
    if err != nil {
    	return 0, err
	}
	fmt.Println("data= ",resp.Data)
	fmt.Println("error= ",resp.Error)
	fmt.Println("code=", resp.Code)
    fmt.Println("reqId= ",resp.RequestId)
	return 0, nil

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
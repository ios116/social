package users

import (
	"context"
	"social/internal/domain/exceptions"
)
// Subscribe method save to bd for subscribe a user to another user
func (p *UserStorage) Subscribe(ctx context.Context, userId int64, subscribeId int64) (int64, error) {
	query := "INSERT INTO subscribers (user_id, subscriber_id) VALUES(:user_id,:subscriber_id)"
	data := map[string]interface{}{
		"user_id":       userId,
		"subscriber_id": subscribeId,
	}
	res, err := p.Db.NamedExecContext(ctx, query, data)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
// UnSubscribe method remove records from bd - unsubscribe user
func (p *UserStorage) UnSubscribe(ctx context.Context, userId int64, subscribeId int64) error {
	query := "DELETE FROM subscribers WHERE user_id=:user_id and subscriber_id=:subscriber_id"
	data := map[string]interface{}{
		"user_id":       userId,
		"subscriber_id": subscribeId,
	}
	res, err := p.Db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return exceptions.ObjectDoesNotExist
	}
	return nil
}

package infrastructure

import "learn-golang/syntax/day_5/domain"

type UserRepoDB struct{}

func (r *UserRepoDB) Save(user domain.User) error {
	return nil
}

func (r *UserRepoDB) FindByID(id int) (domain.User, error) {
	return domain.User{}, nil
}

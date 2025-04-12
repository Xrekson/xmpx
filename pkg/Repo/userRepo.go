package repo

import (
	"context"

	"github.com/xrekson/auction/cmd"
	"github.com/xrekson/auction/pkg/model"
)

type userRepo struct{}

func (r *userRepo) Createuser(user *model.User) error {
	ctx := context.Background()
	_, err := cmd.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *userRepo) Getuser(id int64) (*model.User, error) {
	user := new(model.User)
	ctx := context.Background()
	err := cmd.DB.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	return user, err
}

func (r *userRepo) Updateuser(user *model.User) error {
	ctx := context.Background()
	_, err := cmd.DB.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
	return err
}

func (r *userRepo) Deleteuser(id int64) error {
	ctx := context.Background()
	_, err := cmd.DB.NewDelete().Model(&model.User{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *userRepo) GetAllusers() ([]model.User, error) {
	var users []model.User
	ctx := context.Background()
	err := cmd.DB.NewSelect().Model(&users).Order("id ASC").Scan(ctx)
	return users, err
}

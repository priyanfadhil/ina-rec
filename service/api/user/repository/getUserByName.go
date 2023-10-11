package repository

import (
	"context"
	"fmt"

	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

func (r *UserRepository) GetOneUserByName(ctx context.Context, db *gorm.DB, name string) (user modelUser.User, err error) {
	res := db.Where("name = ?", name).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return user, err
}

package repository

import (
	"context"
	"fmt"

	userModel "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

func (r *UserRepository) DeleteUser(ctx context.Context, db *gorm.DB, id int) error {
	var (
		result userModel.User
	)
	query := fmt.Sprintf(`DELETE from public.users where id = %d`, id)

	q := db.Raw(query).Find(&result)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

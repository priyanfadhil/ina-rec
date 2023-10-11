package repository

import (
	"context"
	"fmt"

	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/gorm"
)

func (r *UserRepository) GetUserById(ctx context.Context, db *gorm.DB, id int) (modelUser.GetUserResponse, error) {
	var (
		result modelUser.GetUserResponse
	)
	query := fmt.Sprintf(`SELECT * FROM public.users where id = %d`, id)
	q := db.Raw(query).Find(&result)
	if q.Error != nil {
		return modelUser.GetUserResponse{}, q.Error
	}
	if q.RowsAffected == 0 {
		return modelUser.GetUserResponse{}, fmt.Errorf("data keranjang kosong")
	}

	return result, nil
}

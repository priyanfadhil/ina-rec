package repository

import (
	"context"
	"fmt"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
	"gorm.io/gorm"
)

func (r *TaskRepository) GetTaskById(ctx context.Context, db *gorm.DB, id int) (taskModel.TaskResponse, error) {
	var (
		result taskModel.TaskResponse
	)
	query := fmt.Sprintf(`SELECT * FROM public.tasks where id = %d`, id)
	q := db.Raw(query).Find(&result)
	if q.Error != nil {
		return taskModel.TaskResponse{}, q.Error
	}
	if q.RowsAffected == 0 {
		return taskModel.TaskResponse{}, fmt.Errorf("data keranjang kosong")
	}

	return result, nil
}

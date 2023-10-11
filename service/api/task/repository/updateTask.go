package repository

import (
	"context"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
	"gorm.io/gorm"
)

func (r *TaskRepository) UpdateTask(ctx context.Context, db *gorm.DB, insert *taskModel.CreateTaskRequest) error {
	err := db.WithContext(ctx).Where("id = ?", insert.ID).Updates(insert).Error

	if err != nil {
		return err
	}

	return nil
}

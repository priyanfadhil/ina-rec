package repository

import (
	"context"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
	"gorm.io/gorm"
)

func (r *TaskRepository) CreateTask(ctx context.Context, db *gorm.DB, insert *taskModel.ModelTask) error {
	err := db.
		Create(&insert).Error

	if err != nil {
		return err
	}

	return nil
}

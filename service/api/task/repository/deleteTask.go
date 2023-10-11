package repository

import (
	"context"
	"fmt"

	taskModel "github.com/priyanfadhil/ina-rec/service/model/task"
	"gorm.io/gorm"
)

func (r *TaskRepository) DeleteTask(ctx context.Context, db *gorm.DB, id int) error {
	var (
		result taskModel.ModelTask
	)
	query := fmt.Sprintf(`DELETE from public.tasks where id = %d`, id)

	q := db.Raw(query).Find(&result)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

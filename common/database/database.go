package database

import (
	"fmt"
	"time"

	modelTask "github.com/priyanfadhil/ina-rec/service/model/task"
	modelUser "github.com/priyanfadhil/ina-rec/service/model/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager interface {
	GetMaster() *gorm.DB
	StartTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) *gorm.DB
	RollbackTransaction(tx *gorm.DB) *gorm.DB

	Initialize(dsn string, maxIdleConns int, maxOpenConns int) error

	Migrate() error // Add this function for migration
}

func NewDatabaseManager() DatabaseManager {
	return &databaseManager{}
}

type databaseManager struct {
	Master *gorm.DB
}

func (dbManager *databaseManager) Initialize(dsn string, maxIdleConns int, maxOpenConns int) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	dbManager.Master = db

	// Call the Migrate function here to automatically migrate tables
	if err := dbManager.Migrate(); err != nil {
		return err
	}

	return nil
}

func (dbManager *databaseManager) GetMaster() *gorm.DB {
	if dbManager.Master == nil {
		return nil
	}
	return dbManager.Master
}

func (dbManager *databaseManager) StartTransaction() *gorm.DB {
	return dbManager.Master.Begin()
}

func (dbManager *databaseManager) CommitTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (dbManager *databaseManager) RollbackTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (dbManager *databaseManager) Migrate() error {
	// Call separate migration functions for each table
	if err := dbManager.MigrateUsersTable(); err != nil {
		return err
	}
	if err := dbManager.MigrateTasksTable(); err != nil {
		return err
	}
	return nil
}

func (dbManager *databaseManager) MigrateUsersTable() error {
	// Define your user model here using Gorm tags
	// Example: type User struct { ... }

	// Migrate the 'users' table
	if err := dbManager.Master.AutoMigrate(&modelUser.User{}); err != nil {
		return err
	}

	return nil
}

func (dbManager *databaseManager) MigrateTasksTable() error {
	// Define your task model here using Gorm tags
	// Example: type Task struct { ... }

	// Migrate the 'tasks' table
	if err := dbManager.Master.AutoMigrate(&modelTask.ModelTask{}); err != nil {
		return err
	}

	return nil
}

func PostgresURI(dbUserName, dbPassword, dbAddress, dbName string) string {
	return fmt.Sprintf(`postgres://%s:%s@%s/%s?sslmode=disable`,
		dbUserName, dbPassword, dbAddress, dbName)
}

package database

import (
	"fmt"
	"religion/config"
	"religion/internal/domain/entities"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(&entities.User{}, &entities.Post{}, &entities.Comment{}, &entities.Like{}, &entities.Follower{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

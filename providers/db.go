package providers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabase interface {
}

type Database struct {
	Connection *gorm.DB
	Config     IConfig
}

// NewDatabase creates a new Database instance with a connection to the
// database using the provided config. It runs migrations for the User,
// OTP and Profile models.
func NewDatabase(config IConfig) IDatabase {
	dsn := config.Get("dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting db")
	}

	// db.AutoMigrate(&models.User{})


	return &Database{
		Connection: db,
		Config:     config,
	}
}

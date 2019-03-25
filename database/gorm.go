package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetGormConn function
func GetGormConn(host, user, dbName, password string, port int) (*gorm.DB, error) {
	return gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password,
	))
}

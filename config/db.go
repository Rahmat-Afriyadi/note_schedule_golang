package config

import (
	"fmt"
	"note_scheduler/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// postgresql
// func SetupDatabaseConnection() *gorm.DB {

// 	dbHost := "localhost"
// 	dbPort := "3306"
// 	dbName := "note_scheduler_golang"
// 	dbUser := "root"

// 	fmt.Println(dbHost)

// 	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=Local", dbUser, dbHost, dbPort, dbName)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to create a connection to database")
// 	}
// 	entities := []interface{}{
// 		&entity.Activity{},
// 		&entity.Schedule{},

// 	}
// 	errMigrate := db.AutoMigrate(entities...)
// 	if errMigrate != nil {
// 		panic("Failed to auto migrate")
// 	}
// 	return db
// }

// mysql
func SetupDatabaseConnection() *gorm.DB {

	dbHost := "localhost"
	dbPort := "3306"
	dbName := "note_scheduler_golang"
	dbUser := "root"

	fmt.Println(dbHost)

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&loc=Local", dbUser, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	entities := []interface{}{
		&entity.Activity{},
		&entity.Schedule{},

	}
	errMigrate := db.AutoMigrate(entities...)
	if errMigrate != nil {
		panic("Failed to auto migrate")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("there is error db close")
	}
	dbSQL.Close()
}

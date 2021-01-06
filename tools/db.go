package tools

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// init ...
func init() {

	once.Do(func() {
		_, err := connect()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("[INFO][Database connection successfully!]")

	})

}

// GetConnection ...
func GetConnection() *gorm.DB {
	return db
}

// connect -> returns a new instance of gorm DB or an error if the connection fails
func connect() (*gorm.DB, error) {

	var err error
	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Silent,
			},
		),
	}

	driver := os.Getenv("DB_DRIVER")
	env := os.Getenv("APP_ENVIRONMENT")

	if len(driver) == 0 {
		driver = "sqlite3"
	}

	if strings.ToLower(env) != "production" {

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color

			},
		)

		gormConfig.Logger = newLogger

	}

	switch strings.ToUpper(driver) {
	case "PGSQL":

		db, err = gorm.Open(postgres.Open(getConnectionString(driver)), gormConfig)
		if err != nil {
			return db, err
		}

	case "MYSQL":

		db, err = gorm.Open(mysql.Open(getConnectionString(driver)), gormConfig)
		if err != nil {
			return db, err
		}

	default:

		fmt.Println("[Warning][Service running with SQLLITE3 Database]")

		// github.com/mattn/go-sqlite3
		db, err = gorm.Open(sqlite.Open(getConnectionString(driver)), gormConfig)
		if err != nil {
			return db, err
		}

	}

	return db, nil
}

func getConnectionString(driver string) string {

	var stringconnection string

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	switch strings.ToUpper(driver) {
	case "PGSQL":

		stringconnection = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbName,
			dbPassword,
		)

	case "MYSQL":

		stringconnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
		)

	default:

		// Default driver is sqlite3
		stringconnection = os.Getenv("DB_SQLITE_PATH")

	}

	return stringconnection

}

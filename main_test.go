package nullable_test

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v\n", err)
		os.Exit(1)
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	dialect := os.Getenv("GORM_DIALECT")
	switch dialect {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=UTC"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "mysql-legacy":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:3307)/gorm?charset=utf8&parseTime=True&loc=UTC"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "mariadb":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:3308)/gorm?charset=utf8&parseTime=True&loc=UTC"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "postgres", "postgres_simple":
		log.Printf("testing %v...", dialect)
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
		}
		if dialect == "postgres" {
			db, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
		} else {
			db, err = gorm.Open(postgres.New(postgres.Config{DSN: dbDSN, PreferSimpleProtocol: true}), &gorm.Config{})
		}
	////// Uncomment this to test Microsoft SQL Server
	// case "mssql":
	// 	log.Println("testing mssql...")
	// 	if dbDSN == "" {
	// 		dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	// 	}
	// 	db, err = gorm.Open(sqlserver.Open(dbDSN), &gorm.Config{})
	default:
		log.Println("testing sqlite3...")
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func SupportedDriver(dialectors ...string) bool {
	for _, dialect := range dialectors {
		if DB.Dialector.Name() == dialect {
			return true
		}
	}
	return false
}

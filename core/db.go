package core

import (
	"github.com/sandisuryadi36/sansan-dashboard/libs"
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBMain    *gorm.DB
	DBMainSQL *sql.DB
)

func StartDBConnection() {
	log.Printf("Starting Db Connections...")

	InitDBMain()

}

func InitDBMain() {
	log.Printf("Main Db - Connecting")
	var err error
	DBMain, err = gorm.Open(postgres.Open(libs.GetEnv("DB_DSN", "")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connect to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL, err = DBMain.DB()
	if err != nil {
		log.Fatalf("Error cannot initiate connection to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL.SetMaxIdleConns(0)
	DBMainSQL.SetMaxOpenConns(0)

	err = DBMainSQL.Ping()
	if err != nil {
		log.Fatalf("Cannot ping DB main: %v", err)
		os.Exit(1)
		return
	}

	log.Printf("Main Db - Connected")
}

func CloseDBMain() {
	log.Print("Closing DB Main Connection ... ")
	if err := DBMainSQL.Close(); err != nil {
		log.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	log.Println("Closing DB Main Success")
}

func MigrateDB() error {
	InitDBMain()
	defer CloseDBMain()

	migrator := DBMain.Migrator()
	if migrator.HasTable(
		// List table ORM from proto gorm
		// &authv1.AuthORM{},
		nil,
	) {
		log.Println("Table already exists, no migration needed")
		return nil
	}

	log.Println("Migration process begin...")
	if err := DBMain.AutoMigrate(
		// List table from proto gorm
		// &authv1.AuthORM{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
		os.Exit(1)
	}

	log.Println("Migration process finished...")

	return nil
}

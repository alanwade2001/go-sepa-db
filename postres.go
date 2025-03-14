package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	utils "github.com/alanwade2001/go-sepa-utils"
)

type Persist struct {
	DB         *gorm.DB
	schemaName string
}

func NewPersist(schemaName string) *Persist {

	service := &Persist{
		schemaName: schemaName,
	}

	service.Connect()

	return service
}

func (p *Persist) Connect() error {

	host := utils.Getenv("DB_HOST", "0.0.0.0")
	port := utils.Getenv("DB_PORT", "5432")
	user := utils.Getenv("DB_USER", "postgres")
	password := utils.Getenv("DB_PASSWORD", "postgres")
	name := utils.Getenv("DB_NAME", "postgres")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)

	log.Println("DSN:" + dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   p.schemaName + ".", // schema name
			SingularTable: false,
		},
	})

	// Get a database handle.
	if err != nil {
		return err
	}

	p.DB = db

	d, _ := db.DB()
	pingErr := d.Ping()
	if pingErr != nil {
		return err
	}
	fmt.Println("Connected!")

	return nil
}

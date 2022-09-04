package handlers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_Handler_New_Returns_Successfully(t *testing.T) {

	db, _, _ := sqlmock.New()
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gromDb, _ := gorm.Open(dialector, &gorm.Config{})

	expectedHandler := handler{DB: gromDb}

	returnedHandler := New(gromDb)

	assert.Equal(t, expectedHandler, returnedHandler)
}

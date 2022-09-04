package handlers

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_AddBook_Returns_Successfully(t *testing.T) {

	db, mock, _ := sqlmock.New()
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gromDb, _ := gorm.Open(dialector, &gorm.Config{})

	h := handler{DB: gromDb}

	sqlmock.
		NewRows([]string{"id", "title", "author", "description"})

	// expect transaction begin
	mock.ExpectBegin()
	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "books" ("title","author","description") VALUES ($1,$2,$3) RETURNING "id"`)).
		WithArgs("Percy Jackson", "Rick Riordan", "The novels...")

	body := []byte(`{"title": "Percy Jackson", "author": "Rick Riordan", "description": "The novels..." }`)

	req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(body))
	respRec := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(h.AddBook)
	httpHandler.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusCreated, respRec.Code)

	respBody, _ := ioutil.ReadAll(respRec.Body)
	expectedBody := `"Created"`
	assert.Equal(t, expectedBody, string(respBody[:len(respBody)-1]))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fail()
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func Test_AddBook_Returns_Fail_to_Ready_Body(t *testing.T) {

	db, _, _ := sqlmock.New()
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gromDb, _ := gorm.Open(dialector, &gorm.Config{})

	h := handler{DB: gromDb}

	req := httptest.NewRequest("POST", "/book", errReader(0))
	respRec := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(h.AddBook)

	assert.PanicsWithValue(t, "test error\n", func() { httpHandler.ServeHTTP(respRec, req) })
}

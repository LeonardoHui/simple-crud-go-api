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
	"github.com/gorilla/mux"
	"github.com/simple-crud-go-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_UpdateBook_Returns_Successfully(t *testing.T) {

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

	book := models.Book{ID: 256, Title: "Game of Phones", Author: "Unknown", Description: "Unknown"}
	row := sqlmock.
		NewRows([]string{"id", "title", "author", "description"}).
		AddRow(book.ID, book.Title, book.Author, book.Description)

	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(256).
		WillReturnRows(row)

	mock.ExpectBegin()
	mock.ExpectExec(
		regexp.QuoteMeta(`UPDATE "books" SET "title"=$1,"author"=$2,"description"=$3 WHERE "id" = $4`)).
		WithArgs("Game of Phones", "G.G.R.MARTINS", "Fiction", 256)

	body := []byte(`{"title": "Game of Phones", "author": "G.G.R.MARTINS", "description": "Fiction" }`)

	req := httptest.NewRequest("PUT", "/books/256", bytes.NewBuffer(body))
	respRec := httptest.NewRecorder()

	//This test we must use the gorilla-mux due the variable in our route
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", h.UpdateBook)
	router.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusOK, respRec.Code)

	respBody, _ := ioutil.ReadAll(respRec.Body)
	expectedBody := `"Updated"`
	assert.Equal(t, expectedBody, string(respBody[:len(respBody)-1]))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fail()
	}
}

func Test_UpdateBook_Returns_Fail_To_FindBook(t *testing.T) {

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

	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(256).
		WillReturnError(errors.New("Test Error"))

	mock.ExpectBegin()
	mock.ExpectRollback()

	body := []byte(`{"title": "Game of Phones", "author": "G.G.R.MARTINS", "description": "Fiction" }`)

	req := httptest.NewRequest("PUT", "/books/256", bytes.NewBuffer(body))
	respRec := httptest.NewRecorder()

	//This test we must use the gorilla-mux due the variable in our route
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", h.UpdateBook)
	router.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusOK, respRec.Code)

	respBody, _ := ioutil.ReadAll(respRec.Body)
	expectedBody := `"Updated"`
	assert.Equal(t, expectedBody, string(respBody[:len(respBody)-1]))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fail()
	}
}

func Test_UpdateBook_Returns_Fail_to_Ready_Body(t *testing.T) {

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

	req := httptest.NewRequest("PUT", "/books/256", errReader(0))
	respRec := httptest.NewRecorder()

	//This test we must use the gorilla-mux due the variable in our route
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", h.UpdateBook)

	assert.PanicsWithValue(t, "test error\n", func() { router.ServeHTTP(respRec, req) })
}

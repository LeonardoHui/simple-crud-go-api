package handlers

import (
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

func Test_GetBook_Returns_Successfully(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening sqlmock: %s", err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening gorm: %s", err)
	}

	h := handler{DB: gormDb}

	book := models.Book{ID: 256, Title: "Game of Phones", Author: "Unknown", Description: "Unknown"}
	row := sqlmock.
		NewRows([]string{"id", "title", "author", "description"}).
		AddRow(book.ID, book.Title, book.Author, book.Description)

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(book.ID).
		WillReturnRows(row)

	req := httptest.NewRequest("GET", "/books/256", nil)
	rr := httptest.NewRecorder()

	//This test we must use the gorilla-mux due the variable in our route
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", h.GetBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := ioutil.ReadAll(rr.Body)
	expectedBody := `{"id":256,"title":"Game of Phones","author":"Unknown","description":"Unknown"}`
	assert.Equal(t, expectedBody, string(body[:len(body)-1])) //Removing new line

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fail()
	}
}

func Test_GetBook_Returns_Fail_To_FindBook(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening sqlmock: %s", err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening gorm: %s", err)
	}

	h := handler{DB: gormDb}

	book := models.Book{ID: 256, Title: "Game of Phones", Author: "Unknown", Description: "Unknown"}

	mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(book.ID).
		WillReturnError(errors.New("Test Error"))

	req := httptest.NewRequest("GET", "/books/256", nil)
	rr := httptest.NewRecorder()

	//This test we must use the gorilla-mux due the variable in our route
	router := mux.NewRouter()
	router.HandleFunc("/books/{id}", h.GetBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := ioutil.ReadAll(rr.Body)
	expectedBody := `{"id":0,"title":"","author":"","description":""}`
	assert.Equal(t, expectedBody, string(body[:len(body)-1])) //Removing new line

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fail()
	}
}

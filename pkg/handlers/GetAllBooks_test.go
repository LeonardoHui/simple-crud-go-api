package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/simple-crud-go-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_GetAllBooks_returnsSuccessfuly(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gromDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("[gorm open] %s", err)
	}

	h := handler{DB: gromDb}

	book := models.Book{ID: 1, Title: "Happy Pot", Author: "Myself", Description: "Parody"}
	rows := sqlmock.
		NewRows([]string{"id", "title", "author", "description"}).
		AddRow(book.ID, book.Title, book.Author, book.Description)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books"`)).WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(h.GetAllBooks)
	httpHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := ioutil.ReadAll(rr.Body)
	expectedBody := `[{"id":1,"title":"Happy Pot","author":"Myself","description":"Parody"}]`
	assert.Equal(t, expectedBody, string(body[:len(body)-1])) //Removing new line

}

func Test_GetAllBooks_returnsEmptyBook(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gromDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("[gorm open] %s", err)
	}

	h := handler{DB: gromDb}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "books"`)).WillReturnError(errors.New("Error expected"))

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	httpHandler := http.HandlerFunc(h.GetAllBooks)
	httpHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := ioutil.ReadAll(rr.Body)
	assert.Equal(t, "null\n", string(body)) //Removing new line

}

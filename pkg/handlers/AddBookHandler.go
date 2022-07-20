package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/simple-crud-go-api/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// book.ID = rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}

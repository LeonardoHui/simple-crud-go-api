package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/simple-crud-go-api/pkg/models"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var UpdateBook models.Book
	json.Unmarshal(body, &UpdateBook)

	// for index, book := range mocks.Books {
	// 	if book.ID == id {
	// 		book.Title = UpdateBook.Title
	// 		book.Author = UpdateBook.Author
	// 		book.Description = UpdateBook.Description

	// 		mocks.Books[index] = book

	// 		w.WriteHeader(http.StatusOK)
	// 		w.Header().Add("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode("Updated")
	// 	}
	// }

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = UpdateBook.Title
	book.Author = UpdateBook.Author
	book.Description = UpdateBook.Description

	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}

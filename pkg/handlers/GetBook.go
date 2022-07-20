package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/simple-crud-go-api/pkg/models"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// for _, book := range mocks.Books {
	// 	if book.ID == id {
	// 		w.WriteHeader(http.StatusOK)
	// 		w.Header().Add("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(book)
	// 		break
	// 	}
	// }

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

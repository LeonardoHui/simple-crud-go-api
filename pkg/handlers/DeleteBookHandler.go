package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/simple-crud-go-api/pkg/models"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// for index, book := range mocks.Books {
	// 	if book.ID == id {
	// 		mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode("Deleted")
	// 		break
	// 	}
	// }

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}

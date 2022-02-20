package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ispiroglu/go-bookstore/pkg/models"
	"github.com/ispiroglu/go-bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	resp, _ := json.Marshal(newBooks)
	utils.SendResponse(w, resp)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookid"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		log.Fatalln(err)
		return
	}
	bookDetails, _ := models.GetBookByID(ID)
	resp, _ := json.Marshal(bookDetails)
	utils.SendResponse(w, resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createdBook := &models.Book{}
	utils.ParseBody(r, createdBook)
	b := createdBook.CreateBook()
	resp, _ := json.Marshal(b)
	utils.SendResponse(w, resp)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	bookID := vars["bookid"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Could not strcnonv")
	}
	b := models.DeleteBook(ID)
	resp, _ := json.Marshal(b)
	utils.SendResponse(w, resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookID := vars["bookid"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Could not strcnonv")
	}
	bookDetails, db := models.GetBookByID(ID)
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(&bookDetails)
	resp, _ := json.Marshal(bookDetails)
	utils.SendResponse(w, resp)
}

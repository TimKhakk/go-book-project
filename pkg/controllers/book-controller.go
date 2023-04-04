package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/timkhakk/go-book-project/pkg/models"
	"github.com/timkhakk/go-book-project/pkg/utils"

	"github.com/gorilla/mux"
)

var NewBooks models.Book

const BOOK_ID_FIELD_NAME = "bookId"

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, _ := json.Marshal(books)

	utils.WriteHeaders(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars[BOOK_ID_FIELD_NAME]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	utils.WriteHeaders(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	book := CreateBook.CreateBook()

	res, _ := json.Marshal(book)

	utils.WriteHeaders(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars[BOOK_ID_FIELD_NAME]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID)

	res, _ := json.Marshal(book)

	utils.WriteHeaders(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBookData := &models.Book{}

	utils.ParseBody(r, updatedBookData)
	vars := mux.Vars(r)

	bookId := vars[BOOK_ID_FIELD_NAME]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	booksDetails, db := models.GetBookById(ID)

	if updatedBookData.Name != "" {
		booksDetails.Name = updatedBookData.Name
	}
	if updatedBookData.Author != "" {
		booksDetails.Author = updatedBookData.Author
	}
	if updatedBookData.Publication != "" {
		booksDetails.Publication = updatedBookData.Publication
	}
	db.Save(&booksDetails)

	res, _ := json.Marshal(booksDetails)

	utils.WriteHeaders(w, res)
}

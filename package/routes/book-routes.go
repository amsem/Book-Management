package routes

import (
	"github.com/gorilla/mux"
	"github.com/amsem/Book-Managment/package/controlers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controlers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controlers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controlers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controlers.DeleteBook).Methods("DELETE")
}
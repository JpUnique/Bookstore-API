package routes

import (
	"github.com/JpUnique/Verbatim/pkg/controllers"
	"github.com/gorilla/mux"
)

const id = "/book/{bookId}"

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc(id, controllers.GetBookByID).Methods("GET")
	router.HandleFunc(id, controllers.UpdateBook).Methods("PUT")
	router.HandleFunc(id, controllers.DeleteBook).Methods("DELETE")
}

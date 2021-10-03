package main

import (
	"github.com/gorilla/mux"
	"waleedElgazar.com/notes-with-golang-mysql/controller"
	"net/http"
)


func main() {
	router:=mux.NewRouter()
	router.HandleFunc("/",controller.Index)
	router.HandleFunc("/add",controller.Add)
	router.HandleFunc("/search/",controller.SearchNotes)
	router.HandleFunc("/delete/",controller.DeleteNotes)
	router.HandleFunc("/getnotes/",controller.Getnotes)
	http.ListenAndServe(":6969",router)

}


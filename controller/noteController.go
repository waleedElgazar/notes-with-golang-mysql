package controller

import (
	"fmt"
	"html/template"
	"waleedElgazar.com/notes-with-golang-mysql/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r*http.Request){
	http.ServeFile(w,r,"view/index.html")
}

func Add(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}
	noteType:=r.FormValue("type")
	noteBody:=r.FormValue("note")
	currentTime := time.Now()
	var date string
	date=currentTime.Format("2006.01.02 15:04:05")
	var notes db.Note
	notes.Type=noteType
	notes.Body=noteBody
	notes.Date=date
	db.AddNote(notes)
	var Notes []db.Note
	Notes=db.FetchNotes()
	t, err := template.ParseFiles("view/get.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, map[string]interface{} {
		"Notes": Notes,
	}) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func SearchNotes(w http.ResponseWriter,r *http.Request){
	var Notes []db.Note
	noteType:=r.FormValue("search")
	Notes=db.SearchNote(noteType)
	t, err := template.ParseFiles("view/get.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, map[string]interface{} {
		"Notes": Notes,
	}) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func DeleteNotes(w http.ResponseWriter,r *http.Request){
	var Notes []db.Note
	id:=r.FormValue("delete")
	fmt.Println(id)
	delId,err:=strconv.Atoi(id)
	if err != nil {
		fmt.Println("can't cast the id",err)
	}
	db.DeleteNote(delId)
	Notes=db.FetchNotes()
	t, err := template.ParseFiles("view/get.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, map[string]interface{} {
		"Notes": Notes,
	}) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func Getnotes(w http.ResponseWriter,r *http.Request){

	var Notes []db.Note

	Notes=db.FetchNotes()
	t, err := template.ParseFiles("view/get.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, map[string]interface{} {
		"Notes": Notes,
	}) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

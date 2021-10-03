package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Config()*sql.DB{
	driver:="mysql"
	cred:="root:00@/notes"
	db,err:=sql.Open(driver,cred)
	if err != nil {
		fmt.Println("error while establishing connection",err)
		return nil
	}
	return db
}

type Note struct {
	IdNote		int			`json:"id_note"`
	Type 		string		`json:"type"`
	Body		string		`json:"body"`
	Date		string		`json:"date"`
}



func AddNote(note  Note){
	db:= Config()
	defer db.Close()
	query:="INSERT INTO note set noteType=?, noteBody=?, noteDate=?"
	insert,err:=db.Prepare(query)
	if err !=nil{
		fmt.Println("error inserting note",err)
	}
	_,err=insert.Exec(note.Type,note.Body,note.Date)
	if err !=nil{
		fmt.Println("error inserting note",err)
	}
}

func DeleteNote(id int){
	db:= Config()
	defer db.Close()
	query:="DELETE  FROM note WHERE idNote=?"
	_,err:=db.Exec(query,id)
	if err != nil {
		fmt.Println("error while deleting",err)
	}
}

func FetchNotes()[] Note{
	db:= Config()
	defer db.Close()
	var notes [] Note
	query:="SELECT idNote,noteType, noteBody, noteDate FROM note"
	result,err:=db.Query(query)
	if err != nil {
		fmt.Println("error while getting notes",err)
	}
	var notetype,body,date string
	var idNote int
	for result.Next() {
		err:=result.Scan(&idNote,&notetype,&body,&date)
		if err != nil {
			fmt.Println("error while scanning returned notes",err)
		}

		notes=append(notes, Note{
			IdNote: idNote,
			Type: notetype,
			Date: date,
			Body: body,
		})
	}
	return notes
}


func SearchNote(text string)[] Note{
	db:= Config()
	defer db.Close()
	var notes [] Note
	query:="SELECT idNote,noteType, noteBody, noteDate FROM note Where noteType=?"
	result,err:=db.Query(query,text)
	if err != nil {
		fmt.Println("error while getting notes",err)
	}
	var notetype,body,date string
	var idNote int
	for result.Next() {
		err:=result.Scan(&idNote,&notetype,&body,&date)
		if err != nil {
			fmt.Println("error while scanning returned notes",err)
		}

		notes=append(notes, Note{
			IdNote: idNote,
			Type: notetype,
			Date: date,
			Body: body,
		})
	}
	return notes
}


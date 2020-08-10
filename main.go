package main

import (
//    "database/sql"
//"io"
//"net/http"
//"net/url"
//"os"

"log"
"os"
	"net/http"
  //  "text/template"
"controller/crudcontroller"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./src"))
	mux.Handle("/src/", http.StripPrefix("/src/", fs))
    log.Println("Server started on: http://localhost:8080")
    mux.HandleFunc("/", crudcontroller.Index)
    mux.HandleFunc("/show",crudcontroller.Show )
    mux.HandleFunc("/new", crudcontroller.New)
    mux.HandleFunc("/edit",crudcontroller.Edit)
    mux.HandleFunc("/insert",crudcontroller.Insert)
    mux.HandleFunc("/update",crudcontroller.Update)
    mux.HandleFunc("/delete",crudcontroller.Delete)
    http.ListenAndServe(":"+port, mux)
}
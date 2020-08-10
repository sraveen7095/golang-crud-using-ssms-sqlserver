package crudcontroller

import (
	"config"
//    "database/sql"
    "log"
    "net/http"

"html/template"
    _ "github.com/go-sql-driver/mysql"
)

type Student struct {
    Id    int
    Name  string
    City string
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("src/views/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
        db := config.Connstr()

    rows, err := db.Query("SELECT * FROM Student ORDER BY id ")
    if err != nil {
        panic(err.Error())
    }
    std := Student{}
    res := []Student{}
    for rows.Next() {
        var id int
        var name, city string
        err = rows.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        std.Id = id
        std.Name = name
        std.City = city
        res = append(res, std)
    }
    tpl.ExecuteTemplate(w, "index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
        db := config.Connstr()

    nId := r.URL.Query().Get("id")
    rows, err := db.Query("SELECT * FROM Student WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    std := Student{}
    for rows.Next() {
        var id int
        var name, city string
        err = rows.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        std.Id = id
        std.Name = name
        std.City = city
    }
    tpl.ExecuteTemplate(w, "show", std)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "new", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
        db := config.Connstr()

    nId := r.URL.Query().Get("id")
    rows, err := db.Query("SELECT * FROM Student WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    std := Student{}
    for rows.Next() {
        var id int
        var name, city string
        err = rows.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        std.Id = id
        std.Name = name
        std.City = city
    }
    tpl.ExecuteTemplate(w, "edit", std)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := config.Connstr()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        insForm, err := db.Prepare("INSERT INTO Student(name, city) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city)
        log.Println("INSERT: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
        db := config.Connstr()

    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Student SET name=?, city=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city, id)
        log.Println("UPDATE: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
        db := config.Connstr()

    std := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Student WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(std)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

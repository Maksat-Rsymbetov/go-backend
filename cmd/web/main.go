package main

import (
	"database/sql"
	"flag"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"os"
	"website.maksat.com/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	products      *models.ProductModel
	users         *models.UserModel
	templateCache map[string]*template.Template
}

func main() {
	// Flag needed for choosing a server port
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Error loggers (they show errors on the terminal)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// The actual database connection pool
	// dataSource := flag.String("dsn", "web:pass@/ecommerce?parseTime=true", "MySQL data source name")
	// db, err := openDB(*dataSource)
	db, err := openDB("./internal/database/db.sqlite.db")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//Template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	//Application struct instance needed to unify error handling
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		products:      &models.ProductModel{DB: db},
		users:         &models.UserModel{DB: db},
		templateCache: templateCache,
	}

	//Server which uses our logger function and contains the server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	// Show Error messages
	infoLog.Printf("Startnig server on port %s", *addr)
	errorLog.Fatal(srv.ListenAndServe())
	// errorLog.Fatal(http.ListenAndServe(*addr, mux))
}

// Return an sql.DB connection pool
func openDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

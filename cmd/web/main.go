// npx nodemon --exec go run ./cmd/web --signal SIGTERM -e go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"snippetbox.khalidjameer.com/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependencies for the // web application.
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

type fileWriter struct{}

func (f *fileWriter) Write(p []byte) (n int, err error) {
	logFile := "log.txt"
	os.Create("hi")
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failure to open %s", logFile)
		return 0, err
	}
	defer file.Close()

	bytesWritten, writeErr := file.Write(p)
	if bytesWritten < len(p) {
		return 0, fmt.Errorf("failure to write all data to %s", logFile)
	}

	if writeErr != nil {
		log.Printf("Error writing to file %s: %v\n", logFile, writeErr)
		return 0, writeErr
	}

	return bytesWritten, nil
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@tcp(localhost)/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	//mux is the part of the app that guides requests
	//to the url that matches their path
	writer := &fileWriter{}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, writer), nil))

	slog.SetDefault(logger)

	logger.Info("starting server", slog.String("addr", *addr))

	// To keep the main() function tidy I've put the code for creating a connection // pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
	}

	err = http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

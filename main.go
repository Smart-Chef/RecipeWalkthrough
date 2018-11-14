package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gchaincl/dotsql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	database *sql.DB
	queries  *dotsql.DotSql
)

func init() {
	// Setup logger
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err = sql.Open("sqlite3", "./smart-chef.db")
	dot, err := dotsql.LoadFromFile("setup.sql")

	dot.Exec(database, "create-smart-chef")
	queries, err = dotsql.LoadFromFile("queries.sql")
}

func main() {
	// Setup Mux
	r := mux.NewRouter()
	var wait time.Duration

	// Log endpoints Method + URI
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// https://golang.org/pkg/net/http/#Request
			log.Info(r.Method + " " + r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	})

	// Route handles & endpoints
	CreateRoutes(r, AllRoutes[:], "/api")

	// Walk all the routes
	r.Walk(RouteWalker)

	// Start server
	srv := &http.Server{
		Handler: r,
		Addr:    ":8001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run the state machine
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	//Info.Println("shutting down")
	os.Exit(0)
}

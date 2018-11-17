package main

import (
	"context"
	"database/sql"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/gchaincl/dotsql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	// Flags
	walkFlg    bool // Walk and print the routes - default: false
	verboseFlg bool // Verbose logging - default: false
	setupFlg   bool // Setup DB - default: false

	// Database
	database *sql.DB
	queries  *dotsql.DotSql

	// CurrentRecipe tracker
	CurrentRecipe *RecipeInfo
)

func init() {
	// Handle cli flags
	flag.BoolVar(&walkFlg, "walk", false, "Walk the routes")
	flag.BoolVar(&verboseFlg, "verbose", false, "Display verbose logs")
	flag.BoolVar(&setupFlg, "setup", false, "Run setup")

	flag.Parse()

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

	// Setup Database
	database, err = sql.Open("sqlite3", "./smart-chef.db")
	if err != nil {
		log.Error(err.Error())
		log.Fatal("Error opening smart-chef.db")
	}

	if setupFlg {
		dot, err := dotsql.LoadFromFile("setup.sql")
		if err != nil {
			log.Error(err.Error())
			log.Fatal("Error reading setup.sql file")
		}

		_, err = dot.Exec(database, "create-smart-chef")
		if err != nil {
			log.Error(err.Error())
			log.Fatal("Error running setup.sql file")
		}
	}

	queries, err = dotsql.LoadFromFile("queries.sql")
	if err != nil {
		log.Error(err.Error())
		log.Fatal("Error loading queries.sql")
	}

	// Create recipe object
	CurrentRecipe = &RecipeInfo{
		ID:         null.IntFrom(-1),
		TotalSteps: 0,
	}
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
	if walkFlg {
		r.Walk(RouteWalker)
	}

	// Start server
	srv := &http.Server{
		Handler: r,
		Addr:    ":8001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Starting recipe-walkthrough server at %s", srv.Addr)
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

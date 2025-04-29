package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = 1 

type config struct {
	port int 
	env string 
}

type application struct {
	config config 
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Env(dev|staging|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: mux, 
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s:", cfg.env, cfg.port)
	err := server.ListenAndServe()
	logger.Fatal(err)
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status available")
	fmt.Fprintf(w, "env: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
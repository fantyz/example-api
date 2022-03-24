package main

import (
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	API struct {
		Addr string `default:"0.0.0.0:8000"`
	}
	Log struct {
		Level string `default:"INFO"`
		JSON  bool   `default:"false"`
	}
}

func main() {
	// initial setup of logging
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// setup config
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal(errors.Wrap(err, "Unable to process config"))
	}

	// configure logging
	if c.Log.JSON {
		log.Formatter = &logrus.JSONFormatter{}
	}
	lvl, err := logrus.ParseLevel(c.Log.Level)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "Unable to parse log level (level=%s)", log.Level))
	}
	log.Infof("Setting log level (level=%s)", lvl)
	log.SetLevel(lvl)

	// router
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/heartbeat", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprintf(w, "1") }).Methods(http.MethodGet)
	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		log.Info("Hello, Log!")
		fmt.Fprintf(w, "Hello, World!")
	}).Methods(http.MethodGet)
	loggingRouter := handlers.LoggingHandler(os.Stdout, r)

	// start server
	s := http.Server{
		Addr:     c.API.Addr,
		ErrorLog: stdlog.New(log.WriterLevel(logrus.WarnLevel), "", 0),
		Handler:  loggingRouter,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(errors.Wrap(err, "ListenAndServe failed"))
	}
}

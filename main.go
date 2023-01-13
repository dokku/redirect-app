package main

import (
	"fmt"

	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

var (
	logger             = GetLogger()
	port               = GetPort()
	redirectURL        = GetRedirectURL()
	redirectStatusCode = GetRedirectStatusCode()
)

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf("%v %v %v %v %v %v %v %v", r.Header["User-Agent"], r.Method, r.Host, r.URL, r.Proto, r.Response, r.RemoteAddr, r.Header)
	l.handler.ServeHTTP(w, r)
}

func GetLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger
}

func GetPort() int {
	value := os.Getenv("PORT")
	if value == "" {
		value = "5000"
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Fatal().Err(err).Str("port", value).Msg("Unable to parse PORT environment variable")
	}

	return i
}

func GetRedirectStatusCode() int {
	value := os.Getenv("REDIRECT_STATUS_CODE")
	if value == "" {
		value = "302"
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Fatal().Err(err).Str("status_code", value).Msg("Unable to parse REDIRECT_STATUS_CODE environment variable")
	}

	return i
}

func GetRedirectURL() *url.URL {
	value := os.Getenv("REDIRECT_URL")
	if value == "" {
		logger.Fatal().Msg("Missing REDIRECT_URL environment variable")
	}

	parsed, err := url.ParseRequestURI(value)
	if err != nil {
		logger.Fatal().Err(err).Str("redirect_url", value).Msg("Unable to parse REDIRECT_URL")
	}

	return parsed
}

func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, redirectURL.String(), redirectStatusCode)
}

func main() {
	router := httprouter.New()
	router.GET("/", Handler)
	srv := &http.Server{
		Handler:      &Logger{router},
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	logger.Info().Int("port", port).Msg("Starting server")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal().Err(err).Send()
	}

	logger.Info().Msg("Server stopped")
}

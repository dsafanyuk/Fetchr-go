package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dsafanyuk/fetchr-go/config"
	"github.com/dsafanyuk/fetchr-go/pkg/database/psql"
	"github.com/dsafanyuk/fetchr-go/pkg/middleware"
	"github.com/dsafanyuk/fetchr-go/pkg/user"
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	config := config.GetConfig()

	var userRepo user.UserRepository
	pconn := postgresConnection(config)
	defer pconn.Close()
	userRepo = psql.NewPostgresUserRepository(pconn)

	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", userHandler.Get).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	router.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	router.HandleFunc("/users", userHandler.Create).Methods("POST")

	http.Handle("/", accessControl(middleware.Logging(router)))

	errs := make(chan error, 2)
	logrus.Info("Listening server mode on port :3000")
	errs <- http.ListenAndServe(":3000", nil)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
	logrus.Errorf("terminated %s", <-errs)

}

func postgresConnection(config *config.Config) *sqlx.DB {
	logrus.Info("Connecting to PostgreSQL DB")
	dbURI := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)
	db := sqlx.MustConnect(config.DB.Dialect, dbURI)

	return db
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

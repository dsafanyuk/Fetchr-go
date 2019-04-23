package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dsafanyuk/fetchr-go/app/handler"
	"github.com/dsafanyuk/fetchr-go/config"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)
	a.DB = sqlx.MustConnect(config.DB.Dialect, dbURI)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the users
	a.Get("/users", a.GetAllUsers)
	a.Post("/users", a.CreateUser)
	a.Get("/users/{userID}", a.GetUser)
	// a.Put("/users/{userID}", a.UpdateUser)
	a.Delete("/users/{userID}", a.DeleteUser)

	// a.Get("/orders", a.GetAllOrders)
	// a.Post("/orders", a.CreateOrder)
	// a.Get("/orders/{orderID}", a.GetOrder)
	// a.Put("/orders/{orderID}", a.UpdateOrder)
	// a.Delete("/order/{orderID}", a.DeleteOrder)

	// a.Get("/products", a.GetAllProducts)
	// a.Post("/products", a.CreateProduct)
	// a.Get("/products/{productID}", a.GetProduct)
	// a.Put("/products/{productID}", a.UpdateProduct)
	// a.Delete("/product/{productID}", a.DeleteProduct)

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
** Users Handlers
 */
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.DB, w, r)
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	handler.CreateUser(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.DB, w, r)
}

// func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateUser(a.DB, w, r)
// }

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	handler.DeleteUser(a.DB, w, r)
}

// /*
// ** Orders Handlers
//  */
// func (a *App) GetAllOrders(w http.ResponseWriter, r *http.Request) {
// 	handler.GetAllOrders(a.DB, w, r)
// }

// func (a *App) CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	handler.CreateOrder(a.DB, w, r)
// }

// func (a *App) GetOrder(w http.ResponseWriter, r *http.Request) {
// 	handler.GetOrder(a.DB, w, r)
// }

// func (a *App) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateOrder(a.DB, w, r)
// }

// func (a *App) DeleteOrder(w http.ResponseWriter, r *http.Request) {
// 	handler.DeleteOrder(a.DB, w, r)
// }

// /*
// ** Products Handlers
//  */
// func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
// 	handler.GetAllProducts(a.DB, w, r)
// }

// func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.CreateProduct(a.DB, w, r)
// }

// func (a *App) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.GetProduct(a.DB, w, r)
// }

// func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateProduct(a.DB, w, r)
// }

// func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	handler.DeleteProduct(a.DB, w, r)
// }

// Run the app on its router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

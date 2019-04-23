package handler

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/dsafanyuk/fetchr-go/app/model"
// 	"github.com/gorilla/mux"
// 	"github.com/jmoiron/sqlx"
// )

// func GetAllProducts(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	products := []model.Product{}
// 	db.Find(&products)
// 	respondJSON(w, http.StatusOK, products)
// }

// func CreateProduct(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	product := model.Product{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&product); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&product).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, product)
// }

// func GetProduct(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	productID, _ := strconv.ParseInt(vars["productID"], 0, 64)
// 	product := getProductOr404(db, productID, w, r)
// 	if product == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, product)
// }

// func UpdateProduct(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	productID, _ := strconv.ParseInt(vars["productID"], 0, 64)
// 	product := getProductOr404(db, productID, w, r)
// 	if product == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&product); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&product).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, product)
// }

// func DeleteProduct(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	productID, _ := strconv.ParseInt(vars["productID"], 0, 64)
// 	product := getProductOr404(db, productID, w, r)
// 	if product == nil {
// 		return
// 	}
// 	if err := db.Delete(&product).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// // getProductOr404 gets a product instance if exists, or respond the 404 error otherwise
// func getProductOr404(db *sqlx.DB, productID int64, w http.ResponseWriter, r *http.Request) *model.Product {
// 	product := model.Product{}
// 	if err := db.First(&product, productID).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &product
// }

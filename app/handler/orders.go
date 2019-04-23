package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/dsafanyuk/fetchr-go/app/model"
// 	"github.com/gorilla/mux"
// 	"github.com/jmoiron/sqlx"
// )

// func GetAllOrders(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	orders := []model.Order{}
// 	db.Find(&orders)
// 	respondJSON(w, http.StatusOK, orders)
// }

// // func CreateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// // 	order := model.Order{}

// // 	decoder := json.NewDecoder(r.Body)
// // 	if err := decoder.Decode(&order); err != nil {
// // 		respondError(w, http.StatusBadRequest, err.Error())
// // 		return
// // 	}
// // 	defer r.Body.Close()

// // 	if err := db.Create(&order).Error; err != nil {
// // 		respondError(w, http.StatusInternalServerError, err.Error())
// // 		return
// // 	}
// // 	respondJSON(w, http.StatusCreated, order)
// // }

// // func GetOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)

// // 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// // 	fmt.Println(orderID)
// // 	order := getOrderOr404(db, orderID, w, r)
// // 	if order == nil {
// // 		return
// // 	}
// // 	respondJSON(w, http.StatusOK, order)
// // }

// // func UpdateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)

// // 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// // 	order := getOrderOr404(db, orderID, w, r)
// // 	if order == nil {
// // 		return
// // 	}

// // 	decoder := json.NewDecoder(r.Body)
// // 	if err := decoder.Decode(&order); err != nil {
// // 		respondError(w, http.StatusBadRequest, err.Error())
// // 		return
// // 	}
// // 	defer r.Body.Close()

// // 	if err := db.Create(&order).Error; err != nil {
// // 		respondError(w, http.StatusInternalServerError, err.Error())
// // 		return
// // 	}
// // 	respondJSON(w, http.StatusOK, order)
// // }

// // func DeleteOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// // 	vars := mux.Vars(r)

// // 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// // 	order := getOrderOr404(db, orderID, w, r)
// // 	if order == nil {
// // 		return
// // 	}
// // 	if err := db.Delete(&order).Error; err != nil {
// // 		respondError(w, http.StatusInternalServerError, err.Error())
// // 		return
// // 	}
// // 	respondJSON(w, http.StatusNoContent, nil)
// // }

// // // getOrderOr404 gets a order instance if exists, or respond the 404 error otherwise
// // func getOrderOr404(db *sqlx.DB, orderID int64, w http.ResponseWriter, r *http.Request) *model.Order {
// // 	order := model.Order{}
// // 	if err := db.First(&order, orderID).Error; err != nil {
// // 		respondError(w, http.StatusNotFound, err.Error())
// // 		return nil
// // 	}
// // 	return &order
// // }

// func CreateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	order := model.Order{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&order); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, order)
// }

// func GetOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	fmt.Println(orderID)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, order)
// }

// func UpdateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&order); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, order)
// }

// func DeleteOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}
// 	if err := db.Delete(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// // getOrderOr404 gets a order instance if exists, or respond the 404 error otherwise
// func getOrderOr404(db *sqlx.DB, orderID int64, w http.ResponseWriter, r *http.Request) *model.Order {
// 	order := model.Order{}
// 	if err := db.First(&order, orderID).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &order
// }

// func CreateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	order := model.Order{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&order); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, order)
// }

// func GetOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	fmt.Println(orderID)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, order)
// }

// func UpdateOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&order); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Create(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, order)
// }

// func DeleteOrder(db *sqlx.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	orderID, _ := strconv.ParseInt(vars["orderID"], 0, 64)
// 	order := getOrderOr404(db, orderID, w, r)
// 	if order == nil {
// 		return
// 	}
// 	if err := db.Delete(&order).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// // getOrderOr404 gets a order instance if exists, or respond the 404 error otherwise
// func getOrderOr404(db *sqlx.DB, orderID int64, w http.ResponseWriter, r *http.Request) *model.Order {
// 	order := model.Order{}
// 	if err := db.First(&order, orderID).Error; err != nil {
// 		respondError(w, http.StatusNotFound, err.Error())
// 		return nil
// 	}
// 	return &order
// }

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getLatestOrders(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	input := vars["user_uuid"]

	var orders []Orders

	getLatestOrders := queries["getLatestOrders"]

	if err := dbConn.db.Raw(getLatestOrders, input).Scan(&orders); err != nil {
		log.Println(err)
	}

	payload := struct {
		LatestOrders []Orders
	}{
		LatestOrders: orders,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func updateOrders(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	method := vars["method"]

	var orders []Orders
	var emailOrder []EmailOrder

	err := json.NewDecoder(r.Body).Decode(&orders)
	if err != nil {
		log.Println("YOUR ERROR: ", err)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range orders {

		if err := dbConn.db.Where("amount = ? and user_uuid = ? and item_id = ? and store_id = ? and created_at::date = current_date",
			orders[i].Amount,
			orders[i].UserUUID,
			orders[i].ItemID,
			orders[i].StoreID).First(&orders[i]).Error; err != nil {
			if err := dbConn.db.Create(&orders[i]).Error; err != nil {
				log.Println(err)
				viewRender.Text(w, http.StatusBadRequest, "Error! Couldn't submit form.")
				return
			}
		} else {
			if err := dbConn.db.Model(&orders[i]).UpdateColumn("updated_at", time.Now()).Error; err != nil {
				log.Println(err)
				viewRender.Text(w, http.StatusBadRequest, "Error! Couldn't submit form.")
				return
			}
		}

	}

	if method == "send" {

		getOrderEmail := queries["getOrderEmail"]

		// now := time.Now()
		// then := now.AddDate(0, 0, -12)

		if err := dbConn.db.Raw(getOrderEmail, orders[0].UserUUID).Scan(&emailOrder); err != nil {
			log.Println(err)
		}

		var users []user

		if err := dbConn.db.Where("role = ?", "admin").Find(&users).Error; err != nil {
			log.Println(err)
			viewRender.Text(w, http.StatusBadRequest, "Error! Couldn't submit form.")
			return
		}

		payloadEmail := struct {
			EmailOrder  []EmailOrder
			LastUpdated string
		}{
			EmailOrder:  emailOrder,
			LastUpdated: time.Now().Format("Mon Jan _2 15:04:05 2006"),
		}

		for _, user := range users {
			sendOrdersEmail(user.UserEmail, payloadEmail)
		}

	}

	viewRender.Text(w, http.StatusCreated, "Success!")

}

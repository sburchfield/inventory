package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	var orders []Orders
	var emailOrder []EmailOrder

	err := json.NewDecoder(r.Body).Decode(&orders)
	if err != nil {
		log.Println("YOUR ERROR: ", err)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range orders {

		if err := dbConn.db.Create(&orders[i]).Error; err != nil {

			log.Println(err)
			viewRender.Text(w, http.StatusBadRequest, "Error! Couldn't submit form.")
			return

		}

	}

	getOrderEmail := queries["getOrderEmail"]

	// now := time.Now()
	// then := now.AddDate(0, 0, -12)

	if err := dbConn.db.Raw(getOrderEmail, orders[0].UserUUID).Scan(&emailOrder); err != nil {
		log.Println(err)
	}

	for i := range emailOrder {
		emailOrder[i].FormatTime = emailOrder[i].UpdatedAt.Format("Mon Jan _2 15:04:05 2006")
	}

	payloadEmail := struct {
		EmailOrder []EmailOrder
	}{
		EmailOrder: emailOrder,
	}

	// TODO: send email to all admins

	sendOrdersEmail("saburchfield@gmail.com", payloadEmail)

	viewRender.Text(w, http.StatusCreated, "Success!")

}

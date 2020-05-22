package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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

	now := time.Now()
	then := now.AddDate(0, 0, -12)

	if err := dbConn.db.Raw(getOrderEmail, then, now, orders[0].UserUUID).Scan(&emailOrder); err != nil {
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

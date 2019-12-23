
package main

import (
	// "fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
}

func home(w http.ResponseWriter, r *http.Request) {

	session, err := sessCookieStore.Get(r, "inventory-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSessionActive(session) {

		payload := struct {
			U            *user
		}{
			U:           getUserFromSession(r),
		}


		viewRender.HTML(w, http.StatusOK, "home", payload)
		return
	}

	payload := struct {
		ErrMsg      string
	}{
		ErrMsg:      "Session timedout, please login again",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

func items(w http.ResponseWriter, r *http.Request) {

	session, err := sessCookieStore.Get(r, "inventory-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSessionActive(session) {

		var item []Items
		var deletedItems []Items
		var categories []Categories

		if err := dbConn.db.Find(&categories).Scan(&categories).Error; err != nil {

			log.Println("Error with retrieving categories: ", err)

		}

		if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

			log.Println("Error with retrieving items: ", err)

		}

		if err := dbConn.db.Raw("SELECT * FROM inventory.items WHERE deleted_at IS NOT null").Scan(&deletedItems).Error; err != nil {

			log.Println("Error with retrieving deleted items: ", err)

		}

		payload := struct {
			U            *user
			Message      string
			Items        []Items
			DeletedItems []Items
			Categories   []Categories
		}{
			U:           getUserFromSession(r),
			Message:     "",
			Items: 			item,
			DeletedItems: deletedItems,
			Categories: categories,
		}

		viewRender.HTML(w, http.StatusOK, "items", payload)
		return
	}

	payload := struct {
		ErrMsg      string
	}{
		ErrMsg:      "Session timedout, please login again",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

func stores(w http.ResponseWriter, r *http.Request) {

	session, err := sessCookieStore.Get(r, "inventory-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSessionActive(session) {

		var stores []Stores
		var deletedStores []Stores

		if err := dbConn.db.Find(&stores).Scan(&stores).Error; err != nil {

			log.Println("Error with retrieving items: ", err)

		}

		if err := dbConn.db.Raw("SELECT * FROM inventory.stores WHERE deleted_at IS NOT null").Scan(&deletedStores).Error; err != nil {

			log.Println("Error with retrieving deleted items: ", err)

		}

		payload := struct {
			U            *user
			Message      string
			Stores        []Stores
			DeletedStores []Stores
		}{
			U:           getUserFromSession(r),
			Message:     "",
			Stores: 			stores,
			DeletedStores: deletedStores,
		}

		viewRender.HTML(w, http.StatusOK, "stores", payload)
		return
	}

	payload := struct {
		ErrMsg      string
	}{
		ErrMsg:      "Session timedout, please login again",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

func categories(w http.ResponseWriter, r *http.Request) {

	session, err := sessCookieStore.Get(r, "inventory-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSessionActive(session) {

		var categories []Categories
		var deletedCategories []Categories

		if err := dbConn.db.Find(&categories).Scan(&categories).Error; err != nil {

			log.Println("Error with retrieving items: ", err)

		}

		if err := dbConn.db.Raw("SELECT * FROM inventory.categories WHERE deleted_at IS NOT null").Scan(&deletedCategories).Error; err != nil {

			log.Println("Error with retrieving deleted items: ", err)

		}

		payload := struct {
			U            *user
			Message      string
			Categories        []Categories
			DeletedCategories []Categories
		}{
			U:           getUserFromSession(r),
			Message:     "",
			Categories: 			categories,
			DeletedCategories: deletedCategories,
		}

		viewRender.HTML(w, http.StatusOK, "categories", payload)
		return
	}

	payload := struct {
		ErrMsg      string
	}{
		ErrMsg:      "Session timedout, please login again",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

func users(w http.ResponseWriter, r *http.Request) {

	session, err := sessCookieStore.Get(r, "inventory-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if isSessionActive(session) {

		var users []user
		var deletedUsers []user

		if err := dbConn.db.Find(&users).Scan(&users).Error; err != nil {

			log.Println("Error with retrieving users: ", err)

		}

		if err := dbConn.db.Raw("SELECT * FROM inventory.users WHERE deleted_at IS NOT null").Scan(&deletedUsers).Error; err != nil {

			log.Println("Error with retrieving deleted items: ", err)

		}

		payload := struct {
			U            *user
			Message      string
			Users        []user
			DeletedUsers []user
		}{
			U:           getUserFromSession(r),
			Message:     "",
			Users: 			users,
			DeletedUsers: deletedUsers,
		}

		viewRender.HTML(w, http.StatusOK, "users", payload)
		return
	}

	payload := struct {
		ErrMsg      string
	}{
		ErrMsg:      "Session timedout, please login again",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

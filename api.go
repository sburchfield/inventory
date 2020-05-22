package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	// "fmt"

	"gnardex/gosecrets"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func updateItems(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var formvalues = Items{
		ItemName:  strings.ToLower(r.FormValue("item_name")),
		ItemCost:  r.FormValue("item_cost"),
		ItemPrice: r.FormValue("item_price"),
		Category:  strings.ToLower(r.FormValue("category")),
	}

	if err := dbConn.db.Save(&formvalues).Error; err != nil {

		log.Println("Error with items table update: ", err)

		return

	}

	var item []Items

	if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

		log.Println("Error with retrieving items: ", err)

	}

	payload := struct {
		Message string
		Items   []Items
	}{
		Items:   item,
		Message: "Form successfully updated",
	}

	viewRender.JSON(w, http.StatusOK, payload)

	return

}

func removeItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	itemID := vars["item_id"]

	var item []Items

	if err := dbConn.db.Where("id = ?", itemID).Delete(item).Error; err != nil {

		log.Println("Error with deleting items: ", err)

	}

	if err := dbConn.db.Raw("Select * FROM inventory.items WHERE deleted_at IS NOT NULL").Scan(&item).Error; err != nil {

		log.Println("Error with retrieving items: ", err)

	}

	payload := struct {
		Message  string
		Items    []Items
		RemoveBy string
	}{
		Items:    item,
		Message:  "Item Deleted",
		RemoveBy: itemID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func restoreItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// item_id := vars["item_id"]
	itemID, err := strconv.Atoi(vars["item_id"])
	if err != nil {
		log.Println("string to int error", err)
	}

	var item []Items

	if err := dbConn.db.Exec("UPDATE inventory.items SET deleted_at = NULL WHERE id = ?", itemID).Error; err != nil {

		log.Println("Error with deleting items: ", err)

	}

	if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

		log.Println("Error with retrieving items: ", err)

	}

	payload := struct {
		Message  string
		Items    []Items
		RemoveBy int
	}{
		Items:    item,
		Message:  "Item Restored",
		RemoveBy: itemID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func updateStores(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var formvalues = Stores{
		StoreName:   strings.ToLower(r.FormValue("store_name")),
		Address:     r.FormValue("address"),
		PhoneNumber: r.FormValue("phone_number"),
		City:        r.FormValue("city"),
		State:       r.FormValue("state"),
		ZipCode:     r.FormValue("zip_code"),
	}

	if err := dbConn.db.Save(&formvalues).Error; err != nil {

		log.Println("Error with items table update: ", err)

		return

	}

	var stores []Stores

	if err := dbConn.db.Find(&stores).Scan(&stores).Error; err != nil {

		log.Println("Error with retrieving items: ", err)

	}

	payload := struct {
		Message string
		Stores  []Stores
	}{
		Stores:  stores,
		Message: "Form successfully updated",
	}

	viewRender.JSON(w, http.StatusOK, payload)

	return

}

func removeStore(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeID := vars["store_id"]

	var store []Stores

	if err := dbConn.db.Where("id = ?", storeID).Delete(store).Error; err != nil {

		log.Println("Error with deleting stores: ", err)

	}

	if err := dbConn.db.Raw("Select * FROM inventory.stores WHERE deleted_at IS NOT NULL").Scan(&store).Error; err != nil {

		log.Println("Error with retrieving stores: ", err)

	}

	payload := struct {
		Message  string
		Stores   []Stores
		RemoveBy string
	}{
		Stores:   store,
		Message:  "Store Deleted",
		RemoveBy: storeID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func restoreStore(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeID, err := strconv.Atoi(vars["store_id"])
	if err != nil {
		log.Println("string to int error", err)
	}

	var store []Stores

	if err := dbConn.db.Exec("UPDATE inventory.stores SET deleted_at = NULL WHERE id = ?", storeID).Error; err != nil {

		log.Println("Error with deleting items: ", err)

	}

	if err := dbConn.db.Find(&store).Scan(&store).Error; err != nil {

		log.Println("Error with retrieving items: ", err)

	}

	payload := struct {
		Message  string
		Stores   []Stores
		RemoveBy int
	}{
		Stores:   store,
		Message:  "Store Restored",
		RemoveBy: storeID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func selectStores(w http.ResponseWriter, r *http.Request) {

	var store []Stores

	if err := dbConn.db.Find(&store).Scan(&store).Error; err != nil {

		log.Println("Error with retrieving stores: ", err)

	}

	payload := struct {
		Stores []Stores
	}{
		Stores: store,
	}

	viewRender.JSON(w, http.StatusOK, payload)
}

func selectItems(w http.ResponseWriter, r *http.Request) {

	var item []Items

	if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

		log.Println("Error with retrieving stores: ", err)

	}

	payload := struct {
		Items []Items
	}{
		Items: item,
	}

	viewRender.JSON(w, http.StatusOK, payload)
}

func updateCategories(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var formvalues = Categories{
		Description: strings.ToLower(r.FormValue("description")),
		Category:    strings.ToLower(r.FormValue("category")),
	}

	if err := dbConn.db.Save(&formvalues).Error; err != nil {

		log.Println("Error with Categories table update: ", err)

		return

	}

	var categories []Categories

	if err := dbConn.db.Find(&categories).Scan(&categories).Error; err != nil {

		log.Println("Error with retrieving Categories: ", err)

	}

	payload := struct {
		Message    string
		Categories []Categories
	}{
		Categories: categories,
		Message:    "Form successfully updated",
	}

	viewRender.JSON(w, http.StatusOK, payload)

	return

}

func removeCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryID := vars["category_id"]
	categoryName := vars["category"]

	var category []Categories
	var item []Items

	if err := dbConn.db.Where("id = ?", categoryID).Delete(category).Error; err != nil {

		log.Println("Error with deleting categories: ", err)

	}

	if err := dbConn.db.Where("category = ?", categoryName).Delete(item).Error; err != nil {

		log.Println("Error with deleting items when deleting a category: ", err)

	}

	if err := dbConn.db.Raw("Select * FROM inventory.categories WHERE deleted_at IS NOT NULL").Scan(&category).Error; err != nil {

		log.Println("Error with retrieving categories: ", err)

	}

	payload := struct {
		Message    string
		Categories []Categories
		RemoveBy   string
	}{
		Categories: category,
		Message:    "Category Deleted",
		RemoveBy:   categoryID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func restoreCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// item_id := vars["item_id"]
	categoryID, err := strconv.Atoi(vars["category_id"])
	categoryName := vars["category"]
	if err != nil {
		log.Println("string to int error", err)
	}

	var category []Categories

	if err := dbConn.db.Exec("UPDATE inventory.categories SET deleted_at = NULL WHERE id = ?", categoryID).Error; err != nil {

		log.Println("Error with restoring categories: ", err)

	}

	if err := dbConn.db.Exec("UPDATE inventory.items SET deleted_at = NULL WHERE category = ?", categoryName).Error; err != nil {

		log.Println("Error with restoring items when deleting a category: ", err)

	}

	if err := dbConn.db.Find(&category).Scan(&category).Error; err != nil {

		log.Println("Error with retrieving categories: ", err)

	}

	payload := struct {
		Message    string
		Categories []Categories
		RemoveBy   int
	}{
		Categories: category,
		Message:    "Item Restored",
		RemoveBy:   categoryID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func signupAction(w http.ResponseWriter, r *http.Request) {

	pw := gosecrets.GeneratePassword()
	userUUID := uuid.New().String()

	var users []user

	r.ParseForm()

	payload := struct {
		Label string
	}{
		Label: "",
	}

	//Generate hashed password
	ps, err := gosecrets.GetPasswordHash(pw)
	if err != nil {
		log.Println(err)
		payload.Label = "There was an error with assigning the user a password"
		viewRender.HTML(w, http.StatusOK, "users", payload)
		return

	}

	var u = user{
		UserUUID:     userUUID,
		UserEmail:    strings.ToLower(r.FormValue("username")),
		Username:     strings.ToLower(r.FormValue("username")),
		FirstName:    strings.ToLower(r.FormValue("first_name")),
		LastName:     strings.ToLower(r.FormValue("last_name")),
		Role:         strings.ToLower(r.FormValue("role")),
		PasswordHash: ps,
		Status:       "active",
		ResetTime:    nil,
	}

	//Check if the user is already signed up
	un := u.Username
	var count int

	if err := dbConn.db.Model(&user{}).Where("username = ?", un).Count(&count).Error; err != nil {
		log.Println(err)

		viewRender.Text(w, http.StatusOK, "Sorry! There was an error in submitting the form.")
		return

	}

	if count > 0 {
		log.Println("User already exists")
		viewRender.Text(w, http.StatusOK, "User with this email ("+un+") already present in the system.")
		return

	}

	//Check user provided email
	if err := checkmail.ValidateFormat(un); err != nil {
		log.Println("Not a valid user email")
		log.Println(err)
		viewRender.Text(w, http.StatusOK, "Please provide a valid email for registration.")
		return
	}

	//create user

	if err := dbConn.db.Create(&u).Error; err != nil {

		log.Println(err)
		viewRender.Text(w, http.StatusOK, "Sorry! There was an error in submitting the form")
		return

	}

	emailPayload := struct {
		Password string
		Username string
	}{
		Password: pw,
		Username: un,
	}

	//Send signup email
	if err := sendSignupEmail(un, emailPayload); err != nil {
		viewRender.Text(w, http.StatusOK, "Signup complete. You can login now, but due to some internal issues unable to send the confirmation email.")
		return

	}

	if err := dbConn.db.Find(&users).Scan(&users).Error; err != nil {

		log.Println("Error with retrieving user: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Retrieving users list.")

	}

	new_payload := struct {
		Message  string
		U        []user
		RemoveBy string
	}{
		U:        users,
		Message:  "Success! User has been added!",
		RemoveBy: userUUID,
	}

	viewRender.JSON(w, http.StatusOK, new_payload)

}

func removeUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userUUID := vars["user_uuid"]

	var u []user

	if err := dbConn.db.Where("user_uuid = ?", userUUID).Delete(u).Error; err != nil {

		log.Println("Error with deleting user: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Deleting user.")

	}

	if err := dbConn.db.Raw("Select * FROM inventory.users WHERE deleted_at IS NOT NULL").Scan(&u).Error; err != nil {

		log.Println("Error with retrieving categories: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Retrieving users list.")

	}

	payload := struct {
		Message  string
		U        []user
		RemoveBy string
	}{
		U:        u,
		Message:  "User Deleted",
		RemoveBy: userUUID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func restoreUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userUUID := vars["user_uuid"]

	var u []user

	if err := dbConn.db.Exec("UPDATE inventory.users SET deleted_at = NULL WHERE user_uuid = ?", userUUID).Error; err != nil {

		log.Println("Error with deleting user: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Restoring user.")

	}

	if err := dbConn.db.Find(&u).Scan(&u).Error; err != nil {

		log.Println("Error with retrieving user: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Retrieving users list.")

	}

	payload := struct {
		Message  string
		U        []user
		RemoveBy string
	}{
		U:        u,
		Message:  "Store Restored",
		RemoveBy: userUUID,
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

func updateRole(w http.ResponseWriter, r *http.Request) {

	var u user

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	if err := dbConn.db.Exec("UPDATE inventory.users SET role = ? WHERE username = ?", u.Role, u.Username).Error; err != nil {

		log.Println("Error with deleting user: ", err)
		viewRender.Text(w, http.StatusOK, "Error! Updating user.")

	}

	payload := struct {
		Message string
		User    string
	}{
		User:    u.Role,
		Message: "User Updated!",
	}

	viewRender.JSON(w, http.StatusOK, payload)

}

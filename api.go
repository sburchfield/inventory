package main

import(
  "log"
  "net/http"
  "strings"
  "strconv"
  "encoding/json"
  "time"
  // "fmt"

  "github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func updateItems (w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  var formvalues = Items{
    ItemName: strings.ToLower(r.FormValue("item_name")),
    ItemCost: r.FormValue("item_cost"),
    ItemPrice: r.FormValue("item_price"),
    Category: strings.ToLower(r.FormValue("category")),
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
    Message      string
    Items        []Items
  }{
    Items: 			 item,
    Message:     "Form successfully updated",
  }

  viewRender.JSON(w, http.StatusOK, payload)

  return

}

func removeItem ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
	item_id := vars["item_id"]

  var item []Items

  if err :=  dbConn.db.Where("id = ?", item_id).Delete(item).Error; err != nil {

    log.Println("Error with deleting items: ", err)

  }

  if err := dbConn.db.Raw("Select * FROM inventory.items WHERE deleted_at IS NOT NULL").Scan(&item).Error; err != nil {

    log.Println("Error with retrieving items: ", err)

  }

  payload := struct {
    Message      string
    Items        []Items
    RemoveBy     string
  }{
    Items: 			 item,
    Message:     "Item Deleted",
    RemoveBy:    item_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}

func restoreItem ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
	// item_id := vars["item_id"]
  item_id, err := strconv.Atoi(vars["item_id"])
  if err != nil{
    log.Println("string to int error", err)
  }

  var item []Items

  if err :=  dbConn.db.Exec("UPDATE inventory.items SET deleted_at = NULL WHERE id = ?", item_id).Error; err != nil {

    log.Println("Error with deleting items: ", err)

  }

  if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

    log.Println("Error with retrieving items: ", err)

  }

  payload := struct {
    Message      string
    Items        []Items
    RemoveBy     int
  }{
    Items: 			 item,
    Message:     "Item Restored",
    RemoveBy:    item_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}


func updateStores (w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  var formvalues = Stores{
    StoreName: strings.ToLower(r.FormValue("store_name")),
    Address: r.FormValue("address"),
    PhoneNumber: r.FormValue("phone_number"),
    City: r.FormValue("city"),
    State: r.FormValue("state"),
    ZipCode: r.FormValue("zip_code"),
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
    Message      string
    Stores       []Stores
  }{
    Stores: 			stores,
    Message:     "Form successfully updated",
  }

  viewRender.JSON(w, http.StatusOK, payload)

  return

}

func removeStore ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
	store_id := vars["store_id"]

  var store []Stores

  if err :=  dbConn.db.Where("id = ?", store_id).Delete(store).Error; err != nil {

    log.Println("Error with deleting stores: ", err)

  }

  if err := dbConn.db.Raw("Select * FROM inventory.stores WHERE deleted_at IS NOT NULL").Scan(&store).Error; err != nil {

    log.Println("Error with retrieving stores: ", err)

  }

  payload := struct {
    Message      string
    Stores       []Stores
    RemoveBy     string
  }{
    Stores: 			 store,
    Message:     "Store Deleted",
    RemoveBy:    store_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}

func restoreStore ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
  store_id, err := strconv.Atoi(vars["store_id"])
  if err != nil{
    log.Println("string to int error", err)
  }

  var store []Stores

  if err :=  dbConn.db.Exec("UPDATE inventory.stores SET deleted_at = NULL WHERE id = ?", store_id).Error; err != nil {

    log.Println("Error with deleting items: ", err)

  }

  if err := dbConn.db.Find(&store).Scan(&store).Error; err != nil {

    log.Println("Error with retrieving items: ", err)

  }

  payload := struct {
    Message      string
    Stores        []Stores
    RemoveBy     int
  }{
    Stores: 			 store,
    Message:     "Store Restored",
    RemoveBy:    store_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}

func selectStores( w http.ResponseWriter, r *http.Request ){

  var store []Stores

  if err := dbConn.db.Find(&store).Scan(&store).Error; err != nil {

    log.Println("Error with retrieving stores: ", err)

  }

  payload := struct {
    Stores       []Stores
  }{
    Stores: 		 store,
  }

  viewRender.JSON(w, http.StatusOK, payload)
}

func selectItems( w http.ResponseWriter, r *http.Request ){

  var item []Items

  if err := dbConn.db.Find(&item).Scan(&item).Error; err != nil {

    log.Println("Error with retrieving stores: ", err)

  }

  payload := struct {
    Items       []Items
  }{
    Items: 		 item,
  }

  viewRender.JSON(w, http.StatusOK, payload)
}

func updateOrders( w http.ResponseWriter, r *http.Request ){

  var orders []Orders
  var emailOrder []EmailOrder

  err := json.NewDecoder(r.Body).Decode(&orders)
  if err != nil {
      log.Println("YOUR ERROR: ",err)
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

  if err := dbConn.db.Raw(getOrderEmail, then, now, orders[0].UserUuid).Scan(&emailOrder);
  err != nil {
      log.Println(err)
    }

    log.Println(emailOrder)

    // defer rows.Close()
    // for rows.Next() {
    //   dbConn.db.ScanRows(rows, &emailOrder)
    //   log.Println(emailOrder)
    // }

    payloadEmail := struct {
      EmailOrder   []EmailOrder
    }{
      EmailOrder: emailOrder,
    }

    sendOrdersEmail("saburchfield@gmail.com", payloadEmail)





  viewRender.Text(w, http.StatusCreated, "Success!")


}

func updateCategories (w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  var formvalues = Categories{
    Description: strings.ToLower(r.FormValue("description")),
    Category: strings.ToLower(r.FormValue("category")),
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
    Message      string
    Categories        []Categories
  }{
    Categories: 			 categories,
    Message:     "Form successfully updated",
  }

  viewRender.JSON(w, http.StatusOK, payload)

  return

}

func removeCategory ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
  category_id := vars["category_id"]
	category_name := vars["category"]

  var category []Categories
  var item []Items

  if err :=  dbConn.db.Where("id = ?", category_id).Delete(category).Error; err != nil {

    log.Println("Error with deleting categories: ", err)

  }

  if err :=  dbConn.db.Where("category = ?", category_name).Delete(item).Error; err != nil {

    log.Println("Error with deleting items when deleting a category: ", err)

  }

  if err := dbConn.db.Raw("Select * FROM inventory.categories WHERE deleted_at IS NOT NULL").Scan(&category).Error; err != nil {

    log.Println("Error with retrieving categories: ", err)

  }

  payload := struct {
    Message      string
    Categories        []Categories
    RemoveBy     string
  }{
    Categories: 			 category,
    Message:     "Category Deleted",
    RemoveBy:    category_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}

func restoreCategory ( w http.ResponseWriter, r *http.Request ){

  vars := mux.Vars(r)
	// item_id := vars["item_id"]
  category_id, err := strconv.Atoi(vars["category_id"])
  category_name := vars["category"]
  if err != nil{
    log.Println("string to int error", err)
  }

  var category []Categories

  if err :=  dbConn.db.Exec("UPDATE inventory.categories SET deleted_at = NULL WHERE id = ?", category_id).Error; err != nil {

    log.Println("Error with restoring categories: ", err)

  }

  if err :=  dbConn.db.Exec("UPDATE inventory.items SET deleted_at = NULL WHERE category = ?", category_name).Error; err != nil {

    log.Println("Error with restoring items when deleting a category: ", err)

  }

  if err := dbConn.db.Find(&category).Scan(&category).Error; err != nil {

    log.Println("Error with retrieving categories: ", err)

  }

  payload := struct {
    Message      string
    Categories        []Categories
    RemoveBy     int
  }{
    Categories: 			 category,
    Message:     "Item Restored",
    RemoveBy:    category_id,
  }

  viewRender.JSON(w, http.StatusOK, payload)

}

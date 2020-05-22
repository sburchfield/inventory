package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var muxRouter *mux.Router
var adminMuxRouter *mux.Router
var webMuxRouter *mux.Router
var apiMuxRouter *mux.Router

func defineRoutes() {

	s1 := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets")))
	muxRouter.PathPrefix("/assets/").Handler(s1)

	muxRouter.HandleFunc("/", index)

	muxRouter.HandleFunc("/login", login)
	muxRouter.HandleFunc("/loginaction", loginAction)

	muxRouter.HandleFunc("/passwordresetrequest", passwordResetRequest)
	muxRouter.HandleFunc("/passwordresetrequestaction", passwordResetRequestAction).Methods("POST")
	muxRouter.HandleFunc("/passwordreset/{user_uuid}", passwordReset)
	muxRouter.HandleFunc("/passwordresetaction", passwordResetAction).Methods("POST")
	muxRouter.HandleFunc("/passwordresetsuccess", passwordResetSuccess)

	muxRouter.HandleFunc("/signup", signup)
	muxRouter.HandleFunc("/signupaction", signupAction).Methods("POST")

	webMuxRouter.HandleFunc("/home", home)
	adminMuxRouter.HandleFunc("/admin/items", items)
	adminMuxRouter.HandleFunc("/admin/stores", stores)
	adminMuxRouter.HandleFunc("/admin/categories", categories)
	adminMuxRouter.HandleFunc("/admin/users", users)

	apiMuxRouter.HandleFunc("/api/updateCategories", updateCategories)
	apiMuxRouter.HandleFunc("/api/removeCategory/{category_id}/{category}", removeCategory)
	apiMuxRouter.HandleFunc("/api/restoreCategory/{category_id}/{category}", restoreCategory)

	apiMuxRouter.HandleFunc("/api/removeUser/{user_uuid}", removeUser)
	apiMuxRouter.HandleFunc("/api/restoreUser/{user_uuid}", restoreUser)

	apiMuxRouter.HandleFunc("/api/updateItems", updateItems)
	apiMuxRouter.HandleFunc("/api/removeItem/{item_id}", removeItem)
	apiMuxRouter.HandleFunc("/api/restoreItem/{item_id}", restoreItem)

	apiMuxRouter.HandleFunc("/api/updateStores", updateStores)
	apiMuxRouter.HandleFunc("/api/removeStore/{store_id}", removeStore)
	apiMuxRouter.HandleFunc("/api/restoreStore/{store_id}", restoreStore)

	apiMuxRouter.HandleFunc("/api/updateRole", updateRole).Methods("POST")

	apiMuxRouter.HandleFunc("/api/select/stores", selectStores)
	apiMuxRouter.HandleFunc("/api/select/items", selectItems)

	apiMuxRouter.HandleFunc("/api/updateOrders", updateOrders).Methods("POST")
	apiMuxRouter.HandleFunc("/api/getLatestOrders/{user_uuid}", getLatestOrders)

	muxRouter.HandleFunc("/logout", logout)

}

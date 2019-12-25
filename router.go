
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var muxRouter *mux.Router
var secureMuxRouter *mux.Router

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

	muxRouter.HandleFunc("/home", home)
	secureMuxRouter.HandleFunc("/admin/items", items)
	secureMuxRouter.HandleFunc("/admin/stores", stores)
	secureMuxRouter.HandleFunc("/admin/categories", categories)
	secureMuxRouter.HandleFunc("/admin/users", users)

	muxRouter.HandleFunc("/updateCategories", updateCategories)
	muxRouter.HandleFunc("/removeCategory/{category_id}/{category}", removeCategory)
	muxRouter.HandleFunc("/restoreCategory/{category_id}/{category}", restoreCategory)

	muxRouter.HandleFunc("/removeUser/{user_uuid}", removeUser)
	muxRouter.HandleFunc("/restoreUser/{user_uuid}", restoreUser)

	muxRouter.HandleFunc("/updateItems", updateItems)
	muxRouter.HandleFunc("/removeItem/{item_id}", removeItem)
	muxRouter.HandleFunc("/restoreItem/{item_id}", restoreItem)

	muxRouter.HandleFunc("/updateStores", updateStores)
	muxRouter.HandleFunc("/removeStore/{store_id}", removeStore)
	muxRouter.HandleFunc("/restoreStore/{store_id}", restoreStore)

	muxRouter.HandleFunc("/select/stores", selectStores)
	muxRouter.HandleFunc("/select/items", selectItems)

	muxRouter.HandleFunc("/updateOrders", updateOrders).Methods("POST")



	muxRouter.HandleFunc("/logout", logout)

}

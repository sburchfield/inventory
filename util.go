
package main

import (
  "errors"
  "net/http"
  "log"
	"github.com/gorilla/sessions"
)


func handleCriticalError(err error) {

  	log.Println(err.Error())

  }

func handleLoginError(w http.ResponseWriter, r *http.Request){

  http.Redirect(w, r, "/login", 302)
  return

}

func isSessionActive(session *sessions.Session) bool {

	if session.Values["active"] == "on" {

		return true

	}

	return false

}

func getUserFromSession(r *http.Request) *user {

	session, err := sessCookieStore.Get(r, "inventory-session")
	if err != nil {

		handleCriticalError(err)
		return &user{}

	}

	active := session.Values["active"].(string)
	if active != "on" {

		handleCriticalError(errors.New("Session is inactive"))
		return &user{}

	}

	return session.Values["user"].(*user)

}

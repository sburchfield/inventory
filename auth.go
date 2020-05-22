package main

import (
	"net/http"
	// "strings"
	"log"
	"strconv"
	"time"

	// "github.com/badoux/checkmail"
	"gnardex/gosecrets"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := sessCookieStore.Get(r, "inventory-session")
	if isSessionActive(session) {
		http.Redirect(w, r, "/home", 303)
		return
	}

	payload := struct {
		ErrMsg string
	}{
		ErrMsg: "",
	}

	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
}

func loginAction(w http.ResponseWriter, r *http.Request) {

	var u user

	un := r.FormValue("username")
	pw := r.FormValue("password")

	if err := dbConn.db.Where("status = ? AND username = ?", "active", un).First(&u).Error; err != nil {

		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Password reset has been requested. Please follow the email sent to your address. If you did not request a Password Reset please contact support@pyaanalytics.com",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	found, u := auth(un, pw)
	if !found {

		log.Println("username and pw not found")
		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Invalid username or password.",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return
	}

	session, _ := sessCookieStore.Get(r, "inventory-session")
	session.Values["active"] = "on"
	session.Values["user"] = u

	if err := session.Save(r, w); err != nil {

		handleCriticalError(err)
		viewRender.Text(w, http.StatusInternalServerError, "Invalid credentials, please try again.")
		return

	}
	http.Redirect(w, r, "/", 303)
	// viewRender.HTML(w, http.StatusOK, "<p>Logged In</p>", "")

}

func auth(un, ps string) (bool, user) {

	var u user

	if err := dbConn.db.Where("username = ?", un).
		First(&u).Error; err != nil {

		return false, u

	}
	return gosecrets.CompareHashWithPassword(u.PasswordHash, ps), u

}

func pwResetCheck(pwResetStatus string) bool {

	var u user

	if err := dbConn.db.Where("status = ?", pwResetStatus).First(&u).Error; err != nil {

		return false

	}
	return true

}

func checkUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, err := sessCookieStore.Get(r, "inventory-session")
	if err != nil {
		handleLoginError(w, r)
		return
	}

	if !isSessionActive(session) {
		handleLoginError(w, r)
		return
	}

	val := session.Values["user"]

	// var u = &user{}
	_, ok := val.(*user)
	if ok != true {
		handleLoginError(w, r)
		return
	}

	next(w, r)

}

func checkAdmin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, err := sessCookieStore.Get(r, "inventory-session")
	if err != nil {
		handleLoginError(w, r)
		return
	}

	if !isSessionActive(session) {
		handleLoginError(w, r)
		return
	}

	val := session.Values["user"]

	var u = &user{}
	u, ok := val.(*user)
	if ok != true {
		handleLoginError(w, r)
		return
	}

	if u.Role != "admin" {
		handleLoginError(w, r)
		return
	}

	next(w, r)

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := sessCookieStore.Get(r, "inventory-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["user"] = ""
	session.Values["active"] = ""
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func signup(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		Label     string
		FirstName string
		LastName  string
		Username  string
	}{
		Label:     "",
		FirstName: "",
		LastName:  "",
		Username:  "",
	}
	viewRender.HTML(w, http.StatusOK, "signup", payload, noLayout)

}

func passwordResetRequest(w http.ResponseWriter, r *http.Request) {

	payload := struct {
		ErrMsg string
	}{
		ErrMsg: "",
	}

	viewRender.HTML(w, http.StatusOK, "passwordResetRequest", payload, noLayout)

}

func passwordResetRequestAction(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")

	payload := struct {
		ErrMsg string
	}{
		ErrMsg: "",
	}

	var u user

	getUser := queries["getUser"]

	if err := dbConn.db.Raw(getUser, username).Scan(&u).Error; err != nil {

		payload.ErrMsg = "Username invalid or not found"
		viewRender.HTML(w, http.StatusOK, "passwordResetRequest", payload, noLayout)
		return

	}

	passwordResetUUID := uuid.New().String()

	passwordResetHash, err := gosecrets.GetPasswordHash(passwordResetUUID)
	if err != nil {
		log.Println(err)
		return
	}

	requestReset := queries["requestReset"]

	if err := dbConn.db.Exec(requestReset, passwordResetHash, time.Now().UTC(), username).Error; err != nil {
		log.Println(err)
		payload.ErrMsg = "Password request/Code is invalid"
		viewRender.HTML(w, http.StatusOK, "passwordResetRequest", payload, noLayout)
		return

	}

	payloadEmail := struct {
		Username string
		UserUUID string
		Code     string
		Route    string
	}{
		Username: username,
		UserUUID: u.UserUUID,
		Code:     passwordResetHash,
		Route:    envVars.appPasswordResetDomain,
	}

	sendPasswordResetEmail(u.UserEmail, payloadEmail)

	payload.ErrMsg = "An email was sent to the email-id on file, please follow the instructions within " +
		strconv.FormatUint(uint64(envVars.appPasswordResetLinkExpiryTime), 10) + " mintues to reset your password."
	viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)

}

func passwordReset(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userUUID := vars["user_uuid"]

	code := r.URL.Query().Get("code")

	var user user
	getUserByUUID := queries["getUserByUUID"]
	if err := dbConn.db.Raw(getUserByUUID, userUUID).
		Scan(&user).
		Error; err != nil {

		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Sorry! Could not process request.",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	if user.PasswordResetHash != code {

		log.Printf("Invalid Security Code. Attempted Code: %s", code)
		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Access Forbidden",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	now := time.Now().UTC()
	duration := now.Sub(*user.ResetTime)
	minutes := duration.Minutes()
	if minutes > float64(envVars.appPasswordResetLinkExpiryTime) {

		log.Print("Link expired")
		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Password reset link expired, please request another.",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	payload := passResetMessage{
		UserUUID: userUUID,
		Code:     code,
		Message:  "",
	}
	viewRender.HTML(w, http.StatusOK, "passwordReset", payload, noLayout)

}

func passwordResetAction(w http.ResponseWriter, r *http.Request) {

	userUUID := r.PostFormValue("user_uuid")
	securityCode := r.PostFormValue("securityCode")
	newPassword := r.PostFormValue("newPassword")
	confirmPassword := r.PostFormValue("confirmNewPassword")

	if newPassword != confirmPassword {

		log.Println("password confimation does not match")
		payload := passResetMessage{
			UserUUID: userUUID,
			Code:     securityCode,
			Message:  "New password and confirm password don not match, please try again.",
		}
		viewRender.HTML(w, http.StatusOK, "passwordReset", payload, noLayout)
		return

	}

	var user user
	getUserByUUID := queries["getUserByUUID"]
	if err := dbConn.db.Raw(getUserByUUID, userUUID).
		Scan(&user).
		Error; err != nil {

		log.Printf("Invalid user. Attempted id: %s", userUUID)

		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Invalid Access!",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	if user.PasswordResetHash != securityCode {

		log.Printf("Invalid Security Code. Attempted Code: %s", securityCode)
		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Invalid Access!",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	now := time.Now()
	duration := now.Sub(*user.ResetTime)
	minutes := duration.Minutes()
	if minutes > float64(envVars.appPasswordResetLinkExpiryTime) {

		log.Println("Link expired -1")
		payload := struct {
			ErrMsg string
		}{
			ErrMsg: "Password reset link expired, please reqest a new one.",
		}
		viewRender.HTML(w, http.StatusOK, "login", payload, noLayout)
		return

	}

	if gosecrets.CompareHashWithPassword(user.PasswordHash, newPassword) {

		log.Println("Old Password and new password should not be the same.")
		payload := passResetMessage{
			UserUUID: userUUID,
			Code:     securityCode,
			Message:  "Old Password and new password should not be the same, please try again.",
		}
		viewRender.HTML(w, http.StatusOK, "passwordReset", payload, noLayout)
		return

	}

	pass, err := gosecrets.GetPasswordHash(newPassword)
	if err != nil {

		payload := passResetMessage{
			UserUUID: userUUID,
			Code:     securityCode,
			Message:  err.Error(),
		}
		viewRender.HTML(w, http.StatusOK, "passwordReset", payload, noLayout)
		return

	}

	resetPass := queries["resetPass"]
	if err := dbConn.db.Exec(resetPass, pass, userUUID).
		Error; err != nil {

		payload := passResetMessage{
			UserUUID: userUUID,
			Code:     securityCode,
			Message:  "Sorry! Could not process request",
		}
		viewRender.HTML(w, http.StatusOK, "passwordReset", payload, noLayout)
		return

	}

	http.Redirect(w, r, "/passwordresetsuccess", 303)

}

func passwordResetSuccess(w http.ResponseWriter, r *http.Request) {

	viewRender.HTML(w, http.StatusOK, "passwordResetSuccess", "", noLayout)

}

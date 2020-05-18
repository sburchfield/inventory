
package main

import (
	"log"
	"fmt"
	"net"
	"os"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	mailer "gnardex/mailer"
)

func (ev *envVariables) init() {

	ev.appMode     = os.Getenv("APP_MODE")
	ev.appPort     = os.Getenv("APP_PORT")
	ev.dbSchema    = os.Getenv("DB_SCHEMA")
	ev.dbHost      = os.Getenv("DB_HOST")
	ev.dbName      = os.Getenv("DB_NAME")
	ev.dbUsername  = os.Getenv("DB_USER")
	ev.dbPassword  = os.Getenv("DB_PASSWORD")
	ev.dbSSLMode   = os.Getenv("DB_SSLMODE")
	ev.appPasswordResetDomain = os.Getenv("APP_PASSWORD_RESET_DOMAIN")

	if ev.appPasswordResetDomain == "" {

		log.Println("Password reset domain not set!")

	}
	// password reset email link exipry time in minutes
	ev.appPasswordResetLinkExpiryTime = 30

	log.Println("Server mode:", ev.appMode)

   if ev.appMode == "local" {
  		getLocalIp(ev)
  	}

}

var envVars envVariables

func init() {

	envVars = envVariables{}
	envVars.init()

	muxRouter = mux.NewRouter()

	adminMuxRouter = mux.NewRouter()
	webMuxRouter = mux.NewRouter()
	apiMuxRouter = mux.NewRouter()
	
	muxRouter.PathPrefix("/admin").Handler(negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
		negroni.HandlerFunc(checkAdmin),
		negroni.Wrap(adminMuxRouter),
	))

	muxRouter.PathPrefix("/home").Handler(negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
		negroni.HandlerFunc(checkUser),
		negroni.Wrap(webMuxRouter),
	))

	muxRouter.PathPrefix("/api").Handler(negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
		negroni.HandlerFunc(checkUser),
		negroni.Wrap(apiMuxRouter),
	))


	defineRoutes()
	setupRender()
	setupCookieStore()
	registerBinaryData()

	//init mailer
	tags := []string{
		"signup",
		"password_reset",
		"order",
	}
	err := mailer.InitializeTemplates(tags)
	if err != nil {

		log.Println("Unable to initialize mailer package,", err)

	}

	dbConn = dbConnection{}
	dbConn.initDB()

}

func getLocalIp(ev *envVariables){
	addrs, err := net.InterfaceAddrs()

	if err != nil {
					fmt.Println(err)
	}

	var currentIP string

	for _, address := range addrs {
					if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
									if ipnet.IP.To4() != nil {
													currentIP = ipnet.IP.String()
									}
					}
	}

	fmt.Println("\nYou can now view this site in the browser")
	fmt.Print("Local:               http://localhost:",ev.appPort,"\n")
	fmt.Print("On Your Network:     http://",currentIP,":",ev.appPort,"\n\n")
}

package app_config

import "os"

var PORT = ":8080"
var STATIC_DIR = "./public"
var STATIC_ROUTE = "/public"

func InitialAppConig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticRoute := os.Getenv("STATIC_ROUTE")
	if staticRoute != "" {
		STATIC_ROUTE = staticRoute
	}

	staticDir := os.Getenv("STATIC_DIR")
	if staticDir != "" {
		STATIC_DIR = staticDir
	}

}

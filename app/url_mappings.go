package app

import ("github.com/nubesFilius/bselling-go-users-api/controllers")

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
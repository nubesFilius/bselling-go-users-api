package app

import ("github.com/nubesFilius/bselling-go-users-api/controllers/users"
"github.com/nubesFilius/bselling-go-users-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
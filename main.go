package main

import server "dawnseek.com/gin-starter/server"

// @title Gin Swagger Example API
// @version 1.0
// @description This is a Gin Starter template for API server development
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http https
func main() {
	server.Start()
}

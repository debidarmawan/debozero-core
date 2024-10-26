package main

import (
	"strconv"
	"sync"

	"debozero-core/config"
	"debozero-core/constants"
	"debozero-core/server"
)

//	@title			DeboZero Core Service
//	@version		1.0
//	@description	This is an API documentation of DeboZero Core Backend Service
//	@contact.name	DeboZero Tech Team
//	@contact.url
//	@contact.email	debidarmawan1998@gmail.com

//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-API-KEY
func main() {
	config.Init()

	maxPool, _ := strconv.Atoi(config.GetEnv(constants.DbMaxPool))
	db := config.ConnectDatabase(maxPool)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()
		server.ServeHTTP(db)
	}()

	waitGroup.Wait()
}

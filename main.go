package main

import (
	"strconv"
	"sync"

	"github.com/debidarmawan/debozero-core/config"
	"github.com/debidarmawan/debozero-core/constants"
	"github.com/debidarmawan/debozero-core/server"
)

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

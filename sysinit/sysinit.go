package sysinit

import (
	"time"
	cache "github.com/patrickmn/go-cache"
	"xcms/utils"
)

func init() {
	//init cache
	utils.Cache = cache.New(60*time.Minute, 120*time.Minute)

	//init db
	initDB()
}
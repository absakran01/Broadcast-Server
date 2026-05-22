package server

import (
	"log"
	"time"

	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/fiber/v2"
)

func Start(host string, port string) {
	cache = ttlcache.NewCache()
	cache.SetTTL(ttl)

	serverApp := fiber.New()
	Routes(serverApp)

	go printCacheStatsEvery5Seconds()

	if err := serverApp.Listen(host + ":" + port); err != nil {
		panic(err)
	}

}


func printCacheStatsEvery5Seconds() {
	for {
		time.Sleep(5 * time.Second)
		if cache != nil {
			log.Printf("Cache Stats - Item Count: %d", cache.Count())
		}
	}
}
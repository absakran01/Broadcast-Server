package server

import (
	"github.com/ReneKroon/ttlcache"
	"github.com/gofiber/fiber/v2"
)

func Start(host string, port string) {
	cache = ttlcache.NewCache()
	cache.SetTTL(ttl)

	serverApp := fiber.New()
	Routes(serverApp)

	if err := serverApp.Listen(host + ":" + port); err != nil {
		panic(err)
	}

}



package redis

import (
	"log"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/axelboberg/lnkshrtnr/internal/random"
)

var ctx = context.Background()
var rdb *redis.Client

func Setup () {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func Set (key string, val string) {
	if rdb == nil {
		Setup()
	}
	rdb.Set(ctx, key, val, 0)
}

func SetRandom (val string) string {
	if rdb == nil {
		Setup()
	}

	key := random.String62(16)
	exists := rdb.Exists(ctx, key)

	if exists.Val() == 1 {
		log.Println("[Redis] Generated key exists, trying again")
		return SetRandom(val)
	}

	Set(key, val)
	return key
}

func Get (key string) (val string, ok bool) {
	if rdb == nil {
		Setup()
	}
	
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return val, true
}
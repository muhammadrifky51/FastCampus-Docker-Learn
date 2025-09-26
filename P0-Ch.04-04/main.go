package main

import(
	"fmt"
	"net/http"
	"context"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

var ctx = context.Background()

func Handler(w http.ResponseWriter, r *http.Request) {
	val, err:= client.Get(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Counter: %s", val)
}

func addCounter(w http.ResponseWriter, r *http.Request) {
	_,err := client.Incr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
}

func resetCounter(w http.ResponseWriter, r *http.Request) {
	err := client.Set(ctx, "counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
}

func SetUpRedisClient() {
	fmt.Println("Setting up Redis client...")
	client = redis.NewClient(&redis.Options{
		Addr: "redis-server:6379",
		Password: "", // no password set
		DB: 0, // use default DB
	})
}

func SetKey(key string, value int) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	SetUpRedisClient()
	SetKey("counter", 1)

	http.HandleFunc("/", Handler)
	http.HandleFunc("/add", addCounter)
	http.HandleFunc("/reset", resetCounter)
	http.ListenAndServe(":8080", nil)
}
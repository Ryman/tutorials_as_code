package main

import (
	"fmt"
	"log"
	"time"

	"github.com/muesli/cache2go"
)

// Keys & values in cache2go can be off arbitrary types, e.g. a struct.
type myStruct struct {
	text     string
	moreData []byte
}

func allKeys(cache *cache2go.CacheTable) {
	keyIndex := 0
	keyCount := cache.Count()
	keyList := make([]string, keyCount)
	cache.Foreach(func(key interface{}, item *cache2go.CacheItem) {
		keyList[keyIndex] = fmt.Sprintf("%q", key)
		keyIndex += 1
		log.Println(key, " : ", item.Key(), " ~ ", item.Data(), len(keyList))
	})
	log.Println(keyList)
}

func main() {
	// Accessing a new cache table for the first time will create it.
	cache := cache2go.Cache("myCache")

	// We will put a new item in the cache. It will expire after
	// not being accessed via Value(key) for more than 5 seconds.
	val := myStruct{"This is a test!", []byte{}}
	cache.Add("someKey", 5*time.Second, &val)
	cache.Add("otherKey", 5*time.Second, "what")
	cache.Add("anotherKey", 5*time.Second, "say what")

	allKeys(cache)

	// Let's retrieve the item from the cache.
	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache.
	time.Sleep(6 * time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires.
	cache.Add("someKey", 0, &val)

	// cache2go supports a few handy callbacks and loading mechanisms.
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	// And wipe the entire cache table.
	cache.Flush()
}
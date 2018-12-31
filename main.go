package main

import (
	"cache"
	"fmt"
)

func main (){
	key := "test"
	var val = cache.Get(key)
	fmt.Println(val)
}
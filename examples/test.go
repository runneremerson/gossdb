package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/*
		gossdb.AuthPassword = ""
		gossdb.Encoding = true
		err := ssdb.Start()
		if err != nil {
			log.Println(err)
			return
		}
		client, err := ssdb.Client()
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer client.Close()
		client.Set("a", "hello1")
		client.Set("b", "hello2")
		client.Set("keys", "hello")
		v, err := client.Hget("set", "key")
		log.Println(v, err)
	*/
	test3()
}

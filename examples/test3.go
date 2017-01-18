package main

import (
	"fmt"
	"github.com/runneremerson/gossdb"
)

func test3() {
	ip := "127.0.0.1"
	port := 8888

	pool, err := gossdb.NewPool(&gossdb.Config{
		Host: ip,
		Port: port,
	})
	if err != nil {
		fmt.Errorf("error new pool %v", err)
		return
	}
	gossdb.Encoding = true

	c, err := pool.NewClient()
	if err != nil {
		fmt.Errorf("new client err=%v", err)
		return
	}

	c.Zset("booklist_order_by_price", "id1", 1)
	c.Zset("booklist_order_by_price", "id2", 2)
	c.Zset("booklist_order_by_price", "id3", 1)
	c.Zset("booklist_order_by_price", "id4", 3)
	c.Zset("booklist_order_by_price", "id5", 51)
	c.Zset("booklist_order_by_price", "id6", 18)
	c.Zset("booklist_order_by_price", "id7", 15)
	c.Zset("booklist_order_by_price", "id8", 22)
	c.Zset("booklist_order_by_price", "id9", 11)
	c.Zset("booklist_order_by_price", "id10", 10)

	fmt.Println("will zscan---------------------------------")
	keys, scores, err := c.Zscan("booklist_order_by_price", "", "", "", 1000)
	fmt.Printf("keys = %v \nscores = %v\n", keys, scores)
	fmt.Printf("keys len = %v \nscores  len = %v\n", len(keys), len(scores))
	for i, k := range keys {
		fmt.Printf("%v : %v - %v\n", i, k, scores[i])
	}
	fmt.Println("will zrscan---------------------------------")
	keys, scores, err = c.Zrscan("booklist_order_by_price", "", "", "", 1000)
	fmt.Printf("keys = %v \nscores = %v\n", keys, scores)
	fmt.Printf("keys len = %v \nscores  len = %v\n", len(keys), len(scores))
	for i, k := range keys {
		fmt.Printf("%v : %v - %v\n", i, k, scores[i])
	}

	c.Zset("zsettest", "zsettest_count", 0)
	c.Zincr("zsettest", "zsettest_count", 1)
	c.Zincr("zsettest", "zsettest_count", 1)
	count, err := c.Zget("zsettest", "zsettest_count")
	if err != nil {
		fmt.Println("zsettest_count error = ", err.Error())
	} else {
		fmt.Printf("zsettest_count = %v\n", count)
	}
}

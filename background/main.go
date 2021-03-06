package main

//TODO add a limitation on the memory of the redis (usage with info), relative to the total
//TODO fix recalculation each launch

import (
	"github.com/go-redis/redis"
	"log"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	Cash *redis.Client
)

type Tribonachi struct {
	Number string
	N      string
}

func init() {
	Cash = redis.NewClient(&redis.Options{
		Addr:     "db:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := Cash.Ping().Err()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	wg.Add(2)
	everySecond := time.NewTicker(1 * time.Second)

	checkLoad(everySecond)

	done := make(chan bool)
	sendRedis := make(chan Tribonachi)

	go trib(sendRedis, done, 0)
	go worker(sendRedis, everySecond, done)

	wg.Wait()

}

func trib(sendRedis chan Tribonachi, done chan bool, int int) {
	a := big.NewInt(0)
	b := big.NewInt(0)
	c := big.NewInt(1)
	tmp := big.NewInt(0)
	var i uint = 0
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(109999), nil)
	for a.Cmp(&limit) < 0 {
		tmp = a.Add(a, b)
		a.Add(c, tmp)
		a, b, c = b, c, a
		sendRedis <- makeTrib(a, i)
	}
	done <- true
	wg.Done()
}

func makeTrib(a *big.Int, i uint) Tribonachi {
	var t Tribonachi
	t.Number = a.String()
	t.N = strconv.FormatUint(uint64(i), 10)
	return t
}

func checkLoad(ticker *time.Ticker) {
	for {
		<-ticker.C
		str, err := Cash.Info("persistence").Result()
		if err != nil {
			log.Println(err)
		}
		if !strings.Contains(str, "loading:1") {
			return
		}
	}
}

func worker(sendRedis chan Tribonachi, ticker *time.Ticker, done chan bool) {
	var arr []Tribonachi
L:
	for {
		select {
		case a := <-sendRedis:
			arr = append(arr, a)
		case <-ticker.C:
			if len(arr) == 0 {
				continue
			}
			for _, a := range arr {
				err := Cash.Set(a.N, a.Number, 0).Err()
				if err != nil {
					log.Println(err)
					continue L
				}
			}
			arr = nil
		case <-done:
			for _, a := range arr {
				err := Cash.Set(a.N, a.Number, 0).Err()
				if err != nil {
					log.Println(err)
				}
			}
			wg.Done()
		}

	}
}

package main

import (
	"os"
	"sync"

	reserveNumber "github.com/devpablocristo/growuphr/reserve-number/api"
)

const defaultReserveNumberPort = "8080"

//const defautUserPort = "8081"

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	reserveNumberPort := os.Getenv("RESERVE-NUMBER_PORT")
	if reserveNumberPort == "" {
		reserveNumberPort = defaultReserveNumberPort
	}

	// userPort := os.Getenv("reserve-number_PORT")
	// if userPort == "" {
	// 	userPort = defaultreserve-numberPort
	// }

	wg.Add(2)
	go reserveNumber.LoadData(&wg)
	go reserveNumber.StartApi(&wg, reserveNumberPort)
	// go user.LoadData(&wg)
	// go user.StartApi(&wg, userPort)

	wg.Wait()
}

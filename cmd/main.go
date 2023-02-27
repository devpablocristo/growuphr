package main

import (
	"os"
	"sync"

	numberManager "github.com/devpablocristo/growuphr/number-manager/api"
)

const (
	defaultReserveNumberPort = "8080"
	port1                    = "8081"
	port2                    = "8082"
)

//const defautUserPort = "8081"

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	reserveNumberPort := os.Getenv("number-manager_PORT")
	if reserveNumberPort == "" {
		reserveNumberPort = defaultReserveNumberPort
	}

	// userPort := os.Getenv("number-manager_PORT")
	// if userPort == "" {
	// 	userPort = defaultnumber-managerPort
	// }

	wg.Add(3)
	//go reserveNumber.LoadData(&wg)
	go numberManager.StartApi(&wg, reserveNumberPort)
	go numberManager.StartApi(&wg, port1)
	go numberManager.StartApi(&wg, port2)

	// go user.LoadData(&wg)
	// go user.StartApi(&wg, userPort)

	wg.Wait()
}

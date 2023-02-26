package api

import (
	"sync"

	application "github.com/devpablocristo/growuphr/reserve-number/application"
	mapdb "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driven-adapter/repository/mapdb"
	handler "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/handler"
	//numsrv "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/number-service"
	//usrsrv "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/user-service"
)

// func LoadData(wg *sync.WaitGroup) {
// 	numbers := []domain.Number{
// 		{UUID: "1", Number: 12},
// 		{UUID: "2", Number: 134},
// 		{UUID: "3", Number: 11242},
// 	}

// 	bs, err := json.Marshal(numbers)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(string(bs))
// }

func StartApi(wg *sync.WaitGroup, port string) {
	//defer wg.Done()

	// db := postgres.ConnectToDB()
	// defer db.Close()

	mdb := mapdb.NewMapDB()
	// uss := usrsrv.NewUserService(mdb)
	// nus := numsrv.NewNumberService(mdb)
	rns := application.NewReserveNumberService(mdb)
	han := handler.NewHandler(rns)
	rou := Router(han)

	HttpServer(port, rou)
}

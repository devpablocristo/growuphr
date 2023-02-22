package api

import (
	"encoding/json"
	"fmt"
	"sync"

	application "github.com/devpablocristo/growuphr/reserve-number/application"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
	mapdb "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driven-adapter/repository/mapdb"
	chihandler "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/handler/chi"
)

func LoadData(wg *sync.WaitGroup) {
	n1 := domain.Number{
		UUID:   "1",
		Number: 12,
	}

	n2 := domain.Number{
		UUID:   "1",
		Number: 12,
	}

	n3 := domain.Number{
		UUID:   "1",
		Number: 12,
	}

	numbers := []domain.Number{
		n1,
		n2,
		n3,
	}

	bs, err := json.Marshal(numbers)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bs))
}

func StartApi(wg *sync.WaitGroup, port string) {
	//defer wg.Done()

	// db := postgres.ConnectToDB()
	// defer db.Close()

	mdb := mapdb.NewMapDB()
	nse := application.NewNumberService(mdb)
	han := chihandler.NewChiHandler(nse)
	rou := ChiRouter(han)

	HttpServer(port, rou)
}

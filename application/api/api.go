package api

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-foward/abroad/application/agency/controller"
	"github.com/labstack/echo"
)

// API entrypoint
type API struct {
	*echo.Echo
	AgencyController controller.AgencyControl
}

// NewAPI creates a new application
func NewAPI(agencyController controller.AgencyControl) *API {
	return &API{
		Echo:             echo.New(),
		AgencyController: agencyController,
	}
}

func (api API) printBanner() {
	api.Echo.HideBanner = true
	bytes, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

// Run start http server
func (api API) Run() {
	api.enableMiddlewares()
	api.enableAgencyHandlers()

	api.printBanner()
	log.Fatal(api.Start(":8080"))
}

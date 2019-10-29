package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/my-ui/internal/app"
)

func main() {
	config, err := app.LoadConfig("")
	if err != nil {
		return
	}
	app.InitEndpoints(config)

	err = app.UpdateMapName()
	if err != nil {
		return
	}

	errs := make(chan error, 2)
	listenForInterrupt(errs)
	startHTTPServer(errs, app.Configuration.Server.Port)
	// readings, err := app.GetReadingByDeviceInTimeRange("Random-Boolean-Device", 1471806984000, 1572062272000)
	// if err != nil {
	// 	return
	// }
	// log.Println("readings:\n", string(readings))

	// vd, err := app.GetValueDescriptorByName("RandomValue_Bool")
	// if err != nil {
	// 	return
	// }
	// log.Println("valuedescriptor:\n", string(vd))

	// vds, err := app.GetValueDescriptorByDeviceName("Random-Boolean-Device")
	// if err != nil {
	// 	return
	// }
	// log.Println("list valuedescriptor:\n", string(vds))

	<-errs
	os.Exit(0)
}

func listenForInterrupt(errChan chan error) {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}

func startHTTPServer(errChan chan error, port int) {
	go func() {
		log.Println("EdgeX UI Server Listen At :" + strconv.Itoa(port))
		r := app.LoadRestRoutes()
		errChan <- http.ListenAndServe(":"+strconv.Itoa(port), r)
	}()
}

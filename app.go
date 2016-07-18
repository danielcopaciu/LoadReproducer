package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"net/http"
	"log"
	"io"
	"strconv"
)

type Config struct {
	port         int
	nativeIdsUrl string
	endpointUrls []string
}

const logPattern = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC

var infoLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func initLogs(infoHandle io.Writer, warnHandle io.Writer, errorHandle io.Writer) {
	infoLogger = log.New(infoHandle, "INFO  - ", logPattern)
	warnLogger = log.New(warnHandle, "WARN  - ", logPattern)
	errorLogger = log.New(errorHandle, "ERROR - ", logPattern)
}

func main() {
	app := cli.App("", "")

	port := app.Int(cli.IntOpt{
		Name:   "port",
		Value:  8080,
		Desc:   "application port",
		EnvVar: "PORT",
	})

	//nativeIdsUrl := app.String(cli.StringOpt{
	//	Name:   "nativeIdsUrl",
	//	Value:  "",
	//	Desc:   "ids endpoint address",
	//	EnvVar: "NATIVE_IDS_URL",
	//})

	endpointUrls := app.Strings(cli.StringsOpt{
		Name:   "endpointUrls",
		Value:  nil,
		Desc:   "endpoints urls list",
		EnvVar: "ENDPOINT_URLS",
	})

	initLogs(os.Stdout, os.Stdout, os.Stderr)

	app.Action = func() {
		config := Config{*port, "", *endpointUrls}

		for _, url := range config.endpointUrls {
			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			status := response.StatusCode
			stringsStatus := strconv.Itoa(status);
			switch status {
			case 500:
				errorLogger.Println("Returned status is " + stringsStatus)
			case 404:
				warnLogger.Println("Returned status is " + stringsStatus)
			default:
				infoLogger.Println("Returned status is " + stringsStatus)
			}

		}

	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("[%v]", err)
	}
}



package main

import (
	"fmt"
	"flag"
	//	"errors"
	"log"
	"os"
)
import(
	"github.com/libamperix-go/amperix"
)

type cmdlineArgs struct {
	url      *string
	port     *string
	sensorid *string
	token    *string
	debug    *bool
	verbose  *bool

}

var config       = amperix.NewConfig()
var serverconfig = amperix.NewServerConfig()

var args = cmdlineArgs{
	url:      flag.String("url",      serverconfig.BaseUrl, "API-Url"),
	port:     flag.String("port",     serverconfig.Port,    "API-Port"),
	sensorid: flag.String("sensorid", "", "Sensorid"),
	token:    flag.String("token",    "", "Token"),
	debug:    flag.Bool("debug",      false, "Debugmode"),
	verbose:  flag.Bool("verbose",    false, "Verbose"),
}


func init() {
	flag.Parse()

	bailIfMissing := func(value *string, flag string) {
		if *value == "" {
			log.Fatal(flag + " missing")
			os.Exit(1)
		}
	}

	bailIfMissing(args.sensorid, "-sensorid")
	bailIfMissing(args.token, "-token")

	serverconfig.SetBaseUrl(*args.url)
	serverconfig.SetPort(*args.port)

	config.SetSensorId(*args.sensorid)
	config.SetToken(*args.token)
	config.SetDebug(*args.debug)
	config.SetVerbose(*args.verbose)
}

func main() {


	ws := amperix.NewWebService(config)
	ret, err := ws.GetValues()
	if(err != nil) {
		if(config.GetDebug()) {	fmt.Printf("N: url: '%s'\n", ret) }
	} else {
		log.Fatal("Query failed")
		log.Fatal(err)
	}
}


// https://api.mysmartgrid.de:8443/sensor/
//  1c98d7d063e2408a91fabe1e02e0f299
// ?start=1427970000
// &end=1427980000
// &resolution=15min
// &unit=eurperyear
// curl -H "X-Version: 1.0" -H "Accept: application/json" -H "X-Token: f7c0126bb0c8c0a9eacc8eb1166bc3f9"  -k "https://api.mysmartgrid.de:8443/sensor/1c98d7d063e2408a91fabe1e02e0f299?start=1427970000&end=1427980000&end=1427980000&resolution=15min&unit=watt"

/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

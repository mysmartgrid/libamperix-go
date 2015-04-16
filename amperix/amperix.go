package amperix

import (
	"fmt"
	"io/ioutil"
//	"io"
//	"strings"
	"log"
	"net/http"
	"crypto/tls"
	"time"
"math"
	"strconv"
	)

	import "encoding/json"

type nanfloat float64

func (p *nanfloat) UnmarshalJSON(data []byte) error {
  if ( string(data) == "\"-nan\"" ) {
    *p = nanfloat(math.NaN())
		return nil
	} 

	f, err := strconv.ParseFloat(string(data), 64)
	*p = nanfloat(f)
	return err
}


type Measurement struct {
	Time  time.Time
	Value float64
}

func (p *Measurement) UnmarshalJSON(data []byte) error {
	var arr [2]nanfloat
			
	//fmt.Printf("R: '%s'\n\n",  data)
	if err := json.Unmarshal(data, &arr); err != nil {
		//fmt.Printf("Unmarshal - error %s\n", err)
		//fmt.Printf("\t %f\n", arr[0])
		return err
	}

	ms := int64(arr[0])
	//p.Time = time.Unix(ms/1000, (ms%1000)*1e6)
	p.Time  = time.Unix(ms, 0)
	p.Value = float64(arr[1])
	return nil
}


type WebService struct {
	
	// Configuration
	config *Config

	// http connector
  client *http.Client
}

func NewWebService(cfg *Config)(*WebService) {
	tlscfg := &tls.Config{InsecureSkipVerify: true}
	tr := &http.Transport{
    TLSClientConfig: tlscfg,
	}

	ws := &WebService{config: cfg, 
										client: &http.Client{Transport: tr },
	}

	return ws
}

func (ws *WebService) GetValues() ([]Measurement) {
	t := time.Now()
	
//	start := 1427970000
			s := ( t.Unix() / 60 ) * 60 - ( 15 * 60)

	start := s
	end := 1427980000
		var url string
		url = fmt.Sprintf("%s/sensor/%s?start=%d&end=%d&resolution=%s&unit=%s",
										 ws.config.GetBaseurl(), ws.config.GetSensorId(),
										 start, end,
										 ws.config.GetInterval(), ws.config.GetUnit())
		url = fmt.Sprintf("%s/sensor/%s?interval=%s&unit=%s",
										 ws.config.GetBaseurl(), ws.config.GetSensorId(),
										 ws.config.GetInterval(), ws.config.GetUnit())

		//res, err := http.Get(url)
	fmt.Printf("N: url: '%s'\n", url)
	var timeseries []Measurement
		err := ws.run_query(url, &timeseries)
	if err != nil {
		log.Fatal(err)
	}

	return timeseries
}

func (ws *WebService) run_query(buffer string, timeseries *[]Measurement) error {

	req, err := http.NewRequest("GET", buffer, nil)
	if err != nil {
		fmt.Printf("NewRequest error\n")
		log.Fatal(err)
		return err
	}
  req.Header.Add("X-Version", "1.0")
  req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Token", ws.config.GetToken())

	ws.client.Do(req)
	resp, err := ws.client.Do(req)
		//fmt.Printf("client do done\n")

	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
		
	if err = json.Unmarshal(robots, &timeseries); err != nil {
		return err
	}
	
	if(ws.config.GetVerbose()) {
		fmt.Printf("'%+v'\n\n", timeseries)
	}
	return nil
}

package amperix

import (
	"bytes"
	"crypto/hmac"
	//        "crypto/rand"
	//        "crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"hash"
	"fmt"
	"io/ioutil"
	//	"io"
	//"strings"
	"log"
	"net/http"
	"time"
	//"math"
	//	"strconv"
)

import "encoding/json"

type WebService struct {
	
	// Server Configuration
	serverconfig *ServerConfig

	// User Configuration
	config *Config

	// http connector
	client *http.Client
}

func NewWebService(cfg *Config)(*WebService) {
	serverconfig := NewServerConfig()

	tlscfg := &tls.Config{InsecureSkipVerify: serverconfig.InsecureSkipVerify}
	tr := &http.Transport{
		TLSClientConfig: tlscfg,
	}

	ws := &WebService{config: cfg, 
		serverconfig: serverconfig,
		client: &http.Client{Transport: tr },
	}

	return ws
}

func createMac(h hash.Hash, value []byte) []byte {
	h.Write(value)
	return h.Sum(nil)
}
func (ws *WebService) digest_message(data []byte, key string)(string, error) {
	hashKey := ([]byte)(key)
	var h func() hash.Hash
	mac := createMac(hmac.New(h, hashKey), data[:len(data)-1])


	return (string)(mac), nil
}

func (ws *WebService) GetValues() ([]Measurement, error) {
	t := time.Now()
	
	s := ( t.Unix() / 60 ) * 60 - ( 15 * 60)
	start := s
	end   := t.Unix()

	var url string
	url = fmt.Sprintf("%s:%s/sensor/%s?start=%d&end=%d&resolution=%s&unit=%s",
		ws.serverconfig.BaseUrl, ws.serverconfig.Port, ws.config.GetSensorId(),
		start, end,
		ws.config.GetInterval(), ws.config.GetUnit())
	url = fmt.Sprintf("%s:%s/sensor/%s?interval=%s&unit=%s",
		ws.serverconfig.BaseUrl, ws.serverconfig.Port, ws.config.GetSensorId(),
		ws.config.GetInterval(), ws.config.GetUnit())

	fmt.Printf("N: url: '%s'\n", url)
	var timeseries []Measurement
	err := ws.run_query(url, &timeseries)
	if err != nil {
		return nil,err
	}

	return timeseries, nil
}

func (ws *WebService) PushValues(timeseries []Measurement) error {
	var buf bytes.Buffer
	writtenAny := false
	buf.WriteString(`{"measurements":[`)
	for _, value := range timeseries {
		if writtenAny {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("[%v,%v]", value.Time.Unix(), value.Value))
		writtenAny = true
	}
	buf.WriteString(`]}`)

	var url string
	url = fmt.Sprintf("%s:%s/sensor/%s",
		ws.serverconfig.BaseUrl, ws.serverconfig.Port,ws.config.GetSensorId())

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return err
	}
	req.Header.Add("X-Version", "1.0")
	req.Header.Add("Accept", "application/json,text/html")
	req.Header.Add("Content-type", "application/json")

	var digest []byte
	digest = createDigest(buf.Bytes(), ([]byte)(ws.config.Key))
	req.Header.Add("X-Digest", hex.EncodeToString(digest))

	resp, err := ws.client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Result Status: %d\n", resp.StatusCode)
	fmt.Printf("Result Status: %s\n", GetError(resp.StatusCode))
	

	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Print("Result:\n")
	fmt.Printf(string(robots))
	if err != nil {
		//log.Fatal(err)
		return err
	}
	
	return nil
}

func (ws *WebService) run_query(buffer string, timeseries *[]Measurement) error {

	req, err := http.NewRequest("GET", buffer, nil)
	if err != nil {
		fmt.Printf("NewRequest error\n")
		return err
	}
	req.Header.Add("User-Agent", "libamperix-go")
	req.Header.Add("X-Version", "1.0")
	req.Header.Add("Accept", "application/json,text/html")
	req.Header.Add("X-Token", ws.config.GetToken())
	req.Header.Add("Content-type", "application/json")

	ws.client.Do(req)
	resp, err := ws.client.Do(req)
	//fmt.Printf("client do done\n")

	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//log.Fatal(err)
		return err
	}
	
	fmt.Printf(string(robots))
	if err = json.Unmarshal(robots, &timeseries); err != nil {
		log.Fatal(err)
		return err
	}
	
	if(ws.config.GetVerbose()) {
		fmt.Printf("'%+v'\n\n", timeseries)
	}
	return nil
}
/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

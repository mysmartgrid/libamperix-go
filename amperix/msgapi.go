package amperix

import (
	"fmt"
	"time"
	"math"
	"crypto/hmac"
	//"crypto/rand"
	//"crypto/sha256"
	"crypto/sha1"
	//"encoding/hex"
	"encoding/json"
	//"encoding/base64"
	//"strconv"
)

type wsClientDevice struct {}

type v1DeviceCmdUpdateArgs struct {
	Values map[string][]Measurement `json:"values"`
}

type Measurement struct {
	Time  time.Time
	Value float64
}

func (p *Measurement) UnmarshalJSON(data []byte) error {
	var arr [2]float64
	
	fmt.Printf("R: '%s'\n\n",  data)
	if err := json.Unmarshal(data, &arr); err != nil {
		//fmt.Printf("Unmarshal - error %s\n", err)
		//fmt.Printf("\t %f\n", arr[0])
		ms := int64(arr[0])
		p.Time  = time.Unix(ms, 0)
		p.Value=math.NaN()
		return nil
	}

	ms := int64(arr[0])
	//p.Time = time.Unix(ms/1000, (ms%1000)*1e6)
	p.Time  = time.Unix(ms, 0)
	p.Value=arr[1]
	return nil
}

func (p *Measurement) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]float64{float64(jsTime(p.Time)),
		p.Value})
	//return json.Marshal([2](float64){jsonTime(p.Time),
	//	jsonValue(p.Value)})
}

func (p *Measurement) Print() {
	fmt.Printf("%d : %f\n", p.Time.Unix(), p.Value)
}

func jsonTime(time time.Time) ([]byte) {
	s := fmt.Sprintf("%d", time.Unix())
	return ([]byte)(s)
}
func jsonValue(value float64) []byte {
	s := fmt.Sprintf("%f", value)
	return ([]byte)(s)
}

func jsTime(time time.Time) int64 {
	//return 1000*time.Unix() + int64(time.Nanosecond()/1e6)
	return time.Unix()///1000000
}

//type Hash interface {
//        Hash
//        Sum() uint64
//}


/*
func (c *wsClientDevice) Update(values map[string][]Measurement) error {
        cmd := v1MessageOut{
                Command: "update",
                Args:    v1DeviceCmdUpdateArgs{values},
        }

        return c.executeCommand(&cmd)
}

func (api *wsUserAPI) sendUpdate(values map[string]map[string][]Measurement) error {
        return api.dispatch.WriteJSON(v1MessageOut{Command: "update", Args: values})
}
*/
func createDigest(bv []byte, key []byte) ([]byte) {
	mac := hmac.New(sha1.New, ([]byte)(key))
	mac.Write(bv[:])
	return mac.Sum(nil)
}
/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

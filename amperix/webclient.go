package amperix

type MsgApi_error struct {
	code int
	text string
}

var msapi_errors = [](MsgApi_error){
	{ 200, "OK"},
	{ 400, "Bad request"},
	{ 401, "Unauthorized"},
	{ 403, "Forbidden"},
	{ 404, "Not Found" },
	{ 405, "Method Not Allowed" },
	{ 406, "Not acceptable" },
	{ 470, "Invalid Timestamp" },
	{ 471, "Invalid Unit" },
	{ 472, "Invalid Measurement" },
	{ 473, "Invalid Object Type" },
	{ 474, "Invalid Object Id" },
	{ 475, "Invalid Key" },
	{ 476, "Invalid Time Period" },
	{ 477, "Invalid Event" },
	{ 478, "Unupgradable Firmware" },
	{ 479, "Invalid Sensor External Id" },
	{ 480, "Invalid Characters" },
	{ 500, "Internal Server Error" },
	{ 501, "Not Implemented" },
}

func GetError(errno int) (string) {
	for _, err := range msapi_errors {
		if(err.code==errno) {
			return err.text
		}
	}
	return "Invalid errno"
}

/*
import {
	"net/http"
        "crypto/hmac"
        "crypto/rand"
        "crypto/sha256"
        "crypto/tls"
}

type Secret struct {
}
*/



/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

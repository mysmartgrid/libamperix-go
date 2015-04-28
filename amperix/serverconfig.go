package amperix

type ServerConfig struct {
	BaseUrl		 string
	Port             string
	InsecureSkipVerify bool
	Version		 string
	Debug		 bool
	Verbose		 bool
}

func NewServerConfig() (*ServerConfig) {
	cfg := &ServerConfig{
		BaseUrl: "https://dev3-api.mysmartgrid.de",
		Port:    "8443", 
		InsecureSkipVerify: true,
	}
	return cfg
}

func (c* ServerConfig) SetBaseUrl(BaseUrl string) {
	c.BaseUrl=BaseUrl
}

func (c* ServerConfig) SetPort(Port string) {
	c.Port=Port
}

/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

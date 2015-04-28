package amperix

type Config struct {
	SensorId	 string
	Token		 string
	Key 		 string
	Unit		 string
	Interval	 string
	FormatterType    string
	FilterType	 string
	OutFilename	 string
	Version		 string
	Debug		 bool
	Verbose		 bool
}

func NewConfig() (*Config) {
	cfg := &Config{
		Unit: "watt",
		Interval: "15min",
		
	}
	return cfg
}

func (c*Config) GetSensorId() string {
	return c.SensorId
}

func (c* Config) GetToken() string {
	return c.Token
}

func (c* Config) GetUnit() string {
	return c.Unit
}

func (c* Config) GetInterval() string {
	return c.Interval
}

func (c* Config) GetVersion() string {
	return c.Version
}

func (c* Config) GetDebug() bool {
	return c.Debug
}

func (c* Config) GetVerbose() bool {
	return c.Verbose
}


func (c*Config) SetSensorId(SensorId string) {
	c.SensorId=SensorId
}

func (c* Config) SetToken(Token string) {
	c.Token=Token
}
func (c* Config) SetKey(Key string) {
	c.Key=Key
}

func (c* Config) SetUnit(Unit string) {
	c.Unit=Unit
}

func (c* Config) SetInterval(Interval string) {
	c.Interval=Interval
}

func (c* Config) SetVersion(Version string) {
	c.Version=Version
}

func (c* Config) SetDebug(Debug bool) {
	c.Debug=Debug
}

func (c* Config) SetVerbose(Verbose bool) {
	c.Verbose=Verbose
}
/*
 * Local variables:
 *  tab-width: 2
 *  c-indent-level: 2
 *  c-basic-offset: 2
 * End:
 */

package config

var (
	Config = make(map[string]map[string]info)
)

type info struct {
	Key string
	I   string
}

func NewReConfig() (map[string]map[string]info, error) {
	return Config, LoadDefaultConf(&Config, "frame", "overwrite")
}

//func (c *info) getClusterKey() string {
//	return c.ClusterKey
//}
//func (c *info) setClusterKey(key string) {
//	c.ClusterKey = key
//}

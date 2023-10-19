package config

var (
	BizRouteConfig = make(map[string]map[string]ClusterInfo)
)

type ClusterInfo struct {
	ClusterKey string
	Index      string
}

func NewReConfig() (map[string]map[string]ClusterInfo, error) {
	return BizRouteConfig, LoadDefaultConf(&BizRouteConfig, "frame", "overwrite")
}

//func (c *ClusterInfo) getClusterKey() string {
//	return c.ClusterKey
//}
//func (c *ClusterInfo) setClusterKey(key string) {
//	c.ClusterKey = key
//}

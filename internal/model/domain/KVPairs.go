package domain

// KVPairs 路由查询参数键值对
type KVPairs struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

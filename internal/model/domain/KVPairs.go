package domain

// KVPairs 路由查询参数键值对，这里假设 KVPairs 是一个键值对结构
type KVPairs struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

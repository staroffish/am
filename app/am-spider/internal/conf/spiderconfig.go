package conf

import "sync"

type SpiderConfig struct {
	spiderParameter *SpiderParameter
	rwLock          *sync.RWMutex
}

func NewSpiderConfig(spiderConfig *SpiderParameter) *SpiderConfig {
	return &SpiderConfig{
		spiderParameter: spiderConfig,
		rwLock:          &sync.RWMutex{},
	}
}

func (s *SpiderConfig) GetConfig() *SpiderParameter {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.spiderParameter
}

func (s *SpiderConfig) SetConfig(config *SpiderParameter) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	s.copySpiderConfig(config)
}

func (s *SpiderConfig) copySpiderConfig(config *SpiderParameter) {
	s.spiderParameter.Method = config.Method
	s.spiderParameter.Proxy = config.Proxy
	s.spiderParameter.Url = config.Url
	s.spiderParameter.Interval = config.Interval
}

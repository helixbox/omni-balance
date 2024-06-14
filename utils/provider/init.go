package provider

import (
	"github.com/pkg/errors"
	"omni-balance/utils/configs"
	"sync"
)

type InitFunc func(conf configs.Config, noInit ...bool) (Provider, error)

var (
	providers = make(map[configs.LiquidityProviderType][]InitFunc)
	m         sync.Mutex
)

func Register(providerType configs.LiquidityProviderType, provider InitFunc) {
	m.Lock()
	defer m.Unlock()
	providers[providerType] = append(providers[providerType], provider)
}

func ListProviders() map[configs.LiquidityProviderType][]InitFunc {
	var result = make(map[configs.LiquidityProviderType][]InitFunc)
	for k, v := range providers {
		result[k] = v
	}
	return result
}

func ListProvidersByConfig(conf configs.Config) map[configs.LiquidityProviderType][]InitFunc {
	m.Lock()
	defer m.Unlock()
	var (
		providerNames = make(map[string]struct{})
		result        = make(map[configs.LiquidityProviderType][]InitFunc)
	)
	for _, v := range conf.LiquidityProviders {
		providerNames[v.LiquidityName] = struct{}{}
	}
	for providerType, providerInitFuncs := range providers {
		for index, fn := range providerInitFuncs {
			p, _ := fn(conf, true)

			if _, ok := providerNames[p.Name()]; !ok {
				continue
			}
			result[providerType] = append(result[providerType], providers[providerType][index])
		}
	}
	return result
}

func LiquidityProviderTypeAndConf(providerType configs.LiquidityProviderType, conf configs.Config) []InitFunc {
	return ListProvidersByConfig(conf)[providerType]
}

func GetProvider(providerType configs.LiquidityProviderType, name string) (InitFunc, error) {
	for _, fn := range providers[providerType] {
		p, err := fn(configs.Config{}, true)
		if err != nil {
			return nil, err
		}
		if p.Name() != name {
			continue
		}
		return fn, nil
	}
	return nil, errors.New("provider not found")
}

func InitializeBridge(providerInitFunc InitFunc, conf configs.Config, noInit ...bool) (Provider, error) {
	bridge, err := providerInitFunc(conf, noInit...)
	if err != nil {
		return nil, errors.Wrap(err, "init bridge error")
	}
	return bridge, nil
}

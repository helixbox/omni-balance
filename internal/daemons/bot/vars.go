package bot

import (
	"fmt"
	"strings"
)

type IgnoreTokens []IgnoreToken

type IgnoreToken struct {
	Name    string `json:"name"`
	Chain   string `json:"chain"`
	Address string `json:"wallet"`
}

func (i IgnoreTokens) Contains(name, chain, address string) bool {
	key := fmt.Sprintf("%s_%s_%s", name, chain, address)
	for _, token := range i {
		if !strings.EqualFold(key, fmt.Sprintf("%s_%s_%s", token.Name, token.Chain, token.Address)) {
			continue
		}
		return true
	}
	return false
}

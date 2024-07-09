package configs

import (
	"encoding/json"
	"fmt"
	"omni-balance/utils"
	"omni-balance/utils/constant"
	"omni-balance/utils/wallets"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// ProviderType liquidity providersMap type
type ProviderType string
type DbType string

const (
	// CEX centralized exchange
	CEX ProviderType = "CEX"
	// DEX decentralized exchange
	DEX ProviderType = "DEX"
	// Bridge cross-chain bridge
	Bridge ProviderType = "Bridge"
)

const (
	// MySQL mysql
	MySQL DbType = "MySQL"
	// PostgreSQL postgresql
	PostgreSQL DbType = "PostgreSQL"
	// SQLite sqlite
	SQLite DbType = "SQLite"
)

type Config struct {
	Debug bool `json:"debug" yaml:"debug" comment:"Debug mode"`

	// Chains need to be monitored
	Chains    []Chain `json:"chains" yaml:"chains" comment:"Chains"`
	chainsMap map[string]Chain

	// SourceToken used to buy other tokens
	SourceToken    []SourceToken `json:"source_token" yaml:"source_token" comment:"Source token used to buy other tokens"`
	sourceTokenMap map[string]SourceToken

	Providers    []Provider `json:"providers" yaml:"providers" comment:"providers"`
	providersMap map[ProviderType][]Provider

	Wallets []Wallet `json:"wallets" yaml:"wallets" comment:"Wallets need to rebalance"`
	wallets map[string]Wallet

	Db DbConfig `json:"db" yaml:"db" comment:"Database config"`

	TaskInterval map[string]time.Duration `json:"task_interval" yaml:"task_interval"`

	Notice Notice `json:"notice" yaml:"notice" comment:"Notice config. When rebalance success, send notice"`
}

type Notice struct {
	Type     string                 `json:"type" yaml:"type" comment:"Notice type, support: slack"`
	Config   map[string]interface{} `json:"config" yaml:"config" comment:"It depends on the notification type, slack needs ['webhook','channel']"`
	Interval time.Duration          `json:"interval" yaml:"interval" comment:"Same message send interval, minimum interval must be greater than or equal to 1 hour, default 1h"`
}

type Chain struct {
	Id           int      `json:"id" yaml:"id" comment:"Chain id"`
	Name         string   `json:"name" yaml:"name" comment:"Chain name"`
	NativeToken  string   `json:"native_token" comment:"Native token name, if not set, use the 0x0000000000000000000000000000000000000000"`
	RpcEndpoints []string `json:"rpc_endpoints" yaml:"rpc_endpoints" comment:"RPC endpoints"`
	Tokens       []Token  `json:"tokens" yaml:"tokens" comment:"Tokens"`
}

type Wallet struct {
	// BotTypes The type of monitoring, support: balance_on_chain, helix_liquidity, the default is balance_on_chain.
	// If BotTypes is not empty and tokens.BotTypes is empty, then use BotTypes.
	// If BotTypes is empty and tokens.BotTypes is not empty, then use tokens.BotTypes.
	// If multiple types are set in BotTypes, multiple bots will be created according to the types.
	BotTypes      []string      `json:"monitor_types" yaml:"botTypes" comment:"BotTypes The type of monitoring, support: balance_on_chain, helix_liquidity, the default is balance_on_chain.\n If BotTypes is not empty and tokens.BotTypes is empty, then use BotTypes.\n If BotTypes is empty and tokens.BotTypes is not empty, then use tokens.BotTypes.\n If multiple types are set in BotTypes, multiple bots will be created according to the types."`
	Address       string        `json:"address" yaml:"address" comment:"Monitoring address"`
	Operator      Operator      `json:"operator" yaml:"operator" comment:"Used to isolate the monitoring address and the operation address, preventing the leakage of the monitoring address private key. If Operator is empty, it is not enabled. If 'multi_sign_type' is not empty, 'address' is multi sign address, 'operator' is multi sign operator address."`
	MultiSignType string        `json:"multi_sign_type" yaml:"multi_sign_type" comment:"multi sign address type, support: safe. If not empty, 'address' is multi sign address, 'operator' is multi sign operator address"`
	Tokens        []WalletToken `json:"tokens" yaml:"tokens" comment:"Tokens to be monitored"`
	PrivateKey    string        `json:"private_key" yaml:"private_key" comment:"'address' private key. If 'operator' is not empty, private_key is the operator's private key"`
}

type Operator struct {
	// Address Operator address
	Address       common.Address `json:"address" yaml:"address" comment:"Operator address"`
	Operator      common.Address `json:"operator" yaml:"operator" comment:"Used to isolate the monitoring address and the operation address, preventing the leakage of the monitoring address private key. If Operator is empty, it is not enabled. If 'multi_sign_type' is not empty, 'address' is multi sign address, 'operator' is multi sign operator address"`
	PrivateKey    string         `json:"private_key" yaml:"private_key" comment:"'address' private key. If 'operator' is not empty, private_key is the operator's private key"`
	MultiSignType string         `json:"multi_sign_type" yaml:"multi_sign_type" comment:"MultiSign type, ex: safe"`
}

type WalletToken struct {
	Name      string              `json:"name" yaml:"name" comment:"Token name"`
	Amount    decimal.Decimal     `json:"amount" yaml:"amount" comment:"The number of each rebalance"`
	Threshold decimal.Decimal     `json:"threshold" yaml:"threshold" comment:"Threshold when the token balance is less than the threshold, the rebalance will be triggered"`
	Chains    []string            `json:"chains" yaml:"chains" comment:"The chains need to be monitored"`
	BotTypes  map[string][]string `json:"botTypes" yaml:"botTypes" comment:"BotTypes The monitoring type of this token on the specified chain, supporting balance_on_chain, helix_liquidity, and the default is balance_on_chain.\n The key is the chain name, and the value is the monitoring types on this chain.Example: {\"arbitrum\":[\"balance_on_chain\", \"helix_liquidity\"]}"`
}

type CrossChain struct {
	Providers   []string `json:"providers" yaml:"providers" comment:"use providers to cross chain"`
	TargetChain string   `json:"target_chain" yaml:"target_chain" comment:"target chain name"`
}

type Token struct {
	Name            string `json:"name" yaml:"name" comment:"Token name"`
	ContractAddress string `json:"contract_address" yaml:"contract_address" comment:"Token contract address"`
	// decimals token decimals
	Decimals int32 `json:"decimals" yaml:"decimals" comment:"Token decimals"`
}

// SourceToken used to buy other tokens
type SourceToken struct {
	Name   string   `json:"name" yaml:"name" comment:"Token name"`
	Chains []string `json:"chains" yaml:"chains" comment:"Chain name"`
}

// TargetToken use SourceToken to buy TargetToken
type TargetToken struct {
	Name      string          `json:"name" yaml:"name" comment:"Token name"`
	Amount    decimal.Decimal `json:"amount" yaml:"amount" comment:"The number of each rebalance"`
	Threshold decimal.Decimal `json:"threshold" yaml:"threshold" help:"Threshold when the token balance is less than the threshold, the rebalance will be triggered"`
}

// Provider liquidity providersMap
type Provider struct {
	// Type liquidity providersMap type
	Type ProviderType `json:"type" yaml:"type" comment:"Type liquidity providersMap type"`
	// Name liquidity providersMap name
	Name string `json:"name" yaml:"name" comment:"providersMap name"`
	// Config liquidity providersMap config, depend on the type
	Config map[string]interface{} `json:"config" yaml:"config" comment:"Config liquidity providersMap config, depend on the type"`
}

type DbConfig struct {
	// Type db type
	Type DbType `json:"type"`
	// MySQL mysql config
	MySQL *MysqlConfig `json:"mysql,omitempty" yaml:"MYSQL" comment:"MYSQL config"`
	// PostgreSQL postgresql config
	PostgreSQL *MysqlConfig `json:"postgresql,omitempty" yaml:"POSTGRESQL" comment:"POSTGRESQL config"`
	// SQLite sqlite config
	SQLite *Sqlite `json:"sqlite,omitempty" yaml:"SQLITE" comment:"SQLITE config"`
}

type MysqlConfig struct {
	// Host mysql host
	Host string `json:"host" yaml:"host"`
	// Port mysql port
	Port string `json:"port" yaml:"port"`
	// User mysql user
	User string `json:"user" yaml:"user"`
	// Password mysql password
	Password string `json:"password" yaml:"password"`
	// Database mysql database
	Database string `json:"database" yaml:"database"`
}

type Sqlite struct {
	// Path sqlite path
	Path string `json:"path"`
}

func (c Chain) GetToken(tokenName string) Token {
	for _, v := range c.Tokens {
		if !strings.EqualFold(v.Name, tokenName) {
			continue
		}
		return v
	}
	return Token{}
}

func (c *Config) Init() *Config {
	c.chainsMap = make(map[string]Chain)
	oldName2NewName := make(map[string]string)

	for index, v := range c.Chains {
		newName := constant.GetChainName(v.Id)
		if newName == "" {
			panic(fmt.Sprintf("chain id %d not found", v.Id))
		}
		oldName2NewName[v.Name] = newName
		c.Chains[index].Name = newName
		c.chainsMap[newName] = c.Chains[index]
	}

	c.sourceTokenMap = make(map[string]SourceToken)
	for index, v := range c.SourceToken {
		var chains []string
		for _, v := range v.Chains {
			if newName, ok := oldName2NewName[v]; ok {
				chains = append(chains, newName)
				continue
			}
			panic(fmt.Sprintf("chain %s not found", v))
		}
		c.SourceToken[index].Chains = chains
		c.sourceTokenMap[v.Name] = c.SourceToken[index]
	}

	c.providersMap = make(map[ProviderType][]Provider)
	for _, v := range c.Providers {
		c.providersMap[v.Type] = append(c.providersMap[v.Type], v)
	}

	c.wallets = make(map[string]Wallet)
	for walletIndex, v := range c.Wallets {
		for index, t := range v.Tokens {
			var (
				chains   []string
				BotTypes = make(map[string][]string)
			)
			for _, v := range t.Chains {
				if newName, ok := oldName2NewName[v]; ok {
					chains = append(chains, newName)
					continue
				}
				panic(fmt.Sprintf("chain %s not found in chains configs", v))
			}
			for name := range t.BotTypes {
				if newName, ok := oldName2NewName[name]; ok {
					BotTypes[newName] = t.BotTypes[name]
					continue
				}
				panic(fmt.Sprintf("chain %s not found", name))
			}
			c.Wallets[walletIndex].Tokens[index].Chains = chains
			c.Wallets[walletIndex].Tokens[index].BotTypes = BotTypes
		}
		c.wallets[v.Address] = c.Wallets[walletIndex]
	}
	return c
}

func (c *Config) Check() error {
	if len(c.Chains) == 0 {
		return errors.New("chains must be set")
	}
	for chainIndex, v := range c.Chains {
		if len(v.RpcEndpoints) == 0 {
			return errors.New("rpc_endpoints must be set")
		}
		if len(v.Tokens) == 0 {
			return errors.New("tokens must be set")
		}
		for index, v := range v.Tokens {
			if v.ContractAddress == "" {
				return errors.Errorf("chainsMap[%d]tokens[%d]contract_address must be set", chainIndex, index)
			}
			if v.Name == "" {
				return errors.Errorf("chainsMap[%d]tokens[%d]name must be set", chainIndex, index)
			}
			if v.Decimals == 0 {
				return errors.Errorf("chainsMap[%d]tokens[%d]decimals must be set", chainIndex, index)
			}
		}
	}

	if len(c.SourceToken) == 0 {
		return errors.New("source_token must be set")
	}
	for index, v := range c.SourceToken {
		for chainIndex, chain := range v.Chains {
			if _, ok := c.chainsMap[chain]; !ok {
				return errors.Errorf("source_token[%d]chainsMap[%d] not in chainsMap", index, chainIndex)
			}
			var ok bool
			for _, token := range c.chainsMap[chain].Tokens {
				if strings.EqualFold(token.Name, v.Name) {
					ok = true
				}
			}
			if !ok {
				return errors.Errorf("source_token[%d] token name not in chainsMap", index)
			}
		}
	}

	if len(c.Providers) == 0 {
		return errors.New("liquidity_providers must be set")
	}
	for index, v := range c.Providers {
		if v.Type == "" {
			return errors.Errorf("liquidity_providers[%d]type must be set", index)
		}

		if v.Name == "" {
			return errors.Errorf("liquidity_providers[%d]liquidity_name must be set", index)
		}
	}

	if len(c.Wallets) == 0 {
		return errors.New("wallets must be set")
	}
	for index, v := range c.Wallets {
		if v.Address == "" {
			return errors.Errorf("wallets[%d]address must be set", index)
		}
		if len(v.Tokens) == 0 {
			return errors.Errorf("wallets[%d]tokens must be set", index)
		}
		for tokenIndex, token := range v.Tokens {
			if token.Name == "" {
				return errors.Errorf("wallets[%d]tokens[%d]name must be set", index, tokenIndex)
			}
			if token.Amount.IsZero() {
				return errors.Errorf("wallets[%d]tokens[%d]amount must be set", index, tokenIndex)
			}
			if token.Threshold.IsZero() {
				return errors.Errorf("wallets[%d]tokens[%d]threshold must be set", index, tokenIndex)
			}

			if len(token.Chains) == 0 {
				return errors.Errorf("wallets[%d]tokens[%d]chainsMap must be set", index, tokenIndex)
			}
			for chainIndex, chain := range token.Chains {
				if _, ok := c.chainsMap[chain]; !ok {
					return errors.Errorf("wallets[%d]tokens[%d]chainsMap[%d] not in chainsMap", index, tokenIndex, chainIndex)
				}
				var ok bool
				for _, chainToken := range c.chainsMap[chain].Tokens {
					if strings.EqualFold(chainToken.Name, chainToken.Name) {
						ok = true
					}
				}
				if !ok {
					return errors.Errorf("wallets[%d]tokens[%d] token name not in chainsMap", index, tokenIndex)
				}
			}
		}
	}

	if c.Db.Type == "" {
		return errors.New("db type must be set")
	}
	if c.Db.Type == MySQL && c.Db.MySQL == nil {
		return errors.New("mysql config must be set")
	}
	if c.Db.Type == PostgreSQL && c.Db.PostgreSQL == nil {
		return errors.New("postgresql config must be set")
	}
	if c.Db.Type == SQLite && c.Db.SQLite == nil {
		return errors.New("sqlite config must be set")
	}
	return nil
}

func (c *Config) GetProvidersConfig(name string, providerType ProviderType, dest interface{}) error {
	for _, provider := range c.providersMap[providerType] {
		if !strings.EqualFold(provider.Name, name) {
			continue
		}

		conf, err := json.Marshal(provider.Config)
		if err != nil {
			return err
		}
		return json.Unmarshal(conf, dest)
	}
	return errors.Errorf("providersMap %s not found", name)
}

func (c *Config) GetChainConfig(chainName string) Chain {
	chain := c.chainsMap[chainName]
	if chain.Name == "" {
		logrus.Panicf("chain %s not found", chainName)
	}
	return chain
}

func (c *Config) GetTokenThreshold(wallet, tokenName, chain string) decimal.Decimal {
	for _, token := range c.wallets[wallet].Tokens {
		if !strings.EqualFold(token.Name, tokenName) {
			continue
		}
		if utils.InArray(chain, token.Chains) {
			return token.Threshold
		}
	}
	panic(fmt.Sprintf("token %s not found on chain %s", tokenName, chain))
}

func (c *Config) GetWallet(wallet string) wallets.Wallets {
	return wallets.NewWallets(wallets.WalletConfig{
		PrivateKey: c.wallets[wallet].PrivateKey,
		Address:    common.HexToAddress(c.wallets[wallet].Address),
		Operator: wallets.Operator{
			Address:       c.wallets[wallet].Operator.Address,
			Operator:      c.wallets[wallet].Operator.Operator,
			PrivateKey:    c.wallets[wallet].Operator.PrivateKey,
			MultiSignType: c.wallets[wallet].Operator.MultiSignType,
		},
		MultiSignType: c.wallets[wallet].MultiSignType,
	})
}

func (c *Config) GetWalletConfig(wallet string) Wallet {
	return c.wallets[wallet]
}

func (c *Config) GetTokenInfoOnChain(tokenName, chainName string) Token {
	for _, token := range c.chainsMap[chainName].Tokens {
		if !strings.EqualFold(token.Name, tokenName) {
			continue
		}
		return token
	}
	panic(fmt.Sprintf("token %s not found on chain %s", tokenName, chainName))
}

func (c *Config) GetTokenInfoOnChainByAddress(tokenAddress, chainName string, force ...bool) Token {
	for _, token := range c.chainsMap[chainName].Tokens {
		if !strings.EqualFold(token.ContractAddress, tokenAddress) {
			continue
		}
		return token
	}
	if len(force) != 0 && force[0] {
		return Token{}
	}
	panic(fmt.Sprintf("token %s not found on chain %s", tokenAddress, chainName))
}

func (c *Config) GetTokenPurchaseAmount(wallet, tokenName, chain string) decimal.Decimal {
	for _, token := range c.wallets[wallet].Tokens {
		if !strings.EqualFold(token.Name, tokenName) {
			continue
		}
		if utils.InArray(chain, token.Chains) {
			return token.Amount
		}
	}
	panic(fmt.Sprintf("token %s not found on chain %s", tokenName, chain))
}

func (c *Config) GetTokenAddress(tokenName, chainName string) string {
	address := c.GetTokenInfoOnChain(tokenName, chainName).ContractAddress
	if address == "" {
		logrus.Fatalf("token %s not found on chain %s", tokenName, chainName)
	}
	return address
}

func (c *Config) GetSourceTokenNamesByChain(chainName string) []string {
	var result []string
	for _, v := range c.SourceToken {
		if !utils.InArray(chainName, v.Chains) {
			continue
		}
		result = append(result, v.Name)
	}
	if len(result) == 0 {
		panic(fmt.Sprintf("source token not found on chain %s", chainName))
	}
	return result
}

func (w Wallet) CheckPrivateKey() error {
	if w.PrivateKey == "" {
		return errors.New("private key must be set")
	}
	key, err := crypto.HexToECDSA(w.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "privateKey")
	}
	if !strings.EqualFold(w.Address, crypto.PubkeyToAddress(key.PublicKey).Hex()) {
		return errors.New("private key and address not match")
	}
	return nil
}

func (c *Config) GetTaskInterval(name string, defaultInterval time.Duration) time.Duration {
	if len(c.TaskInterval) == 0 {
		return defaultInterval
	}
	if _, ok := c.TaskInterval[name]; !ok {
		return defaultInterval
	}
	return c.TaskInterval[name]
}

func (c *Config) IsNativeToken(chainName, tokenName string) bool {
	if strings.EqualFold(c.chainsMap[chainName].NativeToken, tokenName) {
		return true
	}
	for _, token := range c.chainsMap[chainName].Tokens {
		if strings.EqualFold(token.Name, tokenName) {
			return strings.EqualFold(token.ContractAddress, constant.ZeroAddress.Hex())
		}
	}
	return false
}

package wallets

var (
	wallets = make(map[string]func(conf WalletConfig) Wallets)
)

func Register(name string, wallet func(conf WalletConfig) Wallets) {
	wallets[name] = wallet
}

func Get(name string) func(conf WalletConfig) Wallets {
	return wallets[name]
}

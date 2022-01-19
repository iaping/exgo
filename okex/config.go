package okex

type Config struct {
	APIKey, SecretKey, Passphrase string
	Simulated, China              bool
}

package config

type Config struct {
	Telegram TelegramCredentials `yaml:"telegram_bot"`
	LogLevel int
}
type TelegramCredentials struct {
	Token    string `yaml:"token"`
	ProxyURL string `yaml:"proxyURL"`
}

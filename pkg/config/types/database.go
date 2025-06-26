package types
 
type Database struct {
	Host string `env:"PG_HOST"`
	Port string `env:"PG_PORT"`
	User string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	DbName string `env:"PG_DBNAME"`
	TimeZone string `env:"PG_TIME_ZONE"`
	SSLmode string `env:"PG_SSL_MODE"`
}
package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
)

func init() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8090/auth/google/callback",
		ClientID:     "GetConfig().GoogleClientId",
		ClientSecret: "GetConfig().GoogleClientSecret",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type Config struct {
	AppName        string `envconfig:"APP_NAME"`
	AppPort        string `envconfig:"APP_PORT"`
	ContextTimeout string `envconfig:"CONTEXT_TIMEOUT"`

	DbName             string `envconfig:"DB_NAME"`
	DbUser             string `envconfig:"DB_USERNAME"`
	DbPassword         string `envconfig:"DB_PASSWORD"`
	DbHost             string `envconfig:"DB_HOST"`
	DbPort             string `envconfig:"DB_PORT"`
	DbMaxIdleConns     int    `envconfig:"db_max_idle_conns" default:"50"`
	DbMaxOpenConns     int    `envconfig:"db_max_open_conns" default:"100"`
	JWTPublicKey       string `envconfig:"JWT_PUBLIC_KEY" default:"123"`
	GoogleClientId     string `envconfig:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `envconfig:"GOOGLE_CLIENT_SECRET"`
}

var once sync.Once
var instance Config

func GetConfig() Config {
	once.Do(func() {
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return instance
}

// SetConfig set custom config
func SetConfig(cfg Config) {
	instance = cfg
}

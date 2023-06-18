package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-sql-driver/mysql"
	traqoauth2 "github.com/ras0q/traq-oauth2"
)

type SessionKey string

const (
	SessionName string = "traq-oauth2-example"

	CodeVerifierKey SessionKey = "code_verifier"
	TokenKey        SessionKey = "access_token"
	TraqIDKey       SessionKey = "user_id"
)

func getEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func AppAddr() string {
	return getEnv("APP_ADDR", ":8080")
}

func MySQL() *mysql.Config {
	return &mysql.Config{
		User:   getEnv("DB_USER", "root"),
		Passwd: getEnv("DB_PASSWORD", "pass"),
		Net:    getEnv("DB_NET", "tcp"),
		Addr: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
		),
		DBName:               getEnv("DB_NAME", "backend_sample"),
		AllowNativePasswords: true,
	}
}

func TraqOAuth2() *traqoauth2.Config {
	return traqoauth2.NewConfig(
		getEnv("TRAQ_CLIENT_ID", ""),
		getEnv("TRAQ_REDIRECT_URL", ""),
	)
}

func ClientURL() *url.URL {
	u, _ := url.Parse(getEnv("CLIENT_URL", "http://localhost:3000"))
	return u
}

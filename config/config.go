package config

import (
	"os"
)

var (
	AppEnv                 string
	Httpport               string
	Service                string
	DataBaseDriver         string
	ColleagueApiConnection string
)

func Read() {
	if appEnv := os.Getenv("APP_ENV"); appEnv == "" {
		AppEnv = "test"
	} else {
		AppEnv = appEnv
	}
	if httpport := os.Getenv("HTTP_PORT"); httpport == "" {
		Httpport = "8001"
	} else {
		Httpport = httpport
	}
	if dataBaseDriver := os.Getenv("DATABASE_DRIVER"); dataBaseDriver == "" {
		DataBaseDriver = "sqlite3"
	} else {
		DataBaseDriver = dataBaseDriver
	}
	if colleagueApiConnection := os.Getenv("COLLEAGUE_API_CONNECTION"); colleagueApiConnection == "" {
		ColleagueApiConnection = ":memory:"
	} else {
		ColleagueApiConnection = colleagueApiConnection
	}
	Service = "colleague-api"
}

func ReadForTest() {
	AppEnv = "test"
	Httpport = "8001"
	Service = "colleague-api"
	DataBaseDriver = "sqlite3"
	ColleagueApiConnection = ":memory:"
}

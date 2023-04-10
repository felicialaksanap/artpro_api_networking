package config

import "github.com/tkanos/gonfig"

// buat configuration server, sekarang pakai lokal nanti
// diganti di config.json
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}

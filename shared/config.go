package shared

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"path/filepath"
)

//Configs contains application configurations for all application modes
type Configs struct {
	Debug   Config
	Release Config
	Test    Config
}

//Config contains application configuration for active application mode
type Config struct {
	Public        string `json:"public"`
	Domain        string `json:"domain"`
	SessionSecret string `json:"session_secret"`
	CsrfSecret    string `json:"csrf_secret"`
	Ssl           bool   `json:"ssl"`
	SignupEnabled bool   `json:"signup_enabled"` //always set to false in release mode (config.json)
	Database      DatabaseConfig
	Oauth         OauthConfig
}

//DatabaseConfig contains database connection info
type DatabaseConfig struct {
	Host     string
	Name     string //database name
	User     string
	Password string
}

//OauthConfig contains oauth login info
type OauthConfig struct {
	Facebook OauthApp
	Google   OauthApp
	Linkedin OauthApp
	Vk       OauthApp
}

//OauthApp contains oauth application data
type OauthApp struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
	Page         string `json:"page"`  //page id, mainly for facebook atm
	Token        string `json:"token"` //page token, mainly for facebook atm. Read http://stackoverflow.com/questions/17197970/facebook-permanent-page-access-token
}

var (
	config *Config
)

//loadConfig unmarshals config for current application mode
func loadConfig() {
	wd := Getwd()
	data, err := ioutil.ReadFile(filepath.Join(wd, "config", "config.json"))
	if err != nil {
		panic(err)
	}
	configs := &Configs{}
	err = json.Unmarshal(data, configs)
	if err != nil {
		panic(err)
	}
	switch GetMode() {
	case DebugMode:
		config = &configs.Debug
	case ReleaseMode:
		config = &configs.Release
	default:
		config = &configs.Test
	}
	if !path.IsAbs(config.Public) {
		config.Public = path.Join(wd, config.Public)
	}
}

//GetConfig returns actual config
func GetConfig() *Config {
	return config
}

//PublicPath returns path to application public folder
func PublicPath() string {
	return config.Public
}

//UploadsPath returns path to public/uploads folder
func UploadsPath() string {
	return path.Join(config.Public, "uploads")
}

package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "settings/tests.json",
}

type Settings struct {
	Version           				string
	Description        				string
	PrivateKeyPath        			string
	PublicKeyPath         			string
	JWTAccessTokenExpiration    	int
	JWTRefreshTokenExpiration 		int
	ServerPort            			string
}

var settings Settings = Settings{}
var env = ""

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting production environment due to lack of GO_ENV value")
		env = "production"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	configFile := GetExecDirectory() + "/" + environments[env]
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}

	fmt.Println("Version: ", settings.Version)
	fmt.Println("Config File: ", configFile)
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}

// GetExecDirectoy will get current executable directory
func GetExecDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

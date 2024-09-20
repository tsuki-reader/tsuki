package core

import (
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/ini.v1"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Anilist struct {
		ClientID string
	}
	Files struct {
		Config   string
		Database string
	}
	Logger LoggerInterface
}

func (c *Config) GetServerAddress() string {
	return c.Server.Host + ":" + strconv.Itoa(c.Server.Port)
}

var CONFIG *Config

// TODO: Add tests. We will need to point the config file to test/dump/config.ini so maybe
// pass in an optional config dir to this function
func SetupConfig() {
	CONFIG = &Config{}
	CONFIG.Logger = NewLogger()

	configDir := GetConfigDir()
	configFilePath, err := initConfig(configDir)
	if err != nil {
		CONFIG.Logger.Fatal("Could not initialise config in directory: " + configDir)
	}

	cfg, err := ini.Load(configFilePath)
	if err != nil {
		CONFIG.Logger.Fatal("Could not load config file: " + configFilePath)
	}

	serverSection := cfg.Section("Server")
	host := serverSection.Key("host").MustString("127.0.0.1")
	port := serverSection.Key("port").MustInt(1337)

	anilistSection := cfg.Section("Anilist")
	clientId := anilistSection.Key("client_id").MustString("21156")

	CONFIG.Server.Host = host
	CONFIG.Server.Port = port
	CONFIG.Anilist.ClientID = clientId
	CONFIG.Files.Config = configFilePath
	CONFIG.Files.Database = filepath.Join(configDir, "tsuki.db")
}

func GetConfigDir() string {
	dataDir, err := os.UserConfigDir()
	if err != nil {
		dataDir = ""
	}
	return filepath.Join(dataDir, "Tsuki")
}

func initConfig(configDir string) (string, error) {
	// Make the config dir if it doesn't exist
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}

	configFilePath := filepath.Join(configDir, "config.ini")
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		// Create a base config file as it does not exist
		if err := createBaseConfig(configFilePath); err != nil {
			return "", err
		}
	}

	return configFilePath, nil
}

// Create a config file with default values
func createBaseConfig(configFilePath string) error {
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	cfg, err := ini.Load(configFilePath)
	if err != nil {
		return err
	}

	serverSection := cfg.Section("Server")
	serverSection.Key("host").SetValue("127.0.0.1")
	serverSection.Key("port").SetValue("1337")

	anilistSection := cfg.Section("Anilist")
	anilistSection.Key("client_id").SetValue("21156")

	if err := cfg.SaveTo(configFilePath); err != nil {
		return err
	}

	return nil
}

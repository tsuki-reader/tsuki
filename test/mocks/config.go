package mocks

import (
	"tsuki/core"
)

func BuildMockConfig(logger *MockLogger, level int) {
	filePrefix := filepathPrefix(level)
	config := &core.Config{
		Server: struct {
			Host      string
			Port      int
			SecretKey string
		}{
			Host:      "127.0.0.1",
			Port:      1337,
			SecretKey: "mock_key",
		},
		Anilist: struct{ ClientID string }{
			ClientID: "12345",
		},
		Files: struct {
			Config   string
			Database string
		}{
			Config:   filePrefix + "test/dump/config.ini",
			Database: filePrefix + "test/dump/tsuki.db",
		},
		Logger: logger,
	}
	core.CONFIG = config
}

func filepathPrefix(level int) string {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "../"
	}
	return prefix
}

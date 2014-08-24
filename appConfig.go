// Package appConfig ~/.{appName}/configの書き込み,作成,読み込みを行うパッケージ
package appConfig

import (
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

type AppConfig struct {
	ConfigFileName string
	AppName        string
}

// NewAppConfig create AppConfig.
func NewAppConfig(appName string) *AppConfig {
	return &AppConfig{
		ConfigFileName: "config",
		AppName:        appName,
	}
}

// configファイルを作成する中身は空.
func (a AppConfig) WriteAppConfig(data []byte) error {
	if err := createAppConfigDir(a.AppName); err != nil {
		return err
	}

	filePath, err := createAppConfigFilePath(a.AppName, a.ConfigFileName)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, os.FileMode(0644))
}

// configファイルを読み込む[]byte.
func (a AppConfig) ReadAppConfig() ([]byte, error) {
	filePath, err := createAppConfigFilePath(a.AppName, a.ConfigFileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(filePath)
}

// ~/.{appName}ディレクトリを作成
// すでに存在する場合スルー
func createAppConfigDir(appName string) error {
	dirPath, err := createAppConfigDirPath(appName)
	if err != nil {
		return err
	}

	// check
	if _, err := ioutil.ReadDir(dirPath); err == nil {
		return nil
	}

	// create dir
	return os.Mkdir(dirPath, os.FileMode(0755))
}

func createAppConfigDirPath(appName string) (string, error) {
	// home
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dirName := "." + appName
	dirPath := strings.Join([]string{usr.HomeDir, dirName}, "/")
	return dirPath, nil
}

func createAppConfigFilePath(appName, configFileName string) (string, error) {
	// home
	dirPath, err := createAppConfigDirPath(appName)
	if err != nil {
		return "", err
	}

	filePath := strings.Join([]string{dirPath, configFileName}, "/")
	return filePath, nil
}

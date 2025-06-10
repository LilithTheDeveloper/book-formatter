// package config

// import (
// 	"errors"
// 	"go/types"
// 	"os"
// 	"path/filepath"
// 	"gopkg.in/yaml.v3"
// )


// func GetAppDataDirectory() (string, error) {
// 	app_directory_name := "bookformatter"
// 	home, err := os.UserHomeDir()
	

// 	if err != nil {
// 		return "", errors.New("Unable to get user home directory: " + err.Error())
// 	}

// 	if err = os.MkdirAll(filepath.Join(home, app_directory_name), os.ModePerm); err != nil {
// 		return "", errors.New("Unable to create application data directory: " + err.Error())
// 	}

// 	return filepath.Join(home, app_directory_name), nil
// }

// // func load_config(path string) (*types.Config, error) {
// // 	data, err := os.ReadFile(path)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	var cfg types.Config
// // 	err = yaml.Unmarshal(data, &cfg)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	return &cfg, nil
// // }
package config

import "github.com/joho/godotenv"

//NewConfig return config
func NewConfig() (map[string]string, error) {
	cfg, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}
	return cfg, nil

}

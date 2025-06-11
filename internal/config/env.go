package config

import "github.com/joho/godotenv"

func GetEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load env file")
	}

	return nil
}

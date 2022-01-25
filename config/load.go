/**
 * Taken from https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
 */

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

func GetPath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		logger.Errorf("Error with Loading Env %s", err.Error())
		err := fmt.Errorf("error with loading env", err.Error())
		return "", err
	}
	return path, nil
}

func LoadEnv(name string) error {
	path, err := os.UserHomeDir()
	if err != nil {
		logger.Errorf("Error with Loading Env %s", err.Error())
		err := fmt.Errorf("error with loading env", err.Error())
		return err
	}
	// loads values from .env into the system
	// name := "/.lintasarta/config-dbaas.env"
	if err := godotenv.Load(path + name); err != nil {
		logger.Errorf("No .env file in %s called %s. Application dismissed", path, name)
		err := fmt.Errorf("no .env file in %s called %s. application dismissed", path, name)
		return err
	}
	return nil
}

// GetEnv : Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// GetEnvAsInt : Simple helper function to read an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// GetEnvAsBool : Helper to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

// GetEnvAsSlice : Helper to read an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	val := strings.Split(valStr, sep)
	return val
}

package constants

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	GlobalConfig   []byte
	AppName        string
	AppVersion     string
	RedisKeyPrefix string
	Debug          bool
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		return
	}
	Debug = os.Getenv("DEBUG") == "true"
	if Debug {
		os.Setenv("DEBUG", "true")
	} else {
		os.Setenv("DEBUG", "false")
	}

	AppName = os.Getenv("APP_NAME")
	if AppName == "" {
		AppName = "logistics"
	}
	AppVersion = os.Getenv("APP_VERSION")
	if AppVersion == "" {
		AppVersion = "1.0.0"
	}
	RedisKeyPrefix = os.Getenv("REDIS_KEY_PREFIX")
	if RedisKeyPrefix == "" {
		RedisKeyPrefix = "logistics"
	}

	WeChatKeyExpireStr := os.Getenv("WECHAT_KEY_EXPIRE")
	if WeChatKeyExpireStr != "" {
		if val, err := strconv.Atoi(WeChatKeyExpireStr); err == nil {
			WeChatKeyExpire = val
		}
	}

	ObtenationIterationsStr := os.Getenv("OBTENATION_ITERATIONS")
	if ObtenationIterationsStr != "" {
		if val, err := strconv.Atoi(ObtenationIterationsStr); err == nil {
			ObtenationIterations = val
		}
	}
}

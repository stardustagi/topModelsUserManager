package jwt

import (
	"fmt"
	"logistics/libs/logs"
	"os"

	"go.uber.org/zap"
)

var (
	appName    = os.Getenv("APP_NAME")
	appVersion = os.Getenv("APP_VERSION")
)

func AdminTokenKey(id string) string {
	logs.Infof("%s:admin:token:%s", appName, id)
	return fmt.Sprintf("%s:admin:token:%s", appName, id)
}

func UserTokenKey(id string) string {
	logs.Info("user token", zap.String("id", id))
	return fmt.Sprintf("%s:user:token:%s", appName, id)
}

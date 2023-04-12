package dotenv

import (
	"glucovie/pkg/logger"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func GetEnvironmentVariable(key string) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../.env")
	err := godotenv.Load(dir)
	if err != nil {
		logger.Log.Error("Failed to load .env file", zap.Any("error", err))
		dir = path.Join(path.Dir(filename), "../../.env_dev")
		err = godotenv.Load(dir)
		if err != nil {
			logger.Log.Error("Failed to load .env_dev file", zap.Any("error", err))
			return ""
		} else {
			return os.Getenv(key)
		}
	}
	return os.Getenv(key)
}

// Package app 应用信息
package app

import (
	"huango/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimeZone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimeZone)
}

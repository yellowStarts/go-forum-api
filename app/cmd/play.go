package cmd

import (
	"huango/pkg/console"
	"huango/pkg/redis"
	"time"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context.",
	Run:   runPlay,
}

// 调试完成后记得清楚测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 存进去 redis 中
	redis.Redis.Set("hello", "hi from redis.", 10*time.Second)
	// 从 redis 里取出
	console.Success(redis.Redis.Get("hello"))
}

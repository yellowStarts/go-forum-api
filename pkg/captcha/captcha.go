// Package captcha 处理图片验证码逻辑
package captcha

import (
	"huango/pkg/app"
	"huango/pkg/config"
	"huango/pkg/redis"
	"sync"

	"github.com/mojocn/base64Captcha"
)

type Capthca struct {
	Base64Capthca *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Capthca

// NewCaptcha 单例模式获取
func NewCaptcha() *Capthca {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Capthca{}

		// 使用全局 Redis 对象，并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha:",
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 高
			config.GetInt("captcha.width"),       // 宽
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Capthca = base64Captcha.NewCaptcha(driver, &store)
	})
	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Capthca) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Capthca.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Capthca) VerifyCaptcha(id string, answer string) (match bool) {
	// 方便本地和 API 自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Capthca.Verify(id, answer, false)
}
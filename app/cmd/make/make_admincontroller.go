package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeAdminController = &cobra.Command{
	Use:   "admincontroller",
	Short: "Create admin controller, exmaple: make apicontroller user",
	Run:   runMakeAdminController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAdminController(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	// 组件目标目录
	filePath := fmt.Sprintf("app/http/controllers/admin/%s_controller.go", model.TableName)

	// 从模板中创建文件（做好变量替换）
	createFileFromStub(filePath, "admincontroller", model)
}

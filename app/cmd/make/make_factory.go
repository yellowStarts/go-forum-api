package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, exmaple: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeFactory(cmd *cobra.Command, args []string) {

	model := makeModelFromString(args[0])

	// 组件目标目录
	filePath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	// 从模板中创建文件（做好变量替换）
	createFileFromStub(filePath, "factory", model)
}

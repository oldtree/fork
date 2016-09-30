//加载通用的配置文件
package conf

import (
	"fork/tools"
)

var version = "0.0.1"

func init() {
	tools.Info(version)
}

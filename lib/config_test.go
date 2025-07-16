package lib

import (
	"fmt"
	"testing"
)

// 基础测试用例
func TestIntMinBasic(t *testing.T) {
	serverConfig := LoadServerConfig()

	fmt.Printf("Loaded config: %+v\n", serverConfig) // 打印配置信息

}

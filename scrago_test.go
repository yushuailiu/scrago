// 报名设置为_test结尾，限制测试代码只能访问包内导出的符号
package scrago_test

import (
	"testing"
	"net/http"
	"fmt"
)

func TestScrago(t *testing.T) {
	req,err := http.NewRequest("POST", "http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
}

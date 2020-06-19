package httpkit

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_post2(t *testing.T) {

	httpC := NewHTTPClient(12, 2)
	robot := "https:xx"
	message := `{
				  "msgtype": "markdown",
				  "markdown": {
					"title": "xx",
					"text": "@xx，` + "haha" + `"
				  },
				  "at": {
					"atMobiles": [
					  "xx"
					],
					"isAtAll": false
				  }
				}`
	var buf io.Reader
	buf = strings.NewReader(message)
	str, err := httpC.Post2(robot, "application/json;charset=utf-8", buf)

	fmt.Println(str, err)

}

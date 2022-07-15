package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stdReader = bufio.NewReader(os.Stdin)

// GetStdinStringValue 从标准输入获取内容.
//
// 要求前缀(prefix)，与默认值(omit)
func GetStdinStringValue(prefix string, omit string) string {
	var value string
	// 交互前缀Q/A
	fmt.Print(prefix)

	value, _ = stdReader.ReadString('\n')
	value = strings.TrimSpace(value)

	if len(value) == 0 {
		value = omit
	}

	return value
}

// GetStdinNumberValue 从标准输入获取内容.
// 要求前缀(prefix)，与默认值(omit)
func GetStdinNumberValue(prefix string, omit int) int {
	var value = GetStdinStringValue(prefix, strconv.Itoa(omit))

	number, err := strconv.Atoi(value)
	if err != nil {
		number = omit
	}

	return number
}

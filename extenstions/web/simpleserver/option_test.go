package simpleserver

import (
	"fmt"
	"strconv"
	"testing"
)

func TestKK(t *testing.T) {
	var speed string = "1024kk"
	fmt.Println(strconv.Atoi(speed))
}

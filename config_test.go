package loadConfig

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig()
	config.loadConfig("./config.ini")
	value := config.GetConfig("server")
	fmt.Println(value)
}

package loadConfig

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	config := NewConfig()
	config.LoadConfig("./config.ini")
	value := config.GetConfig("server")
	host := config.GetConfig("host")
	fmt.Println(value, host)
}

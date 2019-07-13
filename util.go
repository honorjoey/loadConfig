package loadConfig

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) (*[]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("open the file failed!")
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, 4096*1000)
	_, err = file.Read(buf)
	if err != nil {
		log.Println("read the file failed!")
		return nil, err
	}
	slice := strings.Split(string(buf), "\n")

	return &slice, nil
}

func readConfig(filename string) *map[string]string {
	var ConfigMap = make(map[string]string, 100)
	configs, err := ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	slice := *configs
	for _, v := range slice {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "#") {
			continue
		}
		if len(v) == 0 {
			continue
		}
		config := strings.Split(v, "=")
		ConfigMap[config[0]] = config[1]
	}
	return &ConfigMap
}

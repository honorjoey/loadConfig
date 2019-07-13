# loadConfig
### 1.Introduction
这是一个读取配置文件的项目。
### 2.Example
配置文件示例
```config.ini```
```
#config
server=127.0.0.1
host=3306
```
用该项目读取配置文件
```test.go```
```golang
config := NewConfig()
config.LoadConfig("./config.ini")
value := config.GetConfig("server")
fmt.Println(value)
```
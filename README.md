# xencoding
这是一个给 json 和 xml 序列化功能增加过滤结构体字段的 Go 语言功能包

xencoding/json/standard 和 xencoding/xml/standard 分别从 Go 语言标准库 encoding/json、encoding/xml 中复制，并增加了结构体字段过滤功能
* MarshalFilter - 过滤某些字段
* MarshalSelector - 仅仅选择指定字段用来json序列化
* MarshalWithOption - 传递自定义配置。用来同时传递过滤器和选择器

## 用法
```golang
import (
    "github.com/admpub/xencoding/filter"
    json "github.com/admpub/xencoding/json/standard"
)

type User struct{
    Username string
    Password string
}

func main(){
    u := User{
        Username: "Username",
        Password: "Password",
    }
    r, err := json.MarshalFilter(u, filter.Exclude("Password")) // 排除 Password 字段
    if err != nil {
        panic(err)
    }
    println(string(b)) // output: {"Username":"Username"}
}
```

更多使用方式可以参考单元测试代码：  
[json/standard/filter_test.go](json/standard/filter_test.go)   
[xml/standard/filter_test.go](xml/standard/filter_test.go)
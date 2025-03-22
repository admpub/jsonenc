此处代码从标准库 encoding/json 中复制，并增加了结构体字段过滤功能
* MarshalFilter - 过滤某些字段
* MarshalSelector - 仅仅选择指定字段用来json序列化
* MarshalWithOption - 传递自定义配置。用来同时传递过滤器和选择器
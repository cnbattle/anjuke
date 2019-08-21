# 安居客爬取

## 使用

1. 复制`config.simple.toml` 到 `config.toml`

2. 根据个人需要，修改`config.toml`配置

3. `go run anjuke.go`

等待爬取完成，在运行目录会生成一个类似`database-20190821.db` 的`sqlite3`数据文件，根据个人需要再进去处理。


## 下一步计划

- 爬取字段内容自定义
- 增加生成`csv`文件

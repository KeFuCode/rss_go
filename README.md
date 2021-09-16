# Rss数据源的匹配器

# 项目介绍：
借助goroutine，从不同数据源拉取数据，将数据内容与一组搜索项做对比，将匹配的内容显示在终端窗口。

data：包含data.json，包含一组数据源。

matchers：搜索rss源的匹配器。

search：
- default.go，搜素数据用的默认匹配器；
- feed.go，用于读取json数据文件；
- match.go：用于支持不同匹配器的接口；
- search.go，执行搜索的主控制逻辑。

main.go：程序的入口。

# 项目成果：
程序读取文本文件，进行网络调用，解码XML和JSON成为结构化类型数据，并利用Go语言的并发机制保证这些操作的速度。

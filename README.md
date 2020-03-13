# golang实现的大淘客商品入库抓取
多线程抓取：使用goroutine
抓取完毕自动退出：使用waitGroup实现

### 使用方法
1. 下载本库代码
2. 安装依赖 `go get github.com/joho/godotenv`
3. 复制 `.env.example` 文件为 `.env` 并填写API密钥
4. 在 `fetcher/write-db.go` 文件中处理写入数据库的实现代码
5. 运行即可

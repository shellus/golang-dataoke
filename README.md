# golang实现的大淘客(http://www.dataoke.com/)商品入库抓取示例
多线程抓取：使用 goroutine 
抓取完毕自动退出：使用 WaitGroup 实现

### 使用方法1 下载本库代码直接使用
2. 安装依赖 `go get github.com/joho/godotenv`
3. 复制 `.env.example` 文件为 `.env` 并填写API密钥
4. 在 `fetcher/write-db.go` 文件中处理写入数据库的实现代码
5. 运行即可

### 使用方法2 通过import方式使用
2. 安装本库 `go get github.com/shellus/golang-dataoke`
3. 注入api密钥配置
    ```go
        fetcher.SetConfig(fetcher.Config{
            App_key:    os.Getenv("APP_KEY"),
            App_secret: os.Getenv("APP_SECRET"),
        })
    ```
4. 调用
    ```go
    // 获取总页数
    totalPage := GetTotalPage()
    
    // 请求每一页
    for p := 1; p <= totalPage; p++ {
        items, err := Paginator(p)
        if err != nil {
            panic(err.Error())
        }
        // 循环一页中的商品信息
        for item := range items {
            // todo ...
            // item.Title
        }
    }
    
    ```

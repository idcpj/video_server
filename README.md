## 运行
每个模块都如此 运行
```
> cd video_server/api
> go install
> api
```

## 技术栈
### 路由管理
github.com/julienschmidt/httprouter

## session 管理
同时存入 sync.Map 与数据库中

## http 中间件

在 原始的 httpServer 上封装一层
```
r := RegisetHandlers() 
mh := NewMiddleWareHandle(r)

由于 第二个参数 是 ServeHTTP(ResponseWriter, *Request),所以只要满足这个
接口的参数即可实现该接口,所有只要 mh 最后返回的是 ServerHttp 接口即可
http.ListenAndServe(":8000", mh.r) 
```

## 编译
在模块下 
`go install` 
在 go 的`bin`现在查看生成的可执行文件
编译运行
`go install ./web && ./build.sh   && /Users/idcpj/go/bin/video_server_web_ui/web`
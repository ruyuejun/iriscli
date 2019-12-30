## 运行步骤
```
git clone git@github.com:ruyuejun/iriscli.git
cd iriscli
go run main.go
```

## 测试路由
- GET: localhost:3000/simple        # 简单路由
- GET: localhost:3000/simple/qaq    # 带参简单路由
- POST:localhost:3000/book/list     # 完全MVC+gorm
# Go gin

## 快速开始
安装依赖
```bash
go download

```
启动服务
```bash
make start-service
```
打包服务
```bash
make build-servicez
```

## 生成数据库 Model
首先修改 `gen.yaml` 中的数据库配置，然后执行以下命令
```bash
$ make gen-db-model
```

## 支持的功能
- [x] 生成数据库 Model
- [x] Restful API
- [ ] GraphQL API
- [ ] 全局 panic 捕获
- [ ] 定时任务
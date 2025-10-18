---
sidebar_position: 1
---

# 脚本

通常使用`bin/`下的脚本来操作游戏服务器。

## load_env.sh

用于设置环境变量的辅助脚本。

作用：

- 设置[APP_ENV](./global_environment.md#APP_ENV)默认值
    - development。
- 根据`APP_ENV`值读取[.env](./dotenv.md)内容。

使用方法：

```bash
source ./bin/load_env.sh
```

## migrate

实际上是[Atlas](https://atlasgo.io/)的封装。
`./bin/migrate -h`查看具体用法。

## test

运行测试。

```bash
# 运行所有测试
./bin/test ./...

# 只运行用户模型的测试
./bin/test ./models/user_test.go
```

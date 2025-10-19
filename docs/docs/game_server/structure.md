---
sidebar_position: 1
---

# 结构

总体使用MVC结构。
受[Ruby on Rails](https://rubyonrails.org/)影响。

## Model

代表领域模型的Struct。
大部分业务逻辑都应该包装成模型。

部分模型需要持久化到数据库。
基本持久化模型结构参照User模型(`./models/user.go`)。

### 持久化

此项目使用[Ent](https://entgo.io/)作为ORM框架。
调用Ent通常是Model的职责。

## View

View有可能渲染HTML，也可能渲染JSON。
渲染JSON的View被称为Serializer。
基本Serializer结构参照User Serializer(`./serializers/user.go`)。

## Controller

Controller通常和URL1:1对应。
职责是调用Model执行业务逻辑，然后调用View渲染结果。
基本Controller结构参照Authentication Controller(`./controllers/authentication.go`)

### Routes

URL和Controller的映射关系定义在[这里](./routes/router.go)。

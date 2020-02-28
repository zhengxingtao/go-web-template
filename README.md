# go-web-template

## 项目说明：

### 该项目是一个基于go-restful框架的go-web项目脚手架。其中代码结构是我参照java的标准制定的，开箱即用。

## 作为一个web项目我们关心的解决方案如下：
### 数据库
#### 数据库项目中集成了MongoDB，使用的是 gopkg.in/mgo.v2 库。
### Redis
#### Redis我们使用的go-redis库。目前redis的使用场景是用来存储短信验证码。
### 配置文件统一管理
#### 项目中配置文件统一放在application.yml文件中，方便统一维护管理。
### 授权管理
#### 项目集成jwt，jwt的一些详细配置也已经做好了。
### 统一返回定义
#### 项目所有接口遵循restFull风格，返回值格式Json，项目中定义了一个BaseRespnse，用来统一约定返回的数据体。
### api文档管理
#### 项目集成swagger，并且swagger已经配置好授权模式。

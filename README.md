**# golang-example**

## golang微服务示例

### 文件夹说明

- pkg(该文件夹需要抽离，单独构建pkg共有库)

  提供了代码需要的一些基本封装，包括go-micro调用的封装，redis调用的抽象，数据库连接的封装等等

- proto

  该文件夹为proto文件的定义,即接口的定义处

- store

  数据库交互文件

- svc

  服务逻辑具体定义文件

- srv

  各个服务启动的启动模块，用于定义各个服务的启动参数和dockerFIle的定义处

# SCAT
> SCAT(Smart Contract Audit Tasks)是用于自动化审计智能合约的平台系统，集成了主流的自动化检测工具，涵盖 SWC 中的全部漏洞类型，能够自主上传合约进行审计，持久存储审计结果，建立合约仓库。

#### ENV
- ubuntu 20.04
- golang 1.21.5
- docker
```bash
Client:
 Version:    24.0.5
Server:
 Server Version: 24.0.5
```
- docker-compose
```bash
docker-compose version 1.25.0, build unknown
docker-py version: 4.1.0
CPython version: 3.8.10
OpenSSL version: OpenSSL 1.1.1f  31 Mar 2020
```  
## SCAT架构示意图
![scat架构图](./img/scat.png)  

### contract_analyzer 
对外提供合约漏洞检测接口的模块，根据外部请求中的 tool 类型通过grpc转发到对应的 analyze_tool 服务。

### analyze_tool

每个 tool 都形成一个gRPC 服务端，跟指定的tool_name形成某一款tool的gRPC服务端容器。

### pkg （共享库）
pkg 中是各服务共享的一些库
models 中包含一些全域具有共识的模型，如 Chain、Contract、Vulnerability、Tool 等。
apis 中包含各服务接口的模型。
为了保证服务的兼容性，模型中字段通常只能增加而不能减少（容易引发上下游 bug）。

### pb 
gRPC 服务端和客户端的约定描述  

## 如何使用

```bash
cd gin-grpc-scat
bash build.sh
docker-compose build && docker-compose up -d
```

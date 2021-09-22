##### go-active
##### 奇怪的知识

```
网络参考模型ISO:
1.应用层: HTTP FTP 等
2.表示层: 可忽略
3.会话层: 可忽略
4.传输层: 格式化信息流, 提供可靠传输(TCP三次握手 等)
5.网络层: 为数据包选择路由,封装层IP数据包(为数据选择不同IP地址的路由)
6.链路层: 负责接收IP数据包交给物理层 或 接受物理帧抽出IP数据包交给网络层
7.物理层: 物理传输

TCP/IP:
协议族, 包含网络通信各层中各式各样的协议 
如: 应用层HTTP, 传输层TCP UDP, 网络层IP, 链路层SLIP, 物理层ISO2110 等

TCP(与TCP/IP不同, 单指传输层TCP协议): 
握手: 
1. 客户端发送 SYN 报文给服务器端, 进入 SYN_SEND 状态
2. 服务器端收到 SYN 报文, 回应一个 SYN 报文, 进入 SYN_RECV 状态
3. 客户端收到服务器端的 SYN 报文, 回应一个 ACK 报文, 连接成功
挥手: 
1. 客户端发送 FIN 报文给服务器端
2. 服务器端收到 FIN 报文, 回应一个 ACK 报文
3. 服务器在回应 ACK 报文后一段时间，再回应一个 FIN 报文
4. 客户端收到 FIN 报文, 回应一个 ACK 报文, 连接断开

UDP:
不可靠传输, 发送数据后不管是否到达
不能用于音视频传输, 若某帧的I帧丢失, 后续P帧画面无法正常观看

Socket(套接字):
TCP/IP 抽象封装的 应用层网络通信 API
HTTP等协议均由Socket实现
```

```
***terminal 终端 golang交叉编译
SET GOARCH=amd64  (>=win10 $env GOARCH="amd64)
SET GOOS=linux    (>=win10 $env GOOS="linux")
go build            

***terminal 终端 golang部署
上传交叉编译的文件   scp d:/文件 root@198.168.1.1:/root/文件夹
修改文件权限        chmod -R 777 文件
后台运行            nohup ./文件

linux 查看当前目录下文件 ls 或 ls -al
```
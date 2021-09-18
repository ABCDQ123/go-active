##### go-active
```
***terminal 终端 交叉编译
SET GOARCH=amd64  (>=win10 $env GOARCH="amd64)
SET GOOS=linux    (>=win10 $env:GOOS="linux")
go build            

***terminal 终端 部署
上传交叉编译的文件   scp d:/文件 root@198.168.1.1:/root/文件夹
修改文件权限        chmod -R 777 文件
运行               nohup ./文件

查看当前目录下文件 ls
                ls -al
```
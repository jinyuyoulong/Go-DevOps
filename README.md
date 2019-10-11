# Go-DevOps
Go 实现 DevOps demo 测试
DevOps

1. 编写sh脚本
2. 配置 GitHub的 Webhooks 功能

1 deploy.sh
```
kill -9 $(pgrep webserver)
cd ~/Go-DevOps/
git pull git@github.com:jinyuyoulong/Go-DevOps.git
cd fan-webserver/
./fan-webserver &
```
2 webhooks
add webhooks
在Payload URL 中填入
```
http://10.211.55.5:5000
```
Content type 选择 x-www-form-urlencoded

选择 Just the push event.选项执行操作
当github 被push了之后会自动调用上边配置的URL

因为没有外网服务器IP，最终测试还未走完。
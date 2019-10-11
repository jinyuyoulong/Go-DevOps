# 自动部署
kill -9 $(pgrep webserver)
cd ~/Go-DevOps/
git pull git@github.com:jinyuyoulong/Go-DevOps.git
cd fan-webserver/
go build #如果在本地交叉编译的话，不用执行这行代码。
./fan-webserver &
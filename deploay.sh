kill -9 $(pgrep webserver)
cd ~/Go-DevOps/
git pull git@github.com:jinyuyoulong/Go-DevOps.git
cd fan-webserver/
go build
./fan-webserver &
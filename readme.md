# cmd
## git clone
git clone http://oauth2:xxxxxx@gitlab.test.com/monitor/test46.git C:/git/test46
## git pull
git --work-tree=C:/git/test46 --git-dir=C:/git/test46/.git pull http://oauth2:xxxxxx@gitlab.test.com/monitor/test46.git

# config
## gitlab access token
read_api, read_repository, read_registry
## repository
* test46
## directory
linux server repository directory

# build
```
cd C:/workplace/promci
set GOOS=linux
set GOARCH=amd64
go build promci
```

# config file
/etc/promci/promci.yml
# log file
/var/log/promci.log
# run
192.168.1.101
```
/app/promci &
```

# gitlab config
## webhook
http://192.168.1.101:8866/promci?repository=test46

Secret token

yyy-yyy-yyy-yyy

# curl
```
curl -H "X-Gitlab-Token: yyy-yyy-yyy-yyy" http://localhost:8866/promci?repository=test46
```
# test
update pc folder
```
cd C:/git/test46
git add .
git commit -m "kafka"
git push http://gitlab.test.com/monitor/test46.git
```
linux server result
```
tree /etc/test46/
```
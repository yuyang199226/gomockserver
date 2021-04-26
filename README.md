# gomockserver



```cassandraql

git clone git@github.com:yuyang199226/gomockserver.git
cd gomockserver
go build
./gomockserver --config mock.json


# 生成秘钥

```
openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt

```
```

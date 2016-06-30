https openssl generate crt and key
====
##服务器端私钥和证书,忽略校验 https_no_cert
###private key
- openssl genrsa -out server.key 2048   
用于生成服务端私钥文件server.key，后面的参数2048单位是bit，是私钥的长度

###server crt
- openssl req -new -x509 -key server.key -out server.crt -days 365
根据CA自己的私钥生成自签发的数字证书

###public key
- openssl rsa -in server.key -out server.key.public
openssl生成的私钥中包含了公钥的信息，我们可以根据私钥生成公钥：

### server.crt from server.key
- openssl req -new -x509 -key server.key -out server.crt -days 365
根据私钥直接生成自签发的数字证书

##对服务器端证书进行校验 verify_server_cert

###CA.key CA.crt
- openssl genrsa -out ca.key 2048
- openssl req -x509 -new -nodes -key ca.key -subj "/CN=golang.com" -days 5000 -out ca.crt
生成CA私钥和CA证书

### server.key server.csr
- openssl genrsa -out server.key 2048
- openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
生成server端的私钥，生成数字证书请求，并用我们的ca私钥签发server的数字证书


##对客户端证书进行校验 dual_verify_cert
###CA.key CA.crt
- openssl genrsa -out ca.key 2048
- openssl req -x509 -new -nodes -key ca.key -subj "/CN=golang.com" -days 5000 -out ca.crt
生成CA私钥和CA证书

###client.key client.csr client.crt
- openssl genrsa -out client.key 2048
- openssl req -new -key client.key -subj "/CN=golang.com" -out client.csr

- openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000

###client.crt show
- openssl x509 -text -in client.crt -noout

##License&CopyRight
tonybai [go-and-https](http://tonybai.com/2015/04/30/go-and-https/)

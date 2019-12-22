#!/bin/bash -ex

### ローカル環境を一気に整えるためのコマンドです。

# 全てのキャッシュをクリアする
# sudo /usr/local/go/bin/go clean -i -r -cache -testcache

# .bachrc に書いておくと良い
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
export GOROOT="/usr/local/go"

USERNAME=$(whoami)

## golang install
cd /tmp
sudo rm -rf $HOME/go
sudo rm -rf /usr/local/go
sudo rm -rf go
mkdir -p /tmp/ccc
cd /tmp/ccc

wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
echo '692d17071736f74be04a72a06dab9cac1cd759377bd85316e52b2227604c004c  go1.13.4.linux-amd64.tar.gz' | shasum -a256 -c
sudo tar -C /usr/local -xzf 'go1.13.4.linux-amd64.tar.gz'
go version

mkdir $HOME/go || true
sudo chown -R $USERNAME:$USERNAME $HOME/go
sudo chmod -R a=rX,u+w $HOME/go
cd ../
rm -rf /tmp/ccc

## grpc tools

# protoc-gen-go
cd /tmp
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get google.golang.org/grpc
go install github.com/golang/protobuf/protoc-gen-go

# protoc main
cd /tmp
mkdir protoc-dir && cd protoc-dir
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/protoc-3.10.1-linux-x86_64.zip
unzip protoc-3.10.1-linux-x86_64.zip
sudo mv ./bin/protoc /usr/bin/
protoc --version
rm protoc-3.10.1-linux-x86_64.zip

# protoc gen-doc
go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
# protoc --doc_out=/tmp/aaa --doc_opt=html,index.html `find proto -name "*.proto"`

# grpcurl
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
grpcurl --help

## static check tools

# golint install
go get golang.org/x/lint/golint

# staticchek install
cd /tmp
git clone https://github.com/dominikh/go-tools
cd go-tools/cmd/staticcheck
GOOS=linux GOARCH=amd64 go build -o staticcheck staticcheck.go
mv ./staticcheck $HOME/
cd /tmp
rm -rf /tmp/go-tools

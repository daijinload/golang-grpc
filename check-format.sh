#!/bin/bash -ex

DIR=`pwd`

## ソースコードビルドする方法（落としてきたバイナリで正しく動作しない場合に有効）
## goフォルダの権限で文句を言われたため、権限を変えるコマンドも追記します。自分のユーザ名に変えて実行して下さい。
if [[ ! -e "${HOME}/staticcheck" ]]; then
  cd /tmp
  git clone https://github.com/dominikh/go-tools
  cd go-tools/cmd/staticcheck
  GOOS=linux GOARCH=amd64 go build -o staticcheck staticcheck.go
  mv ./staticcheck $HOME/
  cd /tmp
  rm -rf /tmp/go-tools
fi

cd ${DIR}
cd ../$1

export PATH="$PATH:$HOME/go/bin"

# 基本入ってないので入れる
go get golang.org/x/lint/golint
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

go fmt main.go
go vet main.go
go vet -vettool=$(which shadow) main.go
golint main.go | grep -v "don't use MixedCaps" | grep -v "comment on exported type" | grep -v "should have comment" | grep -v "don't use ALL_CAPS in Go names" || true
${HOME}/staticcheck main.go

cd src
go fmt ./...
go vet ./...
go vet -vettool=$(which shadow) ./...
golint ./... | grep -v "don't use MixedCaps" | grep -v "comment on exported type" | grep -v "should have comment" | grep -v "don't use ALL_CAPS in Go names" || true
${HOME}/staticcheck ./...

echo '☆★☆★　check all green. success.　☆★☆★'

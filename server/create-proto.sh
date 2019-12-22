#!/bin/bash -ex

### 注）localのセットアップ(protocの実行など)がされていないと動きません。

# 通常はコマンド一発で出力できるはずですが、なぜかimportパスの指定が
# うまくいかないので小細工してます。
# go moduleへの移行期せいか、protoc側が対応していない？
# それか、モノレポ構成のためかもしれません。
# とりあえず、importパスがうまくいくように2個上のディレクトリを指定してアウトプットしてます。
# もし、↑の小細工が要らない方法があるのであれば、修正おねがいします。
DIR=`pwd`
cd $DIR/../../
OUT_DIR=`pwd`

# クリア（ハードコーディングです。。。）
rm -rf $DIR/src/infrastructure/grpc/gproto
mkdir -p $DIR/src/infrastructure/grpc/gproto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto
find . -name "*.proto" -exec protoc --proto_path . --go_out=plugins=grpc:$OUT_DIR {} \;

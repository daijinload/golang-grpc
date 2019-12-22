#!/bin/bash -ex

### help commands
# go help test
# go help testflag
# go tool cover -help

DIR_NAME=$1

# test実行時のゴミを掃除する
rm -rf /tmp/go-build* /tmp/p2-ap-test-*.sock

cd ../$DIR_NAME

# テスト実行（eオプションを付けているので、エラーがある場合、次のカバレッジ出力はされない）
go test -v -cover -count=1 -failfast p2-server/$DIR_NAME/src/...
echo '☆★☆★　test fail none　☆★☆★'

# coverage avarage output
# totalのカバレッジ率が出ないので、無理やりコマンドで平均値を出力しています
# 存在するテストと対応するソースコードの中での平均値なので、オールコードにするともっと低いっす。。
NUM=`go test -v -cover -count=1 -coverprofile=/tmp/p2-test-cover-$DIR_NAME.out -covermode=count p2-server/$DIR_NAME/src/... | egrep '^ok' | awk -F '\t' '{print $4}' | sed 's/\%//g' | awk '{sum+=$2} END {print sum/NR}'`
echo "☆★☆★　test coverage average=$NUM%　☆★☆★"

# coverage report
go tool cover -html=/tmp/p2-test-cover-$DIR_NAME.out -o /tmp/p2-test-cover-report-$DIR_NAME.html

echo "☆★☆★　output /tmp/p2-test-cover-report-$DIR_NAME.html　☆★☆★"

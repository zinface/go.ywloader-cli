# 默认将构建 ywloader
all:
	go build cli/ywloader.go

# make install 将会安装 ywloader 到 $GOPATH/bin 中
install:
	go install cli/ywloader.go

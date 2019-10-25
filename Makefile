#生成proto文件
# -I proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录
.PHONY:proto
proto:
	@echo "开始生成proto文件"
	@ protoc -I. \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --grpc-gateway_out=. \
        --go_out=plugins=grpc:.\
        --swagger_out=logtostderr=true:. \
        ./proto/*.proto
	@echo "生成proto文件结束"




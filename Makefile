swagger:
	swag init g main.go --output docs
	go mod vendor -v



#check_install:
#	which swagger || GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger
#swagger:
#	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
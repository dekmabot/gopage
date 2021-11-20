
build:
	@GOOS=linux go build -o ./bin/hello main.go

zip:
	@cd bin && zip hello.zip hello

deploy:
	@aws lambda update-function-code --function-name myFunc --zip-file fileb://./bin/hello.zip --region eu-west-1 --no-cli-pager

push: build zip deploy
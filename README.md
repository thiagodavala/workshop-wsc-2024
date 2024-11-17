# Workshop Worldskills Brasil
#53 Cloud Computing - 2024

## How to Build

```
go mod init thiagoogeremias.io/workshopcloud2024
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go build -o server .
```

## How to execute

```
export DYNAMO_TABLE="tabelateste"
export PORT="80"
```

## Runbook
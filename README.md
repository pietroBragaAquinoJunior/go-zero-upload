# go-zero-upload

Projeto backend em go-zero com endpoint de upload de arquivos para a amazon S3.


É necessário que você tenha criado um usuario com permissão para acessar o seu S3 da amazon.
Logue com o seu usuário:
```bash
aws configure
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
[default]
region = us-east-1
```

Mude o nome do bucket para o nome do seu bucket no arquivo zrpcmethodlogic.go
```go
bucket := "go-zero-upload-bucket"
```

Agore clone o repositório e rode os comandos abaixo (no terminal que você rodou o comando aws configure):
```bash
cd zrpc
go run zrpc.go
```

Comandos utilizados para gerar o código:
```bash
goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc
```

![Screenshot from 2024-08-01 17-18-55](https://github.com/user-attachments/assets/cf54549f-5f60-4833-a0e4-1afe157c38e6)

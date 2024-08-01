# go-zero-upload

Projeto backend em go-zero com endpoint de upload de arquivos.

canal: https://www.youtube.com/@pietrobraga523/videos

Para usá-lo basta clonar o repositório e rodar os comandos abaixo:

```bash
cd zrpc
go run zrpc.go
```
Comandos utilizados para gerar o código:
```bash
goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc
```
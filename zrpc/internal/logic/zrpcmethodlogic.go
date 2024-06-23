package logic

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zeromicro/go-zero/core/logx"
	__ "go-zero-upload/common/pb"
	"go-zero-upload/zrpc/internal/svc"
)

type ZrpcMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZrpcMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZrpcMethodLogic {
	return &ZrpcMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZrpcMethodLogic) ZrpcMethod(in *__.ZrpcMethodRequest) (*__.ZrpcMethodResponse, error) {

	// Decodifica a string Base64 para bytes
	data, err := base64.StdEncoding.DecodeString(in.Image)
	if err != nil {
		l.Logger.Errorf("Erro ao decodificar a imagem Base64: %v", err)
		return nil, err
	}

	// Carrega a configuração da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		l.Logger.Errorf("Erro ao carregar a configuração da AWS: %v", err)
		return nil, err
	}

	// Cria um cliente S3
	s3Client := s3.NewFromConfig(cfg)

	// Faz o upload do arquivo para o bucket S3
	bucket := "go-zero-upload-bucket"
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(in.Filename),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		l.Logger.Errorf("Erro ao fazer upload para o S3: %v", err)
		return nil, err
	}

	l.Logger.Infof("Imagem salva e enviada para o bucket S3 com sucesso: %s", in.Filename)

	return &__.ZrpcMethodResponse{}, nil
}

module github.com/videocoin/cloud-profiles

go 1.13

require (
	github.com/gogo/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jinzhu/gorm v1.9.12
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/videocoin/cloud-api v0.3.0
	github.com/videocoin/cloud-pkg v0.0.6
	google.golang.org/grpc v1.27.1
)

replace github.com/videocoin/cloud-api => ../cloud-api

replace github.com/videocoin/cloud-pkg => ../cloud-pkg

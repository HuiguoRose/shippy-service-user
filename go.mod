module github.com/HuiguoRose/shippy-service-user

go 1.13

replace sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.0.0
	github.com/micro/go-plugins/broker/nats/v2 v2.0.1
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb // indirect
	github.com/micro/protoc-gen-micro/v2 v2.0.0 // indirect
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
)

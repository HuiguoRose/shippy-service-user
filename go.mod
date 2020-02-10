module github.com/HuiguoRose/shippy-service-user

go 1.13

replace sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.0.0
	github.com/satori/go.uuid v1.2.0
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
)

module tape

go 1.14

require (
	github.com/Shopify/sarama v1.23.1 // indirect
	github.com/fsouza/go-dockerclient v1.7.1 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hyperledger/fabric v1.4.10
	github.com/hyperledger/fabric-amcl v0.0.0-20200128223036-d1aa2665426a // indirect
	//github.com/hyperledger/fabric-protos-go v0.0.0-20200424173316-dd554ba3746e
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta1.0.20200317135226-e71412ff3db1
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rogpeppe/godef v1.1.2 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.4.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/sykesm/zap-logfmt v0.0.2 // indirect

	// GM
	github.com/zhigui-projects/gm-crypto v0.0.0-20200719051209-13ea42f5b80c
	github.com/zhigui-projects/gm-go v0.0.0-20200510034956-8e4ef670d055
	github.com/zhigui-projects/gm-plugins v0.0.0-20200721031044-dc235c6ce0d5
	golang.org/dl v0.0.0-20210120004500-be2bfd84e4cf // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	golang.org/x/tools/gopls v0.6.1 // indirect
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

//replace github.com/zhigui-projects/fabric v1.4.6 => ../zhigui-projects/fabric

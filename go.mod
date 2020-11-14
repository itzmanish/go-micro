module github.com/itzmanish/go-micro/v2

go 1.13

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.8
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/bitly/go-simplejson v0.5.0
	github.com/bwmarrin/discordgo v0.20.2
	github.com/caddyserver/certmagic v0.10.6
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ef-ds/deque v1.0.4
	github.com/evanphx/json-patch/v5 v5.0.0
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fsouza/go-dockerclient v1.6.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-acme/lego/v3 v3.4.0
	github.com/go-git/go-git/v5 v5.1.0
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee
	github.com/gobwas/ws v1.0.3
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.2.0
	github.com/gorilla/handlers v1.4.2
	github.com/hashicorp/hcl v1.0.0
	github.com/hpcloud/tail v1.0.0
	github.com/imdario/mergo v0.3.11
	github.com/kr/pretty v0.2.1
	github.com/lib/pq v1.3.0
	github.com/lucas-clemente/quic-go v0.20.1
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.38
	github.com/mitchellh/hashstructure v1.0.0
	github.com/nats-io/nats.go v1.9.2
	github.com/nlopes/slack v0.6.1-0.20191106133607-d06c2a2b3249
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	go.etcd.io/bbolt v1.3.5
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/telegram-bot-api.v4 v4.6.4
)

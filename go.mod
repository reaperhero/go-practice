module github.com/reaperhero/go-practice

go 1.18

require (
	contrib.go.opencensus.io/exporter/graphite v0.0.0-20200424223504-26b90655e0ce
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	cuelang.org/go v0.6.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/ClickHouse/clickhouse-go/v2 v2.0.12
	github.com/Shopify/sarama v1.32.0
	github.com/aceld/zinx v1.0.1
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alibaba/sentinel-golang v1.0.4
	github.com/apolloconfig/agollo/v4 v4.1.0
	github.com/bitly/go-simplejson v0.5.0
	github.com/bluele/gcache v0.0.2
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d
	github.com/c-bata/go-prompt v0.2.6
	github.com/cep21/circuit/v3 v3.2.2
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28
	github.com/colinmarc/hdfs v1.1.3
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/creack/pty v1.1.18
	github.com/dtm-labs/dtmcli v1.12.1
	github.com/dustin/go-humanize v1.0.0
	github.com/elastic/go-ucfg v0.8.4
	github.com/fatih/color v1.15.0
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-echarts/go-echarts/v2 v2.2.4
	github.com/go-ini/ini v1.66.4
	github.com/go-ping/ping v1.1.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-resty/resty/v2 v2.8.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gocolly/colly/v2 v2.1.0
	github.com/gogf/gf v1.16.9
	github.com/golang-module/carbon/v2 v2.1.6
	github.com/golang/glog v1.0.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/google/gopacket v1.1.19
	github.com/google/uuid v1.6.0
	github.com/gookit/goutil v0.4.4
	github.com/hashicorp/consul/api v1.12.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/iawia002/lux v0.15.0
	github.com/igm/sockjs-go/v3 v3.0.2
	github.com/jakecoffman/cron v0.0.0-20190106200828-7e2009c226a5
	github.com/jessevdk/go-flags v1.5.0
	github.com/jinzhu/copier v0.3.5
	github.com/jinzhu/gorm v1.9.16
	github.com/jmoiron/sqlx v1.3.5
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/kataras/iris v0.0.2
	github.com/kevwan/mapreduce v1.1.1
	github.com/kyokomi/emoji/v2 v2.2.13
	github.com/labstack/gommon v0.4.2
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/looplab/fsm v0.3.0
	github.com/matcornic/hermes/v2 v2.1.0
	github.com/mattn/go-runewidth v0.0.15
	github.com/mingrammer/commonregex v1.0.1
	github.com/nats-io/nats.go v1.34.1
	github.com/nsqio/go-nsq v1.1.0
	github.com/olivere/elastic/v7 v7.0.31
	github.com/opentrx/seata-golang/v2 v2.0.7
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/orandin/lumberjackrus v1.0.1
	github.com/panjf2000/ants/v2 v2.4.8
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/sftp v1.13.1
	github.com/pkg/term v1.2.0-beta.2
	github.com/rivo/tview v0.0.0-20221128165837-db36428c92d9
	github.com/robfig/cron v1.2.0
	github.com/russross/blackfriday v1.6.0
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/sirupsen/logrus v1.9.3
	github.com/spacemonkeygo/openssl v0.0.0-20181017203307-c2dcc5cca94a
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/spf13/afero v1.8.2
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.13.0
	github.com/square/certstrap v1.2.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/buntdb v1.2.9
	github.com/tjfoc/gmsm v1.4.1
	github.com/txn2/txeh v1.3.0
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.3.0
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/vulcand/oxy v1.4.2
	github.com/wuYin/timewheel v0.0.0-20190326041048-6015c91b972f
	github.com/xanzy/go-gitlab v0.91.1
	go.etcd.io/etcd v3.3.27+incompatible
	go.opencensus.io v0.23.0
	go.uber.org/dig v1.14.0
	golang.org/x/crypto v0.22.0
	golang.org/x/net v0.24.0
	golang.org/x/sync v0.7.0
	golang.org/x/term v0.19.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.33.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.67.0
	gopkg.in/ldap.v2 v2.5.1
	k8s.io/api v0.28.3
	k8s.io/apimachinery v0.28.3
	k8s.io/klog/v2 v2.100.1
	xorm.io/xorm v1.2.5
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v4 v4.1.0 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/PuerkitoBio/goquery v1.8.1 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.1.8 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/aymerick/raymond v2.0.3-0.20180322193309-b565731e1464+incompatible // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cheggaaa/pb/v3 v3.0.8 // indirect
	github.com/cockroachdb/apd/v3 v3.2.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-redis/redis/v8 v8.11.4 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/goccy/go-json v0.7.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gookit/color v1.5.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190910122728-9d188e94fb99 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.2 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/huandu/xstrings v1.4.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/iris-contrib/jade v1.1.4 // indirect
	github.com/iris-contrib/pongo2 v0.0.1 // indirect
	github.com/iris-contrib/schema v0.0.6 // indirect
	github.com/jaytaylor/html2text v0.0.0-20180606194806-57d518f124b0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/blocks v0.0.8 // indirect
	github.com/kataras/golog v0.1.11 // indirect
	github.com/kataras/pio v0.0.13 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.26 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/gomega v1.27.10 // indirect
	github.com/paulmach/orb v0.4.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/segmentio/fasthash v1.0.3 // indirect
	github.com/shirou/gopsutil/v3 v3.24.3 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/tidwall/btree v1.1.0 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/urfave/cli v1.22.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vanng822/css v0.0.0-20190504095207-a21e860bcd04 // indirect
	github.com/vanng822/go-premailer v0.0.0-20191214114701-be27abe028fe // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.0 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.mongodb.org/mongo-driver v1.8.3 // indirect
	go.opentelemetry.io/otel v1.14.0 // indirect
	go.opentelemetry.io/otel/trace v1.14.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/exp v0.0.0-20240404231335-c0f41cb1a7a0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	gonum.org/v1/gonum v0.11.0 // indirect
	google.golang.org/api v0.81.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	xorm.io/builder v0.3.9 // indirect
	xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0

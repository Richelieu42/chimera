module github.com/richelieu-yang/chimera/v3

go 1.22

require (
	fyne.io/fyne/v2 v2.4.4
	github.com/alwindoss/morse v1.0.1
	github.com/andybalholm/brotli v1.1.0
	github.com/apache/pulsar-client-go v0.12.1
	github.com/apache/rocketmq-clients/golang/v5 v5.1.0-rc.1
	github.com/bytedance/sonic v1.11.3
	github.com/coocood/freecache v1.2.4
	github.com/dablelv/cyan v0.0.54
	github.com/davidbyttow/govips/v2 v2.14.0
	github.com/deckarep/golang-set/v2 v2.6.0
	github.com/duke-git/lancet/v2 v2.3.0
	github.com/emersion/go-imap v1.2.1
	github.com/emersion/go-imap-id v0.0.0-20190926060100-f94a56b9ecde
	github.com/emersion/go-message v0.18.0
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.7.0
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/gin-contrib/cors v1.7.1
	github.com/gin-contrib/gzip v1.0.0
	github.com/gin-contrib/pprof v1.4.0
	github.com/gin-contrib/size v1.0.0
	github.com/gin-contrib/sse v0.1.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-kratos/kratos/v2 v2.7.3
	github.com/go-oauth2/oauth2/v4 v4.5.2
	github.com/go-oauth2/redis/v4 v4.1.1
	github.com/go-playground/validator/v10 v10.19.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redsync/redsync/v4 v4.12.1
	github.com/goccy/go-json v0.10.2
	github.com/gogf/gf/v2 v2.6.4
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang-module/carbon/v2 v2.3.10
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/gorilla/securecookie v1.1.2
	github.com/gorilla/sessions v1.2.2
	github.com/gorilla/websocket v1.5.1
	github.com/h2non/bimg v1.1.9
	github.com/hashicorp/consul/api v1.28.2
	github.com/hashicorp/go-hclog v1.6.2
	github.com/hashicorp/go-version v1.6.0
	github.com/hashicorp/golang-lru/v2 v2.0.7
	github.com/hashicorp/raft v1.6.1
	github.com/hashicorp/raft-boltdb v0.0.0-20231211162105-6c830fa4535e
	github.com/hibiken/asynq v0.24.1
	github.com/imroc/req/v3 v3.43.1
	github.com/jackc/pgx/v5 v5.5.5
	github.com/jedib0t/go-pretty/v6 v6.5.6
	github.com/jinzhu/copier v0.4.0
	github.com/joho/godotenv v1.5.1
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/json-iterator/go v1.1.12
	github.com/klauspost/compress v1.17.7
	github.com/klauspost/cpuid/v2 v2.2.7
	github.com/linxGnu/grocksdb v1.8.14
	github.com/lionsoul2014/ip2region/binding/golang v0.0.0-20240322081424-40c7f53fc97c
	github.com/liuzl/gocc v0.0.0-20231231122217-0372e1059ca5
	github.com/magiconair/properties v1.8.7
	github.com/matoous/go-nanoid/v2 v2.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/modern-go/reflect2 v1.0.2
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.5
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/nicksnyder/go-i18n/v2 v2.4.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/otiai10/gosseract/v2 v2.4.1
	github.com/panjf2000/ants/v2 v2.9.1
	github.com/pelletier/go-toml/v2 v2.2.0
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
	github.com/redis/go-redis/v9 v9.5.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.5.0
	github.com/saintfish/chardet v0.0.0-20230101081208-5e3ef4b5456d
	github.com/samber/lo v1.39.0
	github.com/shirou/gopsutil/v3 v3.24.3
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.9.3
	github.com/sony/sonyflake v1.2.0
	github.com/spf13/afero v1.11.0
	github.com/spf13/cast v1.6.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	github.com/tidwall/buntdb v1.3.0
	github.com/tidwall/gjson v1.17.1
	github.com/unidoc/unioffice v1.30.0
	github.com/xuri/excelize/v2 v2.8.1
	github.com/zeromicro/go-zero v1.6.3
	github.com/zeromicro/zero-contrib/logx/logrusx v0.0.0-20231030135404-af9ae855016f
	go.etcd.io/etcd/client/v3 v3.5.13
	go.mongodb.org/mongo-driver v1.14.0
	go.opentelemetry.io/otel v1.24.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.24.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.24.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.24.0
	go.opentelemetry.io/otel/metric v1.24.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.opentelemetry.io/otel/trace v1.24.0
	go.uber.org/atomic v1.11.0
	go.uber.org/automaxprocs v1.5.3
	go.uber.org/zap v1.27.0
	golang.org/x/arch v0.7.0
	golang.org/x/crypto v0.21.0
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8
	golang.org/x/image v0.15.0
	golang.org/x/mobile v0.0.0-20240329164231-365c35338770
	golang.org/x/mod v0.16.0
	golang.org/x/net v0.22.0
	golang.org/x/oauth2 v0.18.0
	golang.org/x/sync v0.6.0
	golang.org/x/sys v0.18.0
	golang.org/x/term v0.18.0
	golang.org/x/text v0.14.0
	golang.org/x/time v0.5.0
	golang.org/x/tools v0.19.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/bigquery v1.2.0
	gorm.io/driver/clickhouse v0.6.0
	gorm.io/driver/mysql v1.5.6
	gorm.io/driver/postgres v1.5.7
	gorm.io/driver/sqlite v1.5.5
	gorm.io/driver/sqlserver v1.5.3
	gorm.io/gorm v1.25.9
	sigs.k8s.io/yaml v1.4.0
)

require (
	cloud.google.com/go v0.112.0 // indirect
	cloud.google.com/go/bigquery v1.58.0 // indirect
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.5 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.6.0 // indirect
	fyne.io/systray v1.10.1-0.20231115130155-104f5ef7839e // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/AthenZ/athenz v1.10.39 // indirect
	github.com/ClickHouse/ch-go v0.58.2 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.15.0 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/adamzy/cedar-go v0.0.0-20170805034717-80a9c64b256d // indirect
	github.com/alibabacloud-go/debug v0.0.0-20190504072949-9472017b5c68 // indirect
	github.com/alibabacloud-go/tea v1.1.17 // indirect
	github.com/alibabacloud-go/tea-utils v1.4.4 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1800 // indirect
	github.com/aliyun/alibabacloud-dkms-gcs-go-sdk v0.2.2 // indirect
	github.com/aliyun/alibabacloud-dkms-transfer-go-sdk v0.1.7 // indirect
	github.com/apache/arrow/go/v12 v12.0.1 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.4.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dchest/siphash v1.2.3 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21 // indirect
	github.com/emersion/go-textwrapper v0.0.0-20200911093747-65d896831594 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fredbi/uri v1.0.0 // indirect
	github.com/fyne-io/gl-js v0.0.0-20220119005834-d2da28d9ccfe // indirect
	github.com/fyne-io/glfw-js v0.0.0-20220120001248-ee7290d23504 // indirect
	github.com/fyne-io/image v0.0.0-20220602074514-4956b0afb3d2 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.6.1 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/go-text/render v0.0.0-20230619120952-35bccb6164b8 // indirect
	github.com/go-text/typesetting v0.1.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/flatbuffers v2.0.8+incompatible // indirect
	github.com/google/pprof v0.0.0-20240227163752-401108e1b7e7 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-msgpack/v2 v2.1.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jsummers/gobmp v0.0.0-20151104160322-e2ba15ffa76e // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/linkedin/goavro/v2 v2.9.8 // indirect
	github.com/liuzl/cedar-go v0.0.0-20170805034717-80a9c64b256d // indirect
	github.com/liuzl/da v0.0.0-20180704015230-14771aad5b1d // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/microsoft/go-mssqldb v1.6.0 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/onsi/ginkgo/v2 v2.16.0 // indirect
	github.com/paulmach/orb v0.10.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/quic-go v0.41.0 // indirect
	github.com/refraction-networking/utls v1.6.3 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c // indirect
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tevino/abool v1.2.0 // indirect
	github.com/tidwall/btree v1.4.2 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/xuri/efp v0.0.0-20231025114914-d1ff6096ae53 // indirect
	github.com/xuri/nfp v0.0.0-20230919160717-d98342af3f05 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/yuin/goldmark v1.5.5 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.13 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.13 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1 // indirect
	go.opentelemetry.io/proto/otlp v1.1.0 // indirect
	go.uber.org/mock v0.4.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.155.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/js/dom v0.0.0-20210725211120-f030747120f2 // indirect
)

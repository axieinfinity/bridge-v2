module github.com/axieinfinity/bridge-v2

replace github.com/ethereum/go-ethereum => github.com/axieinfinity/ronin v1.10.4-0.20230327040419-9bf4895fbd51

// replicating the replace directive on cosmos SDK
replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

go 1.18

require (
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/hashicorp/go-bexpr v0.1.10
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.16
	golang.org/x/crypto v0.6.0 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0
	gorm.io/gorm v1.24.1
)

require (
	github.com/deckarep/golang-set v1.8.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/axieinfinity/bridge-contracts v0.0.0-20230403071733-9fca121e6448
	github.com/axieinfinity/bridge-core v0.1.2-0.20230417033945-8258b81bcbe5
	github.com/axieinfinity/bridge-migrations v0.0.0-20230405070530-7bb3be2afe72
	github.com/axieinfinity/ronin-kms-client v0.0.0-20220805072849-960e04981b70
	github.com/ethereum/go-ethereum v1.10.26
	github.com/stretchr/testify v1.8.1
	gorm.io/driver/postgres v1.4.5
)

require (
	github.com/VictoriaMetrics/fastcache v1.12.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btcsuite/btcd v0.23.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/go-gormigrate/gormigrate/v2 v2.0.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/huin/goupnp v1.0.3 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/karalabe/usb v0.0.2 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/pointerstructure v1.2.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/peterh/liner v1.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/rjeczalik/notify v0.9.3 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/smartcontractkit/chainlink v1.13.3 // indirect
	github.com/sony/gobreaker v0.5.0 // indirect
	github.com/status-im/keycard-go v0.1.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220614013038-64ee5596c38a // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20221107162902-2d387536bcdd // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/sqlite v1.4.3 // indirect
)

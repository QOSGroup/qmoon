module github.com/QOSGroup/qmoon

require (
	cloud.google.com/go v0.36.0 // indirect
	github.com/QOSGroup/qbase v0.2.3-0.20190927065041-32eb90018d34
	github.com/QOSGroup/qos v0.0.8
	github.com/cosmos/cosmos-sdk v0.37.0
	github.com/didip/tollbooth v4.0.0+incompatible
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/gin-gonic/gin v1.9.0
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.1
	github.com/google/go-querystring v1.0.0
	github.com/hashicorp/go-version v1.1.0
	github.com/lib/pq v1.0.0
	github.com/magiconair/properties v1.8.0
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.32.2
	github.com/tidwall/gjson v1.1.5
	github.com/tidwall/match v1.0.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190330180304-aef51cc3777c

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190801041406-cbf593c0f2f3

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

go 1.13

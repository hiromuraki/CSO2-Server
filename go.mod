module github.com/KouKouChan/CSO2-Server

go 1.25.0

require (
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/willf/bitset v1.1.11
	golang.org/x/text v0.41.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.57.0
)

require (
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
)

replace golang.org/x/text v0.3.3 => github.com/golang/text v0.3.3

module go-commons/middleware/auth

go 1.14

replace go-commons/config => ../../config

replace go-commons/mvc/context => ../../mvc/context

replace go-commons/mvc/context/request => ../../mvc/context/request

replace go-commons/mvc/models => ../../mvc/models

replace go-commons/utils => ../../utils

replace go-commons/middleware/jwt => ../jwt

replace go-commons/middleware/casbin => ../casbin

replace go-commons/middleware/auth => ../auth

replace go-commons/exception => ../../exception

replace go-commons/utils/yaml => ../../utils/yaml

require (
	go-commons/config v0.0.0-00010101000000-000000000000
	go-commons/middleware/casbin v0.0.0-00010101000000-000000000000
	go-commons/middleware/jwt v0.0.0-00010101000000-000000000000
	go-commons/mvc/context v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/request v0.0.0-00010101000000-000000000000 // indirect
	go-commons/mvc/models v0.0.0-00010101000000-000000000000 // indirect
	go-commons/utils v0.0.0-00010101000000-000000000000
	github.com/ajg/form v1.5.1 // indirect
	github.com/casbin/casbin v1.9.1 // indirect
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
)

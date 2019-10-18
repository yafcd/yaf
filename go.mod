module yafcd.dev/yaf

go 1.13

require (
	github.com/gin-gonic/gin v1.4.0
	gopkg.in/src-d/go-git.v4 v4.13.1
	yafcd.dev/yaf/sync v0.0.0
)

replace yafcd.dev/yaf/sync => ./pkg/sync

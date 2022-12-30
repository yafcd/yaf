module yafcd.dev/yaf

go 1.13

require (
	github.com/gin-gonic/gin v1.7.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	gopkg.in/src-d/go-git.v4 v4.13.1 // indirect
	yafcd.dev/yaf/sync v0.0.0
)

replace yafcd.dev/yaf/sync => ./pkg/sync

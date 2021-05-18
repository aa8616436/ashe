package ashe

import (
	"github.com/helloteemo/ashe/cache"
	"github.com/helloteemo/ashe/core"
	"github.com/helloteemo/ashe/models"
	"github.com/helloteemo/ashe/stub"
	"github.com/helloteemo/ashe/unit"
	"testing"
)

// Ashe 主函数 其中 Core stub unit 是默认引入的，可以直接使用，但是db redis需要手动引入
type Ashe struct {
	// 核心模块，包括测试对象和断言函数
	*core.Core

	// fake redis
	*cache.MiniRedis

	// fake db
	*models.SqlMock

	// stub
	*stub.GoMonkey

	// simple unit test
	unit.GoConvey
}

func New(t *testing.T) *Ashe {
	c := &Ashe{}
	c.Core = core.New(t)
	c.GoConvey = unit.New(t)
	c.GoMonkey = stub.New(c.Core)
	return c
}

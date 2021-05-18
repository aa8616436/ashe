package unit

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// GoConvey 单元测试简化接口
type GoConvey interface {
	UnitTesting(testName string, f func())
	MiniUnitTesting(testName string, f func())
	SkipTesting(testName string, f func())
}

// goConvey go-convey的一点点封装
type goConvey struct {
	T *testing.T
	convey.C
}

// New new
func New(t *testing.T) GoConvey {
	return &goConvey{T: t}
}

// UnitTesting 对 Convey("test name" string,t *testing.Test,func()) 进行的封装
func (c *goConvey) UnitTesting(testName string, f func()) {
	convey.Convey(testName, c.T, f)
}

// MiniUnitTesting mini
func (c *goConvey) MiniUnitTesting(testName string, f func()) {
	convey.Convey(testName, f)
}

func (c *goConvey) SkipTesting(testName string, f func()) {
	convey.SkipConvey(testName, f)
}

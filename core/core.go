package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Core 核心单元，包括测试对象和断言函数
type Core struct {
	T *testing.T
	*assert.Assertions
}

// New 新建一个core单元
func New(t *testing.T) *Core {
	return NewWithAssert(t, assert.New(t))
}

// NewWithAssert 通过断言函数新建core单元
func NewWithAssert(t *testing.T, assertions *assert.Assertions) *Core {
	return &Core{T: t, Assertions: assertions}
}

// Assert 配置断言函数
func (a *Core) Assert(assertions *assert.Assertions) {
	a.Assertions = assertions
}

// GetAssert 获取断言函数
func (a *Core) GetAssert() *assert.Assertions {
	if a.Assertions == nil {
		a.Assertions = assert.New(a.T)
	}
	return a.Assertions
}

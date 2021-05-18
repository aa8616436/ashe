package ashe

import (
	"github.com/helloteemo/ashe/cache"
	"github.com/helloteemo/ashe/models"
)

type T int

const (
	Redis T = iota
	DB
)

func (c *Ashe) Use(tArr ...T) *Ashe {
	for _, t := range tArr {
		switch t {
		case Redis:
			c.MiniRedis = cache.New(c.Core)
		case DB:
			c.SqlMock = models.New(c.Core)
		}
	}
	return c
}

func (c *Ashe) Close() {
	if !c.IsEmptyMiniRedis() {
		c.MiniRedis.Close()
	}
	if !c.IsEmptySqlMock() {
		c.SqlMock.Close()
	}
}

func (c *Ashe) IsEmptyMiniRedis() bool {
	return c.MiniRedis == nil
}

func (c *Ashe) IsEmptySqlMock() bool {
	return c.Sqlmock == nil
}

func (c *Ashe) Drop(t T) *Ashe {
	switch t {
	case Redis:
		if c.MiniRedis != nil {
			c.MiniRedis.Close()
		}
		c.MiniRedis = nil
	case DB:
		if c.SqlMock != nil {
			c.SqlMock.Close()
		}
		c.SqlMock = nil
	}
	return c
}

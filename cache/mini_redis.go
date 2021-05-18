package cache

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/garyburd/redigo/redis"
	"github.com/helloteemo/ashe/core"
	"time"
)

// MiniRedis mini redis
type MiniRedis struct {
	*core.Core
	mr *miniredis.Miniredis
}

// New 新建一个MiniRedis测试
func New(t *core.Core) *MiniRedis {
	return &MiniRedis{Core: t}
}

func (m *MiniRedis) GetRedisDataController() *miniredis.Miniredis {
	return m.mr
}

// GetRedisConn 获取一个redis连接
func (m *MiniRedis) GetRedisConn() redis.Conn {
	m.begin()
	conn, err := getLocalRedisConn(m.mr.Addr())
	m.GetAssert().Nil(err)
	m.GetAssert().NotNil(conn)
	return conn
}

func (m *MiniRedis) begin() {
	m.checkMiniRedis()
}

// GetRedisPool 获取一个redis连接池
func (m *MiniRedis) GetRedisPool() *redis.Pool {
	m.begin()
	return &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", m.mr.Addr())
			m.Nil(err)
			m.NotNil(c)
			return c, nil
		},
	}
}

func (m *MiniRedis) checkMiniRedis() {
	if m.mr == nil {
		var err error
		m.mr, err = miniredis.Run()
		m.GetAssert().Nil(err)
	}
}

func getLocalRedisConn(addr string) (conn redis.Conn, err error) {
	conn, err = redis.Dial(`tcp`, addr)
	return
}

// Close 关闭
func (m *MiniRedis) Close() {
	if m.mr != nil {
		m.mr.Close()
	}
}

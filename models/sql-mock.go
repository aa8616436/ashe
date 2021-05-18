package models

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/helloteemo/ashe/core"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

// SqlMock sql mock
type SqlMock struct {
	*core.Core
	g      *gorm.DB
	x      *sqlx.DB
	mockDB *sql.DB
	sqlmock.Sqlmock
}

// New 新建一个MiniRedis测试
func New(t *core.Core) *SqlMock {
	mock := &SqlMock{Core: t}
	mock.init()
	return mock
}

func (m *SqlMock) init() {
	m.initDB()
	m.initGorm()
	m.initSQLX()
}

// GetGormConn 获取一个gorm连接
func (m *SqlMock) GetGormConn() *gorm.DB {
	return m.g
}

func (m *SqlMock) GetSqlXConn() *sqlx.DB {
	return m.x
}

func (m *SqlMock) GetMock() sqlmock.Sqlmock {
	return m.Sqlmock
}

func (m *SqlMock) Close() {
	if m.mockDB != nil {
		_ = m.g.Close()
		_ = m.x.Close()
		_ = m.mockDB.Close()
	}
}

func (m *SqlMock) initGorm() {
	var err error
	m.g, err = gorm.Open(`postgres`, m.mockDB)
	m.GetAssert().Nil(err)
	m.GetAssert().NotNil(m.g)
}

func (m *SqlMock) initDB() {
	var err error
	m.mockDB, m.Sqlmock, err = sqlmock.New()
	m.GetAssert().Nil(err)
	m.GetAssert().NotNil(m.mockDB)
	m.GetAssert().NotNil(m.Sqlmock)
}

func (m *SqlMock) initSQLX() {
	m.x = sqlx.NewDb(m.mockDB, `postgres`)
}

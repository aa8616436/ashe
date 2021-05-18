package example

import (
	"github.com/helloteemo/ashe"
	"regexp"
	"testing"
)

func TestFindByUser(t *testing.T) {
	cat := ashe.New(t).Use(ashe.DB)
	cat.UnitTesting(`TestFindByUser`, func() {
		db := cat.GetGormConn()

		rows := cat.NewRows([]string{"id", "user_name"}).AddRow(1, "Frank")
		cat.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(1).WillReturnRows(rows)

		user, err := FindByUser(db, 1)
		cat.Nil(err)
		cat.Equal(user, User{ID: 1, UserName: "Frank"})
	})
}

func TestFindByUserID(t *testing.T) {
	cat := ashe.New(t).Use(ashe.DB)
	cat.UnitTesting(`TestFindByUserID`, func() {
		db := cat.GetSqlXConn()

		rows := cat.NewRows([]string{"id", "user_name"}).AddRow(1, "Frank")
		cat.ExpectQuery(regexp.QuoteMeta(`select`)).WithArgs(1).WillReturnRows(rows)
		user, err := FindByUserID(db, 1)
		cat.Nil(err)
		cat.Equal(user, User{ID: 1, UserName: "Frank"})
	})
}

func TestFindByCache(t *testing.T) {
	cat := ashe.New(t).Use(ashe.Redis)

	cat.UnitTesting(`TestFindByCache`, func() {

		cat.MiniUnitTesting(`key存在`, func() {
			conn := cat.GetRedisConn()
			data := cat.GetRedisDataController()
			defer data.FlushDB() // 清除所有数据，避免有其它单元测试影响

			_ = data.Set(`username:1`, "FRANK")

			username, err := FindByCache(conn, 1)
			cat.Nil(err)
			cat.Equal(username, `FRANK`)
		})

		cat.MiniUnitTesting(`key不存在`, func() {

		})
	})
}

package example

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int    `db:"id"`
	UserName string `db:"user_name"`
}

// FindByUser find
func FindByUser(db *gorm.DB, id int) (u User, err error) {
	err = db.Table(`users`).Where(`id=?`, id).First(&u).Error
	return u, err
}

// FindByUserID find
func FindByUserID(db *sqlx.DB, id int) (u User, err error) {
	err = db.Unsafe().Get(&u, `select id,user_name from users where id =$1`, id)
	return u, err
}

// FindByCache find
func FindByCache(conn redis.Conn, id int) (string, error) {
	username, err := redis.String(conn.Do(`GET`, fmt.Sprintf(`username:%d`, id)))
	return username, err
}

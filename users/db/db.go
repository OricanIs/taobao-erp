package db

import (
	"database/sql"
	"errors"

	. "github.com/goushuyun/taobao-erp/db"
	"github.com/goushuyun/taobao-erp/pb"
	"github.com/wothing/log"
)

func Login(u *pb.User) error {
	query := "select id, name, role, extract(epoch from create_at)::bigint from users where mobile = $1 and password = $2"
	log.Debugf("select id, name, role, extract(epoch from create_at)::bigint from users where mobile = '%s' and password = '%s'", u.Mobile, u.Password)
	err := DB.QueryRow(query, u.Mobile, u.Password).Scan(&u.Id, &u.Name, &u.Role, &u.CreateAt)
	switch {
	case err == sql.ErrNoRows:
		return errors.New("not_found")
	default:
		return err
	}
}

func SaveUser(u *pb.User) error {
	query := "insert into users (mobile, password, name) values($1, $2, $3) returning id, extract(epoch from create_at)::bigint"

	log.Debugf("insert into users (mobile, password, name) values('%s', '%s', '%s') returning id, extract(epoch from create_at)::bigint", u.Mobile, u.Password, u.Name)
	err := DB.QueryRow(query, u.Mobile, u.Password, u.Name).Scan(&u.Id, &u.CreateAt)

	return err
}

func UserExist(u *pb.User) (bool, error) {
	query := "select count(*) from users where mobile = $1"
	var total int64
	err := DB.QueryRow(query, u.Mobile).Scan(&total)

	if err != nil {
		log.Error(err)
		return false, err
	}

	return total > 0, nil
}

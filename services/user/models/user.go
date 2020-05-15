package models

import (
	"github.com/vmihailenco/msgpack/v4"
	"github.com/weslenng/petssenger/services/user/config"
	"github.com/weslenng/petssenger/services/user/redis"
)

type Users struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func AuthUserByID(ID string) (*Users, error) {
	user := &Users{}

	val, err := redis.Client.Get(ID).Bytes()
	if err == nil {
		if err := msgpack.Unmarshal(val, user); err == nil {
			return user, nil
		}
	}

	if err := db.Model(user).Where("id = ?", ID).Select(); err != nil {
		return nil, err
	}

	val, err = msgpack.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = redis.Client.Set(ID, val, config.Default.RedisExpTime).Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(email string) (*Users, error) {
	user := &Users{Email: email}

	_, err := db.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

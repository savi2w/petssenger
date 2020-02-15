package models

import (
	"github.com/vmihailenco/msgpack/v4"
	"github.com/weslenng/petssenger/services/user/config"
	"github.com/weslenng/petssenger/services/user/redis"
)

// Users represents a user structure
type Users struct {
	ID        string
	Email     string
	createdAt string
}

// AuthUserByID retrieve an given user
func AuthUserByID(ID string) (*Users, error) {
	user := &Users{}

	val, err := redis.Client.Get(ID).Bytes()
	if err == nil {
		err = msgpack.Unmarshal(val, user)
		if err == nil {
			return user, nil
		}
	}

	err = db.Model(user).Where("id = ?", ID).Select()
	if err != nil {
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

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Store struct {
	Username string
	RedisClient *redis.Client
}
func (t *Store) Marshal() ([]byte, error) {
	client := t.RedisClient
	account,err:= client.Get(context.Background(),t.Username).Result()
	if err != nil {
		return nil,errors.Wrap(err,"Marshal account from redis")
	}
	return []byte(account),nil
}
func (t *Store) Unmarshal(bytes []byte) error {
	client := t.RedisClient
	err:= client.Set(context.Background(),t.Username,bytes,0).Err()
	if err != nil {
		return errors.Wrap(err,"unmarshall account to redis")
	}
	return nil
}

func (t *Store) Replace(cache cache.Unmarshaler, key string) {
	client := t.RedisClient
	err:= client.Se
	cache.Unmarshal(t)
	panic("implement me")
}

func (t *Store) Export(cache cache.Marshaler, key string) {
	panic("implement me")
}
func main(){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	store:=Store{
		Username:    "fadhilijuma@microsoft.com",
		RedisClient: rdb,
	}
	account:=public.Account{HomeAccountID: "fadhilijuma@microsoft.com"}
	b,err:=json.Marshal(account)
	if err != nil {
		panic(err)
	}

	err=store.Unmarshal(b)
	if err != nil {
		panic(err)
	}

	b,err=store.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}



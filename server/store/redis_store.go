package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisStore() *RedisStore {
	repo := &RedisStore{
		ctx: context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB),
		}),
	}

	return repo
}

func (r *RedisStore) GetMachines() ([]string, error) {
	cur := uint64(100)
	var machines []string
	for cur != 0 {
		m, c, err := r.client.Scan(r.ctx, 0, "machine:*", 10).Result()
		cur = c
		if err != nil {
			return machines, err
		}
		machines = append(machines, m...)
	}
	return machines, nil
}

func (r *RedisStore) SetGame(id string, data string) error {
	_, err := r.client.Set(r.ctx, fmt.Sprintf("game:%s", id), data, 0).Result()
	return err
}

func (r *RedisStore) GetGame(id string) (string, error) {
	return r.client.Get(r.ctx, fmt.Sprintf("game:%s", id)).Result()
}

func (r *RedisStore) SetUser(userID string, gameID string, data string) error {
	_, err := r.client.Set(r.ctx, fmt.Sprintf("game:%s:user:%s", gameID, userID), data, 0).Result()
	return err
}

func (r *RedisStore) GetUser(userID string, gameID string) (string, error) {
	return r.client.Get(r.ctx, fmt.Sprintf("game:%s:user:%s", gameID, userID)).Result()
}

func (r *RedisStore) GetGameUsers(id string) ([]string, error) {
	cur := uint64(100)
	var users []string
	for cur != 0 {
		m, c, err := r.client.Scan(r.ctx, 0, fmt.Sprintf("game:%s:user:*", id), 10).Result()
		cur = c
		if err != nil {
			return users, err
		}
		users = append(users, m...)
	}
	return users, nil
}

func (r *RedisStore) GetMachineUsers(id string) ([]string, error) {
	return r.client.SMembers(r.ctx, fmt.Sprintf("users:%s", id)).Result()
}

func (r *RedisStore) SetMachineUsers(id string, users []string) error {
	if len(users) == 0 {
		return nil
	}
	_, err := r.client.SAdd(r.ctx, fmt.Sprintf("users:%s", id), users).Result()
	return err
}

func (r *RedisStore) SetMachineServices(id string, services map[string]string) error {
	if len(services) == 0 {
		return nil
	}
	_, err := r.client.HSet(r.ctx, fmt.Sprintf("users:%s", id), services).Result()
	return err
}

func (r *RedisStore) IncMachineTicks(id string) (int64, error) {
	ticks, err := r.client.HIncrBy(r.ctx, fmt.Sprintf("machine:%s", id), "ticks", 1).Result()
	return ticks, err
}

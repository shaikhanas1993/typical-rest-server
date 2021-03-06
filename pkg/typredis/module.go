package typredis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Module of redis
func Module() *typapp.Module {
	return typapp.NewModule().
		Provide(typapp.NewConstructor(Connect)).
		Destroy(typapp.NewDestructor(Disconnect)).
		Prepare(typapp.NewPreparation(Ping)).
		Configure(typcfg.NewConfiguration(DefaultConfigName, DefaultConfig))
}

// Connect to redis server
func Connect(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:           cfg.Password,
		DB:                 cfg.DB,
		PoolSize:           cfg.PoolSize,
		DialTimeout:        cfg.DialTimeout,
		ReadTimeout:        cfg.ReadWriteTimeout,
		WriteTimeout:       cfg.ReadWriteTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
		MaxConnAge:         cfg.MaxConnAge,
	})
}

// Ping redis server
func Ping(client *redis.Client) (err error) {
	if err = client.Ping().Err(); err != nil {
		return fmt.Errorf("Redis: Ping: %w", err)
	}
	return
}

// Disconnect from service server
func Disconnect(client *redis.Client) (err error) {
	if err = client.Close(); err != nil {
		return fmt.Errorf("Redis: Disconnect: %w", err)
	}
	return
}

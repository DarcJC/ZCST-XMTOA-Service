package task

import (
	"github.com/go-redis/redis/v8"
	"github.com/mix-go/dotenv"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
)

var queueFactory = redisq.NewFactory()
var RedisClient *redis.Client
var MainQueue taskq.Queue

func init() {
	{
		// 从环境变量中加载redis配置
		redisAddr := dotenv.Getenv("REDIS_ADDR").String("localhost:6379")
		redisPwd := dotenv.Getenv("REDIS_PASSWORD").String("")
		redisDb := dotenv.Getenv("REDIS_DB").Int64(0)

		// 创建redis客户端
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPwd,
			DB:       int(redisDb),
		})
	}

	{
		// 创建任务队列
		MainQueue = queueFactory.RegisterQueue(&taskq.QueueOptions{
			Name:  "OA-Worker",
			Redis: RedisClient,
		})
	}

	MainQueue.String()
}

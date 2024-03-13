package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx     = context.Background()
	client  *redis.Client
	lockKey = "my_lock"
)

func main() {
	// 创建Redis连接
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // Redis密码，如果没有设置则为空
		DB:       0,                // 使用的Redis数据库
	})

	// 获取锁
	lock := acquireLock()
	if !lock.Val() {
		fmt.Println("获取锁失败")
		return
	}
	fmt.Println("获得锁成功： ", lock.Val())
	defer releaseLock(lock)
	go func() {
		// 获取锁
		lock1 := acquireLock()
		if !lock1.Val() {
			fmt.Println("获取锁失败")
			return
		}
	}()
	// 执行需要锁保护的操作
	fmt.Println("获得锁，执行操作")
	time.Sleep(5 * time.Second)

	fmt.Println("操作完成")
}

// 尝试获取锁
func acquireLock() *redis.BoolCmd {
	return client.SetNX(ctx, lockKey, "locked", 30*time.Second)
}

// 释放锁
func releaseLock(lock *redis.BoolCmd) {
	b, err := lock.Result()
	if err != nil {
		fmt.Println("释放锁失败:", err)
	}
	fmt.Println("释放锁", b)
}

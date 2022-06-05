package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/gorm"
	"math"
	"time"
)

func Pagination(value interface{}, meta *model.PaginationMetadata, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalItems int64
	db.Model(&value).Count(&totalItems)

	meta.TotalItem = totalItems
	totalPages := math.Ceil(float64(totalItems) / float64(meta.GetItemPerPage()))
	meta.TotalPage = int64(totalPages)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(meta.GetOffset())).Limit(int(meta.ItemsPerPage))
	}
}

func SaveCache(c *redis.Client, key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := json.Marshal(value)
	if err != nil {
		return
	}

	return c.Set(ctx, key, v, 5*time.Second).Err()
}

func GetCache(c *redis.Client, key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return
	}

	return json.Unmarshal([]byte(v), value)
}

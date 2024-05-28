package internal

import (
	"log"
	conf "stp/internal/pkg/config"
	"strconv"

	"github.com/Pacific73/gorm-cache/cache"
	"github.com/Pacific73/gorm-cache/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDB() {
	// Build DSN string for PostgreSQL connection
	dsn := "port=" + strconv.Itoa(conf.Config.Postgresql.Port) +
		" sslmode=disable" +
		" TimeZone=Asia/Shanghai" +
		" user=" + conf.Config.Postgresql.User +
		" password=" + conf.Config.Postgresql.Password +
		" host=" + conf.Config.Postgresql.Host +
		" dbname=" + conf.Config.Postgresql.Db

	// Open database connection with PostgreSQL
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true, TablePrefix: "baq_"},
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
	})

	// 数据库内存缓存
	cache, _ := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageMemory,
		Tables:               []string{"baq_record"},
		InvalidateWhenUpdate: true,
		CacheTTL:             0,
		CacheMaxItemCnt:      0,
		CacheSize:            16384,
	})
	DB.Use(cache)

	// Migrate struct model to database
	err := DB.AutoMigrate(&record{})
	if err != nil {
		log.Fatal(err)
	}
}

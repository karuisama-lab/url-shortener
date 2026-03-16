package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"runtime"
	"url-shortener/internal/config"
)

type Postgres struct {
	Cfg  *config.Config
	Pool *pgxpool.Pool
	Lg   *slog.Logger
}

func NewConnect(ctx context.Context, lg *slog.Logger, cfg *config.Config) (*Postgres, error) {
	const op = "postgres.NewConnect"
	lg.Debug(fmt.Sprintf("starting %s", op))

	// Формируем DSN
	dsn := makeDSN(&cfg.Postgres)

	// Парсим DSN в конфигурацию пула
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		lg.Error(fmt.Sprintf("%s: failed parse config, error: %v", op, err))
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}

	// Настраиваем пул (пример, добавь свои параметры из cfg)
	maxConns := calculatePoolSize(cfg)
	poolCfg.MaxConns = maxConns
	poolCfg.MinConns = cfg.Postgres.MinConns
	poolCfg.HealthCheckPeriod = cfg.Postgres.HealthCheckPeriod
	poolCfg.MaxConnIdleTime = cfg.Postgres.MaxIdleTime
	poolCfg.MaxConnLifetime = cfg.Postgres.MaxLifetime
	poolCfg.ConnConfig.ConnectTimeout = cfg.Postgres.ConnectTimeout

	// Создаём пул соединений
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		lg.Error(fmt.Sprintf("%s: failed to create pool, error: %s", op, err.Error()))
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	// Проверяем соединение (ping)
	if err := pool.Ping(ctx); err != nil {
		lg.Error(fmt.Sprintf("%s: failed to ping database, error: %s", op, err.Error()))
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	lg.Debug(fmt.Sprintf("%s: connected to database", op))

	return &Postgres{
		Cfg:  cfg,
		Pool: pool,
		Lg:   lg,
	}, nil
}

// makeDSN собирает строку подключения из полей структуры Postgres если это поле есть в yaml-файле
func makeDSN(pg *config.Postgres) string {
	if pg.DSN != "" {
		return pg.DSN
	}
	// Стандартный формат DSN для PostgreSQL если поле dsn в yaml-файле не заполнено
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		pg.Host, pg.Port, pg.DbName, pg.Login, pg.Password, pg.SslMode)
}

// calculatePoolSize вычисляет размер пула в зависимости от окружения
func calculatePoolSize(cfg *config.Config) int32 {
	if cfg.Postgres.MaxConns != nil {
		return *cfg.Postgres.MaxConns
	}
	cpuCount := runtime.NumCPU()
	switch cfg.Env {
	case "local", "dev":
		return int32(3*cpuCount + 10)
	case "prod":
		return int32(cpuCount + 5)
	default:
		return int32(cpuCount + 5)
	}
}

// Close закрывает пул соединений (можно вызвать при завершении приложения)
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func calculatedPoolSize(cfg *config.Config) int32 {
	if cfg.MaxConns != nil {
		return *cfg.MaxConns
	}
	cpuCount := runtime.NumCPU()
	switch cfg.Env {
	case "local", "dev":
		return int32(3*cpuCount + 10)
	case "prod":
		return int32(cpuCount + 5)
	default:
		return int32(cpuCount + 5)
	}
}

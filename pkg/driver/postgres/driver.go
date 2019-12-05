package postgres

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"

	// Load postgres lib
	_ "github.com/lib/pq"
)

const (
	MaxOpenConnections    = 100
	MaxIdleConnections    = 100
	MaxConnectionLifetime = 100
)

type Config struct {
	Name     string
	User     string
	Password string
	Hostname string
	Port     string
}

func cnxString(name, user, password, hostname, port string) string {
	s := fmt.Sprintf(
		"dbname=%v user=%v sslmode=disable",
		name,
		user,
	)

	if password != "" {
		s = s + fmt.Sprintf(" password=%v", password)
	}
	if hostname != "" {
		s = s + fmt.Sprintf(" host=%v", hostname)
	}
	if port != "" {
		s = s + fmt.Sprintf(" port=%v", port)
	}

	return s
}

// New .
func New(c *Config) (*xorm.Engine, error) {
	if c == nil {
		return nil, fmt.Errorf("nil configuration")
	}

	str := cnxString(c.Name, c.User, c.Password, c.Hostname, c.Port)
	engine, err := xorm.NewEngine("postgres", str)
	if err != nil {
		return nil, err
	}

	engine.SetMaxOpenConns(MaxOpenConnections)
	engine.SetMaxIdleConns(MaxIdleConnections)
	engine.SetConnMaxLifetime(time.Millisecond * time.Duration(MaxConnectionLifetime))
	return engine, nil
}

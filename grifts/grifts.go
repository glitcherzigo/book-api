package grifts

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/edgedb/edgedb-go"
	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func EdgedbClient() *edgedb.Client {
	ctx := context.Background()

	opts := edgedb.Options{}

	client, err := edgedb.CreateClient(ctx, opts)
	if err != nil {
		log.Panic(err)
	}

	return client
}

func GetConfig(path string) *Config {
	var cfg = new(Config)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil
	}

	return cfg
}

func Settings() *atreugo.Config {
	cfg := GetConfig("../../config/config.yaml")
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)

	return &atreugo.Config{
		Addr: addr,
	}
}

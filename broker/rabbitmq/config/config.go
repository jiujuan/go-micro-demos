package config

import (
    "github.com/BurntSushi/toml"
    "path"
    "sync"
    "runtime"
)

type tomlConfig struct {
    Title    string             `toml:"title"`
    Topics   map[string]topic   `toml:"topics"`
    Rabbitmq rabbitmq            `toml:"rabbitmq"`
    Services map[string]service `toml:"services"`
}

type topic struct {
    Name string
}

type rabbitmq struct {
    Amqp string
    Server string
    Port []int
    Username string
    Password string
}

type service struct {
    Name string
    Info string
}

var (
    cfg *tomlConfig
    once sync.Once
)

func Config() *tomlConfig {
    once.Do(func() {
        curPath := getCurrentPath()
        if _, err := toml.DecodeFile(curPath+"/config.toml", &cfg);err != nil {
            panic(err)
        }
    })
    return cfg
}

func getCurrentPath() string {
    _, filename, _, _ := runtime.Caller(1)
    return path.Dir(filename)
}
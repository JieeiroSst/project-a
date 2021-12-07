package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type Config struct {
	Server          ServerConfig
	Mysql           MysqlConfig
	Secret 		    SecretService
	Constant	    Constant
	RabbitMQ        RabbitMQ
	Redis			Redis
	Twilio          TwilioService
	CasbinMysql     CasbinMysql
	AmazonS3        AmazonS3
	Prometheus      Prometheus
}

type ServerConfig struct {
	PortServer    string
	PprofPort     string
}

type Redis struct {
	Dns           string
}

type RabbitMQ struct {
	Host           string
	Port           string
	User           string
	Password       string
	Exchange       string
	Queue          string
	RoutingKey     string
	ConsumerTag    string
	WorkerPoolSize int
}


type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlSSLMode  bool
	MysqlDriver   string
}

type CasbinMysql struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlSSLMode  bool
	MysqlDriver   string
}

type SecretService struct {
	JwtSecretKey string
}

type Constant struct {
	Rbac          string
}

type TwilioService struct {
	TwilioAccountSid  string
	TwilioAuthToken   string
	TwilioPhoneNumber string
}

type AmazonS3 struct {
	S3Region string
	S3Bucket string
	S3ACL   string
}

type Prometheus struct {
	PrometheusPushgateway string
}

func ReadConf(filename string) (*Config, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	return config, nil
}

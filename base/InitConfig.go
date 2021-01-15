package base

import (
	"database/sql"
	"github.com/beanstalkd/go-beanstalk"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var Conf *Config

type Yaml struct {
	Mysql      Mysql      `yaml:"mysql"`
	BeanStalkd BeanStalkd `yaml:"beanstalkd"`
	Dun        Dun        `yaml:"dun"`
	ShuMei     ShuMei     `yaml:"shumei"`
	Logger     Logger     `yaml:"logger"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type BeanStalkd struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	NetWork string `yaml:"network"`
}

type Dun struct {
	SecretKey             string `yaml:"secret_key"`
	SecretId              string `yaml:"secret_id"`
	BusinessId            string `yaml:"business_id"`
	MomentImageBusinessId string `yaml:"moment_image_business_id"`
	ContentUrl            string `yaml:"content_url"`
	ContentVersion        string `yaml:"content_version"`
	ImageUrl              string `yaml:"image_url"`
	ImageVersion          string `yaml:"image_version"`
}

type ShuMei struct {
	AccessKey string `yaml:"access_key"`
}

type Logger struct {
	AccessLog string `yaml:"access_log"`
	ErrorLog  string `yaml:"error_log"`
}

type Config struct {
	Mysql  *sql.DB
	Bean   *beanstalk.Conn
	Dun    Dun
	ShuMei ShuMei
	Logger Logger
}

func loadYaml(path string) (*Yaml, error) {
	conf := &Yaml{}
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(file).Decode(conf)
	}
	return conf, nil
}

func initMysql(mysql Mysql) (*sql.DB, error) {
	path := strings.Join([]string{mysql.Username, ":", mysql.Password, "@tcp(", mysql.Host, ":", mysql.Port, ")/", mysql.Dbname, "?charset=utf8"}, "")
	DB, _ := sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		return nil, err
	}
	return DB, nil
}

func initDun(dun Dun) Dun {
	return dun
}

func initShuMei(shuMei ShuMei) ShuMei {
	return shuMei
}

func initLogger(logger Logger) Logger {
	return logger
}

func initBeanStalkd(bs BeanStalkd) (*beanstalk.Conn, error) {
	c, err := beanstalk.Dial(bs.NetWork, strings.Join([]string{bs.Host, ":", bs.Port}, ""))
	return c, err
}

func Init(yamlPath string) error {
	yaml, err := loadYaml(yamlPath)
	if err != nil {
		return err
	}
	Db, err := initMysql(yaml.Mysql)
	Bs, err := initBeanStalkd(yaml.BeanStalkd)
	Ns := initDun(yaml.Dun)
	Sm := initShuMei(yaml.ShuMei)
	logger := initLogger(yaml.Logger)
	initXLog(logger.AccessLog, logger.ErrorLog)

	conf := &Config{}
	conf.Mysql = Db
	conf.Bean = Bs
	conf.Dun = Ns
	conf.ShuMei = Sm
	conf.Logger = logger
	Conf = conf
	return nil
}

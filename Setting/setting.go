package Setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Appconfig)

type Appconfig struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	Version    string `mapstructure:"version"`
	StartTime  string `mapstructure:"startTime"`
	MachingeID int64  `mapstructure:"MachingeID"`
	Port       int    `mapstructure:"port"`

	*Logconfig       `mapstructure:"log"`
	*MysqlConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*ElaSearchConfig `mapstructure:"elasticSearch"`
}
type Logconfig struct {
	Level       string `mapstructure:"level"`
	Filename    string `mapstructure:"filename"`
	ErrFilename string `mapstructure:"errfilename"`
	Maxsize     int    `mapstructure:"max_size"`
	Maxage      int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         string `mapstructure:"port"`
	MaxOpenConns string `mapstructure:"max_open_conns"`
	MaxIdleConns string `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type ElaSearchConfig struct {
	Host string `mapstructure:"host"`
}

func Init() (err error) {
	viper.SetConfigFile("conf.yaml")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./Config")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed ,err:%v\n", err)
	}
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal() failed,err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config conf.yaml Change ,Starting")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal() failed,err:%v\n", err)
		}
	})
	return

}

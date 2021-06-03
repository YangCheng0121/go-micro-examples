package config

import "github.com/asim/go-micro/v3/config"

// MysqlConfig 创建结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// GetMysqlFromConsul 获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}
	//获取配置
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		return nil, err
	}
	return mysqlConfig, nil
}
package config

import "os"

func MysqlHost() string {
	return os.Getenv("MYSQL_HOST")
}

func MysqlUser() string {
	return os.Getenv("MYSQL_USER")
}

func MysqlPass() string {
	return os.Getenv("MYSQL_PASSWORD")
}

func MysqlDB() string {
	return os.Getenv("MYSQL_DB")
}

func JwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

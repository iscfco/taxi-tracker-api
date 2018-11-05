package dbconn

import (
	"database/sql"
	"fmt"
	"gbmchallenge/api/config"
	_ "github.com/lib/pq"
	"os"
)

func GetPsqlDBConn() (*sql.DB, error) {
	driverName := "postgres"
	format := "user=%s password=%s host=%s dbname=%s sslmode=disable"
	dbConnConfig := GetPsqlDbConfig()
	dataSourceName := fmt.Sprintf(
		format,
		dbConnConfig.User,
		dbConnConfig.Password,
		dbConnConfig.Server,
		dbConnConfig.DbName,
	)
	fmt.Println("CurrentUser:", os.Getenv("USER"))
	fmt.Println(dataSourceName)
	return getDBConnection(&driverName, &dataSourceName)
}

func GetPsqlDbConfig() ServerConfig {
	switch config.DbConnEnv {
	case config.Pro:
		return ServerConfig{
			Server:   os.Getenv("PSQL_HOST"),
			DbName:     os.Getenv("PSQL_DBNAME"),
			User:     os.Getenv("PSQL_USER"),
			Password: os.Getenv("PSQL_PWD"),
		}
	case config.Loc:
		return ServerConfig{
			Server:   "localhost",
			DbName:     "taxi_traker",
			User:     "postgres",
			Password: "12345Ab...",
		}
	}
	return ServerConfig{}
}

func getDBConnection(driverName, dataSourceName *string) (*sql.DB, error) {
	db, err := sql.Open(*driverName, *dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

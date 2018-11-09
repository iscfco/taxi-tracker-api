package dbconn

import (
	"database/sql"
	"fmt"
	"taxi-tracker-api/api/config"
	_ "github.com/lib/pq"
	"os"
)

func GetPsqlDBConn() (*sql.DB, error) {
	driverName := "postgres"
	format := "user=%s password=%s host=%s dbname=%s sslmode=disable"
	dbConnConfig := getPsqlDbConfig()
	dataSourceName := fmt.Sprintf(
		format,
		dbConnConfig.User,
		dbConnConfig.Password,
		dbConnConfig.Server,
		dbConnConfig.DbName,
	)
	return getDBConnection(&driverName, &dataSourceName)
}

func getPsqlDbConfig() ServerConfig {
	switch config.DbConnEnv {
	case config.Production:
		return ServerConfig{
			Server:   os.Getenv("PSQL_HOST"),
			DbName:   os.Getenv("PSQL_DBNAME"),
			User:     os.Getenv("PSQL_USER"),
			Password: os.Getenv("PSQL_PWD"),
		}
	case config.Local:
		return ServerConfig{
			Server:   "localhost",
			DbName:   "taxi_traker",
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

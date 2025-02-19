package configuration

import "fmt"

const(
	//AWS RDS for PostgreSQL
	POSTGRES_AWS = "pg_aws"
	//GCP Cloud SQL for PostgreSQL
	POSTGRES_GCP = "pg_gcp"
	//On-Premise (or locally hosted)
	POSTGRES_ON_PREMISE = "pg_local"
)

type WebAppDBConfig struct {
	ContextName string
	Provider string
	URL string
	User string
	Password string
	DatabaseName string
	CreateConnectionTimeout int
}


func (wdb WebAppDBConfig) String() string {
	return fmt.Sprintf("ContextName: %s, Provider: %s, URL: %s, User: %s, Password: %s, DatabaseName: %s, CreateConnectionTimeout: %d",
		wdb.ContextName, wdb.Provider,
		wdb.URL, wdb.User, wdb.Password, wdb.DatabaseName, wdb.CreateConnectionTimeout)

}

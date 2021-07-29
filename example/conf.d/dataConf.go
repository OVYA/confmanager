package main

// struct of clog app configuration
type AppConfStruct struct {
	Dbs Dbs `json:"databases"`
}

// type of application databases config
type Dbs map[string]*Db

// struct of one database config
type Db struct {
	MigrDir            string `json:"migrDir"` // optionnel
	DriverName         string `json:"driverName"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	DbName             string `json:"dbName"`
	User               string `json:"user"`
	Password           string `json:"password"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	SSLMode            bool   `json:"SSLMode"`
}

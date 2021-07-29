package main

// initialisation of configuration
func initConfig() interface{} {

	conf := &AppConfStruct{
		Dbs: map[string]*Db{
			"clog": {
				MigrDir:    "",
				DriverName: "postgres",
			},
		},
	}

	return conf
}

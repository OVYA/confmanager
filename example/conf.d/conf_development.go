package main

func init() {

	conf := AppConf.(*AppConfStruct)

	conf.Dbs["clog"].DriverName = "driver3"
	conf.Dbs["clog"].MigrDir = "clog/migdir"

}

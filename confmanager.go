package confmanager

import (
	"bytes"
	"fmt"
	"github.com/OVYA/confmanager/app/log"
	"github.com/OVYA/confmanager/app/util"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	ENV_APP_ROOT_PATH = "APP_ROOT_PATH"
	ENV_APP_ENV       = "APP_ENV"

	CONF_DIRECTORY_NAME  = "conf.d"
	GONAME_GENERATE_CONF = "generateConf.go"
	GONAME_DATA_CONF     = "dataConf.go"
	GONAME_CONF_APP      = "conf_app.go"

	MSG_KO = "Config file app.json is not generated."
)

// init configuration application
// return false if initialization failed
// param configDirectory : the path of configuration directory, if it is null then configDirectory = ${APP_ROOT_PATh}/conf.d
func Init(configDirectory *string) bool {

	// check configuration directory
	confDirRelPath, confDirFullPath := checkConfigDirectory(configDirectory)
	if confDirRelPath == nil {
		return false
	}

	// check generateConf.go
	if !checkGoGenerateConf(confDirRelPath) {
		return false
	}

	// check dataConf.go
	if !checkGoDataConf(confDirRelPath) {
		return false
	}

	// check conf_app.go
	if !checkGoConfApp(confDirRelPath) {
		return false
	}

	// check conf_${APP_ENV}.go
	existEnvConf, envConfFileName := checkGoConfEnv(confDirRelPath)

	// run generate conf
	runGenerateConf(existEnvConf, envConfFileName, confDirRelPath, confDirFullPath)

	return true
}

// return the configuration directory
func GetConfigDirectory() string {
	return os.Getenv(ENV_APP_ROOT_PATH) + string(os.PathSeparator) + CONF_DIRECTORY_NAME
}

// ~  private ----------------------------------------------------------------------------------------------------------

// check config directory, return its relative path and fullPath if the directory exist
func checkConfigDirectory(configDirectory *string) (*string, *string) {

	var confDirRelPath string

	if configDirectory == nil {
		confDirRelPath = GetConfigDirectory()
	} else {
		confDirRelPath = *configDirectory
	}

	exist, confDirFullPath := checkExist(confDirRelPath, util.IsExistDir)

	if !exist {

		log.Info.Println("Not found config directory : " + confDirFullPath)
		log.Info.Println(`=> check APP_ROOT_PATH env variable. APP_ROOT_PATH="` + os.Getenv(ENV_APP_ROOT_PATH) + `"`)
		log.Info.Println(MSG_KO)

		return nil, nil
	}

	log.Info.Println("Found config directory : " + confDirFullPath)

	return &confDirRelPath, &confDirFullPath
}

// check if generateConf.go file exists
func checkGoGenerateConf(confDir *string) bool {

	generateConfRelPath := *confDir + string(os.PathSeparator) + GONAME_GENERATE_CONF
	exist, fullPath := checkExist(generateConfRelPath, util.IsExistFile)

	if !exist {

		log.Info.Println("Not found in conf directory : " + fullPath)
		log.Info.Println(MSG_KO)

		return false
	}

	log.Info.Println("Found in conf directory : " + fullPath)

	return true
}

// check if conf_app.go file exists
func checkGoConfApp(confDir *string) bool {

	appConfRelPath := *confDir + string(os.PathSeparator) + GONAME_CONF_APP
	exist, fullPath := checkExist(appConfRelPath, util.IsExistFile)

	if !exist {

		log.Info.Println("Not found in conf directory : " + fullPath)
		log.Info.Println(MSG_KO)

		return false
	}

	log.Info.Println("Found in conf directory : " + fullPath)

	return true
}

// check if dataConf.go file exists
func checkGoDataConf(confDir *string) bool {

	dataConfRelPath := *confDir + string(os.PathSeparator) + GONAME_DATA_CONF
	exist, fullPath := checkExist(dataConfRelPath, util.IsExistFile)

	if !exist {

		log.Info.Println("Not found in conf directory : " + fullPath)
		log.Info.Println(MSG_KO)

		return false
	}

	log.Info.Println("Found in conf directory : " + fullPath)

	return true
}

// check conf_${APP_ENV}.go : return true if overload conf found
func checkGoConfEnv(confDir *string) (bool, string) {

	envConfFileName := "conf_" + os.Getenv(ENV_APP_ENV) + ".go"
	envConfRelPath := *confDir + string(os.PathSeparator) + envConfFileName

	existEnvConf, fullPath := checkExist(envConfRelPath, util.IsExistFile)

	if !existEnvConf {

		log.Info.Println("No configuration overload found : " + fullPath + " doesn't exist.")
		log.Info.Println(`If you want overload configuration, check APP_ENV variable. APP_ENV="` + os.Getenv(ENV_APP_ENV) + `"`)

	} else {

		log.Info.Println("Found configuration overlaod : " + fullPath)
	}

	return existEnvConf, envConfFileName
}

// run generateConf.go
func runGenerateConf(existEnvConf bool, envConfFileName string, confDirRelPath *string, confDirFullPath *string) bool {

	var cmd *exec.Cmd
	if !existEnvConf {
		cmd = exec.Command("go", "run", GONAME_GENERATE_CONF, GONAME_DATA_CONF, GONAME_CONF_APP)
	} else {
		cmd = exec.Command("go", "run", GONAME_GENERATE_CONF, GONAME_DATA_CONF, GONAME_CONF_APP, envConfFileName)
	}
	cmd.Dir = *confDirRelPath

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		log.Info.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}

	log.Info.Println("Configuration file app.json generated in " + *confDirFullPath)

	return true
}

// check if the path exist, exist test handle by isExit function
func checkExist(path string, isExist func(string) bool) (bool, string) {

	fullPath, err := filepath.Abs(path)

	if err != nil {
		log.Info.Println(err.Error())
		return false, fullPath
	}

	if isExist(fullPath) {
		return true, fullPath
	} else {
		return false, fullPath
	}
}

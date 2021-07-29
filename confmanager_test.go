package confmanager

import (
	"os"
	"testing"
)

func init() {
	os.Setenv(ENV_APP_ROOT_PATH, "example")
	os.Setenv(ENV_APP_ENV, "development")
}

func TestConf1(t *testing.T) {

	var confPath = "example/conf.d"

	if !Init(&confPath) {
		t.Fail()
	}
}

func TestConf2(t *testing.T) {

	if !Init(nil) {
		t.Fail()
	}
}

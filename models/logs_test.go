package models

import(
	"testing"
)

func TestLog(t *testing.T){
	log := Log{ Id:1, Log_level:"Debug" }
	if log.Id != 1{
		t.Errorf("Expected Id to be 1 got %d",log.Id)
	}
}

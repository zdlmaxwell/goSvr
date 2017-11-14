package logger

import (
	"log"
	"fmt"
	"testing"
)


func TestLog(t *testing.T) {
	svrLog, err := NewLogger("debug", "debug", log.Lmicroseconds, MB * 10)
	if err != nil {
		fmt.Printf("========err:%v\n", err)
	}
	defer svrLog.Close()	
	for i:= 0; i < 10; i ++{
		svrLog.Debug("+++will not print222============")
	}
}

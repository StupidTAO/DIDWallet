package log

import "testing"

func TestLogInit(t *testing.T) {
	err := LogInit()
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestInfo(t *testing.T) {
	LogInit()
	Info("%s", "caohaitao")
}

func TestWarning(t *testing.T) {
	LogInit()
	Warning("%s", "caohaitao")
}

func TestDebug(t *testing.T) {
	LogInit()
	Debug("%s", "caohaitao")
}

func TestError(t *testing.T) {
	LogInit()
	Error("%s", "caohaitao")
}

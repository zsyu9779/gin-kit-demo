package util

import (
	"fmt"
	"testing"
)


func TestFormattedTimeNow(t *testing.T) {
	fmt.Println(FormattedTimeNow())
	fmt.Println(FormattedTimeNow().Format("2006-01-02 15:04:05.000"))
	t.Log(FormattedTimeNow().Format("2006-01-02 15:04:05.000"))
	t.Log(FormattedTimeNow())
}
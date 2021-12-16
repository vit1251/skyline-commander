package util

import "testing"

func TestGetStatus(t *testing.T) {
	var myStatFs MyStatFs

	err := GetStatus(&myStatFs, "/")
	t.Logf("myStatFs = %+v err = %+v", myStatFs, err)

}

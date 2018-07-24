package go_get_ip

import "testing"

func TestGetIPV4(t *testing.T) {
	_, err := GetIPV4()
	if err != nil {
		t.Error(err)
	}
	//Not going to log the ip in the tests as I dont know where the test logs may end up
}

func TestGetIPV6(t *testing.T) {
	_, err := GetIPV6()
	if err != nil {
		t.Error(err)
	}
	//Not going to log the ip in the tests as I dont know where the test logs may end up
}
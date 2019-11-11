package cuv

import "testing"

func TestIP(t *testing.T) {
	ip := LocalIP()

	t.Log("localip: ", ip)
}

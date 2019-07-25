package ontap

import (
	"fmt"
	"testing"
)

func TestCluster(t *testing.T) {
	fmt.Println("Test: Get Cluster Info")

	c, err := NewClient("apiuser", "foobar11", "10.10.10.111", true, true)
	if err != nil {
		t.Error(err)
	}

	data, err := c.GetCluster()
	if err != nil {
		t.Error(err)
	}

	if data.Version.Generation != 9 {
		t.Error("Major Version not 9 as expected")
	}

}


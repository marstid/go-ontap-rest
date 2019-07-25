package ontap

import (
	"fmt"
	"testing"
)

func TestSvm(t *testing.T) {
	fmt.Println("Test: Storage VM")

	c, err := NewClient("apiuser", "foobar11", "10.10.10.111", true, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Creating Storage VM...")
	err = c.CreateStorageVM("test-svm", "Provisioned by api", "")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Looking up UUID...")
	id, err := c.GetStorageVmUUIDByName("test-svm")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(id)

	fmt.Println("Get Storage VM by UUID...")
	svm, err := c.GetStorageVM(id)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(svm.Name)
	// Output: test-svm


	fmt.Println("Deleting Storage VM...")
	err = c.DeleteStorageVM(id)
	if err != nil {
		t.Error(err)
	}

}

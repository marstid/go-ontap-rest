package ontap

import (
	"fmt"
	"testing"
)

func TestVolume(t *testing.T) {
	fmt.Println("Test: Volume")

	// Connect's to Netapp virtual simulator
	c, err := NewClient("apiuser", "foobar11", "10.10.10.111", true, true)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Creating Aggregate...")
	err = c.CreateAggregate("simsan-01","vtaggr", 5)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Creating Storage VM...")
	err = c.CreateStorageVM("vtsvm", "Test Comment","")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Creating Volume...")
	err = c.CreateVolume("vtvolume", "Test comment", "vtsvm",  []string{"vtaggr"}, 20971520)
	if err != nil {
		t.Error(err)
	}


	fmt.Println("Get Volume UUID...")
	id,err := c.GetVolumeUUID("vtvolume")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Delete Volume...")
	err = c.DeleteVolume(id)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Looking SVM up UUID...")
	id, err = c.GetStorageVmUUIDByName("vtsvm")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Deleting Storage VM...")
	err = c.DeleteStorageVM(id)
	if err != nil {
		t.Error(err)
	}


	id, err = c.GetAggregateUUID("vtaggr")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Deleting Aggregate...")
	err = c.DeleteAggregate(id)
	if err != nil {
		t.Error(err)
	}
}


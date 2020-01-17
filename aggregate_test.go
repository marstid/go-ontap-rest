package ontap

import (
	"fmt"
	"testing"
)

func TestAggr(t *testing.T) {
	fmt.Println("Test: Get Aggrs Info")

	c, err := NewClient("apiuser", "apipassword", "10.10.10.111", true, true)
	if err != nil {
		t.Error(err)
	}

	_, err = c.GetAggregates()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Creating Aggregate...")
	err = c.CreateAggregate("simsan-01","foobar", 5)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Looking up UUID...")

	id, err := c.GetAggregateUUID("foobar")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Deleting Aggregate...")
	err = c.DeleteAggregate(id)
	if err != nil {
		t.Error(err)
	}


}

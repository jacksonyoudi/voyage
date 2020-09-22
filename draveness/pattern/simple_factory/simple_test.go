package simple_factory

import "testing"

func TestHelloAPI_Say(t *testing.T) {
	api := NewAPI(1)
	n := api.Say("xiix")
	if n != "hixiix" {
		t.Fatalf("err")
	}

}

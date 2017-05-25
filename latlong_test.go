package latlong

import (
	"fmt"
	"testing"

	"github.com/coderconvoy/workers"
)

func TestMain(t *testing.T) {
	addrs := []string{
		"BA5 8NP", "BS1 5SP", "BA3 4PQ", "BA1 3TJ",
		"BA7 1PQ", "BA3 4LJ", "BA2 4SD", "BA4 5DO",
		"BS7 1PQ", "BS3 4LJ", "BS2 4SD", "BS4 5DO",
		"CR7 1PQ", "CR3 4LJ", "CR2 4SD", "CR4 5DO",
		"YA7 1PQ", "YA3 4LJ", "YA2 4SD", "YA4 5DO",
	}

	locs := make([]LatLong, len(addrs))

	doloc := func(addr string, i int) func() {
		return func() {
			loc, err := GetLatLong(addr)
			if err != nil {
				fmt.Println("Could not set loc for ", addr, i)
			}
			locs[i] = loc
		}
	}
	wg := workers.New(10)
	for i, v := range addrs {
		wg.Add(doloc(v, i))
	}
	wg.Wait()
	for i := 0; i < len(locs); i++ {
		fmt.Println(locs[i])
		fmt.Printf("Add %d:%s,%v\n", i, addrs[i], locs[i])
	}
}

func TestLatLong(t *testing.T) {
	loc, err := GetLatLong("BA5 2ER")
	if err != nil {
		t.Logf("Error %s", err)
		t.Fail()
	}
	fmt.Println("Loc : ", loc)
}
func TestBody(t *testing.T) {

	gres, err := Get("https://maps.googleapis.com/maps/api/geocode/json",
		"key", GEOKEY,
		"address", "BA5 2ER",
	)

	if err != nil {
		t.Logf("Error %s", err)
		t.Fail()
	}
	if len(gres) < 100 {
		t.Logf("Expected Longer result")
		t.Fail()
	}
	//fmt.Println(string(gres))

}

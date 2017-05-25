package latlong

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type LatLong struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}

func GetLatLong(add string) (LatLong, error) {
	gres, err := Get("https://maps.googleapis.com/maps/api/geocode/json",
		"key", GEOKEY,
		"address", add,
	)
	if err != nil {
		return LatLong{}, errors.Wrap(err, "No Response body")
	}

	jres := gjson.Get(string(gres), "results.0.geometry.location")
	//	fmt.Println("JRES:", jres)
	res := LatLong{}
	err = json.Unmarshal([]byte(jres.Raw), &res)
	if err != nil {
		return res, errors.Wrap(err, "Could not marshal location,: "+jres.Raw)
	}
	return res, nil

}

// Get is a generic Web Get function that builds a client, and grabs web information.
// params. url --the URL, args, string pairs of key value for sending in the URL
// returns a []byte of the response and an error, if anything went wrong

func Get(url string, args ...string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not make request")
	}
	q := req.URL.Query()
	for i := 0; i+1 < len(args); i += 2 {
		q.Add(args[i], args[i+1])
	}
	req.URL.RawQuery = q.Encode()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not get response")
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

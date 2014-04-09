/*
	LibBring

	Description: A Go wrapper around some of the most commonly used functions of Bring.no's API.


	Copyright (C) 2014  Brian C. Tomlinson

	Contact: brian.tomlinson@linux.com

	This program is free software; you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation; either version 2 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License along
	with this program; if not, write to the Free Software Foundation, Inc.,
	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
*/
package libbring

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	version         string = "1.0"
	baseTrackingUrl string = "http://sporing.bring.no/sporing.json?q="
	basePickUpUrl   string = "http://fraktguide.bring.no/fraktguide/api/pickuppoint"
	data            map[string]interface{}
)

// request takes a url as an argument and returns the JSON data
// relevant to the request or an error if the request is not successful.
func request(url string) (map[string]interface{}, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// Track takes a trackingNumber string as an argument, sets up the request
// url, and then returns the JSON data of the request on success.
func Track(trackingNumber string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%v%v", baseTrackingUrl, trackingNumber)
	resp, err := request(url)
	return resp, err
}

// PickByPostalCode takes a postalCode as an argument, sets up the request
// url, and then returns the JSON data of the request on success.
func PickByPostalCode(postalCode string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%v/postalcode/%v.json", basePickUpUrl, postalCode)
	resp, err := request(url)
	return resp, err
}

// PickByPostalCodeWithQuery takes two strings, postalCode and query, as arguments,
// sets up the request url, and then returns the JSON data of the request on success.
//
// NOTE: The returned results will all contain the query string in their names on match.
// No string match, no results.
func PickByPostalCodeWithQuery(postalCode, query string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%v/postalcode/%v.json?searchForText=%v", basePickUpUrl, postalCode, query)
	resp, err := request(url)
	return resp, err
}

// PickByLocation takes the latitude and longitude as arguments, sets up the request
// url, and then returns the JSON data of the request on success.
func PickByLocation(lat, lon float64) (map[string]interface{}, error) {

	url := fmt.Sprintf("%v/location/%v/%v.json", basePickUpUrl, lat, lon)
	resp, err := request(url)
	return resp, err
}

// PickById takes a pickuplocation id as an argument, sets up the request
// url, and then returns the JSON data of the request on success.
func PickById(id string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%v/id/%v.json", basePickUpUrl, id)
	resp, err := request(url)
	return resp, err
}

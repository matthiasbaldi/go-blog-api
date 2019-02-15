// this packages allows to fetch all swiss communities from the github source
// source: https://raw.githubusercontent.com/secanis/communityindex-data/master/communes.json
package communityindex

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// A Commune represents a single commune with all their attributes
type Commune struct {
	Id         string
	Name       string
	Population Population
}

// The population represents all corespondent attributes
type Population struct {
	Total         string
	Change        string
	Concentration string
	Foreigner     string
}

type Response struct {
	Version  string
	Communes []Commune
}

// With the SearchByName function you can search over all communes by a communename
func SearchByName(communeName string) (*Commune, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/secanis/communityindex-data/master/communes.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	// _, err = io.Copy(os.Stdout, resp.Body)
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, errors.New(resp.Status)
	}

	var commune Commune
	for _, item := range r.Communes {
		if strings.ToLower(item.Name) == strings.ToLower(communeName) {
			commune = item
		}
	}

	return &commune, nil
}

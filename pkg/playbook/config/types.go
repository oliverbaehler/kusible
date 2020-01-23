/*
Copyright © 2019 Michael Gruener

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package config implements the playbook config format
*/
package config

// Config holds a list of plays
type Config struct {
	Plays []*Play `json:"plays"`
}

// Play defines which charts are deployed from which
// repositories to which targets
type Play struct {
	Name   string   `json:"name"`
	Groups []string `json:"groups"`
	Charts []*Chart `json:"charts"`
	Repos  []*Repo  `json:"repos"`
}

// Chart holds all information to deploy a helm chart
type Chart struct {
	Name      string                 `json:"name"`
	Repo      string                 `json:"repo"`
	Chart     string                 `json:"chart"`
	Version   string                 `json:"version"`
	Namespace string                 `json:"namespace"`
	Values    map[string]interface{} `json:"values"`
}

// Repo represents a helm chart repository
type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
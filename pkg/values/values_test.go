/*
Copyright © 2021 Bedag Informatik AG

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

package values

import (
	"io/ioutil"
	"testing"

	"github.com/bedag/kusible/pkg/wrapper/ejson"
	"gotest.tools/assert"
	"sigs.k8s.io/yaml"
)

func TestValues(t *testing.T) {
	loadYaml := func(path string) (map[string]interface{}, error) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		var result map[string]interface{}
		err = yaml.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	tests := map[string]struct {
		input    string
		skipEval bool
		expected string
	}{
		"file":           {input: "file/spruce-eval.yml", skipEval: false, expected: "file/spruce-eval.expected.yml"},
		"file-skip-eval": {input: "file/spruce-eval.yml", skipEval: true, expected: "file/spruce-eval.yml"},
		"dir":            {input: "file", skipEval: false, expected: "file/spruce-eval.expected.yml"},
		"dir-skip-eval":  {input: "file", skipEval: true, expected: "file/spruce-eval.yml"},
	}

	ejsonSettings := ejson.Settings{
		KeyDir: "testdata/keydir",
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			d, err := New("testdata/"+tc.input, []string{}, tc.skipEval, ejsonSettings)
			assert.NilError(t, err)
			got := d.Map()
			assert.NilError(t, err)
			delete(got, "_public_key")

			want, err := loadYaml("testdata/" + tc.expected)
			assert.NilError(t, err)
			assert.DeepEqual(t, want, got)
		})
	}
}

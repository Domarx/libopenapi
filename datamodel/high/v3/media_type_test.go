// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
    "github.com/pb33f/libopenapi/datamodel"
    "github.com/pb33f/libopenapi/datamodel/low"
    "github.com/pb33f/libopenapi/datamodel/low/v3"
    "github.com/pb33f/libopenapi/index"
    "github.com/stretchr/testify/assert"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "strings"
    "testing"
)

func TestMediaType_MarshalYAML(t *testing.T) {
    // load the petstore spec
    data, _ := ioutil.ReadFile("../../../test_specs/petstorev3.json")
    info, _ := datamodel.ExtractSpecInfo(data)
    var err []error
    lowDoc, err = v3.CreateDocumentFromConfig(info, &datamodel.DocumentConfiguration{})
    if err != nil {
        panic("broken something")
    }

    // create a new document and extract a media type object from it.
    d := NewDocument(lowDoc)
    mt := d.Paths.PathItems["/pet"].Put.RequestBody.Content["application/json"]

    // render out the media type
    yml, _ := mt.Render()

    // the rendered output should be a ref to the media type.
    op := `schema:
    $ref: '#/components/schemas/Pet'`

    assert.Equal(t, op, strings.TrimSpace(string(yml)))

    // modify the media type to have an example
    mt.Example = "testing a nice mutation"

    op = `schema:
    $ref: '#/components/schemas/Pet'
example: testing a nice mutation`

    yml, _ = mt.Render()

    assert.Equal(t, op, strings.TrimSpace(string(yml)))

}

func TestMediaType_Examples(t *testing.T) {
    yml := `examples:
    pbjBurger:
        summary: A horrible, nutty, sticky mess.
        value:
            name: Peanut And Jelly
            numPatties: 3
    cakeBurger:
        summary: A sickly, sweet, atrocity
        value:
            name: Chocolate Cake Burger
            numPatties: 5`

    var idxNode yaml.Node
    _ = yaml.Unmarshal([]byte(yml), &idxNode)
    idx := index.NewSpecIndexWithConfig(&idxNode, index.CreateOpenAPIIndexConfig())

    var n v3.MediaType
    _ = low.BuildModel(idxNode.Content[0], &n)
    _ = n.Build(idxNode.Content[0], idx)

    r := NewMediaType(&n)

    rend, _ := r.Render()
    assert.Equal(t, yml, strings.TrimSpace(string(rend)))
}

package main

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestFetchNotebookIds(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	id := "2CG4GYWN1"
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", endpoint+"/api/notebook", httpmock.NewStringResponder(200, `{"body": [ { "id": "`+id+`", "name": "name" } ], "message": "", "status": "OK" }`))

	notebooks, err := fetchNotebookIds(endpoint)
	if err != nil {
		t.Fatal(err)
	}

	if len(notebooks) != 1 {
		t.Fatalf("notebooks length is not 1: %d. %s", len(notebooks), notebooks)
	}

	if notebooks[0] != id {
		t.Fatalf(`notebooks[0] is not "%s". notebooks[0] is "%s"`, id, notebooks[0])
	}
}

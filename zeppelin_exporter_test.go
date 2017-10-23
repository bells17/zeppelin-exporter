package main

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestFetchNotebookIds(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	id := "2CG4GYWN1"
	res := `{
		"body": [
			{
				"id": "` + id + `",
				"name": "name"
			}
		],
		"message": "",
		"status": "OK"
	}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", endpoint+"/api/notebook", httpmock.NewStringResponder(200, res))

	notebookIds, err := fetchNotebookIds(endpoint)
	if err != nil {
		t.Fatal(err)
	}

	if len(notebookIds) != 1 {
		t.Fatalf("notebookIds length is not 1: %d", len(notebookIds))
	}

	if notebookIds[0] != id {
		t.Fatalf(`notebookIds[0] is not "%s". notebookIds[0] is "%s"`, id, notebookIds[0])
	}
}

func TestExportNotebooks(t *testing.T) {
	endpoint := "http://127.0.0.1:8080"
	notebookIds := []string{"2CG4GYWN1"}
	// A "body" value is need an encoded json string
	body := `{\r\n\t\"body\": {\r\n\t\t\"paragraphs\": [\r\n\t\t\t{\r\n\t\t\t\t\"text\": \"%md This is my new paragraph in my new note\",\r\n\t\t\t\t\"dateUpdated\": \"Jan 8, 2016 4:49:38 PM\",\r\n\t\t\t\t\"config\": {\r\n\t\t\t\t\t\"enabled\": true\r\n\t\t\t\t},\r\n\t\t\t\t\"settings\": {\r\n\t\t\t\t\t\"params\": {},\r\n\t\t\t\t\t\"forms\": {}\r\n\t\t\t\t},\r\n\t\t\t\t\"jobName\": \"paragraph_1452300578795_1196072540\",\r\n\t\t\t\t\"id\": \"20160108-164938_1685162144\",\r\n\t\t\t\t\"dateCreated\": \"Jan 8, 2016 4:49:38 PM\",\r\n\t\t\t\t\"status\": \"READY\",\r\n\t\t\t\t\"progressUpdateIntervalMs\": 500\r\n\t\t\t}\r\n\t\t],\r\n\t\t\"name\": \"source note for export\",\r\n\t\t\"id\": \"2B82H3RR1\",\r\n\t\t\"angularObjects\": {},\r\n\t\t\"config\": {},\r\n\t\t\"info\": {}\r\n\t},\r\n\t\"message\": \"\",\r\n\t\"status\": \"OK\"\r\n}`
	res := `{
		"body": "` + body + `",
		"message": "",
		"status": "OK"
	}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", endpoint+"/api/notebook/export/"+notebookIds[0],
		httpmock.NewStringResponder(200, res))

	notebooks, err := exportNotebooks(endpoint, notebookIds)
	if err != nil {
		t.Fatal(err)
	}

	if len(notebooks) != 1 {
		t.Fatalf("notebooks length is not 1: %d", len(notebooks))
	}
}

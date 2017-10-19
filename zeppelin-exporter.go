package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Options struct {
	Host     string `long:"host" description:"Zeppelin host" required:"true"`
	Port     int    `short:"p" long:"port" description:"port" default:"8080"`
	Protocol string `long:"protocol" description:"protocol" default:"http"`
}

type noteBookListResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Body    []Notebook `json:body`
}

type Notebook struct {
	ID   string `json:"id"`
	Name string `json:name`
}

func main() {
	opts := Options{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		os.Exit(1)
	}

	url := fmt.Sprintf("%s://%s:%d", opts.Protocol, opts.Host, opts.Port)
	res, err := http.Get(url + "/api/notebook")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonBytes := ([]byte)(string(b))
	data := new(noteBookListResponse)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	if data.Status != "OK" {
		fmt.Println("Fetch notebook list response status: " + data.Status)
		return
	}

	notebooks := []string{}
	for _, notebook := range data.Body {
		res, err := http.Get(url + "/api/notebook/export/" + notebook.ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var notebook interface{}
		err = json.Unmarshal(([]byte)(string(b)), &notebook)
		if err != nil {
			fmt.Println("JSON Unmarshal error:", err)
			return
		}

		var notebookBody interface{}
		notebookBody = notebook.(map[string]interface{})["body"]
		str, ok := notebookBody.(string)
		if !ok {
			fmt.Println("Error stringify")
			return
		}
		notebooks = append(notebooks, str)
	}

	fmt.Println("[" + strings.Join(notebooks[:], ",") + "]")
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var BuildVersion string

type Options struct {
	Host     string `long:"host" description:"Zeppelin host" default:"127.0.0.1"`
	Port     int    `short:"p" long:"port" description:"port" default:"8080"`
	Protocol string `long:"protocol" description:"protocol" default:"http"`
	Version  bool   `long:"version" description:"print version"`
}

type NoteBooksResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Body    []Notebook `json:"body"`
}

type Notebook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	opts := getOptions()
	if opts.Version {
		printVersion()
		os.Exit(0)
	}

	endpoint := fmt.Sprintf("%s://%s:%d", opts.Protocol, opts.Host, opts.Port)
	notebookIds, err := fetchNotebookIds(endpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	notebooks, err := exportNotebooks(endpoint, notebookIds)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("[" + strings.Join(notebooks[:], ",") + "]")
}

func getOptions() Options {
	opts := Options{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		os.Exit(1)
	}
	return opts
}

func printVersion() {
	fmt.Printf(`zeppelin-exporter %s
Compiler: %s %s
`,
		BuildVersion,
		runtime.Compiler,
		runtime.Version())
}

func fetchNotebookIds(endpoint string) ([]string, error) {
	res, err := http.Get(endpoint + "/api/notebook")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsonBytes := ([]byte)(string(b))
	data := new(NoteBooksResponse)
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		return nil, err
	}

	notebookIds := []string{}
	for _, notebook := range data.Body {
		notebookIds = append(notebookIds, notebook.ID)
	}
	return notebookIds, nil
}

func exportNotebooks(endpoint string, notebookIds []string) ([]string, error) {
	notebooks := []string{}
	for _, notebookId := range notebookIds {
		res, err := http.Get(endpoint + "/api/notebook/export/" + notebookId)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var notebook interface{}
		err = json.Unmarshal(([]byte)(string(b)), &notebook)
		if err != nil {
			return nil, err
		}

		var notebookBody interface{}
		notebookBody = notebook.(map[string]interface{})["body"]
		str, ok := notebookBody.(string)
		if !ok {
			return nil, errors.New("Error stringify")
		}
		notebooks = append(notebooks, str)
	}

	return notebooks, nil
}

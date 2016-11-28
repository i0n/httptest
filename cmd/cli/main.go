package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/i0n/httptest/formatjson"
	"github.com/i0n/httptest/utils"
	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
)

var (
	//Version string used to compile version number
	Version string
	//GitCommit string used to compile git commit hash
	GitCommit string

	versionFlag    = flag.Bool("version", false, "The current version")
	gitCommitFlag  = flag.Bool("git-commit", false, "The current git commit")
	configFilePath = flag.String("config-file", "./httptest.yml", "The path to the config file containing your tests")
)

// Manifest struct representing entire JSON manifest
type Manifest struct {
	APIVersion    int    `json:"api_version" yaml:"api_version"`
	ManifestName  string `json:"manifest_name" yaml:"manifest_name"`
	HostName      string `json:"host_name" yaml:"host_name"`
	HostPort      int    `json:"host_port" yaml:"host_port"`
	Paths         []Path `json:"paths" yaml:"paths"`
	manifestType  string
	totalFailures int
}

// Path struct represents a single path to be tested
type Path struct {
	Path           string `json:"path" yaml:"path"`
	Method         string `json:"method" yaml:"method"`
	ContentType    string `json:"content_type" yaml:"content_type"`
	ResponseStatus int    `json:"response_status" yaml:"response_status"`
	ResponseBody   string `json:"response_body" yaml:"response_body"`
	failures       []map[string]string
}

// Failures returns a slice of strings representing test failures for the path
func (p Path) Failures() []map[string]string {
	return p.failures
}

func main() {
	flag.Parse()

	if *versionFlag == true {
		fmt.Println(Version)
		os.Exit(0)
	}

	if *gitCommitFlag == true {
		fmt.Println(GitCommit)
		os.Exit(0)
	}

	// allowed http methods can be found at: http://www.iana.org/assignments/http-methods/http-methods.xhtml
	var allowedMethods = []string{
		"ACL",
		"BASELINE-CONTROL",
		"BIND",
		"CHECKIN",
		"CHECKOUT",
		"CONNECT",
		"COPY",
		"DELETE",
		"GET",
		"HEAD",
		"LABEL",
		"LINK",
		"LOCK",
		"MERGE",
		"MKACTIVITY",
		"MKCALENDAR",
		"MKCOL",
		"MKREDIRECTREF",
		"MKWORKSPACE",
		"MOVE",
		"OPTIONS",
		"ORDERPATCH",
		"PATCH",
		"POST",
		"PRI",
		"PROPFIND",
		"PROPPATCH",
		"PUT",
		"REBIND",
		"REPORT",
		"SEARCH",
		"TRACE",
		"UNBIND",
		"UNCHECKOUT",
		"UNLINK",
		"UNLOCK",
		"UPDATE",
		"UPDATEREDIRECTREF",
		"VERSION-CONTROL",
	}

	fmt.Printf("\n%v\n\n", color.BlueString("Beginning httptest..."))

	configFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	configFileType := filepath.Ext(*configFilePath)
	manifest := Manifest{}

	switch configFileType {
	case ".yml":
		fmt.Printf("YAML config detected\n")
		manifest.manifestType = "yaml"
		yaml.Unmarshal(configFile, &manifest)
	case ".yaml":
		fmt.Printf("YAML config detected\n")
		manifest.manifestType = "yaml"
		yaml.Unmarshal(configFile, &manifest)
	case ".json":
		fmt.Printf("JSON config detected\n")
		manifest.manifestType = "json"
		json.Unmarshal(configFile, &manifest)
	default:
		panic("Unknown config file type")
	}

	url := "http://" + manifest.HostName + ":" + strconv.Itoa(manifest.HostPort)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	for i, p := range manifest.Paths {

		// Test for failures...

		testPathMethod(p.Method, allowedMethods, &p)

		res, body, err := makeRequest(netClient, url, p)
		if err != nil {
			fmt.Printf("Error calling %v\n", url)
			fmt.Println(err)
			os.Exit(1)
		}

		testPathContentType(res, &p)
		testPathStatusCode(res, &p)
		testPathResponseBody(res, body, &p)

		testInfo := ""

		switch manifest.manifestType {
		case "yaml":
			yaml, err := yaml.Marshal(p)
			if err != nil {
				panic(err)
			}
			testInfo = string(yaml)
		case "json":
			json, err := formatjson.Marshal(p)
			if err != nil {
				panic(err)
			}
			testInfo = string(json)
		default:
			panic("Unknown manifest type")
		}

		fmt.Println("----------------------------------------------")
		fmt.Printf("Test: %v\n", i+1)
		fmt.Printf("\n%v\n", testInfo)
		if len(p.Failures()) != 0 {
			manifest.totalFailures += len(p.Failures())
			fmt.Printf("\nFailures:\n")
			for _, val := range p.Failures() {
				fmt.Printf("%v\n", color.RedString(val["failure"]))
				if val["expected"] != "" {
					fmt.Printf("  Expected: %v\n", val["expected"])
				}
				if val["result"] != "" {
					fmt.Printf("  Result:   %v\n", val["result"])
				}
			}
			color.Unset()
		} else {

		}
		fmt.Println("----------------------------------------------")

	}

	if manifest.totalFailures > 0 {
		fmt.Printf("%v\n", color.RedString("FAIL!"))
		fmt.Printf("%v total failures\n", color.RedString(strconv.Itoa(manifest.totalFailures)))
		os.Exit(1)
	}
	fmt.Printf("%v\n", color.GreenString("PASS!"))

}

func makeRequest(netClient *http.Client, url string, p Path) (http.Response, string, error) {
	b := bytes.NewBufferString(p.ResponseBody)
	req, _ := http.NewRequest(p.Method, url+p.Path, b)
	req.Header.Add("Content-Type", p.ContentType)
	res, err := netClient.Do(req)
	if err != nil {
		return http.Response{}, "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return http.Response{}, "", err
	}
	return *res, string(body), nil
}

func testPathMethod(method string, allowedMethods []string, p *Path) {
	if !utils.StringInSlice(p.Method, allowedMethods) {
		m := make(map[string]string)
		m["failure"] = "method"
		p.failures = append(p.failures, m)
	}
}

func testPathContentType(res http.Response, p *Path) {
	if res.Header.Get("Content-Type") != p.ContentType {
		m := make(map[string]string)
		m["failure"] = "content_type"
		m["expected"] = p.ContentType
		m["result"] = res.Header.Get("Content-Type")
		p.failures = append(p.failures, m)
	}
}

func testPathStatusCode(res http.Response, p *Path) {
	if res.StatusCode != p.ResponseStatus {
		m := make(map[string]string)
		m["failure"] = "response_status"
		m["expected"] = strconv.Itoa(p.ResponseStatus)
		m["result"] = strconv.Itoa(res.StatusCode)
		p.failures = append(p.failures, m)
	}
}

func testPathResponseBody(res http.Response, body string, p *Path) {
	if p.ResponseBody != "" && string(body) != p.ResponseBody {
		m := make(map[string]string)
		m["failure"] = "response_body"
		m["expected"] = p.ResponseBody
		m["result"] = string(body)
		p.failures = append(p.failures, m)
	}
}

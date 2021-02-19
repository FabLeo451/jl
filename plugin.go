/**
 * JLogic command line utility
 * Written by Fabio Leone
 */

package main

import (
    "fmt"
    "os"
    //"flag"
    //"crypto/tls"
    "net/http"
    "net/url"
    "encoding/json"
    "io/ioutil"
    "github.com/lensesio/tableprinter"
    "path/filepath"
    "strconv"
    "strings"
)

type Plugin struct {
    Name        string `header:"name"`
    Version     string `header:"version"`
    Description string `header:"description"`
}

/**
 * List plugins
 */
func ListPlugins(args []string) int {
    var plugin []Plugin

    client := GetClient()
    url := fmt.Sprintf("https://%s:%d/plugins", host, port)

    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Got error %s", err.Error())
	}
	req.SetBasicAuth(user, password)

    //req.Header.Add("Authorization","Basic " + basicAuth("username1","password123"))
    response, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        return 1
    }
    defer response.Body.Close()

    //fmt.Println("Response status:", response.Status)
    if response.StatusCode != 200 {
        PrintResponseError(response)
        return 1
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err.Error())
    }

    json.Unmarshal(body, &plugin)

    printer := tableprinter.New(os.Stdout)
    printer.Print(plugin)

    return 0
}

/**
 * Copmile a program
 */
func InstallPlugin(args []string) int {
    if (len(args) < 2) {
        fmt.Fprintln(os.Stderr, "Missing file path")
        return 1
    }

    filepath, _ := filepath.Abs(args[1])
    fmt.Println("Installing", filepath, "...")

    data := url.Values{}
    data.Set("jar", filepath)

    client := GetClient()
    url := fmt.Sprintf("https://%s:%d/plugin/install", host, port)

    req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
    if err != nil {
    	fmt.Println("Got error %s", err.Error())
    }
    req.SetBasicAuth(user, password)
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    //req.Header.Add("Authorization","Basic " + basicAuth("username1","password123"))
    response, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        return 1
    }
    defer response.Body.Close()

    //fmt.Println("Response status:", response.Status)
    if response.StatusCode != 200 {
        PrintResponseError(response)
        return 1
    }

    fmt.Println("Plugin successfully installed")

    return 0
}

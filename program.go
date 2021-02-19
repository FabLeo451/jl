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
    "encoding/json"
    "io/ioutil"
    "github.com/lensesio/tableprinter"
)

type Program struct {
    Id      string `header:"id"`
    Name    string `header:"name"`
    Version string `header:"version"`
    Status  string `header:"status"`
}

/**
 * List programs
 */
func ListPrograms(args []string) int {
    var programs []Program

    client := GetClient()
    url := fmt.Sprintf("https://%s:%d/programs", host, port)

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
        //fmt.Fprintf(os.Stderr, "Error %d: %s\n", response.StatusCode, http.StatusText(response.StatusCode))
        PrintResponseError(response)
        return 1
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err.Error())
    }

    json.Unmarshal(body, &programs)

    printer := tableprinter.New(os.Stdout)
    printer.Print(programs)

    return 0
}

/**
 * Copmile a program
 */
func CopmileProgram(args []string) int {
    if (len(args) < 2) {
        fmt.Fprintln(os.Stderr, "Missing program id")
        return 1
    }

    client := GetClient()
    url := fmt.Sprintf("https://%s:%d/program/%s/compile", host, port, args[1])

    req, err := http.NewRequest("POST", url, nil)
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

    fmt.Println("Successfully compiled")

    return 0
}

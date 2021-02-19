/**
 * JLogic command line utility
 * Written by Fabio Leone
 */

package main

import (
    "fmt"
    "os"
    "crypto/tls"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type JLResponse struct {
    Status  int     `json:"status"`
    Message string  `json:"message"`
}

func GetClient() *http.Client {
    transCfg := &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
    }
    client := &http.Client{Transport: transCfg};

    return(client)
}

func PrintResponseError(response *http.Response) {
    body, err := ioutil.ReadAll(response.Body)

    if err != nil {
        panic(err.Error())
    }

    //fmt.Println(string(body))

    jo := JLResponse{}
    jsonErr := json.Unmarshal(body, &jo)

    if jsonErr != nil {
        //fmt.Println(jsonErr)
        fmt.Fprintf(os.Stderr, "Error %d: %s\n", response.StatusCode, http.StatusText(response.StatusCode))
    } else {
        fmt.Println(jo.Message)
    }
}

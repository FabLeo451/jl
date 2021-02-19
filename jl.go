/**
 * JLogic command line utility
 * Written by Fabio Leone
 */

package main

import (
    "fmt"
    "os"
    "flag"
    "path"
)

var version = "1.0.0"

var port int
var host, user, password string
var flagVersion = false

type Callback func([]string) int

type Command struct {
    f Callback
    args string
    help string
}

func main() {
    mapCommands := make(map[string]Command)
    mapCommands["programs"] = Command{ ListPrograms, "", "List programs" }
    mapCommands["compile"] = Command{ CopmileProgram, "ID", "Compile program with the given id" }
    mapCommands["plugins"] = Command{ ListPlugins, "", "List installed plugins" }
    mapCommands["install"] = Command{ InstallPlugin, "JARFILE", "Install a plugin from a jar file" }

    flag.StringVar(&user, "u", "", "Set user")
    flag.StringVar(&password, "p", "", "Set password")
    flag.StringVar(&host, "H", "localhost", "Set server host")
    flag.IntVar(&port, "P", 8443, "Set server port")
    flag.BoolVar(&flagVersion, "v", false, "Show version")

    flag.Usage = func() {
        fmt.Printf("%s %s\n", path.Base(os.Args[0]), version)
        fmt.Printf("Usage: %s [options] command [arguments]\n", path.Base(os.Args[0]))

        fmt.Println("\nCommands:")

        for key, element := range mapCommands {
            fmt.Printf("  %-8s %-10s %s\n", key, element.args, element.help)
        }

        fmt.Println("\nOptions:")

        flag.VisitAll(func(f *flag.Flag) {
            a, d := "", ""

            if f.Value.String() != "false" {
                d = "(default: " + f.Value.String() +")"
                a = "<value>"
            }
            fmt.Printf("  -%s %-10s %s %s\n", f.Name, a, f.Usage, d) // f.Name, f.Value
        })
    }

    flag.Parse()

    if flagVersion {
        fmt.Println(version)
    }

    if user == "" {
        user = os.Getenv("JLOGIC_USER")
    }

    if password == "" {
        password = os.Getenv("JLOGIC_PASSWORD")
    }

    //fmt.Println("tail:", flag.Args())
    args := flag.Args();

    if len(args) == 0 {
        os.Exit(0)
    }

    var exitValue int = 0;

    if c, found := mapCommands[args[0]]; found {
        exitValue = c.f(args)
    } else {
        fmt.Fprintln(os.Stderr, "Unknown command: ", args[0])
        exitValue = 1
    }

    os.Exit(exitValue)
}

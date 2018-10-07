package main

import (
    "flag"
    "github.com/drewrip/OpenTA/tasc"
    "path/filepath"
    "os"
)

func main(){
    // Defines flags
    defdir,_ := filepath.Abs("." + string(os.PathSeparator))
    buildName := flag.String("o", "Project", "Name of the project after successful build")
    buildDir := flag.String("d", defdir, "Where to build your story to")

    flag.Parse()
    // Sets compile time variables
    tasc.SetInputPath(flag.Arg(0))
    tasc.SetOut(*buildDir)
    tasc.SetName(*buildName)
    tasc.SetBuildPath()
    tasc.Init()
}

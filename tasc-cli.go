package main

import (
    //"fmt"
    "flag"
    "github.com/drewrip/OpenTA/tasc"
)

func main(){
    //buildName := flag.String("o", "Project", "Name of the project after successful build")

    //Uses -s flag to specify source directory of .tam files
    srcDir := flag.String("s", "./", "Directory where the .tam source files are stored")

    flag.Parse()

    // Sets global path to source directory
    tasc.SetPath(*srcDir)
    tasc.Init()
}

package tasc

import (
    "encoding/json"
    "io/ioutil"
    "strings"
    "fmt"
)

//Separates Story and Meta from a given file (not including extension)
func ParseFile(f string) (string, string){

    rawFileText, err := ioutil.ReadFile(Path + f + ".tam")
    if err != nil {
        panic(err)
    }
    fileText := string(rawFileText)
    parts := strings.Split(fileText, "¤")
    return parts[0], parts[1]
}

// Decodes the meta section of a tam file and returns it as an array of Options
func DecodeMeta(m string) []Option {
    var opt []Option
    json.Unmarshal([]byte(m), &opt)
    return opt
}

func (b *Branch)PopulateOptions(){
    // Recursive base case
    if b.TextPath == "END" && b.Choice == "END" {
        return
    }

    _, meta := ParseFile(b.TextPath)
    Opts := DecodeMeta(meta)
    // For each option listed in a tam file, parse into a Branch struct
    for _,x := range Opts{
        stage := Branch{}
        stage.TextPath = x.Result
        stage.Choice = x.Input
        stage.Branches = []Branch{}
        stage.PopulateOptions()
        b.Branches = append(b.Branches, stage)
    }

}

// Runs through entire story and maps out options
func (b *Branch)ParseStory() string {
    b.PopulateOptions()
    f, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return string(f)
}

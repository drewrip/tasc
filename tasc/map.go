package tasc

import (
    "encoding/json"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

type Branch struct {
    //Name of the file that contains the story and options for progression (ex. moveRight.tam -> "moveRight")
    TextPath string `json:"textpath"`
    //Option (ex. "Move Right") the user picks to navigate down this branch of the story
    Choice string `json:"choice"`
    //Options for progession from the current branch
    Branches []Branch `json:"branches"`
}

// Go struct represents the options defined in tam files, only used during initial parsing/mapping
type Option struct{
    Input string
    Result string
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
        fmt.Println("PLACE: " + x.Result)
        stage.Choice = x.Input
        stage.Branches = []Branch{}
        hash := sha256.Sum256([]byte(x.Result))
        stage.TextPath = hex.EncodeToString(hash[:])
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

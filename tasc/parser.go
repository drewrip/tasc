package tasc

import (
    "encoding/json"
    "io/ioutil"
    "strings"
)

type Info struct{
    // Story info subject to change based on community needs. Much of this info is up to the devs of the players as to whether
    // or not they would like to incorporate it.
    Title string // This is the story's name, what it should be called
    Description string // A blurb about what the story is about
    Author string // Can be one name or a list, but is always in string form
    AgeRec int // Author's age recommendation
    Disclaimer string // Any disclaimer, that you would like to include
    Tags []string // Tags for the story
}

//Separates Story and Meta from a given file (not including extension)
func ParseFile(f string) (string, string){

    rawFileText, err := ioutil.ReadFile(InputPath + f + ".tam")
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

func GetInfo(p string) Info {
    rawInfo, err := ioutil.ReadFile(p + "info.json")
    if err != nil {
        panic(err)
    }
    inf := Info{}
    json.Unmarshal(rawInfo, inf)
    return inf
}

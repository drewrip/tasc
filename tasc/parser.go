package tasc

import (
    "encoding/json"
    "io/ioutil"
    "strings"
)

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

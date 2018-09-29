package tasc

import (
    "encoding/json"
    "io/ioutil"
    "regexp"
)

// Go struct represents the options defined in tam files
type Option struct{
    Input string
    Result string
}

// Pulls the meta data out of the branch page
func (b *Branch) GetMeta() string{

    //Accounting for the base branch case
    if b.TextPath == "END" && b.Choice == "END" {
        return "END"
    }

    fileText, err := ioutil.ReadFile(Path + b.TextPath)
    if err != nil {
        panic(err)
    }

    r, _ := regexp.Compile("\\[([^()])*\\]")

    return string(r.Find(fileText))
}

// Decodes the meta section of a tam file and returns it as an array of Options
func (b *Branch) Decode() []Option {
    var opt []Option
    json.Unmarshal([]byte(b.GetMeta()), &opt)
    return opt
}

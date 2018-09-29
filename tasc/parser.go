package tasc

import (
    "fmt"
    "io/ioutil"
    //"regexp"
)

func (b *Branch) GetMeta() string{

    //Accounting for the base branch case
    if b.TextPath == "END" && b.Option == "END" {
        return "END"
    }

    fileText, err := ioutil.ReadFile(Path + b.TextPath)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(fileText))
    return ""
}

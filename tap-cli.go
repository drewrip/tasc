package main

import (
    "flag"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "bufio"
    "os"
    "strings"
)


type Branch struct {
    //Name of the file that contains the story and options for progression (ex. moveRight.tam -> "moveRight")
    TextPath string
    //Option (ex. "Move Right") the user picks to navigate down this branch of the story
    Choice string
    //Options for progession from the current branch
    Branches []Branch
}

const idel = '\n'
var PATH string
func display(p string){
    rawFileText, _ := ioutil.ReadFile(PATH + string(os.PathSeparator) + p + ".tas")
    fmt.Println("====================================================")
    fmt.Println(string(rawFileText))
    fmt.Println("====================================================")
}

func (b *Branch) Update(){
    display(b.TextPath)
    if b.Branches[0].Choice == "END"{
        return
    }
    fmt.Println("\tOptions:\n")
    for i,d := range b.Branches{
        choices := fmt.Sprintf("%d. %s", i+1, d.Choice)
        fmt.Println(choices)
    }

    fmt.Println("\nChoice: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString(idel)
    correct := false
    tempBranch := Branch{}
    for _,d := range b.Branches{
        if d.Choice == strings.TrimRight(input, "\n"){
            correct = true
            tempBranch = d
            break
        }
    }

    if correct{
        fmt.Println()
        tempBranch.Update()
    } else {
        fmt.Println()
        b.Update()
    }
}

func Init(s string){
    PATH = s
    rawJSON, _ := ioutil.ReadFile(PATH + string(os.PathSeparator) + "map.json")
    Story := Branch{}
    json.Unmarshal([]byte(rawJSON), &Story)
    Story.Update()
}

func main(){
    flag.Parse()
    // Sets compile time variables
    story := flag.Arg(0)
    Init(story)
}

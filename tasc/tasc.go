package tasc

import (
    "fmt"
)

// Setting the path for the source .tam files
var Path string
func SetPath(p string){
    Path = p
}

//Creating the initial root branch of the story
func Init(){
    Story := Branch{
            TextPath: "root",
            Choice: "",
            Branches: []Branch{},
        }
    fmt.Println(Story.ParseStory())
}
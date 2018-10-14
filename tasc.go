package tasc

import (
    "os"
)

// Setting the path for the source .tam files
var InputPath string
var Out string
var Name string
var BuildPath string

func SetInputPath(p string){
    InputPath = p
}

func SetOut(p string){
    Out = p
}

func SetName(n string){
    Name = n
}

func SetBuildPath(){
    BuildPath = Out + string(os.PathSeparator) + Name
}


//Creating the initial root branch of the story
func Init(){
    Story := Branch{
            TextPath: "root",
            Choice: "",
            Branches: []Branch{},
        }
    Story.ParseStory()
    CreateDir()
    BuildStories()
    Story.GenMap()
    BuildInfo()
}

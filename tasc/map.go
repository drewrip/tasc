package tasc

import (
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

// Base case for the recursive branch definition, the end of a branch of the story
var BASE Branch = Branch{
    TextPath: "END",
    Choice: "END",
    Branches: nil,
}

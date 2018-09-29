package tasc

import (
)

type Branch struct {
    //Name of the file that contains the story and options for progression (ex. moveRight.json -> "moveRight")
    TextPath string
    //Option (ex. "Move Right") the user picks to navigate down this branch of the story
    Choice string
    //Options for progession from the current branch
    Branches []Branch
}

// Base case for the recursive branch definition, the end of a branch of the story
var BASE Branch = Branch{
    TextPath: "END",
    Choice: "END",
    Branches: nil,
}

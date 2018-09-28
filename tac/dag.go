package OpenTA

import (
    "fmt"
)

type Branch struct {
    //Name of the file that contains the story and options for progression (ex. moveRight.json -> "moveRight")
    TextPath string
    //Option (ex. "Move Right") the user picks to navigate down this branch of the story
    Option string
    //Options for progession from the current branch
    Branches []Branch
}

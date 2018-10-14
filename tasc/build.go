package tasc

import (
    "os"
    "io/ioutil"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
)

func CreateDir(){
    os.Mkdir(BuildPath, 0700)
}

// Generates the compiled stories and the map
func BuildStories(){
    files, err := ioutil.ReadDir(InputPath)
    if err != nil {
        panic(err)
    }
    for _, f := range files{
            if f.Name() == "info.json"{
                continue
            }
            content, _ := ParseFile(f.Name()[:len(f.Name())-4])
            hash := sha256.Sum256([]byte(content))
            fileName := hex.EncodeToString(hash[:])
            f, err := os.Create(BuildPath + string(os.PathSeparator) + fileName + ".tas")
            if err != nil {
                panic(err)
            }
            f.WriteString(content)
            f.Sync()
            f.Close()
    }
}

func BuildInfo(){
    f, err := os.Create(BuildPath + string(os.PathSeparator) + "info.json")
    rawInfo, err := ioutil.ReadFile(InputPath + "info.json")
    if err != nil {
        panic(err)
    }
    f.WriteString(string(rawInfo))
    f.Sync()
    f.Close()
}

func (b *Branch) ReplaceWithHashed(){
    // Recursive base case
    if b.TextPath == "END" && b.Choice == "END" {
        return
    }
    content, _ := ParseFile(b.TextPath)
    hash := sha256.Sum256([]byte(content))
    b.TextPath = hex.EncodeToString(hash[:])
    // For each option listed in a tam file, parse into a Branch struct
    for i := 0; i < len(b.Branches); i++ {
        b.Branches[i].ReplaceWithHashed()
    }
}
func (b *Branch) GenMap(){
    b.ReplaceWithHashed()
    j, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    f, err := os.Create(BuildPath + string(os.PathSeparator) + "map.json")
    if err != nil {
        panic(err)
    }
    f.WriteString(string(j))
    f.Sync()
    f.Close()
}

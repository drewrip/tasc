package tasc

import (
    "os"
    "io/ioutil"
    "crypto/sha256"
    "encoding/hex"
)

func CreateDir(){
    os.Mkdir(BuildPath, 0700)
}

func BuildStories(){
    files, err := ioutil.ReadDir(InputPath)
    if err != nil {
        panic(err)
    }
    for _, f := range files{
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

/*func (b *Branch) GenManifest(){

}*/

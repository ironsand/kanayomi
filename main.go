package main

import (
  "bytes"
  "compress/flate"
  "encoding/gob"
  "fmt"
  "os"
  "strings"
)

// Restore map object    
func SetupDic() map[string]string {
  buf := bytes.NewBuffer(BepDicByteArray)
  in := flate.NewReader(buf)
  dic := make(map[string]string)
  if err := gob.NewDecoder(in).Decode(&dic); err != nil {
    panic(err)
  }
  in.Close()

  return dic
}

func main() {
  dic := SetupDic()
  if len(os.Args) <= 1 {
	  return;
  }
  if r, ok := dic[strings.ToUpper(os.Args[1])]; ok {
    fmt.Print(r)
  }
}
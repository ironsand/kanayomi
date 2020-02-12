package main

import (
  "bytes"
  "compress/flate"
  "encoding/csv"
  "encoding/gob"
  "fmt"
  "os"
  "strings"
)

func main() {
  tbl := make(map[string]string)
  reader := csv.NewReader(os.Stdin)
  reader.Comma = ' '
  reader.FieldsPerRecord = -1
  for r, err := reader.Read(); err == nil; r, err = reader.Read() {
    tbl[r[0]] = strings.Replace(r[1], "トゥ", "ト", -1)
  }

  buf := new(bytes.Buffer)
  out, _ := flate.NewWriter(buf, flate.DefaultCompression)
  if err := gob.NewEncoder(out).Encode(tbl); err != nil {
    panic(err)
  }
  out.Close()

  fmt.Printf("package main\n\nvar BepDicByteArray []byte = []byte{\n\t")
  for i, v := range buf.Bytes() {
    fmt.Printf("%#02x, ", v)
    if (i+1)%10 == 0 {
      fmt.Printf("\n\t")
    }
  }
  fmt.Printf("\n}")
}

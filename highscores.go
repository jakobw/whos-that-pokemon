package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Score struct {
  Name string
  Value string
}

func main() {
  http.Handle("/", http.FileServer(http.Dir("public")))
  http.HandleFunc("/score", submitScoreHandler)
  http.ListenAndServe(":8080", nil)
}

func readScores() []Score {
  var scores []Score
  contents, _ := ioutil.ReadFile("public/scores.json")
  err := json.Unmarshal(contents, &scores)
  if err != nil {
    return []Score{}
  }

  return scores
}

func writeScores(scores []Score) {
  contend, err := json.Marshal(scores)
  if err == nil {
    ioutil.WriteFile("public/scores.json", contend, 0x644)
  }
}

func submitScoreHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    score := Score{
      Name: r.FormValue("name"),
      Value: r.FormValue("score"),
    }
    writeScores(append(readScores(), score))
  } else {
    fmt.Fprint(w, "")
  }
}

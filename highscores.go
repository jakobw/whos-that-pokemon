package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strconv"
)

type Score struct {
  Name string
  Value, Missing, Guessed int
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
    var score Score
    score.Name = r.FormValue("name")
    score.Value, _ = strconv.Atoi(r.FormValue("score"))
    score.Missing, _ = strconv.Atoi(r.FormValue("missing"))
    score.Guessed, _ = strconv.Atoi(r.FormValue("guessed"))

    writeScores(append(readScores(), score))
  } else {
    fmt.Fprint(w, "")
  }
}

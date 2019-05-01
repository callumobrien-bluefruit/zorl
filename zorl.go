package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/rwestlund/gotex"
)

type RequestData struct {
    Document string
}

func main() {
    http.HandleFunc("/", handle)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func handle(w http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 500)
        return
    }

    var data RequestData
    err = json.Unmarshal(body, &data)
    if err != nil {
        log.Println(err)
        http.Error(w, err.Error(), 400)
        return
    }

    pdf, err := render(data.Document)
    if err != nil {
        log.Println(err)
        http.Error(w, "failed to render LaTeX document", 400)
        return
    }

    w.Write(pdf)
}

func render(document string) ([]byte, error) {
    opts := gotex.Options{
        Command: `pdflatex.exe`,
        Runs: 1,
    }
    return gotex.Render(document, opts)
}

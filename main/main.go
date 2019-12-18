package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
)

type headline struct {
    Source string `json:"source.name"`
    Author string `json:"author"`
    Title string `json:"title"`
    Description string `json:"description"`
    Image string `json:"urlToImage"`
    Url string `json:"url"`
    Content string `json:"content"`
}
type headlines struct {
    Headlines []headline `json:"articles"`
}

const HEADLINR_NEWSAPI_KEY = "HEADLINR_NEWSAPI_KEY"

func main() {
    apiKey := os.Getenv(HEADLINR_NEWSAPI_KEY)

    if apiKey == "" {
        fmt.Println("Missing environment variable " + HEADLINR_NEWSAPI_KEY + ".")
        return
    }

    httpClient := &http.Client{}
    resp, _ := httpClient.Get("https://newsapi.org/v2/top-headlines?country=gb&category=general&apiKey=" + apiKey)

    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err)
        }
        body := string(bodyBytes)

        headlinesMap := headlines{}
        err = json.Unmarshal([]byte(body), &headlinesMap)
        if err != nil {
            fmt.Println(err)
            return
        }
        for _, currHeadline := range headlinesMap.Headlines {
            fmt.Printf("%s:\n", currHeadline.Url)
            fmt.Printf("\t\033[1;36m%s\033[0m\n", currHeadline.Title)
        }
    }
}

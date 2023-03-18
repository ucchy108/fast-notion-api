package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type Header struct {
	Method      string
	ContentType string
	Version     string
}

func main() {
	args := os.Args
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	data := fmt.Sprintf(`{
		"parent": { "database_id": "%s" },
		"properties": {
			"名前": {
				"title": [
					{
						"text": {
							"content": "%s"
						}
					}
				]
			}
		}
	}`, os.Getenv("NOTION_DB_ID"), args[1])

	body := []byte(data)
	buf := bytes.NewBuffer(body)

	url := "https://api.notion.com/v1/pages"
	header := Header{
		Method:      "POST",
		ContentType: "application/json",
		Version:     "2022-06-28",
	}

	req, err := http.NewRequest(header.Method, url, buf)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", header.ContentType)
	req.Header.Set("Notion-Version", header.Version)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("NOTION_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

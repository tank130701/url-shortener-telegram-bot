package helpers

import(
	"net/http"
	"encoding/json"
	"bytes"
	"log"
)

type RequestData struct {
	FullURL   string `json:"full_url"`
}

type ResponseData struct {
	ShortURL   string `json:"short_url"`
}

const (
	url = "http://localhost:8000/api/shorten"
)

func GenerateShortLink(link string)string{

	data := RequestData{
		FullURL:  link,
	}
	// Выполняем POST-запрос на указанный URL
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal JSON:", err)
		return ""
	}

	// Отправка POST-запроса с данными в формате JSON
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Failed to send request:", err)
		return ""
	}
	defer resp.Body.Close()

	// Декодирование ответа JSON
	var responseData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		log.Println("Failed to decode JSON response:", err)
		return ""
	}

	shortLink := responseData.ShortURL
	return shortLink
	
}
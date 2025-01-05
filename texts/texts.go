package texts

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const CurrentLanguage string = "ru"

func Text(key string) string {
	files, err := filepath.Glob("./texts/" + CurrentLanguage + ".json")
	if err != nil || len(files) == 0 {
		log.Fatal("file not found")
	}

	content, err := os.ReadFile(files[0])
	if err != nil {
		log.Fatal("cannot read file")
	}

	var data map[string]string
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("cannot parse json")
	}

	value, exists := data[key]
	if !exists {
		return "key not found"
	}

	return value
}

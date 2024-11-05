package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func readJSONFromFile(path string) interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		os.Stderr.WriteString("Error reading file: " + err.Error() + "\n")
		os.Exit(1)
	}
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		os.Stderr.WriteString("Error parsing JSON: " + err.Error() + "\n")
		os.Exit(1)
	}
	return jsonData
}

func readJSONFromStdin() interface{} {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Stderr.WriteString("Error reading stdin: " + err.Error() + "\n")
		os.Exit(1)
	}
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		os.Stderr.WriteString("Error parsing JSON: " + err.Error() + "\n")
		os.Exit(1)
	}
	return jsonData
}

func queryJSONMap(jsonData interface{}, query map[string]interface{}) interface{} {
	result := make(map[string]interface{})
	for key, value := range query {
		if value == nil {
			result[key] = jsonData
		} else {
			result[key] = queryJSON(jsonData, value.(string))
		}
	}
	return result
}

func queryJSONArray(jsonData interface{}, query []interface{}) interface{} {
	result := make([]interface{}, len(query))
	for i, value := range query {
		result[i] = queryJSON(jsonData, value.(string))
	}
	return result
}

func queryJSON(jsonData interface{}, query string) interface{} {
	var queryData interface{}
	err := json.Unmarshal([]byte(query), &queryData)
	if err != nil {
		os.Stderr.WriteString("Error parsing query: " + err.Error() + "\n")
		os.Exit(1)
	}

	var result interface{}
	switch queryData := queryData.(type) {
	case map[string]interface{}:
		result = queryJSONMap(jsonData, queryData)
	case []interface{}:
		result = queryJSONArray(jsonData, queryData)
	default:
		os.Stderr.WriteString("Error: query must be an object or an array\n")
		os.Exit(1)
	}
	return result
}

func prettyPrintJSON(jsonData interface{}) {
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		os.Stderr.WriteString("Error marshalling JSON: " + err.Error() + "\n")
		os.Exit(1)
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	var filePtr, queryPtr string
	flag.StringVar(&filePtr, "f", "", "path to JSON file")
	flag.StringVar(&queryPtr, "q", "", "query to JSON file")
	flag.Parse()

	var jsonData interface{}
	if filePtr != "" {
		jsonData = readJSONFromFile(filePtr)
	} else {
		jsonData = readJSONFromStdin()
	}

	if queryPtr != "" {
		jsonData = queryJSON(jsonData, queryPtr)
	}

	prettyPrintJSON(jsonData)
}

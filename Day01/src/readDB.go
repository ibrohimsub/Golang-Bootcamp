package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Cake represents a cake recipe.
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

// Ingredient represents an ingredient in a cake recipe.
type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

// DBReader is the interface for reading databases.
type DBReader interface {
	ReadDatabase() ([]Cake, error)
}

// XMLReader reads XML databases.
type XMLReader struct {
	filename string
}

// JSONReader reads JSON databases.
type JSONReader struct {
	filename string
}

// ReadDatabase reads the XML database and returns the cakes.
func (r XMLReader) ReadDatabase() ([]Cake, error) {
	xmlFile, err := os.Open(r.filename)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var recipes struct {
		Cakes []Cake `xml:"cake"`
	}

	err = xml.Unmarshal(byteValue, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes.Cakes, nil
}

// ReadDatabase reads the JSON database and returns the cakes.
func (r JSONReader) ReadDatabase() ([]Cake, error) {
	jsonFile, err := os.Open(r.filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var recipes struct {
		Cakes []Cake `json:"cake"`
	}

	err = json.Unmarshal(byteValue, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes.Cakes, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./readDB -f <filename>")
		os.Exit(1)
	}

	filename := os.Args[2]
	extension := getFileExtension(filename)

	var reader DBReader

	switch extension {
	case "xml":
		reader = XMLReader{filename: filename}
	case "json":
		reader = JSONReader{filename: filename}
	default:
		fmt.Println("Unsupported file format")
		os.Exit(1)
	}

	cakes, err := reader.ReadDatabase()
	if err != nil {
		fmt.Println("Error reading database:", err)
		os.Exit(1)
	}

	// Print JSON version when reading XML and vice versa
	var output []byte

	switch extension {
	case "xml":
		output, err = json.MarshalIndent(cakes, "", "    ")
	case "json":
		output, err = xml.MarshalIndent(cakes, "", "    ")
	}

	if err != nil {
		fmt.Println("Error converting database:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}

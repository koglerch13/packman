package lib

import (
	"encoding/xml"
	"os"
)

func LoadConfig(configFilePath string) (*Config, error) {
	data, error := readFile(configFilePath)
	if error != nil {
		return nil, error
	}

	packages := Config{}
	error = xml.Unmarshal(data, &packages)
	if error != nil {
		return nil, error
	}

	return &packages, nil
}

func readFile(configFilePath string) ([]byte, error) {
	configFileData, error := os.ReadFile(configFilePath)
	if error != nil {
		return nil, error
	}

	return configFileData, nil
}

type Config struct {
	XMLName  xml.Name  `xml:"packages"`
	Packages []Package `xml:"package"`
}

type Package struct {
	XMLName          xml.Name `xml:"package"`
	Uri              string   `xml:"uri,attr"`
	Hash             string   `xml:"hash,attr"`
	Destination      string   `xml:"destination,attr"`
	ClearDestination bool     `xml:"clear-destination,attr"`
}

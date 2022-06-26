package benchjson

import (
	"encoding/json"
)

// DataCamel type
type DataCamel struct {
	ID                string   `json:"id"`
	NameNameNome      string   `json:"nameNameNome"`
	FieldOne          string   `json:"fieldOne"`
	FieldTwo          string   `json:"fieldTwo"`
	FieldThree        string   `json:"fieldThree"`
	FieldFour         string   `json:"fieldFour"`
	FieldFive         string   `json:"fieldFive"`
	FieldSix          string   `json:"fieldSix"`
	FieldSeven        string   `json:"fieldSeven"`
	FieldEight        string   `json:"fieldEight"`
	FieldNine         string   `json:"fieldNine"`
	FieldTen          string   `json:"fieldTen"`
	BlaBleBloQuantity uint64   `json:"blaBleBloQuantity"`
	Images            []string `json:"images"`
}

// ToJSON function
func (p *DataCamel) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// FromCamelJSON function
func FromCamelJSON(rawData []byte) (*DataCamel, error) {
	var p DataCamel

	err := json.Unmarshal(rawData, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// DataSnake type
type DataSnake struct {
	ID                string   `json:"id"`
	NameNameNome      string   `json:"name_name_nome"`
	FieldOne          string   `json:"field_one"`
	FieldTwo          string   `json:"field_two"`
	FieldThree        string   `json:"field_three"`
	FieldFour         string   `json:"field_four"`
	FieldFive         string   `json:"field_five"`
	FieldSix          string   `json:"field_six"`
	FieldSeven        string   `json:"field_seven"`
	FieldEight        string   `json:"field_eight"`
	FieldNine         string   `json:"field_nine"`
	FieldTen          string   `json:"field_ten"`
	BlaBleBloQuantity uint64   `json:"bla_ble_blo_quantity"`
	Images            []string `json:"images"`
}

// ToJSON function
func (p *DataSnake) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}

// FromSnakeJSON function
func FromSnakeJSON(rawData []byte) (*DataSnake, error) {
	var p DataSnake

	err := json.Unmarshal(rawData, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

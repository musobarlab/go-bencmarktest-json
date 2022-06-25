package benchjson

import (
	"testing"
)

func BenchmarkSerializeDataCamelToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := &DataCamel{
			ID:                "1",
			NameNameNome:      "Samsung Galaxy S10",
			FieldOne:          "field one",
			FieldTwo:          "field two",
			FieldThree:        "field three",
			BlaBleBloQuantity: 10,
			Images:            []string{"img1", "img2"},
		}

		data.ToJSON()
	}
}

func BenchmarkSerializeDataSnakeToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := &DataSnake{
			ID:                "1",
			NameNameNome:      "Samsung Galaxy S10",
			FieldOne:          "field one",
			FieldTwo:          "field two",
			FieldThree:        "field three",
			BlaBleBloQuantity: 10,
			Images:            []string{"img1", "img2"},
		}

		data.ToJSON()
	}
}

func BenchmarkDeserializeDataCamelFromJSON(b *testing.B) {
	data := &DataCamel{
		ID:                "1",
		NameNameNome:      "Samsung Galaxy S10",
		FieldOne:          "field one",
		FieldTwo:          "field two",
		FieldThree:        "field three",
		BlaBleBloQuantity: 10,
		Images:            []string{"img1", "img2"},
	}

	rawData, _ := data.ToJSON()
	for i := 0; i < b.N; i++ {
		FromCamelJSON(rawData)
	}
}

func BenchmarkDeserializeDataSnakeFromJSON(b *testing.B) {
	data := &DataSnake{
		ID:                "1",
		NameNameNome:      "Samsung Galaxy S10",
		FieldOne:          "field one",
		FieldTwo:          "field two",
		FieldThree:        "field three",
		BlaBleBloQuantity: 10,
		Images:            []string{"img1", "img2"},
	}

	rawData, _ := data.ToJSON()
	for i := 0; i < b.N; i++ {
		FromSnakeJSON(rawData)
	}
}

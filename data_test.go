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
			FieldFour:         "field four",
			FieldFive:         "field five",
			FieldSix:          "field six",
			FieldSeven:        "field seven",
			FieldEight:        "field eight",
			FieldNine:         "field nine",
			FieldTen:          "field ten",
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
			FieldFour:         "field four",
			FieldFive:         "field five",
			FieldSix:          "field six",
			FieldSeven:        "field seven",
			FieldEight:        "field eight",
			FieldNine:         "field nine",
			FieldTen:          "field ten",
			BlaBleBloQuantity: 10,
			Images:            []string{"img1", "img2"},
		}

		data.ToJSON()
	}
}

func BenchmarkDeserializeDataCamelFromJSON(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := &DataCamel{
			ID:                "1",
			NameNameNome:      "Samsung Galaxy S10",
			FieldOne:          "field one",
			FieldTwo:          "field two",
			FieldThree:        "field three",
			FieldFour:         "field four",
			FieldFive:         "field five",
			FieldSix:          "field six",
			FieldSeven:        "field seven",
			FieldEight:        "field eight",
			FieldNine:         "field nine",
			FieldTen:          "field ten",
			BlaBleBloQuantity: 10,
			Images:            []string{"img1", "img2"},
		}

		rawData, _ := data.ToJSON()

		FromCamelJSON(rawData)
	}
}

func BenchmarkDeserializeDataSnakeFromJSON(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := &DataSnake{
			ID:                "1",
			NameNameNome:      "Samsung Galaxy S10",
			FieldOne:          "field one",
			FieldTwo:          "field two",
			FieldThree:        "field three",
			FieldFour:         "field four",
			FieldFive:         "field five",
			FieldSix:          "field six",
			FieldSeven:        "field seven",
			FieldEight:        "field eight",
			FieldNine:         "field nine",
			FieldTen:          "field ten",
			BlaBleBloQuantity: 10,
			Images:            []string{"img1", "img2"},
		}

		rawData, _ := data.ToJSON()

		FromSnakeJSON(rawData)
	}
}

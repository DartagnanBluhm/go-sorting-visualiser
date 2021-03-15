package util

import "encoding/json"

type Payload struct {
	AlgorithmType string        `json:"algorithmType"`
	ArrayLength   int           `json:"arrayLength"`
	Array         []PayloadData `json:"array"`
}

type PayloadData struct {
	Value int `json:"value"`
}

type SortChanges struct {
	FirstIndex  int `json:"first-index"`
	SecondIndex int `json:"second-index"`
	FirstValue  int `json:"first-value"`
	SecondValue int `json:"second-value"`
}

func ConvertJSONtoArray(data []byte) ([]int, string, error) {
	var formData Payload
	err := json.Unmarshal(data, &formData)
	if err != nil {
		return []int{}, "", err
	}
	res := make([]int, formData.ArrayLength)
	for i := range formData.Array {
		res[i] = formData.Array[i].Value
	}
	return res, formData.AlgorithmType, nil
}

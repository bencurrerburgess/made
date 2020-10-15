package logic

import (
	"testing"
)

type TestData struct {
	TestCases []struct{
		TestName string `json:"testCase"`
		Data []int `json:"data"`
		ExpectedValue int `json:"expected"`
	}`json:"testCases"`
}

func TestFileLoad(t *testing.T){
	td := TestData{}
	err := LoadJsonFile("testdata.json", &td)
	if err != nil {
		t.Errorf("%v", err)
	}
	expectedName := "test1"
	if td.TestCases[0].TestName != expectedName {
		t.Errorf("expected first test to be named %v", expectedName)
	}

	expectedValue := td.TestCases[0].ExpectedValue
	actualValue := 0
	for _, num := range td.TestCases[0].Data {
		actualValue += num
	}
	if actualValue != expectedValue {
		t.Errorf("expected %v but got %v. %v", expectedValue, actualValue, td.TestCases[0])
	}
}






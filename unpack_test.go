package unpack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCaseType struct {
	testData       string
	expectedResult string
}

var testCases = []TestCaseType{
	{
		testData:       `a4bc2d5e`,
		expectedResult: `aaaabccddddde`,
	},
	{
		testData:       `a4bc0d5e`,
		expectedResult: `aaaabddddde`,
	},
	{
		testData:       `a04`,
		expectedResult: ``,
	},
	{
		testData:       `abcd`,
		expectedResult: `abcd`,
	},
	{
		testData:       `45`,
		expectedResult: ``,
	},
	{
		testData:       `qwe\4\5`,
		expectedResult: `qwe45`,
	},

	{
		testData:       `qwe\45`,
		expectedResult: `qwe44444`,
	},
	{
		testData:       `qwe\4\5`,
		expectedResult: `qwe45`,
	},
	{
		testData:       `qwe\\5`,
		expectedResult: `qwe\\\\\`,
	},
	{
		testData:       `a4b\4\5c0d5e\52qwe\4\5`,
		expectedResult: `aaaab45ddddde55qwe45`,
	},
}

func TestParser(t *testing.T) {
	for _, testCase := range testCases {
		require.Equal(t, Parser(testCase.testData), testCase.expectedResult)
	}
}

func TestParser2(t *testing.T) {
	for _, testCase := range testCases {
		require.Equal(t, Parser2(testCase.testData), testCase.expectedResult)
	}
}

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parser(`a4b\4\5c0d5e\52\\3qwe\4\5`)
	}
}

func BenchmarkParser2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parser2(`a4b\4\5c0d5e\52\\3qwe\4\5`)
	}
}

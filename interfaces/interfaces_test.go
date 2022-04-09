package interfaces

import (
	"bytes"
	"io"
	"testing"
)

type args struct {
	in io.ReadSeeker
}

func TestMyCopy(t * testing.T) {
	tests := []struct{
		name string
		args args
		expectedOutput string
		expectedError bool
	} {
		{
			name: "Base case",
			args: args{
				in: bytes.NewReader([]byte("example")),
			},
			expectedOutput: "exampleexample",
			expectedError: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := myCopy(testCase.args.in, out); err != nil && testCase.expectedError {
				t.Errorf("copy() error = %v, expectedError = %v", err, testCase.expectedError)
				return
			}
			if output := out.String(); output != testCase.expectedOutput {
				t.Errorf("copy() = %v, expectedOutput = %v", output, testCase.expectedOutput)
				return
			}
		})
	}
}
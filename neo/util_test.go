package neo

import "testing"

func TestGenerateCmd(t *testing.T) {
	type args struct {
		objIn interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := GenerateCmd(tt.args.objIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("GenerateCmd() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

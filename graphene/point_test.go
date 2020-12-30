package graphene

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"testing"
)

func TestNeoPoint(t *testing.T) {
	type args struct {
		p neo4j.Point3D
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "First Test Case",
			args: args{p: neo4j.Point3D{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultCypher := NeoPoint(tt.args.p)
			fmt.Println(resultCypher)
		})
	}
}
package test

import (
	"github.com/SynaesthesiaFlow/stateful-mesh-service/pkg/fixture"
	"github.com/nargetdev/neostate/graphene"
	"testing"
)

func TestFixture(t *testing.T) {
	f := fixture.CreateStrip(.5, 10)

	for p := range f.Points {
		graphene.NeoPoint(p)
	}
}
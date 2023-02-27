package docs

import (
	"github.com/CrestLee/selefra-provider-consul/resources"
	"github.com/selefra/selefra-provider-sdk/doc_gen"
	"testing"
)

func Test(t *testing.T) {
	if err := doc_gen.New(resources.GetSelefraProvider(), "./tables").Run(); err != nil {
		panic(err)
	}
}

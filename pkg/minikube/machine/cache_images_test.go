package machine

import (
	"io/ioutil"
	"testing"

	"k8s.io/minikube/pkg/minikube/constants"
)

func TestGetSrcRef(t *testing.T) {
	for _, image := range constants.LocalkubeCachedImages {
		if _, err := getSrcRef(image); err != nil {
			t.Errorf("Error getting src ref for %s: %s", image, err)
		}
	}
}

func TestGetDstRef(t *testing.T) {
	dst, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Error making temp directory: %s", err)
	}
	for _, image := range constants.LocalkubeCachedImages {
		if _, err := getDstRef(image, dst); err != nil {
			t.Errorf("Error getting src ref for %s: %s", image, err)
		}
	}
}

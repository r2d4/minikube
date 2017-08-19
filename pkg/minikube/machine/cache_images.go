package machine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/bootstrapper"
	"k8s.io/minikube/pkg/minikube/constants"

	"github.com/containers/image/copy"
	"github.com/containers/image/docker"
	"github.com/containers/image/docker/archive"
	"github.com/containers/image/signature"
	"github.com/containers/image/types"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const tempLoadDir = "/tmp"

func CacheAndLoadImagesParallel(cmd bootstrapper.CommandRunner, images []string, cacheDir string) error {
	fmt.Println("Cache dir: ", cacheDir)
	var g errgroup.Group
	for _, image := range images {
		image := image
		g.Go(func() error {
			dst := filepath.Join(cacheDir, image)
			dst = sanitizeCacheDir(dst)
			if err := CacheImage(image, dst); err != nil {
				return errors.Wrapf(err, "caching image %s", dst)
			}
			if err := LoadFromCache(cmd, dst); err != nil {
				return errors.Wrapf(err, "loading image from cache: %s", dst)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "caching images")
	}

	return nil
}

// # ParseReference cannot have a : in the directory path
func sanitizeCacheDir(image string) string {
	return strings.Replace(image, ":", "_", -1)
}

func LoadFromCache(cmd bootstrapper.CommandRunner, src string) error {
	glog.Infoln("Loading image from cache at ", src)
	filename := filepath.Base(src)
	dst := filepath.Join(tempLoadDir, filename)
	f, err := assets.NewFileAsset(src, tempLoadDir, filename, "0777")
	if err != nil {
		return errors.Wrapf(err, "creating copyable file asset: %s", filename)
	}
	if err := cmd.Copy(f); err != nil {
		return errors.Wrap(err, "transferring cached image")
	}

	dockerLoadCmd := "docker load -i " + dst

	if err := cmd.Run(dockerLoadCmd); err != nil {
		return errors.Wrapf(err, "loading docker image: %s", dst)
	}

	return nil
}

func getSrcRef(image string) (types.ImageReference, error) {
	srcRef, err := docker.ParseReference("//" + image)
	if err != nil {
		return nil, errors.Wrap(err, "parsing docker image src ref")
	}
	return srcRef, nil
}

func getDstRef(image, dst string) (types.ImageReference, error) {
	dstRef, err := archive.ParseReference(dst + ":" + image)
	if err != nil {
		return nil, errors.Wrap(err, "parsing docker archive dst ref")
	}
	return dstRef, nil
}

func CacheImage(image, dst string) error {
	glog.Infof("Attempting to cache image: %s at %s\n", image, dst)
	if _, err := os.Stat(dst); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(dst), 0777); err != nil {
		return errors.Wrapf(err, "making cache image directory: %s", dst)
	}

	srcRef, err := getSrcRef(image)
	if err != nil {
		return errors.Wrap(err, "creating docker image src ref")
	}

	dstRef, err := getDstRef(image, dst)
	if err != nil {
		return errors.Wrap(err, "creating docker archive dst ref")
	}

	policy, err := signature.DefaultPolicy(nil)
	if err != nil {
		return errors.Wrap(err, "getting default signature policy")
	}

	policyContext, err := signature.NewPolicyContext(policy)
	if err != nil {
		return errors.Wrap(err, "getting policy context")
	}

	sourceCtx := &types.SystemContext{
		// By default, the image library will try to look at /etc/docker/certs.d
		// As a non-root user, this would result in a permissions error,
		// so, we skip this step by just looking in the minikube home directory.
		DockerCertPath: constants.MakeMiniPath("cache"),
	}

	err = copy.Image(policyContext, dstRef, srcRef, &copy.Options{
		SourceCtx: sourceCtx,
	})
	if err != nil {
		return errors.Wrap(err, "copying image")
	}

	return nil
}

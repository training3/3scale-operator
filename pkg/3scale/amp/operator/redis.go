package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
	v1 "k8s.io/api/core/v1"
)

func (o *OperatorRedisOptionsProvider) GetRedisOptions() (*component.RedisOptions, error) {
	optProv := component.RedisOptionsBuilder{}

	imageProvider := product.CurrentImageProvider()

	optProv.AppLabel(*o.APIManagerSpec.AppLabel)
	optProv.AMPRelease(product.ThreescaleRelease)
	optProv.InsecureImportPolicy(*o.APIManagerSpec.ImageStreamTagImportInsecure)

	if o.APIManagerSpec.Backend != nil && o.APIManagerSpec.Backend.RedisImage != nil {
		optProv.BackendImage(*o.APIManagerSpec.Backend.RedisImage)
	} else {
		optProv.BackendImage(imageProvider.GetBackendRedisImage())
	}

	if o.APIManagerSpec.System != nil && o.APIManagerSpec.System.RedisImage != nil {
		optProv.SystemImage(*o.APIManagerSpec.System.RedisImage)
	} else {
		optProv.SystemImage(imageProvider.GetSystemRedisImage())
	}

	o.setResourceRequirementsOptions(&optProv)

	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create Redis Options - %s", err)
	}
	return res, nil
}

func (o *OperatorRedisOptionsProvider) setResourceRequirementsOptions(b *component.RedisOptionsBuilder) {
	if !*o.APIManagerSpec.ResourceRequirementsEnabled {
		b.SystemRedisContainerResourceRequirements(v1.ResourceRequirements{})
		b.BackendRedisContainerResourceRequirements(v1.ResourceRequirements{})
	}
}

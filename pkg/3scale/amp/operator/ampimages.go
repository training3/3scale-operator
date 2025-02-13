package operator

import (
	"fmt"

	"github.com/3scale/3scale-operator/pkg/3scale/amp/component"
	"github.com/3scale/3scale-operator/pkg/3scale/amp/product"
)

func (o *OperatorAmpImagesOptionsProvider) GetAmpImagesOptions() (*component.AmpImagesOptions, error) {
	optProv := component.AmpImagesOptionsBuilder{}

	imageProvider := product.CurrentImageProvider()

	optProv.AppLabel(*o.APIManagerSpec.AppLabel)
	optProv.AMPRelease(product.ThreescaleRelease)
	if o.APIManagerSpec.Apicast != nil && o.APIManagerSpec.Apicast.Image != nil {
		optProv.ApicastImage(*o.APIManagerSpec.Apicast.Image)
	} else {
		optProv.ApicastImage(imageProvider.GetApicastImage())
	}

	if o.APIManagerSpec.Backend != nil && o.APIManagerSpec.Backend.Image != nil {
		optProv.BackendImage(*o.APIManagerSpec.Backend.Image)
	} else {
		optProv.BackendImage(imageProvider.GetBackendImage())
	}

	if o.APIManagerSpec.System != nil && o.APIManagerSpec.System.Image != nil {
		optProv.SystemImage(*o.APIManagerSpec.System.Image)
	} else {
		optProv.SystemImage(imageProvider.GetSystemImage())
	}

	if o.APIManagerSpec.Zync != nil && o.APIManagerSpec.Zync.Image != nil {
		optProv.ZyncImage(*o.APIManagerSpec.Zync.Image)
	} else {
		optProv.ZyncImage(imageProvider.GetZyncImage())
	}

	if o.APIManagerSpec.Zync != nil && o.APIManagerSpec.Zync.PostgreSQLImage != nil {
		optProv.ZyncDatabasePostgreSQLImage(*o.APIManagerSpec.Zync.PostgreSQLImage)
	} else {
		optProv.ZyncDatabasePostgreSQLImage(imageProvider.GetZyncPostgreSQLImage())
	}

	if o.APIManagerSpec.System != nil && o.APIManagerSpec.System.MemcachedImage != nil {
		optProv.SystemMemcachedImage(*o.APIManagerSpec.System.MemcachedImage)
	} else {
		optProv.SystemMemcachedImage(imageProvider.GetSystemMemcachedImage())
	}

	optProv.InsecureImportPolicy(*o.APIManagerSpec.ImageStreamTagImportInsecure)
	res, err := optProv.Build()
	if err != nil {
		return nil, fmt.Errorf("unable to create AMPImages Options - %s", err)
	}
	return res, nil
}

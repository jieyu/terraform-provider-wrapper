package wrapper

import (
	"github.com/hashicorp/terraform/terraform"
)

type Provider struct {
	terraform.ResourceProvider
	interceptor *Interceptor
}

func NewProvider(
	provider terraform.ResourceProvider,
	interceptor *Interceptor) terraform.ResourceProvider {
	return &Provider{
		ResourceProvider: provider,
		interceptor:      interceptor,
	}
}

func (provider *Provider) GetSchema(
	request *terraform.ProviderSchemaRequest) (*terraform.ProviderSchema, error) {
	if provider.interceptor == nil ||
		provider.interceptor.GetSchema == nil {
		return provider.ResourceProvider.GetSchema(request)
	}

	return provider.interceptor.GetSchema(
		provider.ResourceProvider.GetSchema, request)
}

func (provider *Provider) Input(
	input terraform.UIInput,
	config *terraform.ResourceConfig) (*terraform.ResourceConfig, error) {
	if provider.interceptor == nil ||
		provider.interceptor.Input == nil {
		return provider.ResourceProvider.Input(input, config)
	}
	return provider.interceptor.Input(
		provider.ResourceProvider.Input,
		input,
		config)
}

func (provider *Provider) Validate(
	config *terraform.ResourceConfig) ([]string, []error) {
	if provider.interceptor == nil ||
		provider.interceptor.Validate == nil {
		return provider.ResourceProvider.Validate(config)
	}

	return provider.interceptor.Validate(
		provider.ResourceProvider.Validate,
		config)
}

func (provider *Provider) Configure(
	config *terraform.ResourceConfig) error {
	if provider.interceptor == nil ||
		provider.interceptor.Configure == nil {
		return provider.ResourceProvider.Configure(config)
	}

	return provider.interceptor.Configure(
		provider.ResourceProvider.Configure,
		config)
}

func (provider *Provider) Resources() []terraform.ResourceType {
	if provider.interceptor == nil ||
		provider.interceptor.Resources == nil {
		return provider.ResourceProvider.Resources()
	}

	return provider.interceptor.Resources(
		provider.ResourceProvider.Resources)
}

func (provider *Provider) ValidateResource(
	resourceType string,
	config *terraform.ResourceConfig) ([]string, []error) {
	if provider.interceptor == nil ||
		provider.interceptor.ValidateResource == nil {
		return provider.ResourceProvider.ValidateResource(resourceType, config)
	}

	return provider.interceptor.ValidateResource(
		provider.ResourceProvider.ValidateResource,
		resourceType,
		config)
}

func (provider *Provider) Apply(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState,
	diff *terraform.InstanceDiff) (*terraform.InstanceState, error) {
	if provider.interceptor == nil ||
		provider.interceptor.Apply == nil {
		return provider.ResourceProvider.Apply(info, state, diff)
	}

	return provider.interceptor.Apply(
		provider.ResourceProvider.Apply,
		info,
		state,
		diff)
}

func (provider *Provider) Diff(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState,
	config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	if provider.interceptor == nil ||
		provider.interceptor.Diff == nil {
		return provider.ResourceProvider.Diff(info, state, config)
	}

	return provider.interceptor.Diff(
		provider.ResourceProvider.Diff,
		info,
		state,
		config)
}

func (provider *Provider) Refresh(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState) (*terraform.InstanceState, error) {
	if provider.interceptor == nil ||
		provider.interceptor.Refresh == nil {
		return provider.ResourceProvider.Refresh(info, state)
	}

	return provider.interceptor.Refresh(
		provider.ResourceProvider.Refresh,
		info,
		state)
}

func (provider *Provider) ImportState(
	info *terraform.InstanceInfo,
	id string) ([]*terraform.InstanceState, error) {
	if provider.interceptor == nil ||
		provider.interceptor.ImportState == nil {
		return provider.ResourceProvider.ImportState(info, id)
	}

	return provider.interceptor.ImportState(
		provider.ResourceProvider.ImportState,
		info,
		id)
}

func (provider *Provider) ValidateDataSource(
	dataSourceType string,
	config *terraform.ResourceConfig) ([]string, []error) {
	if provider.interceptor == nil ||
		provider.interceptor.ValidateDataSource == nil {
		return provider.ResourceProvider.ValidateDataSource(dataSourceType, config)
	}

	return provider.interceptor.ValidateDataSource(
		provider.ResourceProvider.ValidateDataSource,
		dataSourceType,
		config)
}

func (provider *Provider) DataSources() []terraform.DataSource {
	if provider.interceptor == nil ||
		provider.interceptor.DataSources == nil {
		return provider.ResourceProvider.DataSources()
	}

	return provider.interceptor.DataSources(
		provider.ResourceProvider.DataSources)
}

func (provider *Provider) ReadDataDiff(
	info *terraform.InstanceInfo,
	config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	if provider.interceptor == nil ||
		provider.interceptor.ReadDataDiff == nil {
		return provider.ResourceProvider.ReadDataDiff(info, config)
	}

	return provider.interceptor.ReadDataDiff(
		provider.ResourceProvider.ReadDataDiff,
		info,
		config)
}

func (provider *Provider) ReadDataApply(
	info *terraform.InstanceInfo,
	diff *terraform.InstanceDiff) (*terraform.InstanceState, error) {
	if provider.interceptor == nil ||
		provider.interceptor.ReadDataApply == nil {
		return provider.ResourceProvider.ReadDataApply(info, diff)
	}

	return provider.interceptor.ReadDataApply(
		provider.ResourceProvider.ReadDataApply,
		info,
		diff)
}

package wrapper

import (
	"github.com/hashicorp/terraform/terraform"
)

type Provider struct {
	terraform.ResourceProvider
}

func NewProvider(provider terraform.ResourceProvider) terraform.ResourceProvider {
	return &Provider{
		ResourceProvider: provider,
	}
}

func (provider *Provider) GetSchema(
	request *terraform.ProviderSchemaRequest) (*terraform.ProviderSchema, error) {
	return provider.ResourceProvider.GetSchema(request)
}

func (provider *Provider) Input(
	input terraform.UIInput,
	config *terraform.ResourceConfig) (*terraform.ResourceConfig, error) {
	return provider.ResourceProvider.Input(input, config)
}

func (provider *Provider) Validate(
	config *terraform.ResourceConfig) ([]string, []error) {
	return provider.ResourceProvider.Validate(config)
}

func (provider *Provider) Configure(
	config *terraform.ResourceConfig) error {
	return provider.ResourceProvider.Configure(config)
}

func (provider *Provider) Resources() []terraform.ResourceType {
	return provider.ResourceProvider.Resources()
}

func (provider *Provider) ValidateResource(
	resourceType string,
	config *terraform.ResourceConfig) ([]string, []error) {
	return provider.ResourceProvider.ValidateResource(resourceType, config)
}

func (provider *Provider) Apply(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState,
	diff *terraform.InstanceDiff) (*terraform.InstanceState, error) {
	return provider.ResourceProvider.Apply(info, state, diff)
}

func (provider *Provider) Diff(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState,
	config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	return provider.ResourceProvider.Diff(info, state, config)
}

func (provider *Provider) Refresh(
	info *terraform.InstanceInfo,
	state *terraform.InstanceState) (*terraform.InstanceState, error) {
	return provider.ResourceProvider.Refresh(info, state)
}

func (provider *Provider) ImportState(
	info *terraform.InstanceInfo,
	id string) ([]*terraform.InstanceState, error) {
	return provider.ResourceProvider.ImportState(info, id)
}

func (provider *Provider) ValidateDataSource(
	dataSourceType string,
	config *terraform.ResourceConfig) ([]string, []error) {
	return provider.ResourceProvider.ValidateDataSource(dataSourceType, config)
}

func (provider *Provider) DataSources() []terraform.DataSource {
	return provider.ResourceProvider.DataSources()
}

func (provider *Provider) ReadDataDiff(
	info *terraform.InstanceInfo,
	config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	return provider.ResourceProvider.ReadDataDiff(info, config)
}

func (provider *Provider) ReadDataApply(
	info *terraform.InstanceInfo,
	diff *terraform.InstanceDiff) (*terraform.InstanceState, error) {
	return provider.ResourceProvider.ReadDataApply(info, diff)
}

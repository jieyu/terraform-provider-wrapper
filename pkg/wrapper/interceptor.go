package wrapper

import (
	"github.com/hashicorp/terraform/terraform"
)

type (
	GetSchemaHandler          func(*terraform.ProviderSchemaRequest) (*terraform.ProviderSchema, error)
	InputHandler              func(terraform.UIInput, *terraform.ResourceConfig) (*terraform.ResourceConfig, error)
	ValidateHandler           func(*terraform.ResourceConfig) ([]string, []error)
	ConfigureHandler          func(*terraform.ResourceConfig) error
	ResourcesHandler          func() []terraform.ResourceType
	StopHandler               func() error
	ValidateResourceHandler   func(string, *terraform.ResourceConfig) ([]string, []error)
	ApplyHandler              func(*terraform.InstanceInfo, *terraform.InstanceState, *terraform.InstanceDiff) (*terraform.InstanceState, error)
	DiffHandler               func(*terraform.InstanceInfo, *terraform.InstanceState, *terraform.ResourceConfig) (*terraform.InstanceDiff, error)
	RefreshHandler            func(*terraform.InstanceInfo, *terraform.InstanceState) (*terraform.InstanceState, error)
	ImportStateHandler        func(*terraform.InstanceInfo, string) ([]*terraform.InstanceState, error)
	ValidateDataSourceHandler func(string, *terraform.ResourceConfig) ([]string, []error)
	DataSourcesHandler        func() []terraform.DataSource
	ReadDataDiffHandler       func(*terraform.InstanceInfo, *terraform.ResourceConfig) (*terraform.InstanceDiff, error)
	ReadDataApplyHandler      func(*terraform.InstanceInfo, *terraform.InstanceDiff) (*terraform.InstanceState, error)

	GetSchemaInterceptor          func(GetSchemaHandler, *terraform.ProviderSchemaRequest) (*terraform.ProviderSchema, error)
	InputInterceptor              func(InputHandler, terraform.UIInput, *terraform.ResourceConfig) (*terraform.ResourceConfig, error)
	ValidateInterceptor           func(ValidateHandler, *terraform.ResourceConfig) ([]string, []error)
	ConfigureInterceptor          func(ConfigureHandler, *terraform.ResourceConfig) error
	ResourcesInterceptor          func(ResourcesHandler) []terraform.ResourceType
	StopInterceptor               func(StopHandler) error
	ValidateResourceInterceptor   func(ValidateResourceHandler, string, *terraform.ResourceConfig) ([]string, []error)
	ApplyInterceptor              func(ApplyHandler, *terraform.InstanceInfo, *terraform.InstanceState, *terraform.InstanceDiff) (*terraform.InstanceState, error)
	DiffInterceptor               func(DiffHandler, *terraform.InstanceInfo, *terraform.InstanceState, *terraform.ResourceConfig) (*terraform.InstanceDiff, error)
	RefreshInterceptor            func(RefreshHandler, *terraform.InstanceInfo, *terraform.InstanceState) (*terraform.InstanceState, error)
	ImportStateInterceptor        func(ImportStateHandler, *terraform.InstanceInfo, string) ([]*terraform.InstanceState, error)
	ValidateDataSourceInterceptor func(ValidateDataSourceHandler, string, *terraform.ResourceConfig) ([]string, []error)
	DataSourcesInterceptor        func(DataSourcesHandler) []terraform.DataSource
	ReadDataDiffInterceptor       func(ReadDataDiffHandler, *terraform.InstanceInfo, *terraform.ResourceConfig) (*terraform.InstanceDiff, error)
	ReadDataApplyInterceptor      func(ReadDataApplyHandler, *terraform.InstanceInfo, *terraform.InstanceDiff) (*terraform.InstanceState, error)
)

type Interceptor struct {
	GetSchema          GetSchemaInterceptor
	Input              InputInterceptor
	Validate           ValidateInterceptor
	Configure          ConfigureInterceptor
	Resources          ResourcesInterceptor
	Stop               StopInterceptor
	ValidateResource   ValidateResourceInterceptor
	Apply              ApplyInterceptor
	Diff               DiffInterceptor
	Refresh            RefreshInterceptor
	ImportState        ImportStateInterceptor
	ValidateDataSource ValidateDataSourceInterceptor
	DataSources        DataSourcesInterceptor
	ReadDataDiff       ReadDataDiffInterceptor
	ReadDataApply      ReadDataApplyInterceptor
}

func NewInterceptor() *Interceptor {
	return &Interceptor{}
}

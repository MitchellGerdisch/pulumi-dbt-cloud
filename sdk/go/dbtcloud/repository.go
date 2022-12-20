// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Repository struct {
	pulumi.CustomResourceState

	// Public key generated by DBT when using `deploy_key` clone strategy
	DeployKey pulumi.StringOutput `pulumi:"deployKey"`
	// Whether we should return the public deploy key
	FetchDeployKey pulumi.BoolPtrOutput `pulumi:"fetchDeployKey"`
	// Git clone strategy for the repository
	GitCloneStrategy pulumi.StringPtrOutput `pulumi:"gitCloneStrategy"`
	// Identifier for the Gitlab project
	GitlabProjectId pulumi.IntPtrOutput `pulumi:"gitlabProjectId"`
	// Whether the repository is active
	IsActive pulumi.BoolPtrOutput `pulumi:"isActive"`
	// Project ID to create the repository in
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Git URL for the repository or <Group>/<Project> for Gitlab
	RemoteUrl pulumi.StringOutput `pulumi:"remoteUrl"`
	// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
	RepositoryCredentialsId pulumi.IntPtrOutput `pulumi:"repositoryCredentialsId"`
	// Repository Identifier
	RepositoryId pulumi.IntOutput `pulumi:"repositoryId"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.RemoteUrl == nil {
		return nil, errors.New("invalid value for required argument 'RemoteUrl'")
	}
	var resource Repository
	err := ctx.RegisterResource("dbtcloud:index/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("dbtcloud:index/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Public key generated by DBT when using `deploy_key` clone strategy
	DeployKey *string `pulumi:"deployKey"`
	// Whether we should return the public deploy key
	FetchDeployKey *bool `pulumi:"fetchDeployKey"`
	// Git clone strategy for the repository
	GitCloneStrategy *string `pulumi:"gitCloneStrategy"`
	// Identifier for the Gitlab project
	GitlabProjectId *int `pulumi:"gitlabProjectId"`
	// Whether the repository is active
	IsActive *bool `pulumi:"isActive"`
	// Project ID to create the repository in
	ProjectId *int `pulumi:"projectId"`
	// Git URL for the repository or <Group>/<Project> for Gitlab
	RemoteUrl *string `pulumi:"remoteUrl"`
	// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
	RepositoryCredentialsId *int `pulumi:"repositoryCredentialsId"`
	// Repository Identifier
	RepositoryId *int `pulumi:"repositoryId"`
}

type RepositoryState struct {
	// Public key generated by DBT when using `deploy_key` clone strategy
	DeployKey pulumi.StringPtrInput
	// Whether we should return the public deploy key
	FetchDeployKey pulumi.BoolPtrInput
	// Git clone strategy for the repository
	GitCloneStrategy pulumi.StringPtrInput
	// Identifier for the Gitlab project
	GitlabProjectId pulumi.IntPtrInput
	// Whether the repository is active
	IsActive pulumi.BoolPtrInput
	// Project ID to create the repository in
	ProjectId pulumi.IntPtrInput
	// Git URL for the repository or <Group>/<Project> for Gitlab
	RemoteUrl pulumi.StringPtrInput
	// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
	RepositoryCredentialsId pulumi.IntPtrInput
	// Repository Identifier
	RepositoryId pulumi.IntPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Whether we should return the public deploy key
	FetchDeployKey *bool `pulumi:"fetchDeployKey"`
	// Git clone strategy for the repository
	GitCloneStrategy *string `pulumi:"gitCloneStrategy"`
	// Identifier for the Gitlab project
	GitlabProjectId *int `pulumi:"gitlabProjectId"`
	// Whether the repository is active
	IsActive *bool `pulumi:"isActive"`
	// Project ID to create the repository in
	ProjectId int `pulumi:"projectId"`
	// Git URL for the repository or <Group>/<Project> for Gitlab
	RemoteUrl string `pulumi:"remoteUrl"`
	// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
	RepositoryCredentialsId *int `pulumi:"repositoryCredentialsId"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Whether we should return the public deploy key
	FetchDeployKey pulumi.BoolPtrInput
	// Git clone strategy for the repository
	GitCloneStrategy pulumi.StringPtrInput
	// Identifier for the Gitlab project
	GitlabProjectId pulumi.IntPtrInput
	// Whether the repository is active
	IsActive pulumi.BoolPtrInput
	// Project ID to create the repository in
	ProjectId pulumi.IntInput
	// Git URL for the repository or <Group>/<Project> for Gitlab
	RemoteUrl pulumi.StringInput
	// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
	RepositoryCredentialsId pulumi.IntPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//          RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//          RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Public key generated by DBT when using `deploy_key` clone strategy
func (o RepositoryOutput) DeployKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.DeployKey }).(pulumi.StringOutput)
}

// Whether we should return the public deploy key
func (o RepositoryOutput) FetchDeployKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.BoolPtrOutput { return v.FetchDeployKey }).(pulumi.BoolPtrOutput)
}

// Git clone strategy for the repository
func (o RepositoryOutput) GitCloneStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.GitCloneStrategy }).(pulumi.StringPtrOutput)
}

// Identifier for the Gitlab project
func (o RepositoryOutput) GitlabProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntPtrOutput { return v.GitlabProjectId }).(pulumi.IntPtrOutput)
}

// Whether the repository is active
func (o RepositoryOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.BoolPtrOutput { return v.IsActive }).(pulumi.BoolPtrOutput)
}

// Project ID to create the repository in
func (o RepositoryOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Git URL for the repository or <Group>/<Project> for Gitlab
func (o RepositoryOutput) RemoteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RemoteUrl }).(pulumi.StringOutput)
}

// Credentials ID for the repository (From the repository side not the DBT Cloud ID)
func (o RepositoryOutput) RepositoryCredentialsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntPtrOutput { return v.RepositoryCredentialsId }).(pulumi.IntPtrOutput)
}

// Repository Identifier
func (o RepositoryOutput) RepositoryId() pulumi.IntOutput {
	return o.ApplyT(func(v *Repository) pulumi.IntOutput { return v.RepositoryId }).(pulumi.IntOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
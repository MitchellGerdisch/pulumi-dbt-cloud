// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("dbtcloud:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	ProjectId int `pulumi:"projectId"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	ConnectionId int `pulumi:"connectionId"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	Name         string `pulumi:"name"`
	ProjectId    int    `pulumi:"projectId"`
	RepositoryId int    `pulumi:"repositoryId"`
	State        int    `pulumi:"state"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	ProjectId pulumi.IntInput `pulumi:"projectId"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ConnectionId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.ConnectionId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) RepositoryId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.RepositoryId }).(pulumi.IntOutput)
}

func (o LookupProjectResultOutput) State() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.State }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}

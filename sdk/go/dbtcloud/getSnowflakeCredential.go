// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnowflakeCredential(ctx *pulumi.Context, args *LookupSnowflakeCredentialArgs, opts ...pulumi.InvokeOption) (*LookupSnowflakeCredentialResult, error) {
	var rv LookupSnowflakeCredentialResult
	err := ctx.Invoke("dbtcloud:index/getSnowflakeCredential:getSnowflakeCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnowflakeCredential.
type LookupSnowflakeCredentialArgs struct {
	CredentialId int `pulumi:"credentialId"`
	ProjectId    int `pulumi:"projectId"`
}

// A collection of values returned by getSnowflakeCredential.
type LookupSnowflakeCredentialResult struct {
	AuthType     string `pulumi:"authType"`
	CredentialId int    `pulumi:"credentialId"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	IsActive   bool   `pulumi:"isActive"`
	NumThreads int    `pulumi:"numThreads"`
	ProjectId  int    `pulumi:"projectId"`
	Schema     string `pulumi:"schema"`
	User       string `pulumi:"user"`
}

func LookupSnowflakeCredentialOutput(ctx *pulumi.Context, args LookupSnowflakeCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupSnowflakeCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnowflakeCredentialResult, error) {
			args := v.(LookupSnowflakeCredentialArgs)
			r, err := LookupSnowflakeCredential(ctx, &args, opts...)
			var s LookupSnowflakeCredentialResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnowflakeCredentialResultOutput)
}

// A collection of arguments for invoking getSnowflakeCredential.
type LookupSnowflakeCredentialOutputArgs struct {
	CredentialId pulumi.IntInput `pulumi:"credentialId"`
	ProjectId    pulumi.IntInput `pulumi:"projectId"`
}

func (LookupSnowflakeCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnowflakeCredentialArgs)(nil)).Elem()
}

// A collection of values returned by getSnowflakeCredential.
type LookupSnowflakeCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupSnowflakeCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnowflakeCredentialResult)(nil)).Elem()
}

func (o LookupSnowflakeCredentialResultOutput) ToLookupSnowflakeCredentialResultOutput() LookupSnowflakeCredentialResultOutput {
	return o
}

func (o LookupSnowflakeCredentialResultOutput) ToLookupSnowflakeCredentialResultOutputWithContext(ctx context.Context) LookupSnowflakeCredentialResultOutput {
	return o
}

func (o LookupSnowflakeCredentialResultOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o LookupSnowflakeCredentialResultOutput) CredentialId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) int { return v.CredentialId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSnowflakeCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSnowflakeCredentialResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

func (o LookupSnowflakeCredentialResultOutput) NumThreads() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) int { return v.NumThreads }).(pulumi.IntOutput)
}

func (o LookupSnowflakeCredentialResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o LookupSnowflakeCredentialResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) string { return v.Schema }).(pulumi.StringOutput)
}

func (o LookupSnowflakeCredentialResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnowflakeCredentialResult) string { return v.User }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnowflakeCredentialResultOutput{})
}
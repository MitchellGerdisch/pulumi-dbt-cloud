// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbtcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUser(ctx *pulumi.Context, args *GetUserArgs, opts ...pulumi.InvokeOption) (*GetUserResult, error) {
	var rv GetUserResult
	err := ctx.Invoke("dbtcloud:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	Email string `pulumi:"email"`
}

// A collection of values returned by getUser.
type GetUserResult struct {
	Email string `pulumi:"email"`
	Id    int    `pulumi:"id"`
}

func GetUserOutput(ctx *pulumi.Context, args GetUserOutputArgs, opts ...pulumi.InvokeOption) GetUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserResult, error) {
			args := v.(GetUserArgs)
			r, err := GetUser(ctx, &args, opts...)
			var s GetUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserResultOutput)
}

// A collection of arguments for invoking getUser.
type GetUserOutputArgs struct {
	Email pulumi.StringInput `pulumi:"email"`
}

func (GetUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type GetUserResultOutput struct{ *pulumi.OutputState }

func (GetUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserResult)(nil)).Elem()
}

func (o GetUserResultOutput) ToGetUserResultOutput() GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) ToGetUserResultOutputWithContext(ctx context.Context) GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetUserResultOutput) Id() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.Id }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserResultOutput{})
}

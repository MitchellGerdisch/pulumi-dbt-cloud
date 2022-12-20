# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[int],
                 token: pulumi.Input[str],
                 host_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[int] account_id: Account identifier for your DBT Cloud implementation
        :param pulumi.Input[str] token: API token for your DBT Cloud
        :param pulumi.Input[str] host_url: URL for your DBT Cloud deployment
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "token", token)
        if host_url is not None:
            pulumi.set(__self__, "host_url", host_url)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[int]:
        """
        Account identifier for your DBT Cloud implementation
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[str]:
        """
        API token for your DBT Cloud
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="hostUrl")
    def host_url(self) -> Optional[pulumi.Input[str]]:
        """
        URL for your DBT Cloud deployment
        """
        return pulumi.get(self, "host_url")

    @host_url.setter
    def host_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[int]] = None,
                 host_url: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the dbtcloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] account_id: Account identifier for your DBT Cloud implementation
        :param pulumi.Input[str] host_url: URL for your DBT Cloud deployment
        :param pulumi.Input[str] token: API token for your DBT Cloud
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the dbtcloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[int]] = None,
                 host_url: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = pulumi.Output.from_input(account_id).apply(pulumi.runtime.to_json) if account_id is not None else None
            __props__.__dict__["host_url"] = host_url
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = token
        super(Provider, __self__).__init__(
            'dbtcloud',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="hostUrl")
    def host_url(self) -> pulumi.Output[Optional[str]]:
        """
        URL for your DBT Cloud deployment
        """
        return pulumi.get(self, "host_url")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        API token for your DBT Cloud
        """
        return pulumi.get(self, "token")


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetEnvironmentResult',
    'AwaitableGetEnvironmentResult',
    'get_environment',
    'get_environment_output',
]

@pulumi.output_type
class GetEnvironmentResult:
    """
    A collection of values returned by getEnvironment.
    """
    def __init__(__self__, credential_id=None, custom_branch=None, dbt_version=None, environment_id=None, id=None, is_active=None, name=None, project_id=None, type=None, use_custom_branch=None):
        if credential_id and not isinstance(credential_id, int):
            raise TypeError("Expected argument 'credential_id' to be a int")
        pulumi.set(__self__, "credential_id", credential_id)
        if custom_branch and not isinstance(custom_branch, str):
            raise TypeError("Expected argument 'custom_branch' to be a str")
        pulumi.set(__self__, "custom_branch", custom_branch)
        if dbt_version and not isinstance(dbt_version, str):
            raise TypeError("Expected argument 'dbt_version' to be a str")
        pulumi.set(__self__, "dbt_version", dbt_version)
        if environment_id and not isinstance(environment_id, int):
            raise TypeError("Expected argument 'environment_id' to be a int")
        pulumi.set(__self__, "environment_id", environment_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_custom_branch and not isinstance(use_custom_branch, bool):
            raise TypeError("Expected argument 'use_custom_branch' to be a bool")
        pulumi.set(__self__, "use_custom_branch", use_custom_branch)

    @property
    @pulumi.getter(name="credentialId")
    def credential_id(self) -> int:
        return pulumi.get(self, "credential_id")

    @property
    @pulumi.getter(name="customBranch")
    def custom_branch(self) -> str:
        return pulumi.get(self, "custom_branch")

    @property
    @pulumi.getter(name="dbtVersion")
    def dbt_version(self) -> str:
        return pulumi.get(self, "dbt_version")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> int:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useCustomBranch")
    def use_custom_branch(self) -> bool:
        return pulumi.get(self, "use_custom_branch")


class AwaitableGetEnvironmentResult(GetEnvironmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentResult(
            credential_id=self.credential_id,
            custom_branch=self.custom_branch,
            dbt_version=self.dbt_version,
            environment_id=self.environment_id,
            id=self.id,
            is_active=self.is_active,
            name=self.name,
            project_id=self.project_id,
            type=self.type,
            use_custom_branch=self.use_custom_branch)


def get_environment(environment_id: Optional[int] = None,
                    project_id: Optional[int] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['environmentId'] = environment_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/getEnvironment:getEnvironment', __args__, opts=opts, typ=GetEnvironmentResult).value

    return AwaitableGetEnvironmentResult(
        credential_id=__ret__.credential_id,
        custom_branch=__ret__.custom_branch,
        dbt_version=__ret__.dbt_version,
        environment_id=__ret__.environment_id,
        id=__ret__.id,
        is_active=__ret__.is_active,
        name=__ret__.name,
        project_id=__ret__.project_id,
        type=__ret__.type,
        use_custom_branch=__ret__.use_custom_branch)


@_utilities.lift_output_func(get_environment)
def get_environment_output(environment_id: Optional[pulumi.Input[int]] = None,
                           project_id: Optional[pulumi.Input[int]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

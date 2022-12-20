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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
    'get_repository_output',
]

@pulumi.output_type
class GetRepositoryResult:
    """
    A collection of values returned by getRepository.
    """
    def __init__(__self__, deploy_key=None, fetch_deploy_key=None, git_clone_strategy=None, gitlab_project_id=None, id=None, is_active=None, project_id=None, remote_url=None, repository_credentials_id=None, repository_id=None):
        if deploy_key and not isinstance(deploy_key, str):
            raise TypeError("Expected argument 'deploy_key' to be a str")
        pulumi.set(__self__, "deploy_key", deploy_key)
        if fetch_deploy_key and not isinstance(fetch_deploy_key, bool):
            raise TypeError("Expected argument 'fetch_deploy_key' to be a bool")
        pulumi.set(__self__, "fetch_deploy_key", fetch_deploy_key)
        if git_clone_strategy and not isinstance(git_clone_strategy, str):
            raise TypeError("Expected argument 'git_clone_strategy' to be a str")
        pulumi.set(__self__, "git_clone_strategy", git_clone_strategy)
        if gitlab_project_id and not isinstance(gitlab_project_id, int):
            raise TypeError("Expected argument 'gitlab_project_id' to be a int")
        pulumi.set(__self__, "gitlab_project_id", gitlab_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if remote_url and not isinstance(remote_url, str):
            raise TypeError("Expected argument 'remote_url' to be a str")
        pulumi.set(__self__, "remote_url", remote_url)
        if repository_credentials_id and not isinstance(repository_credentials_id, int):
            raise TypeError("Expected argument 'repository_credentials_id' to be a int")
        pulumi.set(__self__, "repository_credentials_id", repository_credentials_id)
        if repository_id and not isinstance(repository_id, int):
            raise TypeError("Expected argument 'repository_id' to be a int")
        pulumi.set(__self__, "repository_id", repository_id)

    @property
    @pulumi.getter(name="deployKey")
    def deploy_key(self) -> str:
        return pulumi.get(self, "deploy_key")

    @property
    @pulumi.getter(name="fetchDeployKey")
    def fetch_deploy_key(self) -> Optional[bool]:
        return pulumi.get(self, "fetch_deploy_key")

    @property
    @pulumi.getter(name="gitCloneStrategy")
    def git_clone_strategy(self) -> str:
        return pulumi.get(self, "git_clone_strategy")

    @property
    @pulumi.getter(name="gitlabProjectId")
    def gitlab_project_id(self) -> int:
        return pulumi.get(self, "gitlab_project_id")

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="remoteUrl")
    def remote_url(self) -> str:
        return pulumi.get(self, "remote_url")

    @property
    @pulumi.getter(name="repositoryCredentialsId")
    def repository_credentials_id(self) -> int:
        return pulumi.get(self, "repository_credentials_id")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> int:
        return pulumi.get(self, "repository_id")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            deploy_key=self.deploy_key,
            fetch_deploy_key=self.fetch_deploy_key,
            git_clone_strategy=self.git_clone_strategy,
            gitlab_project_id=self.gitlab_project_id,
            id=self.id,
            is_active=self.is_active,
            project_id=self.project_id,
            remote_url=self.remote_url,
            repository_credentials_id=self.repository_credentials_id,
            repository_id=self.repository_id)


def get_repository(fetch_deploy_key: Optional[bool] = None,
                   project_id: Optional[int] = None,
                   repository_id: Optional[int] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['fetchDeployKey'] = fetch_deploy_key
    __args__['projectId'] = project_id
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('dbtcloud:index/getRepository:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        deploy_key=__ret__.deploy_key,
        fetch_deploy_key=__ret__.fetch_deploy_key,
        git_clone_strategy=__ret__.git_clone_strategy,
        gitlab_project_id=__ret__.gitlab_project_id,
        id=__ret__.id,
        is_active=__ret__.is_active,
        project_id=__ret__.project_id,
        remote_url=__ret__.remote_url,
        repository_credentials_id=__ret__.repository_credentials_id,
        repository_id=__ret__.repository_id)


@_utilities.lift_output_func(get_repository)
def get_repository_output(fetch_deploy_key: Optional[pulumi.Input[Optional[bool]]] = None,
                          project_id: Optional[pulumi.Input[int]] = None,
                          repository_id: Optional[pulumi.Input[int]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...

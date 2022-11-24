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
                 api_key: pulumi.Input[str],
                 endpoint: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_key: The ApiKey used to authenticate requests towards Symbiosis.
        :param pulumi.Input[str] endpoint: Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.
        """
        pulumi.set(__self__, "api_key", api_key)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Input[str]:
        """
        The ApiKey used to authenticate requests towards Symbiosis.
        """
        return pulumi.get(self, "api_key")

    @api_key.setter
    def api_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_key", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the symbiosis package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: The ApiKey used to authenticate requests towards Symbiosis.
        :param pulumi.Input[str] endpoint: Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the symbiosis package. By default, resources use package-wide configuration
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
                 api_key: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if api_key is None and not opts.urn:
                raise TypeError("Missing required property 'api_key'")
            __props__.__dict__["api_key"] = api_key
            __props__.__dict__["endpoint"] = endpoint
        super(Provider, __self__).__init__(
            'symbiosis',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        The ApiKey used to authenticate requests towards Symbiosis.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Endpoint for reaching the symbiosis API. Used for debugging or when accessed through a proxy.
        """
        return pulumi.get(self, "endpoint")

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class NodeBalancer(pulumi.CustomResource):
    client_conn_throttle: pulumi.Output[int]
    """
    Throttle connections per second (0-20). Set to 0 (default) to disable throttling.
    """
    created: pulumi.Output[str]
    hostname: pulumi.Output[str]
    ipv4: pulumi.Output[str]
    ipv6: pulumi.Output[str]
    label: pulumi.Output[str]
    """
    The label of the Linode NodeBalancer
    """
    region: pulumi.Output[str]
    """
    The region where this NodeBalancer will be deployed.  Examples are `"us-east"`, `"us-west"`, `"ap-south"`, etc.  *Changing `region` forces the creation of a new Linode NodeBalancer.*.
    """
    tags: pulumi.Output[list]
    """
    A list of tags applied to this object. Tags are for organizational purposes only.
    """
    transfer: pulumi.Output[dict]
    updated: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, client_conn_throttle=None, label=None, region=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a Linode NodeBalancer resource.  This can be used to create, modify, and delete Linodes NodeBalancers in Linode's managed load balancer service.
        For more information, see [Getting Started with NodeBalancers](https://www.linode.com/docs/platform/nodebalancer/getting-started-with-nodebalancers/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createNodeBalancer).
        
        The Linode Guide, [Create a NodeBalancer with Terraform](https://www.linode.com/docs/applications/configuration-management/create-a-nodebalancer-with-terraform/), provides step-by-step guidance and additional examples.
        
        ## Attributes
        
        This resource exports the following attributes:
        
        * `hostname` - This NodeBalancer's hostname, ending with .nodebalancer.linode.com
        
        * `ipv4` - The Public IPv4 Address of this NodeBalancer
        
        * `ipv6` - The Public IPv6 Address of this NodeBalancer
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] client_conn_throttle: Throttle connections per second (0-20). Set to 0 (default) to disable throttling.
        :param pulumi.Input[str] label: The label of the Linode NodeBalancer
        :param pulumi.Input[str] region: The region where this NodeBalancer will be deployed.  Examples are `"us-east"`, `"us-west"`, `"ap-south"`, etc.  *Changing `region` forces the creation of a new Linode NodeBalancer.*.
        :param pulumi.Input[list] tags: A list of tags applied to this object. Tags are for organizational purposes only.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['client_conn_throttle'] = client_conn_throttle

        __props__['label'] = label

        if region is None:
            raise TypeError('Missing required property region')
        __props__['region'] = region

        __props__['tags'] = tags

        __props__['created'] = None
        __props__['hostname'] = None
        __props__['ipv4'] = None
        __props__['ipv6'] = None
        __props__['transfer'] = None
        __props__['updated'] = None

        super(NodeBalancer, __self__).__init__(
            'linode:index/nodeBalancer:NodeBalancer',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Domain(pulumi.CustomResource):
    axfr_ips: pulumi.Output[list]
    """
    The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
    """
    description: pulumi.Output[str]
    """
    A description for this Domain. This is for display purposes only.
    """
    domain: pulumi.Output[str]
    """
    The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
    """
    expire_sec: pulumi.Output[int]
    """
    The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
    """
    group: pulumi.Output[str]
    """
    The group this Domain belongs to. This is for display purposes only.
    """
    master_ips: pulumi.Output[list]
    """
    The IP addresses representing the master DNS for this Domain.
    """
    refresh_sec: pulumi.Output[int]
    """
    The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
    """
    retry_sec: pulumi.Output[int]
    """
    The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
    """
    soa_email: pulumi.Output[str]
    """
    Start of Authority email address. This is required for master Domains.
    """
    status: pulumi.Output[str]
    """
    Used to control whether this Domain is currently being rendered (defaults to "active").
    """
    tags: pulumi.Output[list]
    """
    A list of tags applied to this object. Tags are for organizational purposes only.
    """
    ttl_sec: pulumi.Output[int]
    """
    'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
    """
    type: pulumi.Output[str]
    """
    If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
    """
    def __init__(__self__, resource_name, opts=None, axfr_ips=None, description=None, domain=None, expire_sec=None, group=None, master_ips=None, refresh_sec=None, retry_sec=None, soa_email=None, status=None, tags=None, ttl_sec=None, type=None, __name__=None, __opts__=None):
        """
        Provides a Linode Domain resource.  This can be used to create, modify, and delete Linode Domains through Linode's managed DNS service.
        For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomain).
        
        The Linode Guide, [Deploy a WordPress Site Using Terraform and Linode StackScripts](https://www.linode.com/docs/applications/configuration-management/deploy-a-wordpress-site-using-terraform-and-linode-stackscripts/), demonstrates the management of Linode Domain resources in the context of Linode Instance running WordPress.
        
        ## Attributes
        
        This resource exports no additional attributes, however `status` may reflect degraded states.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] axfr_ips: The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
        :param pulumi.Input[str] description: A description for this Domain. This is for display purposes only.
        :param pulumi.Input[str] domain: The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
        :param pulumi.Input[int] expire_sec: The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        :param pulumi.Input[str] group: The group this Domain belongs to. This is for display purposes only.
        :param pulumi.Input[list] master_ips: The IP addresses representing the master DNS for this Domain.
        :param pulumi.Input[int] refresh_sec: The amount of time in seconds before this Domain should be refreshed. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        :param pulumi.Input[int] retry_sec: The interval, in seconds, at which a failed refresh should be retried. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        :param pulumi.Input[str] soa_email: Start of Authority email address. This is required for master Domains.
        :param pulumi.Input[str] status: Used to control whether this Domain is currently being rendered (defaults to "active").
        :param pulumi.Input[list] tags: A list of tags applied to this object. Tags are for organizational purposes only.
        :param pulumi.Input[int] ttl_sec: 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        :param pulumi.Input[str] type: If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
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

        __props__['axfr_ips'] = axfr_ips

        __props__['description'] = description

        if domain is None:
            raise TypeError('Missing required property domain')
        __props__['domain'] = domain

        __props__['expire_sec'] = expire_sec

        __props__['group'] = group

        __props__['master_ips'] = master_ips

        __props__['refresh_sec'] = refresh_sec

        __props__['retry_sec'] = retry_sec

        __props__['soa_email'] = soa_email

        __props__['status'] = status

        __props__['tags'] = tags

        __props__['ttl_sec'] = ttl_sec

        if type is None:
            raise TypeError('Missing required property type')
        __props__['type'] = type

        super(Domain, __self__).__init__(
            'linode:index/domain:Domain',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop


# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class DomainRecord(pulumi.CustomResource):
    domain_id: pulumi.Output[int]
    """
    The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.
    """
    name: pulumi.Output[str]
    """
    The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
    """
    port: pulumi.Output[int]
    """
    The port this Record points to.
    """
    priority: pulumi.Output[int]
    """
    The priority of the target host. Lower values are preferred.
    """
    protocol: pulumi.Output[str]
    """
    The protocol this Record's service communicates with. Only valid for SRV records.
    """
    record_type: pulumi.Output[str]
    """
    The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.
    """
    service: pulumi.Output[str]
    """
    The service this Record identified. Only valid for SRV records.
    """
    tag: pulumi.Output[str]
    """
    The tag portion of a CAA record. It is invalid to set this on other record types.
    """
    target: pulumi.Output[str]
    """
    The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
    - - -
    """
    ttl_sec: pulumi.Output[int]
    """
    'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
    """
    weight: pulumi.Output[int]
    """
    The relative weight of this Record. Higher values are preferred.
    """
    def __init__(__self__, resource_name, opts=None, domain_id=None, name=None, port=None, priority=None, protocol=None, record_type=None, service=None, tag=None, target=None, ttl_sec=None, weight=None, __name__=None, __opts__=None):
        """
        Provides a Linode Domain Record resource.  This can be used to create, modify, and delete Linodes Domain Records.
        For more information, see [DNS Manager](https://www.linode.com/docs/platform/manager/dns-manager/) and the [Linode APIv4 docs](https://developers.linode.com/api/v4#operation/createDomainRecord).
        
        ## Attributes
        
        This resource exports no additional attributes.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] domain_id: The ID of the Domain to access.  *Changing `domain_id` forces the creation of a new Linode Domain Record.*.
        :param pulumi.Input[str] name: The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
        :param pulumi.Input[int] port: The port this Record points to.
        :param pulumi.Input[int] priority: The priority of the target host. Lower values are preferred.
        :param pulumi.Input[str] protocol: The protocol this Record's service communicates with. Only valid for SRV records.
        :param pulumi.Input[str] record_type: The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. *Changing `record_type` forces the creation of a new Linode Domain Record.*.
        :param pulumi.Input[str] service: The service this Record identified. Only valid for SRV records.
        :param pulumi.Input[str] tag: The tag portion of a CAA record. It is invalid to set this on other record types.
        :param pulumi.Input[str] target: The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
               - - -
        :param pulumi.Input[int] ttl_sec: 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
        :param pulumi.Input[int] weight: The relative weight of this Record. Higher values are preferred.
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

        if domain_id is None:
            raise TypeError('Missing required property domain_id')
        __props__['domain_id'] = domain_id

        if name is None:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        __props__['port'] = port

        __props__['priority'] = priority

        __props__['protocol'] = protocol

        if record_type is None:
            raise TypeError('Missing required property record_type')
        __props__['record_type'] = record_type

        __props__['service'] = service

        __props__['tag'] = tag

        if target is None:
            raise TypeError('Missing required property target')
        __props__['target'] = target

        __props__['ttl_sec'] = ttl_sec

        __props__['weight'] = weight

        super(DomainRecord, __self__).__init__(
            'linode:index/domainRecord:DomainRecord',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

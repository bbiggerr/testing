# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetRegionResult(object):
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, country=None):
        if country and not isinstance(country, str):
            raise TypeError('Expected argument country to be a str')
        __self__.country = country

async def get_region(country=None, id=None):
    """
    `linode_region` provides details about a specific Linode region.
    """
    __args__ = dict()

    __args__['country'] = country
    __args__['id'] = id
    __ret__ = await pulumi.runtime.invoke('linode:index/getRegion:getRegion', __args__)

    return GetRegionResult(
        country=__ret__.get('country'))

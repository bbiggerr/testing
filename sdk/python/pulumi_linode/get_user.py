# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetUserResult(object):
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, email=None, restricted=None, ssh_keys=None, id=None):
        if email and not isinstance(email, str):
            raise TypeError('Expected argument email to be a str')
        __self__.email = email
        if restricted and not isinstance(restricted, bool):
            raise TypeError('Expected argument restricted to be a bool')
        __self__.restricted = restricted
        if ssh_keys and not isinstance(ssh_keys, list):
            raise TypeError('Expected argument ssh_keys to be a list')
        __self__.ssh_keys = ssh_keys
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_user(username=None):
    """
    Provides information about a Linode user
    
    ## Attributes
    
    The Linode User resource exports the following attributes:
    
    * `ssh_keys` - A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.
    
    * `email` - The email address for this User, for account management communications, and may be used for other communications as configured.
    
    * `restricted` - If true, this User must be granted access to perform actions or access entities on this Account.
    """
    __args__ = dict()

    __args__['username'] = username
    __ret__ = await pulumi.runtime.invoke('linode:index/getUser:getUser', __args__)

    return GetUserResult(
        email=__ret__.get('email'),
        restricted=__ret__.get('restricted'),
        ssh_keys=__ret__.get('sshKeys'),
        id=__ret__.get('id'))

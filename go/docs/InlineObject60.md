# InlineObject60

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACTIVITY_STREAM_ENABLED** | **bool** | Enable capturing activity for the activity stream. | 
**ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC** | **bool** | Enable capturing activity for the activity stream when running inventory sync. | 
**AD_HOC_COMMANDS** | Pointer to **[]string** | List of modules allowed to be used by ad-hoc jobs. | [optional] 
**ALLOW_JINJA_IN_EXTRA_VARS** | **string** | Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where Tower users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to \&quot;template\&quot; or \&quot;never\&quot;. | 
**ALLOWOAUTH2FOREXTERNALUSERS** | Pointer to **bool** | For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off. | [optional] 
**ANSIBLE_FACT_CACHE_TIMEOUT** | Pointer to **int32** | Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed. | [optional] 
**AUTH_BASIC_ENABLED** | **bool** | Enable HTTP Basic Auth for the API Browser. | 
**AUTHLDAP1BINDDN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTHLDAP1BINDPASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTHLDAP1CONNECTIONOPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTHLDAP1DENYGROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTHLDAP1GROUPSEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTHLDAP1GROUPTYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTHLDAP1GROUPTYPEPARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTHLDAP1ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP1REQUIREGROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTHLDAP1SERVERURI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTHLDAP1STARTTLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTHLDAP1TEAMMAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP1USERATTRMAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTHLDAP1USERDNTEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTHLDAP1USERFLAGSBYGROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTHLDAP1USERSEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTHLDAP2BINDDN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTHLDAP2BINDPASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTHLDAP2CONNECTIONOPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTHLDAP2DENYGROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTHLDAP2GROUPSEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTHLDAP2GROUPTYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTHLDAP2GROUPTYPEPARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTHLDAP2ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP2REQUIREGROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTHLDAP2SERVERURI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTHLDAP2STARTTLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTHLDAP2TEAMMAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP2USERATTRMAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTHLDAP2USERDNTEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTHLDAP2USERFLAGSBYGROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTHLDAP2USERSEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTHLDAP3BINDDN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTHLDAP3BINDPASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTHLDAP3CONNECTIONOPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTHLDAP3DENYGROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTHLDAP3GROUPSEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTHLDAP3GROUPTYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTHLDAP3GROUPTYPEPARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTHLDAP3ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP3REQUIREGROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTHLDAP3SERVERURI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTHLDAP3STARTTLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTHLDAP3TEAMMAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP3USERATTRMAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTHLDAP3USERDNTEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTHLDAP3USERFLAGSBYGROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTHLDAP3USERSEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTHLDAP4BINDDN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTHLDAP4BINDPASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTHLDAP4CONNECTIONOPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTHLDAP4DENYGROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTHLDAP4GROUPSEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTHLDAP4GROUPTYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTHLDAP4GROUPTYPEPARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTHLDAP4ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP4REQUIREGROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTHLDAP4SERVERURI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTHLDAP4STARTTLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTHLDAP4TEAMMAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP4USERATTRMAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTHLDAP4USERDNTEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTHLDAP4USERFLAGSBYGROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTHLDAP4USERSEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTHLDAP5BINDDN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTHLDAP5BINDPASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTHLDAP5CONNECTIONOPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTHLDAP5DENYGROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTHLDAP5GROUPSEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTHLDAP5GROUPTYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTHLDAP5GROUPTYPEPARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTHLDAP5ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP5REQUIREGROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTHLDAP5SERVERURI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTHLDAP5STARTTLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTHLDAP5TEAMMAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTHLDAP5USERATTRMAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTHLDAP5USERDNTEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTHLDAP5USERFLAGSBYGROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTHLDAP5USERSEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTH_LDAP_BIND_DN** | Pointer to **string** | DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the Ansible Tower documentation for example syntax. | [optional] 
**AUTH_LDAP_BIND_PASSWORD** | Pointer to **string** | Password used to bind LDAP user account. | [optional] 
**AUTH_LDAP_CONNECTION_OPTIONS** | Pointer to **map[string]interface{}** | Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. \&quot;OPT_REFERRALS\&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. | [optional] 
**AUTH_LDAP_DENY_GROUP** | Pointer to **string** | Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. | [optional] 
**AUTH_LDAP_GROUP_SEARCH** | Pointer to **[]string** | Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. | [optional] 
**AUTH_LDAP_GROUP_TYPE** | Pointer to **string** | The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups | [optional] 
**AUTH_LDAP_GROUP_TYPE_PARAMS** | Pointer to **map[string]interface{}** | Key value parameters to send the chosen group type init method. | [optional] 
**AUTH_LDAP_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping between organization admins/users and LDAP groups. This controls which users are placed into which Tower organizations relative to their LDAP group memberships. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTH_LDAP_REQUIRE_GROUP** | Pointer to **string** | Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login via Tower. Only one require group is supported. | [optional] 
**AUTH_LDAP_SERVER_URI** | Pointer to **string** | URI to connect to LDAP server, such as \&quot;ldap://ldap.example.com:389\&quot; (non-SSL) or \&quot;ldaps://ldap.example.com:636\&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. | [optional] 
**AUTH_LDAP_START_TLS** | Pointer to **bool** | Whether to enable TLS when the LDAP connection is not using SSL. | [optional] 
**AUTH_LDAP_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping between team members (users) and LDAP groups. Configuration details are available in the Ansible Tower documentation. | [optional] 
**AUTH_LDAP_USER_ATTR_MAP** | Pointer to **map[string]interface{}** | Mapping of LDAP user schema to Tower API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the Ansible Tower documentation for additional details. | [optional] 
**AUTH_LDAP_USER_DN_TEMPLATE** | Pointer to **string** | Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. | [optional] 
**AUTH_LDAP_USER_FLAGS_BY_GROUP** | Pointer to **map[string]interface{}** | Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the Ansible Tower documentation for more detail. | [optional] 
**AUTH_LDAP_USER_SEARCH** | Pointer to **[]string** | LDAP search query to find users.  Any user that matches the given pattern will be able to login to Tower.  The user should also be mapped into a Tower organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of \&quot;LDAPUnion\&quot; is possible. See Tower documentation for details. | [optional] 
**AUTOMATION_ANALYTICS_GATHER_INTERVAL** | Pointer to **int32** | Interval (in seconds) between data gathering. | [optional] 
**AUTOMATION_ANALYTICS_LAST_GATHER** | **string** |  | 
**AUTOMATION_ANALYTICS_URL** | Pointer to **string** | This setting is used to to configure data collection for the Automation Analytics dashboard | [optional] 
**AWX_ANSIBLE_CALLBACK_PLUGINS** | Pointer to **[]string** | List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line. | [optional] 
**AWX_COLLECTIONS_ENABLED** | Pointer to **bool** | Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects. | [optional] 
**AWX_ISOLATED_CHECK_INTERVAL** | **int32** | The number of seconds to sleep between status checks for jobs running on isolated instances. | 
**AWX_ISOLATED_CONNECTION_TIMEOUT** | Pointer to **int32** | Ansible SSH connection timeout (in seconds) to use when communicating with isolated instances. Value should be substantially greater than expected network latency. | [optional] 
**AWX_ISOLATED_HOST_KEY_CHECKING** | Pointer to **bool** | When set to True, AWX will enforce strict host key checking for communication with isolated nodes. | [optional] 
**AWX_ISOLATED_LAUNCH_TIMEOUT** | **int32** | The timeout (in seconds) for launching jobs on isolated instances.  This includes the time needed to copy source control files (playbooks) to the isolated instance. | 
**AWX_PROOT_BASE_PATH** | **string** | The directory in which Tower will create new temporary directories for job execution and isolation (such as credential files and custom inventory scripts). | 
**AWX_PROOT_ENABLED** | **bool** | Isolates an Ansible job from protected parts of the system to prevent exposing sensitive information. | 
**AWX_PROOT_HIDE_PATHS** | Pointer to **[]string** | Additional paths to hide from isolated processes. Enter one path per line. | [optional] 
**AWX_PROOT_SHOW_PATHS** | Pointer to **[]string** | List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. | [optional] 
**AWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL** | Pointer to **float32** | Interval (in seconds) between polls for cpu usage. Setting this lower than the default will affect playbook performance. | [optional] 
**AWX_RESOURCE_PROFILING_ENABLED** | Pointer to **bool** | If set, detailed resource profiling data will be collected on all jobs. This data can be gathered with &#x60;sosreport&#x60;. | [optional] 
**AWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL** | Pointer to **float32** | Interval (in seconds) between polls for memory usage. Setting this lower than the default will affect playbook performance. | [optional] 
**AWX_RESOURCE_PROFILING_PID_POLL_INTERVAL** | Pointer to **float32** | Interval (in seconds) between polls for PID count. Setting this lower than the default will affect playbook performance. | [optional] 
**AWX_ROLES_ENABLED** | Pointer to **bool** | Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects. | [optional] 
**AWX_SHOW_PLAYBOOK_LINKS** | Pointer to **bool** | Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself. | [optional] 
**AWX_TASK_ENV** | Pointer to **map[string]interface{}** | Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending. | [optional] 
**CUSTOM_LOGIN_INFO** | Pointer to **string** | If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported. | [optional] 
**CUSTOM_LOGO** | Pointer to **string** | To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported. | [optional] 
**CUSTOM_VENV_PATHS** | Pointer to **[]string** | Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line. | [optional] 
**DEFAULT_INVENTORY_UPDATE_TIMEOUT** | Pointer to **int32** | Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this. | [optional] 
**DEFAULT_JOB_TIMEOUT** | Pointer to **int32** | Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this. | [optional] 
**DEFAULT_PROJECT_UPDATE_TIMEOUT** | Pointer to **int32** | Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this. | [optional] 
**EVENT_STDOUT_MAX_BYTES_DISPLAY** | **int32** | Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. &#x60;stdout&#x60; will end with &#x60;…&#x60; when truncated. | 
**GALAXY_IGNORE_CERTS** | Pointer to **bool** | If set to true, certificate validation will not be done when installing content from any Galaxy server. | [optional] 
**INSIGHTS_TRACKING_STATE** | Pointer to **bool** | Enables Tower to gather data on automation and send it to Red Hat. | [optional] 
**LOGIN_REDIRECT_OVERRIDE** | Pointer to **string** | URL to which unauthorized users will be redirected to log in. If blank, users will be sent to the Tower login page. | [optional] 
**LOG_AGGREGATOR_ENABLED** | Pointer to **bool** | Enable sending logs to external log aggregator. | [optional] 
**LOG_AGGREGATOR_HOST** | Pointer to **string** | Hostname/IP where external logs will be sent to. | [optional] 
**LOG_AGGREGATOR_INDIVIDUAL_FACTS** | Pointer to **bool** | If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing. | [optional] 
**LOG_AGGREGATOR_LEVEL** | Pointer to **string** | Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting) | [optional] 
**LOG_AGGREGATOR_LOGGERS** | Pointer to **[]string** | List of loggers that will send HTTP logs to the collector, these can include any or all of:  awx - service logs activity_stream - activity stream records job_events - callback data from Ansible job events system_tracking - facts gathered from scan jobs. | [optional] 
**LOG_AGGREGATOR_MAX_DISK_USAGE_GB** | Pointer to **int32** | Amount of data to store (in gigabytes) during an outage of the external log aggregator (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting. | [optional] 
**LOG_AGGREGATOR_MAX_DISK_USAGE_PATH** | Pointer to **string** | Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting. | [optional] 
**LOG_AGGREGATOR_PASSWORD** | Pointer to **string** | Password or authentication token for external log aggregator (if required; HTTP/s only). | [optional] 
**LOG_AGGREGATOR_PORT** | Pointer to **int32** | Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator). | [optional] 
**LOG_AGGREGATOR_PROTOCOL** | Pointer to **string** | Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname. | [optional] 
**LOG_AGGREGATOR_RSYSLOGD_DEBUG** | Pointer to **bool** | Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation. | [optional] 
**LOG_AGGREGATOR_TCP_TIMEOUT** | Pointer to **int32** | Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols. | [optional] 
**LOG_AGGREGATOR_TOWER_UUID** | Pointer to **string** | Useful to uniquely identify Tower instances. | [optional] 
**LOG_AGGREGATOR_TYPE** | Pointer to **string** | Format messages for the chosen log aggregator. | [optional] 
**LOG_AGGREGATOR_USERNAME** | Pointer to **string** | Username for external log aggregator (if required; HTTP/s only). | [optional] 
**LOG_AGGREGATOR_VERIFY_CERT** | Pointer to **bool** | Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is \&quot;https\&quot;. If enabled, Tower&#39;s log handler will verify certificate sent by external log aggregator before establishing connection. | [optional] 
**MANAGE_ORGANIZATION_AUTH** | **bool** | Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration. | 
**MAX_FORKS** | Pointer to **int32** | Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied. | [optional] 
**MAX_UI_JOB_EVENTS** | **int32** | Maximum number of job events for the UI to retrieve within a single request. | 
**OAUTH2PROVIDER** | Pointer to **map[string]interface{}** | Dictionary for customizing OAuth 2 timeouts, available items are &#x60;ACCESS_TOKEN_EXPIRE_SECONDS&#x60;, the duration of access tokens in the number of seconds, &#x60;AUTHORIZATION_CODE_EXPIRE_SECONDS&#x60;, the duration of authorization codes in the number of seconds, and &#x60;REFRESH_TOKEN_EXPIRE_SECONDS&#x60;, the duration of refresh tokens, after expired access tokens, in the number of seconds. | [optional] 
**ORG_ADMINS_CAN_SEE_ALL_USERS** | **bool** | Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization. | 
**PROJECT_UPDATE_VVV** | **bool** | Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates. | 
**PROXY_IP_ALLOWED_LIST** | **[]string** | If Tower is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which Tower should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally&#39;) | 
**RADIUS_PORT** | Pointer to **int32** | Port of RADIUS server. | [optional] 
**RADIUS_SECRET** | Pointer to **string** | Shared secret for authenticating to RADIUS server. | [optional] 
**RADIUS_SERVER** | Pointer to **string** | Hostname/IP of RADIUS server. RADIUS authentication is disabled if this setting is empty. | [optional] 
**REDHAT_PASSWORD** | Pointer to **string** | This password is used to retrieve license information and to send Automation Analytics | [optional] 
**REDHAT_USERNAME** | Pointer to **string** | This username is used to retrieve license information and to send Automation Analytics | [optional] 
**REMOTE_HOST_HEADERS** | **[]string** | HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as \&quot;HTTP_X_FORWARDED_FOR\&quot;, if behind a reverse proxy. See the \&quot;Proxy Support\&quot; section of the Adminstrator guide for more details. | 
**SAML_AUTO_CREATE_OBJECTS** | Pointer to **bool** | When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login. | [optional] 
**SCHEDULE_MAX_JOBS** | **int32** | Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created. | 
**SESSIONS_PER_USER** | **int32** | Maximum number of simultaneous logged in sessions a user may have. To disable enter -1. | 
**SESSION_COOKIE_AGE** | **int32** | Number of seconds that a user is inactive before they will need to login again. | 
**SOCIALAUTHAZUREADOAUTH2KEY** | Pointer to **string** | The OAuth2 key (Client ID) from your Azure AD application. | [optional] 
**SOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIALAUTHAZUREADOAUTH2SECRET** | Pointer to **string** | The OAuth2 secret (Client Secret) from your Azure AD application. | [optional] 
**SOCIALAUTHAZUREADOAUTH2TEAMMAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_KEY** | Pointer to **string** | The OAuth2 key (Client ID) from your GitHub developer application. | [optional] 
**SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_ORG_KEY** | Pointer to **string** | The OAuth2 key (Client ID) from your GitHub organization application. | [optional] 
**SOCIAL_AUTH_GITHUB_ORG_NAME** | Pointer to **string** | The name of your GitHub organization, as used in your organization&#39;s URL: https://github.com/&lt;yourorg&gt;/. | [optional] 
**SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_ORG_SECRET** | Pointer to **string** | The OAuth2 secret (Client Secret) from your GitHub organization application. | [optional] 
**SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_SECRET** | Pointer to **string** | The OAuth2 secret (Client Secret) from your GitHub developer application. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_ID** | Pointer to **string** | Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_KEY** | Pointer to **string** | The OAuth2 key (Client ID) from your GitHub organization application. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_SECRET** | Pointer to **string** | The OAuth2 secret (Client Secret) from your GitHub organization application. | [optional] 
**SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS** | Pointer to **map[string]interface{}** | Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the Ansible Tower documentation for more detail. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2KEY** | Pointer to **string** | The OAuth2 key from your web application. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2SECRET** | Pointer to **string** | The OAuth2 secret from your web application. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2TEAMMAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS** | Pointer to **[]string** | Update this setting to restrict the domains who are allowed to login using Google OAuth2. | [optional] 
**SOCIAL_AUTH_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIAL_AUTH_SAML_ENABLED_IDPS** | Pointer to **map[string]interface{}** | Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax. | [optional] 
**SOCIAL_AUTH_SAML_EXTRA_DATA** | Pointer to **[]string** | A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value. | [optional] 
**SOCIAL_AUTH_SAML_ORGANIZATION_ATTR** | Pointer to **map[string]interface{}** | Used to translate user organization membership into Tower. | [optional] 
**SOCIAL_AUTH_SAML_ORGANIZATION_MAP** | Pointer to **map[string]interface{}** | Mapping to organization admins/users from social auth accounts. This setting controls which users are placed into which Tower organizations based on their username and email address. Configuration details are available in the Ansible Tower documentation. | [optional] 
**SOCIAL_AUTH_SAML_ORG_INFO** | **map[string]interface{}** | Provide the URL, display name, and the name of your app. Refer to the Ansible Tower documentation for example syntax. | 
**SOCIAL_AUTH_SAML_SECURITY_CONFIG** | Pointer to **map[string]interface{}** | A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings | [optional] 
**SOCIAL_AUTH_SAML_SP_ENTITY_ID** | Pointer to **string** | The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for Tower. | [optional] 
**SOCIAL_AUTH_SAML_SP_EXTRA** | Pointer to **map[string]interface{}** | A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting. | [optional] 
**SOCIAL_AUTH_SAML_SP_PRIVATE_KEY** | **string** | Create a keypair for Tower to use as a service provider (SP) and include the private key content here. | 
**SOCIAL_AUTH_SAML_SP_PUBLIC_CERT** | **string** | Create a keypair for Tower to use as a service provider (SP) and include the certificate content here. | 
**SOCIAL_AUTH_SAML_SUPPORT_CONTACT** | **map[string]interface{}** | Provide the name and email address of the support contact for your service provider. Refer to the Ansible Tower documentation for example syntax. | 
**SOCIAL_AUTH_SAML_TEAM_ATTR** | Pointer to **map[string]interface{}** | Used to translate user team membership into Tower. | [optional] 
**SOCIAL_AUTH_SAML_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIAL_AUTH_SAML_TECHNICAL_CONTACT** | **map[string]interface{}** | Provide the name and email address of the technical contact for your service provider. Refer to the Ansible Tower documentation for example syntax. | 
**SOCIAL_AUTH_TEAM_MAP** | Pointer to **map[string]interface{}** | Mapping of team members (users) from social auth accounts. Configuration details are available in Tower documentation. | [optional] 
**SOCIAL_AUTH_USER_FIELDS** | Pointer to **[]string** | When set to an empty list &#x60;[]&#x60;, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login. | [optional] 
**STDOUT_MAX_BYTES_DISPLAY** | **int32** | Maximum Size of Standard Output in bytes to display before requiring the output be downloaded. | 
**TACACSPLUS_AUTH_PROTOCOL** | Pointer to **string** | Choose the authentication protocol used by TACACS+ client. | [optional] 
**TACACSPLUS_HOST** | Pointer to **string** | Hostname of TACACS+ server. | [optional] 
**TACACSPLUS_PORT** | Pointer to **int32** | Port number of TACACS+ server. | [optional] 
**TACACSPLUS_SECRET** | Pointer to **string** | Shared secret for authenticating to TACACS+ server. | [optional] 
**TACACSPLUS_SESSION_TIMEOUT** | Pointer to **int32** | TACACS+ session timeout value in seconds, 0 disables timeout. | [optional] 
**TOWER_URL_BASE** | **string** | This setting is used by services like notifications to render a valid url to the Tower host. | 
**UI_LIVE_UPDATES_ENABLED** | **bool** | If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details. | 

## Methods

### NewInlineObject60

`func NewInlineObject60(aCTIVITYSTREAMENABLED bool, aCTIVITYSTREAMENABLEDFORINVENTORYSYNC bool, aLLOWJINJAINEXTRAVARS string, aUTHBASICENABLED bool, aUTOMATIONANALYTICSLASTGATHER string, aWXISOLATEDCHECKINTERVAL int32, aWXISOLATEDLAUNCHTIMEOUT int32, aWXPROOTBASEPATH string, aWXPROOTENABLED bool, eVENTSTDOUTMAXBYTESDISPLAY int32, mANAGEORGANIZATIONAUTH bool, mAXUIJOBEVENTS int32, oRGADMINSCANSEEALLUSERS bool, pROJECTUPDATEVVV bool, pROXYIPALLOWEDLIST []string, rEMOTEHOSTHEADERS []string, sCHEDULEMAXJOBS int32, sESSIONSPERUSER int32, sESSIONCOOKIEAGE int32, sOCIALAUTHSAMLORGINFO map[string]interface{}, sOCIALAUTHSAMLSPPRIVATEKEY string, sOCIALAUTHSAMLSPPUBLICCERT string, sOCIALAUTHSAMLSUPPORTCONTACT map[string]interface{}, sOCIALAUTHSAMLTECHNICALCONTACT map[string]interface{}, sTDOUTMAXBYTESDISPLAY int32, tOWERURLBASE string, uILIVEUPDATESENABLED bool, ) *InlineObject60`

NewInlineObject60 instantiates a new InlineObject60 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject60WithDefaults

`func NewInlineObject60WithDefaults() *InlineObject60`

NewInlineObject60WithDefaults instantiates a new InlineObject60 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACTIVITY_STREAM_ENABLED

`func (o *InlineObject60) GetACTIVITY_STREAM_ENABLED() bool`

GetACTIVITY_STREAM_ENABLED returns the ACTIVITY_STREAM_ENABLED field if non-nil, zero value otherwise.

### GetACTIVITY_STREAM_ENABLEDOk

`func (o *InlineObject60) GetACTIVITY_STREAM_ENABLEDOk() (*bool, bool)`

GetACTIVITY_STREAM_ENABLEDOk returns a tuple with the ACTIVITY_STREAM_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACTIVITY_STREAM_ENABLED

`func (o *InlineObject60) SetACTIVITY_STREAM_ENABLED(v bool)`

SetACTIVITY_STREAM_ENABLED sets ACTIVITY_STREAM_ENABLED field to given value.


### GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC

`func (o *InlineObject60) GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC() bool`

GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC returns the ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC field if non-nil, zero value otherwise.

### GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNCOk

`func (o *InlineObject60) GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNCOk() (*bool, bool)`

GetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNCOk returns a tuple with the ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC

`func (o *InlineObject60) SetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC(v bool)`

SetACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC sets ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC field to given value.


### GetAD_HOC_COMMANDS

`func (o *InlineObject60) GetAD_HOC_COMMANDS() []string`

GetAD_HOC_COMMANDS returns the AD_HOC_COMMANDS field if non-nil, zero value otherwise.

### GetAD_HOC_COMMANDSOk

`func (o *InlineObject60) GetAD_HOC_COMMANDSOk() (*[]string, bool)`

GetAD_HOC_COMMANDSOk returns a tuple with the AD_HOC_COMMANDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAD_HOC_COMMANDS

`func (o *InlineObject60) SetAD_HOC_COMMANDS(v []string)`

SetAD_HOC_COMMANDS sets AD_HOC_COMMANDS field to given value.

### HasAD_HOC_COMMANDS

`func (o *InlineObject60) HasAD_HOC_COMMANDS() bool`

HasAD_HOC_COMMANDS returns a boolean if a field has been set.

### GetALLOW_JINJA_IN_EXTRA_VARS

`func (o *InlineObject60) GetALLOW_JINJA_IN_EXTRA_VARS() string`

GetALLOW_JINJA_IN_EXTRA_VARS returns the ALLOW_JINJA_IN_EXTRA_VARS field if non-nil, zero value otherwise.

### GetALLOW_JINJA_IN_EXTRA_VARSOk

`func (o *InlineObject60) GetALLOW_JINJA_IN_EXTRA_VARSOk() (*string, bool)`

GetALLOW_JINJA_IN_EXTRA_VARSOk returns a tuple with the ALLOW_JINJA_IN_EXTRA_VARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALLOW_JINJA_IN_EXTRA_VARS

`func (o *InlineObject60) SetALLOW_JINJA_IN_EXTRA_VARS(v string)`

SetALLOW_JINJA_IN_EXTRA_VARS sets ALLOW_JINJA_IN_EXTRA_VARS field to given value.


### GetALLOWOAUTH2FOREXTERNALUSERS

`func (o *InlineObject60) GetALLOWOAUTH2FOREXTERNALUSERS() bool`

GetALLOWOAUTH2FOREXTERNALUSERS returns the ALLOWOAUTH2FOREXTERNALUSERS field if non-nil, zero value otherwise.

### GetALLOWOAUTH2FOREXTERNALUSERSOk

`func (o *InlineObject60) GetALLOWOAUTH2FOREXTERNALUSERSOk() (*bool, bool)`

GetALLOWOAUTH2FOREXTERNALUSERSOk returns a tuple with the ALLOWOAUTH2FOREXTERNALUSERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALLOWOAUTH2FOREXTERNALUSERS

`func (o *InlineObject60) SetALLOWOAUTH2FOREXTERNALUSERS(v bool)`

SetALLOWOAUTH2FOREXTERNALUSERS sets ALLOWOAUTH2FOREXTERNALUSERS field to given value.

### HasALLOWOAUTH2FOREXTERNALUSERS

`func (o *InlineObject60) HasALLOWOAUTH2FOREXTERNALUSERS() bool`

HasALLOWOAUTH2FOREXTERNALUSERS returns a boolean if a field has been set.

### GetANSIBLE_FACT_CACHE_TIMEOUT

`func (o *InlineObject60) GetANSIBLE_FACT_CACHE_TIMEOUT() int32`

GetANSIBLE_FACT_CACHE_TIMEOUT returns the ANSIBLE_FACT_CACHE_TIMEOUT field if non-nil, zero value otherwise.

### GetANSIBLE_FACT_CACHE_TIMEOUTOk

`func (o *InlineObject60) GetANSIBLE_FACT_CACHE_TIMEOUTOk() (*int32, bool)`

GetANSIBLE_FACT_CACHE_TIMEOUTOk returns a tuple with the ANSIBLE_FACT_CACHE_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetANSIBLE_FACT_CACHE_TIMEOUT

`func (o *InlineObject60) SetANSIBLE_FACT_CACHE_TIMEOUT(v int32)`

SetANSIBLE_FACT_CACHE_TIMEOUT sets ANSIBLE_FACT_CACHE_TIMEOUT field to given value.

### HasANSIBLE_FACT_CACHE_TIMEOUT

`func (o *InlineObject60) HasANSIBLE_FACT_CACHE_TIMEOUT() bool`

HasANSIBLE_FACT_CACHE_TIMEOUT returns a boolean if a field has been set.

### GetAUTH_BASIC_ENABLED

`func (o *InlineObject60) GetAUTH_BASIC_ENABLED() bool`

GetAUTH_BASIC_ENABLED returns the AUTH_BASIC_ENABLED field if non-nil, zero value otherwise.

### GetAUTH_BASIC_ENABLEDOk

`func (o *InlineObject60) GetAUTH_BASIC_ENABLEDOk() (*bool, bool)`

GetAUTH_BASIC_ENABLEDOk returns a tuple with the AUTH_BASIC_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_BASIC_ENABLED

`func (o *InlineObject60) SetAUTH_BASIC_ENABLED(v bool)`

SetAUTH_BASIC_ENABLED sets AUTH_BASIC_ENABLED field to given value.


### GetAUTHLDAP1BINDDN

`func (o *InlineObject60) GetAUTHLDAP1BINDDN() string`

GetAUTHLDAP1BINDDN returns the AUTHLDAP1BINDDN field if non-nil, zero value otherwise.

### GetAUTHLDAP1BINDDNOk

`func (o *InlineObject60) GetAUTHLDAP1BINDDNOk() (*string, bool)`

GetAUTHLDAP1BINDDNOk returns a tuple with the AUTHLDAP1BINDDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1BINDDN

`func (o *InlineObject60) SetAUTHLDAP1BINDDN(v string)`

SetAUTHLDAP1BINDDN sets AUTHLDAP1BINDDN field to given value.

### HasAUTHLDAP1BINDDN

`func (o *InlineObject60) HasAUTHLDAP1BINDDN() bool`

HasAUTHLDAP1BINDDN returns a boolean if a field has been set.

### GetAUTHLDAP1BINDPASSWORD

`func (o *InlineObject60) GetAUTHLDAP1BINDPASSWORD() string`

GetAUTHLDAP1BINDPASSWORD returns the AUTHLDAP1BINDPASSWORD field if non-nil, zero value otherwise.

### GetAUTHLDAP1BINDPASSWORDOk

`func (o *InlineObject60) GetAUTHLDAP1BINDPASSWORDOk() (*string, bool)`

GetAUTHLDAP1BINDPASSWORDOk returns a tuple with the AUTHLDAP1BINDPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1BINDPASSWORD

`func (o *InlineObject60) SetAUTHLDAP1BINDPASSWORD(v string)`

SetAUTHLDAP1BINDPASSWORD sets AUTHLDAP1BINDPASSWORD field to given value.

### HasAUTHLDAP1BINDPASSWORD

`func (o *InlineObject60) HasAUTHLDAP1BINDPASSWORD() bool`

HasAUTHLDAP1BINDPASSWORD returns a boolean if a field has been set.

### GetAUTHLDAP1CONNECTIONOPTIONS

`func (o *InlineObject60) GetAUTHLDAP1CONNECTIONOPTIONS() map[string]interface{}`

GetAUTHLDAP1CONNECTIONOPTIONS returns the AUTHLDAP1CONNECTIONOPTIONS field if non-nil, zero value otherwise.

### GetAUTHLDAP1CONNECTIONOPTIONSOk

`func (o *InlineObject60) GetAUTHLDAP1CONNECTIONOPTIONSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1CONNECTIONOPTIONSOk returns a tuple with the AUTHLDAP1CONNECTIONOPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1CONNECTIONOPTIONS

`func (o *InlineObject60) SetAUTHLDAP1CONNECTIONOPTIONS(v map[string]interface{})`

SetAUTHLDAP1CONNECTIONOPTIONS sets AUTHLDAP1CONNECTIONOPTIONS field to given value.

### HasAUTHLDAP1CONNECTIONOPTIONS

`func (o *InlineObject60) HasAUTHLDAP1CONNECTIONOPTIONS() bool`

HasAUTHLDAP1CONNECTIONOPTIONS returns a boolean if a field has been set.

### GetAUTHLDAP1DENYGROUP

`func (o *InlineObject60) GetAUTHLDAP1DENYGROUP() string`

GetAUTHLDAP1DENYGROUP returns the AUTHLDAP1DENYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP1DENYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP1DENYGROUPOk() (*string, bool)`

GetAUTHLDAP1DENYGROUPOk returns a tuple with the AUTHLDAP1DENYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1DENYGROUP

`func (o *InlineObject60) SetAUTHLDAP1DENYGROUP(v string)`

SetAUTHLDAP1DENYGROUP sets AUTHLDAP1DENYGROUP field to given value.

### HasAUTHLDAP1DENYGROUP

`func (o *InlineObject60) HasAUTHLDAP1DENYGROUP() bool`

HasAUTHLDAP1DENYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP1GROUPSEARCH

`func (o *InlineObject60) GetAUTHLDAP1GROUPSEARCH() []string`

GetAUTHLDAP1GROUPSEARCH returns the AUTHLDAP1GROUPSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP1GROUPSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP1GROUPSEARCHOk() (*[]string, bool)`

GetAUTHLDAP1GROUPSEARCHOk returns a tuple with the AUTHLDAP1GROUPSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1GROUPSEARCH

`func (o *InlineObject60) SetAUTHLDAP1GROUPSEARCH(v []string)`

SetAUTHLDAP1GROUPSEARCH sets AUTHLDAP1GROUPSEARCH field to given value.

### HasAUTHLDAP1GROUPSEARCH

`func (o *InlineObject60) HasAUTHLDAP1GROUPSEARCH() bool`

HasAUTHLDAP1GROUPSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP1GROUPTYPE

`func (o *InlineObject60) GetAUTHLDAP1GROUPTYPE() string`

GetAUTHLDAP1GROUPTYPE returns the AUTHLDAP1GROUPTYPE field if non-nil, zero value otherwise.

### GetAUTHLDAP1GROUPTYPEOk

`func (o *InlineObject60) GetAUTHLDAP1GROUPTYPEOk() (*string, bool)`

GetAUTHLDAP1GROUPTYPEOk returns a tuple with the AUTHLDAP1GROUPTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1GROUPTYPE

`func (o *InlineObject60) SetAUTHLDAP1GROUPTYPE(v string)`

SetAUTHLDAP1GROUPTYPE sets AUTHLDAP1GROUPTYPE field to given value.

### HasAUTHLDAP1GROUPTYPE

`func (o *InlineObject60) HasAUTHLDAP1GROUPTYPE() bool`

HasAUTHLDAP1GROUPTYPE returns a boolean if a field has been set.

### GetAUTHLDAP1GROUPTYPEPARAMS

`func (o *InlineObject60) GetAUTHLDAP1GROUPTYPEPARAMS() map[string]interface{}`

GetAUTHLDAP1GROUPTYPEPARAMS returns the AUTHLDAP1GROUPTYPEPARAMS field if non-nil, zero value otherwise.

### GetAUTHLDAP1GROUPTYPEPARAMSOk

`func (o *InlineObject60) GetAUTHLDAP1GROUPTYPEPARAMSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1GROUPTYPEPARAMSOk returns a tuple with the AUTHLDAP1GROUPTYPEPARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1GROUPTYPEPARAMS

`func (o *InlineObject60) SetAUTHLDAP1GROUPTYPEPARAMS(v map[string]interface{})`

SetAUTHLDAP1GROUPTYPEPARAMS sets AUTHLDAP1GROUPTYPEPARAMS field to given value.

### HasAUTHLDAP1GROUPTYPEPARAMS

`func (o *InlineObject60) HasAUTHLDAP1GROUPTYPEPARAMS() bool`

HasAUTHLDAP1GROUPTYPEPARAMS returns a boolean if a field has been set.

### GetAUTHLDAP1ORGANIZATIONMAP

`func (o *InlineObject60) GetAUTHLDAP1ORGANIZATIONMAP() map[string]interface{}`

GetAUTHLDAP1ORGANIZATIONMAP returns the AUTHLDAP1ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP1ORGANIZATIONMAPOk

`func (o *InlineObject60) GetAUTHLDAP1ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1ORGANIZATIONMAPOk returns a tuple with the AUTHLDAP1ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1ORGANIZATIONMAP

`func (o *InlineObject60) SetAUTHLDAP1ORGANIZATIONMAP(v map[string]interface{})`

SetAUTHLDAP1ORGANIZATIONMAP sets AUTHLDAP1ORGANIZATIONMAP field to given value.

### HasAUTHLDAP1ORGANIZATIONMAP

`func (o *InlineObject60) HasAUTHLDAP1ORGANIZATIONMAP() bool`

HasAUTHLDAP1ORGANIZATIONMAP returns a boolean if a field has been set.

### GetAUTHLDAP1REQUIREGROUP

`func (o *InlineObject60) GetAUTHLDAP1REQUIREGROUP() string`

GetAUTHLDAP1REQUIREGROUP returns the AUTHLDAP1REQUIREGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP1REQUIREGROUPOk

`func (o *InlineObject60) GetAUTHLDAP1REQUIREGROUPOk() (*string, bool)`

GetAUTHLDAP1REQUIREGROUPOk returns a tuple with the AUTHLDAP1REQUIREGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1REQUIREGROUP

`func (o *InlineObject60) SetAUTHLDAP1REQUIREGROUP(v string)`

SetAUTHLDAP1REQUIREGROUP sets AUTHLDAP1REQUIREGROUP field to given value.

### HasAUTHLDAP1REQUIREGROUP

`func (o *InlineObject60) HasAUTHLDAP1REQUIREGROUP() bool`

HasAUTHLDAP1REQUIREGROUP returns a boolean if a field has been set.

### GetAUTHLDAP1SERVERURI

`func (o *InlineObject60) GetAUTHLDAP1SERVERURI() string`

GetAUTHLDAP1SERVERURI returns the AUTHLDAP1SERVERURI field if non-nil, zero value otherwise.

### GetAUTHLDAP1SERVERURIOk

`func (o *InlineObject60) GetAUTHLDAP1SERVERURIOk() (*string, bool)`

GetAUTHLDAP1SERVERURIOk returns a tuple with the AUTHLDAP1SERVERURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1SERVERURI

`func (o *InlineObject60) SetAUTHLDAP1SERVERURI(v string)`

SetAUTHLDAP1SERVERURI sets AUTHLDAP1SERVERURI field to given value.

### HasAUTHLDAP1SERVERURI

`func (o *InlineObject60) HasAUTHLDAP1SERVERURI() bool`

HasAUTHLDAP1SERVERURI returns a boolean if a field has been set.

### GetAUTHLDAP1STARTTLS

`func (o *InlineObject60) GetAUTHLDAP1STARTTLS() bool`

GetAUTHLDAP1STARTTLS returns the AUTHLDAP1STARTTLS field if non-nil, zero value otherwise.

### GetAUTHLDAP1STARTTLSOk

`func (o *InlineObject60) GetAUTHLDAP1STARTTLSOk() (*bool, bool)`

GetAUTHLDAP1STARTTLSOk returns a tuple with the AUTHLDAP1STARTTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1STARTTLS

`func (o *InlineObject60) SetAUTHLDAP1STARTTLS(v bool)`

SetAUTHLDAP1STARTTLS sets AUTHLDAP1STARTTLS field to given value.

### HasAUTHLDAP1STARTTLS

`func (o *InlineObject60) HasAUTHLDAP1STARTTLS() bool`

HasAUTHLDAP1STARTTLS returns a boolean if a field has been set.

### GetAUTHLDAP1TEAMMAP

`func (o *InlineObject60) GetAUTHLDAP1TEAMMAP() map[string]interface{}`

GetAUTHLDAP1TEAMMAP returns the AUTHLDAP1TEAMMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP1TEAMMAPOk

`func (o *InlineObject60) GetAUTHLDAP1TEAMMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1TEAMMAPOk returns a tuple with the AUTHLDAP1TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1TEAMMAP

`func (o *InlineObject60) SetAUTHLDAP1TEAMMAP(v map[string]interface{})`

SetAUTHLDAP1TEAMMAP sets AUTHLDAP1TEAMMAP field to given value.

### HasAUTHLDAP1TEAMMAP

`func (o *InlineObject60) HasAUTHLDAP1TEAMMAP() bool`

HasAUTHLDAP1TEAMMAP returns a boolean if a field has been set.

### GetAUTHLDAP1USERATTRMAP

`func (o *InlineObject60) GetAUTHLDAP1USERATTRMAP() map[string]interface{}`

GetAUTHLDAP1USERATTRMAP returns the AUTHLDAP1USERATTRMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP1USERATTRMAPOk

`func (o *InlineObject60) GetAUTHLDAP1USERATTRMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1USERATTRMAPOk returns a tuple with the AUTHLDAP1USERATTRMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1USERATTRMAP

`func (o *InlineObject60) SetAUTHLDAP1USERATTRMAP(v map[string]interface{})`

SetAUTHLDAP1USERATTRMAP sets AUTHLDAP1USERATTRMAP field to given value.

### HasAUTHLDAP1USERATTRMAP

`func (o *InlineObject60) HasAUTHLDAP1USERATTRMAP() bool`

HasAUTHLDAP1USERATTRMAP returns a boolean if a field has been set.

### GetAUTHLDAP1USERDNTEMPLATE

`func (o *InlineObject60) GetAUTHLDAP1USERDNTEMPLATE() string`

GetAUTHLDAP1USERDNTEMPLATE returns the AUTHLDAP1USERDNTEMPLATE field if non-nil, zero value otherwise.

### GetAUTHLDAP1USERDNTEMPLATEOk

`func (o *InlineObject60) GetAUTHLDAP1USERDNTEMPLATEOk() (*string, bool)`

GetAUTHLDAP1USERDNTEMPLATEOk returns a tuple with the AUTHLDAP1USERDNTEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1USERDNTEMPLATE

`func (o *InlineObject60) SetAUTHLDAP1USERDNTEMPLATE(v string)`

SetAUTHLDAP1USERDNTEMPLATE sets AUTHLDAP1USERDNTEMPLATE field to given value.

### HasAUTHLDAP1USERDNTEMPLATE

`func (o *InlineObject60) HasAUTHLDAP1USERDNTEMPLATE() bool`

HasAUTHLDAP1USERDNTEMPLATE returns a boolean if a field has been set.

### GetAUTHLDAP1USERFLAGSBYGROUP

`func (o *InlineObject60) GetAUTHLDAP1USERFLAGSBYGROUP() map[string]interface{}`

GetAUTHLDAP1USERFLAGSBYGROUP returns the AUTHLDAP1USERFLAGSBYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP1USERFLAGSBYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP1USERFLAGSBYGROUPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP1USERFLAGSBYGROUPOk returns a tuple with the AUTHLDAP1USERFLAGSBYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1USERFLAGSBYGROUP

`func (o *InlineObject60) SetAUTHLDAP1USERFLAGSBYGROUP(v map[string]interface{})`

SetAUTHLDAP1USERFLAGSBYGROUP sets AUTHLDAP1USERFLAGSBYGROUP field to given value.

### HasAUTHLDAP1USERFLAGSBYGROUP

`func (o *InlineObject60) HasAUTHLDAP1USERFLAGSBYGROUP() bool`

HasAUTHLDAP1USERFLAGSBYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP1USERSEARCH

`func (o *InlineObject60) GetAUTHLDAP1USERSEARCH() []string`

GetAUTHLDAP1USERSEARCH returns the AUTHLDAP1USERSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP1USERSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP1USERSEARCHOk() (*[]string, bool)`

GetAUTHLDAP1USERSEARCHOk returns a tuple with the AUTHLDAP1USERSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP1USERSEARCH

`func (o *InlineObject60) SetAUTHLDAP1USERSEARCH(v []string)`

SetAUTHLDAP1USERSEARCH sets AUTHLDAP1USERSEARCH field to given value.

### HasAUTHLDAP1USERSEARCH

`func (o *InlineObject60) HasAUTHLDAP1USERSEARCH() bool`

HasAUTHLDAP1USERSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP2BINDDN

`func (o *InlineObject60) GetAUTHLDAP2BINDDN() string`

GetAUTHLDAP2BINDDN returns the AUTHLDAP2BINDDN field if non-nil, zero value otherwise.

### GetAUTHLDAP2BINDDNOk

`func (o *InlineObject60) GetAUTHLDAP2BINDDNOk() (*string, bool)`

GetAUTHLDAP2BINDDNOk returns a tuple with the AUTHLDAP2BINDDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2BINDDN

`func (o *InlineObject60) SetAUTHLDAP2BINDDN(v string)`

SetAUTHLDAP2BINDDN sets AUTHLDAP2BINDDN field to given value.

### HasAUTHLDAP2BINDDN

`func (o *InlineObject60) HasAUTHLDAP2BINDDN() bool`

HasAUTHLDAP2BINDDN returns a boolean if a field has been set.

### GetAUTHLDAP2BINDPASSWORD

`func (o *InlineObject60) GetAUTHLDAP2BINDPASSWORD() string`

GetAUTHLDAP2BINDPASSWORD returns the AUTHLDAP2BINDPASSWORD field if non-nil, zero value otherwise.

### GetAUTHLDAP2BINDPASSWORDOk

`func (o *InlineObject60) GetAUTHLDAP2BINDPASSWORDOk() (*string, bool)`

GetAUTHLDAP2BINDPASSWORDOk returns a tuple with the AUTHLDAP2BINDPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2BINDPASSWORD

`func (o *InlineObject60) SetAUTHLDAP2BINDPASSWORD(v string)`

SetAUTHLDAP2BINDPASSWORD sets AUTHLDAP2BINDPASSWORD field to given value.

### HasAUTHLDAP2BINDPASSWORD

`func (o *InlineObject60) HasAUTHLDAP2BINDPASSWORD() bool`

HasAUTHLDAP2BINDPASSWORD returns a boolean if a field has been set.

### GetAUTHLDAP2CONNECTIONOPTIONS

`func (o *InlineObject60) GetAUTHLDAP2CONNECTIONOPTIONS() map[string]interface{}`

GetAUTHLDAP2CONNECTIONOPTIONS returns the AUTHLDAP2CONNECTIONOPTIONS field if non-nil, zero value otherwise.

### GetAUTHLDAP2CONNECTIONOPTIONSOk

`func (o *InlineObject60) GetAUTHLDAP2CONNECTIONOPTIONSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2CONNECTIONOPTIONSOk returns a tuple with the AUTHLDAP2CONNECTIONOPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2CONNECTIONOPTIONS

`func (o *InlineObject60) SetAUTHLDAP2CONNECTIONOPTIONS(v map[string]interface{})`

SetAUTHLDAP2CONNECTIONOPTIONS sets AUTHLDAP2CONNECTIONOPTIONS field to given value.

### HasAUTHLDAP2CONNECTIONOPTIONS

`func (o *InlineObject60) HasAUTHLDAP2CONNECTIONOPTIONS() bool`

HasAUTHLDAP2CONNECTIONOPTIONS returns a boolean if a field has been set.

### GetAUTHLDAP2DENYGROUP

`func (o *InlineObject60) GetAUTHLDAP2DENYGROUP() string`

GetAUTHLDAP2DENYGROUP returns the AUTHLDAP2DENYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP2DENYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP2DENYGROUPOk() (*string, bool)`

GetAUTHLDAP2DENYGROUPOk returns a tuple with the AUTHLDAP2DENYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2DENYGROUP

`func (o *InlineObject60) SetAUTHLDAP2DENYGROUP(v string)`

SetAUTHLDAP2DENYGROUP sets AUTHLDAP2DENYGROUP field to given value.

### HasAUTHLDAP2DENYGROUP

`func (o *InlineObject60) HasAUTHLDAP2DENYGROUP() bool`

HasAUTHLDAP2DENYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP2GROUPSEARCH

`func (o *InlineObject60) GetAUTHLDAP2GROUPSEARCH() []string`

GetAUTHLDAP2GROUPSEARCH returns the AUTHLDAP2GROUPSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP2GROUPSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP2GROUPSEARCHOk() (*[]string, bool)`

GetAUTHLDAP2GROUPSEARCHOk returns a tuple with the AUTHLDAP2GROUPSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2GROUPSEARCH

`func (o *InlineObject60) SetAUTHLDAP2GROUPSEARCH(v []string)`

SetAUTHLDAP2GROUPSEARCH sets AUTHLDAP2GROUPSEARCH field to given value.

### HasAUTHLDAP2GROUPSEARCH

`func (o *InlineObject60) HasAUTHLDAP2GROUPSEARCH() bool`

HasAUTHLDAP2GROUPSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP2GROUPTYPE

`func (o *InlineObject60) GetAUTHLDAP2GROUPTYPE() string`

GetAUTHLDAP2GROUPTYPE returns the AUTHLDAP2GROUPTYPE field if non-nil, zero value otherwise.

### GetAUTHLDAP2GROUPTYPEOk

`func (o *InlineObject60) GetAUTHLDAP2GROUPTYPEOk() (*string, bool)`

GetAUTHLDAP2GROUPTYPEOk returns a tuple with the AUTHLDAP2GROUPTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2GROUPTYPE

`func (o *InlineObject60) SetAUTHLDAP2GROUPTYPE(v string)`

SetAUTHLDAP2GROUPTYPE sets AUTHLDAP2GROUPTYPE field to given value.

### HasAUTHLDAP2GROUPTYPE

`func (o *InlineObject60) HasAUTHLDAP2GROUPTYPE() bool`

HasAUTHLDAP2GROUPTYPE returns a boolean if a field has been set.

### GetAUTHLDAP2GROUPTYPEPARAMS

`func (o *InlineObject60) GetAUTHLDAP2GROUPTYPEPARAMS() map[string]interface{}`

GetAUTHLDAP2GROUPTYPEPARAMS returns the AUTHLDAP2GROUPTYPEPARAMS field if non-nil, zero value otherwise.

### GetAUTHLDAP2GROUPTYPEPARAMSOk

`func (o *InlineObject60) GetAUTHLDAP2GROUPTYPEPARAMSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2GROUPTYPEPARAMSOk returns a tuple with the AUTHLDAP2GROUPTYPEPARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2GROUPTYPEPARAMS

`func (o *InlineObject60) SetAUTHLDAP2GROUPTYPEPARAMS(v map[string]interface{})`

SetAUTHLDAP2GROUPTYPEPARAMS sets AUTHLDAP2GROUPTYPEPARAMS field to given value.

### HasAUTHLDAP2GROUPTYPEPARAMS

`func (o *InlineObject60) HasAUTHLDAP2GROUPTYPEPARAMS() bool`

HasAUTHLDAP2GROUPTYPEPARAMS returns a boolean if a field has been set.

### GetAUTHLDAP2ORGANIZATIONMAP

`func (o *InlineObject60) GetAUTHLDAP2ORGANIZATIONMAP() map[string]interface{}`

GetAUTHLDAP2ORGANIZATIONMAP returns the AUTHLDAP2ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP2ORGANIZATIONMAPOk

`func (o *InlineObject60) GetAUTHLDAP2ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2ORGANIZATIONMAPOk returns a tuple with the AUTHLDAP2ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2ORGANIZATIONMAP

`func (o *InlineObject60) SetAUTHLDAP2ORGANIZATIONMAP(v map[string]interface{})`

SetAUTHLDAP2ORGANIZATIONMAP sets AUTHLDAP2ORGANIZATIONMAP field to given value.

### HasAUTHLDAP2ORGANIZATIONMAP

`func (o *InlineObject60) HasAUTHLDAP2ORGANIZATIONMAP() bool`

HasAUTHLDAP2ORGANIZATIONMAP returns a boolean if a field has been set.

### GetAUTHLDAP2REQUIREGROUP

`func (o *InlineObject60) GetAUTHLDAP2REQUIREGROUP() string`

GetAUTHLDAP2REQUIREGROUP returns the AUTHLDAP2REQUIREGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP2REQUIREGROUPOk

`func (o *InlineObject60) GetAUTHLDAP2REQUIREGROUPOk() (*string, bool)`

GetAUTHLDAP2REQUIREGROUPOk returns a tuple with the AUTHLDAP2REQUIREGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2REQUIREGROUP

`func (o *InlineObject60) SetAUTHLDAP2REQUIREGROUP(v string)`

SetAUTHLDAP2REQUIREGROUP sets AUTHLDAP2REQUIREGROUP field to given value.

### HasAUTHLDAP2REQUIREGROUP

`func (o *InlineObject60) HasAUTHLDAP2REQUIREGROUP() bool`

HasAUTHLDAP2REQUIREGROUP returns a boolean if a field has been set.

### GetAUTHLDAP2SERVERURI

`func (o *InlineObject60) GetAUTHLDAP2SERVERURI() string`

GetAUTHLDAP2SERVERURI returns the AUTHLDAP2SERVERURI field if non-nil, zero value otherwise.

### GetAUTHLDAP2SERVERURIOk

`func (o *InlineObject60) GetAUTHLDAP2SERVERURIOk() (*string, bool)`

GetAUTHLDAP2SERVERURIOk returns a tuple with the AUTHLDAP2SERVERURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2SERVERURI

`func (o *InlineObject60) SetAUTHLDAP2SERVERURI(v string)`

SetAUTHLDAP2SERVERURI sets AUTHLDAP2SERVERURI field to given value.

### HasAUTHLDAP2SERVERURI

`func (o *InlineObject60) HasAUTHLDAP2SERVERURI() bool`

HasAUTHLDAP2SERVERURI returns a boolean if a field has been set.

### GetAUTHLDAP2STARTTLS

`func (o *InlineObject60) GetAUTHLDAP2STARTTLS() bool`

GetAUTHLDAP2STARTTLS returns the AUTHLDAP2STARTTLS field if non-nil, zero value otherwise.

### GetAUTHLDAP2STARTTLSOk

`func (o *InlineObject60) GetAUTHLDAP2STARTTLSOk() (*bool, bool)`

GetAUTHLDAP2STARTTLSOk returns a tuple with the AUTHLDAP2STARTTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2STARTTLS

`func (o *InlineObject60) SetAUTHLDAP2STARTTLS(v bool)`

SetAUTHLDAP2STARTTLS sets AUTHLDAP2STARTTLS field to given value.

### HasAUTHLDAP2STARTTLS

`func (o *InlineObject60) HasAUTHLDAP2STARTTLS() bool`

HasAUTHLDAP2STARTTLS returns a boolean if a field has been set.

### GetAUTHLDAP2TEAMMAP

`func (o *InlineObject60) GetAUTHLDAP2TEAMMAP() map[string]interface{}`

GetAUTHLDAP2TEAMMAP returns the AUTHLDAP2TEAMMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP2TEAMMAPOk

`func (o *InlineObject60) GetAUTHLDAP2TEAMMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2TEAMMAPOk returns a tuple with the AUTHLDAP2TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2TEAMMAP

`func (o *InlineObject60) SetAUTHLDAP2TEAMMAP(v map[string]interface{})`

SetAUTHLDAP2TEAMMAP sets AUTHLDAP2TEAMMAP field to given value.

### HasAUTHLDAP2TEAMMAP

`func (o *InlineObject60) HasAUTHLDAP2TEAMMAP() bool`

HasAUTHLDAP2TEAMMAP returns a boolean if a field has been set.

### GetAUTHLDAP2USERATTRMAP

`func (o *InlineObject60) GetAUTHLDAP2USERATTRMAP() map[string]interface{}`

GetAUTHLDAP2USERATTRMAP returns the AUTHLDAP2USERATTRMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP2USERATTRMAPOk

`func (o *InlineObject60) GetAUTHLDAP2USERATTRMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2USERATTRMAPOk returns a tuple with the AUTHLDAP2USERATTRMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2USERATTRMAP

`func (o *InlineObject60) SetAUTHLDAP2USERATTRMAP(v map[string]interface{})`

SetAUTHLDAP2USERATTRMAP sets AUTHLDAP2USERATTRMAP field to given value.

### HasAUTHLDAP2USERATTRMAP

`func (o *InlineObject60) HasAUTHLDAP2USERATTRMAP() bool`

HasAUTHLDAP2USERATTRMAP returns a boolean if a field has been set.

### GetAUTHLDAP2USERDNTEMPLATE

`func (o *InlineObject60) GetAUTHLDAP2USERDNTEMPLATE() string`

GetAUTHLDAP2USERDNTEMPLATE returns the AUTHLDAP2USERDNTEMPLATE field if non-nil, zero value otherwise.

### GetAUTHLDAP2USERDNTEMPLATEOk

`func (o *InlineObject60) GetAUTHLDAP2USERDNTEMPLATEOk() (*string, bool)`

GetAUTHLDAP2USERDNTEMPLATEOk returns a tuple with the AUTHLDAP2USERDNTEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2USERDNTEMPLATE

`func (o *InlineObject60) SetAUTHLDAP2USERDNTEMPLATE(v string)`

SetAUTHLDAP2USERDNTEMPLATE sets AUTHLDAP2USERDNTEMPLATE field to given value.

### HasAUTHLDAP2USERDNTEMPLATE

`func (o *InlineObject60) HasAUTHLDAP2USERDNTEMPLATE() bool`

HasAUTHLDAP2USERDNTEMPLATE returns a boolean if a field has been set.

### GetAUTHLDAP2USERFLAGSBYGROUP

`func (o *InlineObject60) GetAUTHLDAP2USERFLAGSBYGROUP() map[string]interface{}`

GetAUTHLDAP2USERFLAGSBYGROUP returns the AUTHLDAP2USERFLAGSBYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP2USERFLAGSBYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP2USERFLAGSBYGROUPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP2USERFLAGSBYGROUPOk returns a tuple with the AUTHLDAP2USERFLAGSBYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2USERFLAGSBYGROUP

`func (o *InlineObject60) SetAUTHLDAP2USERFLAGSBYGROUP(v map[string]interface{})`

SetAUTHLDAP2USERFLAGSBYGROUP sets AUTHLDAP2USERFLAGSBYGROUP field to given value.

### HasAUTHLDAP2USERFLAGSBYGROUP

`func (o *InlineObject60) HasAUTHLDAP2USERFLAGSBYGROUP() bool`

HasAUTHLDAP2USERFLAGSBYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP2USERSEARCH

`func (o *InlineObject60) GetAUTHLDAP2USERSEARCH() []string`

GetAUTHLDAP2USERSEARCH returns the AUTHLDAP2USERSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP2USERSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP2USERSEARCHOk() (*[]string, bool)`

GetAUTHLDAP2USERSEARCHOk returns a tuple with the AUTHLDAP2USERSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP2USERSEARCH

`func (o *InlineObject60) SetAUTHLDAP2USERSEARCH(v []string)`

SetAUTHLDAP2USERSEARCH sets AUTHLDAP2USERSEARCH field to given value.

### HasAUTHLDAP2USERSEARCH

`func (o *InlineObject60) HasAUTHLDAP2USERSEARCH() bool`

HasAUTHLDAP2USERSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP3BINDDN

`func (o *InlineObject60) GetAUTHLDAP3BINDDN() string`

GetAUTHLDAP3BINDDN returns the AUTHLDAP3BINDDN field if non-nil, zero value otherwise.

### GetAUTHLDAP3BINDDNOk

`func (o *InlineObject60) GetAUTHLDAP3BINDDNOk() (*string, bool)`

GetAUTHLDAP3BINDDNOk returns a tuple with the AUTHLDAP3BINDDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3BINDDN

`func (o *InlineObject60) SetAUTHLDAP3BINDDN(v string)`

SetAUTHLDAP3BINDDN sets AUTHLDAP3BINDDN field to given value.

### HasAUTHLDAP3BINDDN

`func (o *InlineObject60) HasAUTHLDAP3BINDDN() bool`

HasAUTHLDAP3BINDDN returns a boolean if a field has been set.

### GetAUTHLDAP3BINDPASSWORD

`func (o *InlineObject60) GetAUTHLDAP3BINDPASSWORD() string`

GetAUTHLDAP3BINDPASSWORD returns the AUTHLDAP3BINDPASSWORD field if non-nil, zero value otherwise.

### GetAUTHLDAP3BINDPASSWORDOk

`func (o *InlineObject60) GetAUTHLDAP3BINDPASSWORDOk() (*string, bool)`

GetAUTHLDAP3BINDPASSWORDOk returns a tuple with the AUTHLDAP3BINDPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3BINDPASSWORD

`func (o *InlineObject60) SetAUTHLDAP3BINDPASSWORD(v string)`

SetAUTHLDAP3BINDPASSWORD sets AUTHLDAP3BINDPASSWORD field to given value.

### HasAUTHLDAP3BINDPASSWORD

`func (o *InlineObject60) HasAUTHLDAP3BINDPASSWORD() bool`

HasAUTHLDAP3BINDPASSWORD returns a boolean if a field has been set.

### GetAUTHLDAP3CONNECTIONOPTIONS

`func (o *InlineObject60) GetAUTHLDAP3CONNECTIONOPTIONS() map[string]interface{}`

GetAUTHLDAP3CONNECTIONOPTIONS returns the AUTHLDAP3CONNECTIONOPTIONS field if non-nil, zero value otherwise.

### GetAUTHLDAP3CONNECTIONOPTIONSOk

`func (o *InlineObject60) GetAUTHLDAP3CONNECTIONOPTIONSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3CONNECTIONOPTIONSOk returns a tuple with the AUTHLDAP3CONNECTIONOPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3CONNECTIONOPTIONS

`func (o *InlineObject60) SetAUTHLDAP3CONNECTIONOPTIONS(v map[string]interface{})`

SetAUTHLDAP3CONNECTIONOPTIONS sets AUTHLDAP3CONNECTIONOPTIONS field to given value.

### HasAUTHLDAP3CONNECTIONOPTIONS

`func (o *InlineObject60) HasAUTHLDAP3CONNECTIONOPTIONS() bool`

HasAUTHLDAP3CONNECTIONOPTIONS returns a boolean if a field has been set.

### GetAUTHLDAP3DENYGROUP

`func (o *InlineObject60) GetAUTHLDAP3DENYGROUP() string`

GetAUTHLDAP3DENYGROUP returns the AUTHLDAP3DENYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP3DENYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP3DENYGROUPOk() (*string, bool)`

GetAUTHLDAP3DENYGROUPOk returns a tuple with the AUTHLDAP3DENYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3DENYGROUP

`func (o *InlineObject60) SetAUTHLDAP3DENYGROUP(v string)`

SetAUTHLDAP3DENYGROUP sets AUTHLDAP3DENYGROUP field to given value.

### HasAUTHLDAP3DENYGROUP

`func (o *InlineObject60) HasAUTHLDAP3DENYGROUP() bool`

HasAUTHLDAP3DENYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP3GROUPSEARCH

`func (o *InlineObject60) GetAUTHLDAP3GROUPSEARCH() []string`

GetAUTHLDAP3GROUPSEARCH returns the AUTHLDAP3GROUPSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP3GROUPSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP3GROUPSEARCHOk() (*[]string, bool)`

GetAUTHLDAP3GROUPSEARCHOk returns a tuple with the AUTHLDAP3GROUPSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3GROUPSEARCH

`func (o *InlineObject60) SetAUTHLDAP3GROUPSEARCH(v []string)`

SetAUTHLDAP3GROUPSEARCH sets AUTHLDAP3GROUPSEARCH field to given value.

### HasAUTHLDAP3GROUPSEARCH

`func (o *InlineObject60) HasAUTHLDAP3GROUPSEARCH() bool`

HasAUTHLDAP3GROUPSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP3GROUPTYPE

`func (o *InlineObject60) GetAUTHLDAP3GROUPTYPE() string`

GetAUTHLDAP3GROUPTYPE returns the AUTHLDAP3GROUPTYPE field if non-nil, zero value otherwise.

### GetAUTHLDAP3GROUPTYPEOk

`func (o *InlineObject60) GetAUTHLDAP3GROUPTYPEOk() (*string, bool)`

GetAUTHLDAP3GROUPTYPEOk returns a tuple with the AUTHLDAP3GROUPTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3GROUPTYPE

`func (o *InlineObject60) SetAUTHLDAP3GROUPTYPE(v string)`

SetAUTHLDAP3GROUPTYPE sets AUTHLDAP3GROUPTYPE field to given value.

### HasAUTHLDAP3GROUPTYPE

`func (o *InlineObject60) HasAUTHLDAP3GROUPTYPE() bool`

HasAUTHLDAP3GROUPTYPE returns a boolean if a field has been set.

### GetAUTHLDAP3GROUPTYPEPARAMS

`func (o *InlineObject60) GetAUTHLDAP3GROUPTYPEPARAMS() map[string]interface{}`

GetAUTHLDAP3GROUPTYPEPARAMS returns the AUTHLDAP3GROUPTYPEPARAMS field if non-nil, zero value otherwise.

### GetAUTHLDAP3GROUPTYPEPARAMSOk

`func (o *InlineObject60) GetAUTHLDAP3GROUPTYPEPARAMSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3GROUPTYPEPARAMSOk returns a tuple with the AUTHLDAP3GROUPTYPEPARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3GROUPTYPEPARAMS

`func (o *InlineObject60) SetAUTHLDAP3GROUPTYPEPARAMS(v map[string]interface{})`

SetAUTHLDAP3GROUPTYPEPARAMS sets AUTHLDAP3GROUPTYPEPARAMS field to given value.

### HasAUTHLDAP3GROUPTYPEPARAMS

`func (o *InlineObject60) HasAUTHLDAP3GROUPTYPEPARAMS() bool`

HasAUTHLDAP3GROUPTYPEPARAMS returns a boolean if a field has been set.

### GetAUTHLDAP3ORGANIZATIONMAP

`func (o *InlineObject60) GetAUTHLDAP3ORGANIZATIONMAP() map[string]interface{}`

GetAUTHLDAP3ORGANIZATIONMAP returns the AUTHLDAP3ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP3ORGANIZATIONMAPOk

`func (o *InlineObject60) GetAUTHLDAP3ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3ORGANIZATIONMAPOk returns a tuple with the AUTHLDAP3ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3ORGANIZATIONMAP

`func (o *InlineObject60) SetAUTHLDAP3ORGANIZATIONMAP(v map[string]interface{})`

SetAUTHLDAP3ORGANIZATIONMAP sets AUTHLDAP3ORGANIZATIONMAP field to given value.

### HasAUTHLDAP3ORGANIZATIONMAP

`func (o *InlineObject60) HasAUTHLDAP3ORGANIZATIONMAP() bool`

HasAUTHLDAP3ORGANIZATIONMAP returns a boolean if a field has been set.

### GetAUTHLDAP3REQUIREGROUP

`func (o *InlineObject60) GetAUTHLDAP3REQUIREGROUP() string`

GetAUTHLDAP3REQUIREGROUP returns the AUTHLDAP3REQUIREGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP3REQUIREGROUPOk

`func (o *InlineObject60) GetAUTHLDAP3REQUIREGROUPOk() (*string, bool)`

GetAUTHLDAP3REQUIREGROUPOk returns a tuple with the AUTHLDAP3REQUIREGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3REQUIREGROUP

`func (o *InlineObject60) SetAUTHLDAP3REQUIREGROUP(v string)`

SetAUTHLDAP3REQUIREGROUP sets AUTHLDAP3REQUIREGROUP field to given value.

### HasAUTHLDAP3REQUIREGROUP

`func (o *InlineObject60) HasAUTHLDAP3REQUIREGROUP() bool`

HasAUTHLDAP3REQUIREGROUP returns a boolean if a field has been set.

### GetAUTHLDAP3SERVERURI

`func (o *InlineObject60) GetAUTHLDAP3SERVERURI() string`

GetAUTHLDAP3SERVERURI returns the AUTHLDAP3SERVERURI field if non-nil, zero value otherwise.

### GetAUTHLDAP3SERVERURIOk

`func (o *InlineObject60) GetAUTHLDAP3SERVERURIOk() (*string, bool)`

GetAUTHLDAP3SERVERURIOk returns a tuple with the AUTHLDAP3SERVERURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3SERVERURI

`func (o *InlineObject60) SetAUTHLDAP3SERVERURI(v string)`

SetAUTHLDAP3SERVERURI sets AUTHLDAP3SERVERURI field to given value.

### HasAUTHLDAP3SERVERURI

`func (o *InlineObject60) HasAUTHLDAP3SERVERURI() bool`

HasAUTHLDAP3SERVERURI returns a boolean if a field has been set.

### GetAUTHLDAP3STARTTLS

`func (o *InlineObject60) GetAUTHLDAP3STARTTLS() bool`

GetAUTHLDAP3STARTTLS returns the AUTHLDAP3STARTTLS field if non-nil, zero value otherwise.

### GetAUTHLDAP3STARTTLSOk

`func (o *InlineObject60) GetAUTHLDAP3STARTTLSOk() (*bool, bool)`

GetAUTHLDAP3STARTTLSOk returns a tuple with the AUTHLDAP3STARTTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3STARTTLS

`func (o *InlineObject60) SetAUTHLDAP3STARTTLS(v bool)`

SetAUTHLDAP3STARTTLS sets AUTHLDAP3STARTTLS field to given value.

### HasAUTHLDAP3STARTTLS

`func (o *InlineObject60) HasAUTHLDAP3STARTTLS() bool`

HasAUTHLDAP3STARTTLS returns a boolean if a field has been set.

### GetAUTHLDAP3TEAMMAP

`func (o *InlineObject60) GetAUTHLDAP3TEAMMAP() map[string]interface{}`

GetAUTHLDAP3TEAMMAP returns the AUTHLDAP3TEAMMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP3TEAMMAPOk

`func (o *InlineObject60) GetAUTHLDAP3TEAMMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3TEAMMAPOk returns a tuple with the AUTHLDAP3TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3TEAMMAP

`func (o *InlineObject60) SetAUTHLDAP3TEAMMAP(v map[string]interface{})`

SetAUTHLDAP3TEAMMAP sets AUTHLDAP3TEAMMAP field to given value.

### HasAUTHLDAP3TEAMMAP

`func (o *InlineObject60) HasAUTHLDAP3TEAMMAP() bool`

HasAUTHLDAP3TEAMMAP returns a boolean if a field has been set.

### GetAUTHLDAP3USERATTRMAP

`func (o *InlineObject60) GetAUTHLDAP3USERATTRMAP() map[string]interface{}`

GetAUTHLDAP3USERATTRMAP returns the AUTHLDAP3USERATTRMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP3USERATTRMAPOk

`func (o *InlineObject60) GetAUTHLDAP3USERATTRMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3USERATTRMAPOk returns a tuple with the AUTHLDAP3USERATTRMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3USERATTRMAP

`func (o *InlineObject60) SetAUTHLDAP3USERATTRMAP(v map[string]interface{})`

SetAUTHLDAP3USERATTRMAP sets AUTHLDAP3USERATTRMAP field to given value.

### HasAUTHLDAP3USERATTRMAP

`func (o *InlineObject60) HasAUTHLDAP3USERATTRMAP() bool`

HasAUTHLDAP3USERATTRMAP returns a boolean if a field has been set.

### GetAUTHLDAP3USERDNTEMPLATE

`func (o *InlineObject60) GetAUTHLDAP3USERDNTEMPLATE() string`

GetAUTHLDAP3USERDNTEMPLATE returns the AUTHLDAP3USERDNTEMPLATE field if non-nil, zero value otherwise.

### GetAUTHLDAP3USERDNTEMPLATEOk

`func (o *InlineObject60) GetAUTHLDAP3USERDNTEMPLATEOk() (*string, bool)`

GetAUTHLDAP3USERDNTEMPLATEOk returns a tuple with the AUTHLDAP3USERDNTEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3USERDNTEMPLATE

`func (o *InlineObject60) SetAUTHLDAP3USERDNTEMPLATE(v string)`

SetAUTHLDAP3USERDNTEMPLATE sets AUTHLDAP3USERDNTEMPLATE field to given value.

### HasAUTHLDAP3USERDNTEMPLATE

`func (o *InlineObject60) HasAUTHLDAP3USERDNTEMPLATE() bool`

HasAUTHLDAP3USERDNTEMPLATE returns a boolean if a field has been set.

### GetAUTHLDAP3USERFLAGSBYGROUP

`func (o *InlineObject60) GetAUTHLDAP3USERFLAGSBYGROUP() map[string]interface{}`

GetAUTHLDAP3USERFLAGSBYGROUP returns the AUTHLDAP3USERFLAGSBYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP3USERFLAGSBYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP3USERFLAGSBYGROUPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP3USERFLAGSBYGROUPOk returns a tuple with the AUTHLDAP3USERFLAGSBYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3USERFLAGSBYGROUP

`func (o *InlineObject60) SetAUTHLDAP3USERFLAGSBYGROUP(v map[string]interface{})`

SetAUTHLDAP3USERFLAGSBYGROUP sets AUTHLDAP3USERFLAGSBYGROUP field to given value.

### HasAUTHLDAP3USERFLAGSBYGROUP

`func (o *InlineObject60) HasAUTHLDAP3USERFLAGSBYGROUP() bool`

HasAUTHLDAP3USERFLAGSBYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP3USERSEARCH

`func (o *InlineObject60) GetAUTHLDAP3USERSEARCH() []string`

GetAUTHLDAP3USERSEARCH returns the AUTHLDAP3USERSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP3USERSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP3USERSEARCHOk() (*[]string, bool)`

GetAUTHLDAP3USERSEARCHOk returns a tuple with the AUTHLDAP3USERSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP3USERSEARCH

`func (o *InlineObject60) SetAUTHLDAP3USERSEARCH(v []string)`

SetAUTHLDAP3USERSEARCH sets AUTHLDAP3USERSEARCH field to given value.

### HasAUTHLDAP3USERSEARCH

`func (o *InlineObject60) HasAUTHLDAP3USERSEARCH() bool`

HasAUTHLDAP3USERSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP4BINDDN

`func (o *InlineObject60) GetAUTHLDAP4BINDDN() string`

GetAUTHLDAP4BINDDN returns the AUTHLDAP4BINDDN field if non-nil, zero value otherwise.

### GetAUTHLDAP4BINDDNOk

`func (o *InlineObject60) GetAUTHLDAP4BINDDNOk() (*string, bool)`

GetAUTHLDAP4BINDDNOk returns a tuple with the AUTHLDAP4BINDDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4BINDDN

`func (o *InlineObject60) SetAUTHLDAP4BINDDN(v string)`

SetAUTHLDAP4BINDDN sets AUTHLDAP4BINDDN field to given value.

### HasAUTHLDAP4BINDDN

`func (o *InlineObject60) HasAUTHLDAP4BINDDN() bool`

HasAUTHLDAP4BINDDN returns a boolean if a field has been set.

### GetAUTHLDAP4BINDPASSWORD

`func (o *InlineObject60) GetAUTHLDAP4BINDPASSWORD() string`

GetAUTHLDAP4BINDPASSWORD returns the AUTHLDAP4BINDPASSWORD field if non-nil, zero value otherwise.

### GetAUTHLDAP4BINDPASSWORDOk

`func (o *InlineObject60) GetAUTHLDAP4BINDPASSWORDOk() (*string, bool)`

GetAUTHLDAP4BINDPASSWORDOk returns a tuple with the AUTHLDAP4BINDPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4BINDPASSWORD

`func (o *InlineObject60) SetAUTHLDAP4BINDPASSWORD(v string)`

SetAUTHLDAP4BINDPASSWORD sets AUTHLDAP4BINDPASSWORD field to given value.

### HasAUTHLDAP4BINDPASSWORD

`func (o *InlineObject60) HasAUTHLDAP4BINDPASSWORD() bool`

HasAUTHLDAP4BINDPASSWORD returns a boolean if a field has been set.

### GetAUTHLDAP4CONNECTIONOPTIONS

`func (o *InlineObject60) GetAUTHLDAP4CONNECTIONOPTIONS() map[string]interface{}`

GetAUTHLDAP4CONNECTIONOPTIONS returns the AUTHLDAP4CONNECTIONOPTIONS field if non-nil, zero value otherwise.

### GetAUTHLDAP4CONNECTIONOPTIONSOk

`func (o *InlineObject60) GetAUTHLDAP4CONNECTIONOPTIONSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4CONNECTIONOPTIONSOk returns a tuple with the AUTHLDAP4CONNECTIONOPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4CONNECTIONOPTIONS

`func (o *InlineObject60) SetAUTHLDAP4CONNECTIONOPTIONS(v map[string]interface{})`

SetAUTHLDAP4CONNECTIONOPTIONS sets AUTHLDAP4CONNECTIONOPTIONS field to given value.

### HasAUTHLDAP4CONNECTIONOPTIONS

`func (o *InlineObject60) HasAUTHLDAP4CONNECTIONOPTIONS() bool`

HasAUTHLDAP4CONNECTIONOPTIONS returns a boolean if a field has been set.

### GetAUTHLDAP4DENYGROUP

`func (o *InlineObject60) GetAUTHLDAP4DENYGROUP() string`

GetAUTHLDAP4DENYGROUP returns the AUTHLDAP4DENYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP4DENYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP4DENYGROUPOk() (*string, bool)`

GetAUTHLDAP4DENYGROUPOk returns a tuple with the AUTHLDAP4DENYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4DENYGROUP

`func (o *InlineObject60) SetAUTHLDAP4DENYGROUP(v string)`

SetAUTHLDAP4DENYGROUP sets AUTHLDAP4DENYGROUP field to given value.

### HasAUTHLDAP4DENYGROUP

`func (o *InlineObject60) HasAUTHLDAP4DENYGROUP() bool`

HasAUTHLDAP4DENYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP4GROUPSEARCH

`func (o *InlineObject60) GetAUTHLDAP4GROUPSEARCH() []string`

GetAUTHLDAP4GROUPSEARCH returns the AUTHLDAP4GROUPSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP4GROUPSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP4GROUPSEARCHOk() (*[]string, bool)`

GetAUTHLDAP4GROUPSEARCHOk returns a tuple with the AUTHLDAP4GROUPSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4GROUPSEARCH

`func (o *InlineObject60) SetAUTHLDAP4GROUPSEARCH(v []string)`

SetAUTHLDAP4GROUPSEARCH sets AUTHLDAP4GROUPSEARCH field to given value.

### HasAUTHLDAP4GROUPSEARCH

`func (o *InlineObject60) HasAUTHLDAP4GROUPSEARCH() bool`

HasAUTHLDAP4GROUPSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP4GROUPTYPE

`func (o *InlineObject60) GetAUTHLDAP4GROUPTYPE() string`

GetAUTHLDAP4GROUPTYPE returns the AUTHLDAP4GROUPTYPE field if non-nil, zero value otherwise.

### GetAUTHLDAP4GROUPTYPEOk

`func (o *InlineObject60) GetAUTHLDAP4GROUPTYPEOk() (*string, bool)`

GetAUTHLDAP4GROUPTYPEOk returns a tuple with the AUTHLDAP4GROUPTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4GROUPTYPE

`func (o *InlineObject60) SetAUTHLDAP4GROUPTYPE(v string)`

SetAUTHLDAP4GROUPTYPE sets AUTHLDAP4GROUPTYPE field to given value.

### HasAUTHLDAP4GROUPTYPE

`func (o *InlineObject60) HasAUTHLDAP4GROUPTYPE() bool`

HasAUTHLDAP4GROUPTYPE returns a boolean if a field has been set.

### GetAUTHLDAP4GROUPTYPEPARAMS

`func (o *InlineObject60) GetAUTHLDAP4GROUPTYPEPARAMS() map[string]interface{}`

GetAUTHLDAP4GROUPTYPEPARAMS returns the AUTHLDAP4GROUPTYPEPARAMS field if non-nil, zero value otherwise.

### GetAUTHLDAP4GROUPTYPEPARAMSOk

`func (o *InlineObject60) GetAUTHLDAP4GROUPTYPEPARAMSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4GROUPTYPEPARAMSOk returns a tuple with the AUTHLDAP4GROUPTYPEPARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4GROUPTYPEPARAMS

`func (o *InlineObject60) SetAUTHLDAP4GROUPTYPEPARAMS(v map[string]interface{})`

SetAUTHLDAP4GROUPTYPEPARAMS sets AUTHLDAP4GROUPTYPEPARAMS field to given value.

### HasAUTHLDAP4GROUPTYPEPARAMS

`func (o *InlineObject60) HasAUTHLDAP4GROUPTYPEPARAMS() bool`

HasAUTHLDAP4GROUPTYPEPARAMS returns a boolean if a field has been set.

### GetAUTHLDAP4ORGANIZATIONMAP

`func (o *InlineObject60) GetAUTHLDAP4ORGANIZATIONMAP() map[string]interface{}`

GetAUTHLDAP4ORGANIZATIONMAP returns the AUTHLDAP4ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP4ORGANIZATIONMAPOk

`func (o *InlineObject60) GetAUTHLDAP4ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4ORGANIZATIONMAPOk returns a tuple with the AUTHLDAP4ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4ORGANIZATIONMAP

`func (o *InlineObject60) SetAUTHLDAP4ORGANIZATIONMAP(v map[string]interface{})`

SetAUTHLDAP4ORGANIZATIONMAP sets AUTHLDAP4ORGANIZATIONMAP field to given value.

### HasAUTHLDAP4ORGANIZATIONMAP

`func (o *InlineObject60) HasAUTHLDAP4ORGANIZATIONMAP() bool`

HasAUTHLDAP4ORGANIZATIONMAP returns a boolean if a field has been set.

### GetAUTHLDAP4REQUIREGROUP

`func (o *InlineObject60) GetAUTHLDAP4REQUIREGROUP() string`

GetAUTHLDAP4REQUIREGROUP returns the AUTHLDAP4REQUIREGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP4REQUIREGROUPOk

`func (o *InlineObject60) GetAUTHLDAP4REQUIREGROUPOk() (*string, bool)`

GetAUTHLDAP4REQUIREGROUPOk returns a tuple with the AUTHLDAP4REQUIREGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4REQUIREGROUP

`func (o *InlineObject60) SetAUTHLDAP4REQUIREGROUP(v string)`

SetAUTHLDAP4REQUIREGROUP sets AUTHLDAP4REQUIREGROUP field to given value.

### HasAUTHLDAP4REQUIREGROUP

`func (o *InlineObject60) HasAUTHLDAP4REQUIREGROUP() bool`

HasAUTHLDAP4REQUIREGROUP returns a boolean if a field has been set.

### GetAUTHLDAP4SERVERURI

`func (o *InlineObject60) GetAUTHLDAP4SERVERURI() string`

GetAUTHLDAP4SERVERURI returns the AUTHLDAP4SERVERURI field if non-nil, zero value otherwise.

### GetAUTHLDAP4SERVERURIOk

`func (o *InlineObject60) GetAUTHLDAP4SERVERURIOk() (*string, bool)`

GetAUTHLDAP4SERVERURIOk returns a tuple with the AUTHLDAP4SERVERURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4SERVERURI

`func (o *InlineObject60) SetAUTHLDAP4SERVERURI(v string)`

SetAUTHLDAP4SERVERURI sets AUTHLDAP4SERVERURI field to given value.

### HasAUTHLDAP4SERVERURI

`func (o *InlineObject60) HasAUTHLDAP4SERVERURI() bool`

HasAUTHLDAP4SERVERURI returns a boolean if a field has been set.

### GetAUTHLDAP4STARTTLS

`func (o *InlineObject60) GetAUTHLDAP4STARTTLS() bool`

GetAUTHLDAP4STARTTLS returns the AUTHLDAP4STARTTLS field if non-nil, zero value otherwise.

### GetAUTHLDAP4STARTTLSOk

`func (o *InlineObject60) GetAUTHLDAP4STARTTLSOk() (*bool, bool)`

GetAUTHLDAP4STARTTLSOk returns a tuple with the AUTHLDAP4STARTTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4STARTTLS

`func (o *InlineObject60) SetAUTHLDAP4STARTTLS(v bool)`

SetAUTHLDAP4STARTTLS sets AUTHLDAP4STARTTLS field to given value.

### HasAUTHLDAP4STARTTLS

`func (o *InlineObject60) HasAUTHLDAP4STARTTLS() bool`

HasAUTHLDAP4STARTTLS returns a boolean if a field has been set.

### GetAUTHLDAP4TEAMMAP

`func (o *InlineObject60) GetAUTHLDAP4TEAMMAP() map[string]interface{}`

GetAUTHLDAP4TEAMMAP returns the AUTHLDAP4TEAMMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP4TEAMMAPOk

`func (o *InlineObject60) GetAUTHLDAP4TEAMMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4TEAMMAPOk returns a tuple with the AUTHLDAP4TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4TEAMMAP

`func (o *InlineObject60) SetAUTHLDAP4TEAMMAP(v map[string]interface{})`

SetAUTHLDAP4TEAMMAP sets AUTHLDAP4TEAMMAP field to given value.

### HasAUTHLDAP4TEAMMAP

`func (o *InlineObject60) HasAUTHLDAP4TEAMMAP() bool`

HasAUTHLDAP4TEAMMAP returns a boolean if a field has been set.

### GetAUTHLDAP4USERATTRMAP

`func (o *InlineObject60) GetAUTHLDAP4USERATTRMAP() map[string]interface{}`

GetAUTHLDAP4USERATTRMAP returns the AUTHLDAP4USERATTRMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP4USERATTRMAPOk

`func (o *InlineObject60) GetAUTHLDAP4USERATTRMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4USERATTRMAPOk returns a tuple with the AUTHLDAP4USERATTRMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4USERATTRMAP

`func (o *InlineObject60) SetAUTHLDAP4USERATTRMAP(v map[string]interface{})`

SetAUTHLDAP4USERATTRMAP sets AUTHLDAP4USERATTRMAP field to given value.

### HasAUTHLDAP4USERATTRMAP

`func (o *InlineObject60) HasAUTHLDAP4USERATTRMAP() bool`

HasAUTHLDAP4USERATTRMAP returns a boolean if a field has been set.

### GetAUTHLDAP4USERDNTEMPLATE

`func (o *InlineObject60) GetAUTHLDAP4USERDNTEMPLATE() string`

GetAUTHLDAP4USERDNTEMPLATE returns the AUTHLDAP4USERDNTEMPLATE field if non-nil, zero value otherwise.

### GetAUTHLDAP4USERDNTEMPLATEOk

`func (o *InlineObject60) GetAUTHLDAP4USERDNTEMPLATEOk() (*string, bool)`

GetAUTHLDAP4USERDNTEMPLATEOk returns a tuple with the AUTHLDAP4USERDNTEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4USERDNTEMPLATE

`func (o *InlineObject60) SetAUTHLDAP4USERDNTEMPLATE(v string)`

SetAUTHLDAP4USERDNTEMPLATE sets AUTHLDAP4USERDNTEMPLATE field to given value.

### HasAUTHLDAP4USERDNTEMPLATE

`func (o *InlineObject60) HasAUTHLDAP4USERDNTEMPLATE() bool`

HasAUTHLDAP4USERDNTEMPLATE returns a boolean if a field has been set.

### GetAUTHLDAP4USERFLAGSBYGROUP

`func (o *InlineObject60) GetAUTHLDAP4USERFLAGSBYGROUP() map[string]interface{}`

GetAUTHLDAP4USERFLAGSBYGROUP returns the AUTHLDAP4USERFLAGSBYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP4USERFLAGSBYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP4USERFLAGSBYGROUPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP4USERFLAGSBYGROUPOk returns a tuple with the AUTHLDAP4USERFLAGSBYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4USERFLAGSBYGROUP

`func (o *InlineObject60) SetAUTHLDAP4USERFLAGSBYGROUP(v map[string]interface{})`

SetAUTHLDAP4USERFLAGSBYGROUP sets AUTHLDAP4USERFLAGSBYGROUP field to given value.

### HasAUTHLDAP4USERFLAGSBYGROUP

`func (o *InlineObject60) HasAUTHLDAP4USERFLAGSBYGROUP() bool`

HasAUTHLDAP4USERFLAGSBYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP4USERSEARCH

`func (o *InlineObject60) GetAUTHLDAP4USERSEARCH() []string`

GetAUTHLDAP4USERSEARCH returns the AUTHLDAP4USERSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP4USERSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP4USERSEARCHOk() (*[]string, bool)`

GetAUTHLDAP4USERSEARCHOk returns a tuple with the AUTHLDAP4USERSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP4USERSEARCH

`func (o *InlineObject60) SetAUTHLDAP4USERSEARCH(v []string)`

SetAUTHLDAP4USERSEARCH sets AUTHLDAP4USERSEARCH field to given value.

### HasAUTHLDAP4USERSEARCH

`func (o *InlineObject60) HasAUTHLDAP4USERSEARCH() bool`

HasAUTHLDAP4USERSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP5BINDDN

`func (o *InlineObject60) GetAUTHLDAP5BINDDN() string`

GetAUTHLDAP5BINDDN returns the AUTHLDAP5BINDDN field if non-nil, zero value otherwise.

### GetAUTHLDAP5BINDDNOk

`func (o *InlineObject60) GetAUTHLDAP5BINDDNOk() (*string, bool)`

GetAUTHLDAP5BINDDNOk returns a tuple with the AUTHLDAP5BINDDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5BINDDN

`func (o *InlineObject60) SetAUTHLDAP5BINDDN(v string)`

SetAUTHLDAP5BINDDN sets AUTHLDAP5BINDDN field to given value.

### HasAUTHLDAP5BINDDN

`func (o *InlineObject60) HasAUTHLDAP5BINDDN() bool`

HasAUTHLDAP5BINDDN returns a boolean if a field has been set.

### GetAUTHLDAP5BINDPASSWORD

`func (o *InlineObject60) GetAUTHLDAP5BINDPASSWORD() string`

GetAUTHLDAP5BINDPASSWORD returns the AUTHLDAP5BINDPASSWORD field if non-nil, zero value otherwise.

### GetAUTHLDAP5BINDPASSWORDOk

`func (o *InlineObject60) GetAUTHLDAP5BINDPASSWORDOk() (*string, bool)`

GetAUTHLDAP5BINDPASSWORDOk returns a tuple with the AUTHLDAP5BINDPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5BINDPASSWORD

`func (o *InlineObject60) SetAUTHLDAP5BINDPASSWORD(v string)`

SetAUTHLDAP5BINDPASSWORD sets AUTHLDAP5BINDPASSWORD field to given value.

### HasAUTHLDAP5BINDPASSWORD

`func (o *InlineObject60) HasAUTHLDAP5BINDPASSWORD() bool`

HasAUTHLDAP5BINDPASSWORD returns a boolean if a field has been set.

### GetAUTHLDAP5CONNECTIONOPTIONS

`func (o *InlineObject60) GetAUTHLDAP5CONNECTIONOPTIONS() map[string]interface{}`

GetAUTHLDAP5CONNECTIONOPTIONS returns the AUTHLDAP5CONNECTIONOPTIONS field if non-nil, zero value otherwise.

### GetAUTHLDAP5CONNECTIONOPTIONSOk

`func (o *InlineObject60) GetAUTHLDAP5CONNECTIONOPTIONSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5CONNECTIONOPTIONSOk returns a tuple with the AUTHLDAP5CONNECTIONOPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5CONNECTIONOPTIONS

`func (o *InlineObject60) SetAUTHLDAP5CONNECTIONOPTIONS(v map[string]interface{})`

SetAUTHLDAP5CONNECTIONOPTIONS sets AUTHLDAP5CONNECTIONOPTIONS field to given value.

### HasAUTHLDAP5CONNECTIONOPTIONS

`func (o *InlineObject60) HasAUTHLDAP5CONNECTIONOPTIONS() bool`

HasAUTHLDAP5CONNECTIONOPTIONS returns a boolean if a field has been set.

### GetAUTHLDAP5DENYGROUP

`func (o *InlineObject60) GetAUTHLDAP5DENYGROUP() string`

GetAUTHLDAP5DENYGROUP returns the AUTHLDAP5DENYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP5DENYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP5DENYGROUPOk() (*string, bool)`

GetAUTHLDAP5DENYGROUPOk returns a tuple with the AUTHLDAP5DENYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5DENYGROUP

`func (o *InlineObject60) SetAUTHLDAP5DENYGROUP(v string)`

SetAUTHLDAP5DENYGROUP sets AUTHLDAP5DENYGROUP field to given value.

### HasAUTHLDAP5DENYGROUP

`func (o *InlineObject60) HasAUTHLDAP5DENYGROUP() bool`

HasAUTHLDAP5DENYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP5GROUPSEARCH

`func (o *InlineObject60) GetAUTHLDAP5GROUPSEARCH() []string`

GetAUTHLDAP5GROUPSEARCH returns the AUTHLDAP5GROUPSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP5GROUPSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP5GROUPSEARCHOk() (*[]string, bool)`

GetAUTHLDAP5GROUPSEARCHOk returns a tuple with the AUTHLDAP5GROUPSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5GROUPSEARCH

`func (o *InlineObject60) SetAUTHLDAP5GROUPSEARCH(v []string)`

SetAUTHLDAP5GROUPSEARCH sets AUTHLDAP5GROUPSEARCH field to given value.

### HasAUTHLDAP5GROUPSEARCH

`func (o *InlineObject60) HasAUTHLDAP5GROUPSEARCH() bool`

HasAUTHLDAP5GROUPSEARCH returns a boolean if a field has been set.

### GetAUTHLDAP5GROUPTYPE

`func (o *InlineObject60) GetAUTHLDAP5GROUPTYPE() string`

GetAUTHLDAP5GROUPTYPE returns the AUTHLDAP5GROUPTYPE field if non-nil, zero value otherwise.

### GetAUTHLDAP5GROUPTYPEOk

`func (o *InlineObject60) GetAUTHLDAP5GROUPTYPEOk() (*string, bool)`

GetAUTHLDAP5GROUPTYPEOk returns a tuple with the AUTHLDAP5GROUPTYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5GROUPTYPE

`func (o *InlineObject60) SetAUTHLDAP5GROUPTYPE(v string)`

SetAUTHLDAP5GROUPTYPE sets AUTHLDAP5GROUPTYPE field to given value.

### HasAUTHLDAP5GROUPTYPE

`func (o *InlineObject60) HasAUTHLDAP5GROUPTYPE() bool`

HasAUTHLDAP5GROUPTYPE returns a boolean if a field has been set.

### GetAUTHLDAP5GROUPTYPEPARAMS

`func (o *InlineObject60) GetAUTHLDAP5GROUPTYPEPARAMS() map[string]interface{}`

GetAUTHLDAP5GROUPTYPEPARAMS returns the AUTHLDAP5GROUPTYPEPARAMS field if non-nil, zero value otherwise.

### GetAUTHLDAP5GROUPTYPEPARAMSOk

`func (o *InlineObject60) GetAUTHLDAP5GROUPTYPEPARAMSOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5GROUPTYPEPARAMSOk returns a tuple with the AUTHLDAP5GROUPTYPEPARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5GROUPTYPEPARAMS

`func (o *InlineObject60) SetAUTHLDAP5GROUPTYPEPARAMS(v map[string]interface{})`

SetAUTHLDAP5GROUPTYPEPARAMS sets AUTHLDAP5GROUPTYPEPARAMS field to given value.

### HasAUTHLDAP5GROUPTYPEPARAMS

`func (o *InlineObject60) HasAUTHLDAP5GROUPTYPEPARAMS() bool`

HasAUTHLDAP5GROUPTYPEPARAMS returns a boolean if a field has been set.

### GetAUTHLDAP5ORGANIZATIONMAP

`func (o *InlineObject60) GetAUTHLDAP5ORGANIZATIONMAP() map[string]interface{}`

GetAUTHLDAP5ORGANIZATIONMAP returns the AUTHLDAP5ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP5ORGANIZATIONMAPOk

`func (o *InlineObject60) GetAUTHLDAP5ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5ORGANIZATIONMAPOk returns a tuple with the AUTHLDAP5ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5ORGANIZATIONMAP

`func (o *InlineObject60) SetAUTHLDAP5ORGANIZATIONMAP(v map[string]interface{})`

SetAUTHLDAP5ORGANIZATIONMAP sets AUTHLDAP5ORGANIZATIONMAP field to given value.

### HasAUTHLDAP5ORGANIZATIONMAP

`func (o *InlineObject60) HasAUTHLDAP5ORGANIZATIONMAP() bool`

HasAUTHLDAP5ORGANIZATIONMAP returns a boolean if a field has been set.

### GetAUTHLDAP5REQUIREGROUP

`func (o *InlineObject60) GetAUTHLDAP5REQUIREGROUP() string`

GetAUTHLDAP5REQUIREGROUP returns the AUTHLDAP5REQUIREGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP5REQUIREGROUPOk

`func (o *InlineObject60) GetAUTHLDAP5REQUIREGROUPOk() (*string, bool)`

GetAUTHLDAP5REQUIREGROUPOk returns a tuple with the AUTHLDAP5REQUIREGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5REQUIREGROUP

`func (o *InlineObject60) SetAUTHLDAP5REQUIREGROUP(v string)`

SetAUTHLDAP5REQUIREGROUP sets AUTHLDAP5REQUIREGROUP field to given value.

### HasAUTHLDAP5REQUIREGROUP

`func (o *InlineObject60) HasAUTHLDAP5REQUIREGROUP() bool`

HasAUTHLDAP5REQUIREGROUP returns a boolean if a field has been set.

### GetAUTHLDAP5SERVERURI

`func (o *InlineObject60) GetAUTHLDAP5SERVERURI() string`

GetAUTHLDAP5SERVERURI returns the AUTHLDAP5SERVERURI field if non-nil, zero value otherwise.

### GetAUTHLDAP5SERVERURIOk

`func (o *InlineObject60) GetAUTHLDAP5SERVERURIOk() (*string, bool)`

GetAUTHLDAP5SERVERURIOk returns a tuple with the AUTHLDAP5SERVERURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5SERVERURI

`func (o *InlineObject60) SetAUTHLDAP5SERVERURI(v string)`

SetAUTHLDAP5SERVERURI sets AUTHLDAP5SERVERURI field to given value.

### HasAUTHLDAP5SERVERURI

`func (o *InlineObject60) HasAUTHLDAP5SERVERURI() bool`

HasAUTHLDAP5SERVERURI returns a boolean if a field has been set.

### GetAUTHLDAP5STARTTLS

`func (o *InlineObject60) GetAUTHLDAP5STARTTLS() bool`

GetAUTHLDAP5STARTTLS returns the AUTHLDAP5STARTTLS field if non-nil, zero value otherwise.

### GetAUTHLDAP5STARTTLSOk

`func (o *InlineObject60) GetAUTHLDAP5STARTTLSOk() (*bool, bool)`

GetAUTHLDAP5STARTTLSOk returns a tuple with the AUTHLDAP5STARTTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5STARTTLS

`func (o *InlineObject60) SetAUTHLDAP5STARTTLS(v bool)`

SetAUTHLDAP5STARTTLS sets AUTHLDAP5STARTTLS field to given value.

### HasAUTHLDAP5STARTTLS

`func (o *InlineObject60) HasAUTHLDAP5STARTTLS() bool`

HasAUTHLDAP5STARTTLS returns a boolean if a field has been set.

### GetAUTHLDAP5TEAMMAP

`func (o *InlineObject60) GetAUTHLDAP5TEAMMAP() map[string]interface{}`

GetAUTHLDAP5TEAMMAP returns the AUTHLDAP5TEAMMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP5TEAMMAPOk

`func (o *InlineObject60) GetAUTHLDAP5TEAMMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5TEAMMAPOk returns a tuple with the AUTHLDAP5TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5TEAMMAP

`func (o *InlineObject60) SetAUTHLDAP5TEAMMAP(v map[string]interface{})`

SetAUTHLDAP5TEAMMAP sets AUTHLDAP5TEAMMAP field to given value.

### HasAUTHLDAP5TEAMMAP

`func (o *InlineObject60) HasAUTHLDAP5TEAMMAP() bool`

HasAUTHLDAP5TEAMMAP returns a boolean if a field has been set.

### GetAUTHLDAP5USERATTRMAP

`func (o *InlineObject60) GetAUTHLDAP5USERATTRMAP() map[string]interface{}`

GetAUTHLDAP5USERATTRMAP returns the AUTHLDAP5USERATTRMAP field if non-nil, zero value otherwise.

### GetAUTHLDAP5USERATTRMAPOk

`func (o *InlineObject60) GetAUTHLDAP5USERATTRMAPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5USERATTRMAPOk returns a tuple with the AUTHLDAP5USERATTRMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5USERATTRMAP

`func (o *InlineObject60) SetAUTHLDAP5USERATTRMAP(v map[string]interface{})`

SetAUTHLDAP5USERATTRMAP sets AUTHLDAP5USERATTRMAP field to given value.

### HasAUTHLDAP5USERATTRMAP

`func (o *InlineObject60) HasAUTHLDAP5USERATTRMAP() bool`

HasAUTHLDAP5USERATTRMAP returns a boolean if a field has been set.

### GetAUTHLDAP5USERDNTEMPLATE

`func (o *InlineObject60) GetAUTHLDAP5USERDNTEMPLATE() string`

GetAUTHLDAP5USERDNTEMPLATE returns the AUTHLDAP5USERDNTEMPLATE field if non-nil, zero value otherwise.

### GetAUTHLDAP5USERDNTEMPLATEOk

`func (o *InlineObject60) GetAUTHLDAP5USERDNTEMPLATEOk() (*string, bool)`

GetAUTHLDAP5USERDNTEMPLATEOk returns a tuple with the AUTHLDAP5USERDNTEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5USERDNTEMPLATE

`func (o *InlineObject60) SetAUTHLDAP5USERDNTEMPLATE(v string)`

SetAUTHLDAP5USERDNTEMPLATE sets AUTHLDAP5USERDNTEMPLATE field to given value.

### HasAUTHLDAP5USERDNTEMPLATE

`func (o *InlineObject60) HasAUTHLDAP5USERDNTEMPLATE() bool`

HasAUTHLDAP5USERDNTEMPLATE returns a boolean if a field has been set.

### GetAUTHLDAP5USERFLAGSBYGROUP

`func (o *InlineObject60) GetAUTHLDAP5USERFLAGSBYGROUP() map[string]interface{}`

GetAUTHLDAP5USERFLAGSBYGROUP returns the AUTHLDAP5USERFLAGSBYGROUP field if non-nil, zero value otherwise.

### GetAUTHLDAP5USERFLAGSBYGROUPOk

`func (o *InlineObject60) GetAUTHLDAP5USERFLAGSBYGROUPOk() (*map[string]interface{}, bool)`

GetAUTHLDAP5USERFLAGSBYGROUPOk returns a tuple with the AUTHLDAP5USERFLAGSBYGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5USERFLAGSBYGROUP

`func (o *InlineObject60) SetAUTHLDAP5USERFLAGSBYGROUP(v map[string]interface{})`

SetAUTHLDAP5USERFLAGSBYGROUP sets AUTHLDAP5USERFLAGSBYGROUP field to given value.

### HasAUTHLDAP5USERFLAGSBYGROUP

`func (o *InlineObject60) HasAUTHLDAP5USERFLAGSBYGROUP() bool`

HasAUTHLDAP5USERFLAGSBYGROUP returns a boolean if a field has been set.

### GetAUTHLDAP5USERSEARCH

`func (o *InlineObject60) GetAUTHLDAP5USERSEARCH() []string`

GetAUTHLDAP5USERSEARCH returns the AUTHLDAP5USERSEARCH field if non-nil, zero value otherwise.

### GetAUTHLDAP5USERSEARCHOk

`func (o *InlineObject60) GetAUTHLDAP5USERSEARCHOk() (*[]string, bool)`

GetAUTHLDAP5USERSEARCHOk returns a tuple with the AUTHLDAP5USERSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHLDAP5USERSEARCH

`func (o *InlineObject60) SetAUTHLDAP5USERSEARCH(v []string)`

SetAUTHLDAP5USERSEARCH sets AUTHLDAP5USERSEARCH field to given value.

### HasAUTHLDAP5USERSEARCH

`func (o *InlineObject60) HasAUTHLDAP5USERSEARCH() bool`

HasAUTHLDAP5USERSEARCH returns a boolean if a field has been set.

### GetAUTH_LDAP_BIND_DN

`func (o *InlineObject60) GetAUTH_LDAP_BIND_DN() string`

GetAUTH_LDAP_BIND_DN returns the AUTH_LDAP_BIND_DN field if non-nil, zero value otherwise.

### GetAUTH_LDAP_BIND_DNOk

`func (o *InlineObject60) GetAUTH_LDAP_BIND_DNOk() (*string, bool)`

GetAUTH_LDAP_BIND_DNOk returns a tuple with the AUTH_LDAP_BIND_DN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_BIND_DN

`func (o *InlineObject60) SetAUTH_LDAP_BIND_DN(v string)`

SetAUTH_LDAP_BIND_DN sets AUTH_LDAP_BIND_DN field to given value.

### HasAUTH_LDAP_BIND_DN

`func (o *InlineObject60) HasAUTH_LDAP_BIND_DN() bool`

HasAUTH_LDAP_BIND_DN returns a boolean if a field has been set.

### GetAUTH_LDAP_BIND_PASSWORD

`func (o *InlineObject60) GetAUTH_LDAP_BIND_PASSWORD() string`

GetAUTH_LDAP_BIND_PASSWORD returns the AUTH_LDAP_BIND_PASSWORD field if non-nil, zero value otherwise.

### GetAUTH_LDAP_BIND_PASSWORDOk

`func (o *InlineObject60) GetAUTH_LDAP_BIND_PASSWORDOk() (*string, bool)`

GetAUTH_LDAP_BIND_PASSWORDOk returns a tuple with the AUTH_LDAP_BIND_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_BIND_PASSWORD

`func (o *InlineObject60) SetAUTH_LDAP_BIND_PASSWORD(v string)`

SetAUTH_LDAP_BIND_PASSWORD sets AUTH_LDAP_BIND_PASSWORD field to given value.

### HasAUTH_LDAP_BIND_PASSWORD

`func (o *InlineObject60) HasAUTH_LDAP_BIND_PASSWORD() bool`

HasAUTH_LDAP_BIND_PASSWORD returns a boolean if a field has been set.

### GetAUTH_LDAP_CONNECTION_OPTIONS

`func (o *InlineObject60) GetAUTH_LDAP_CONNECTION_OPTIONS() map[string]interface{}`

GetAUTH_LDAP_CONNECTION_OPTIONS returns the AUTH_LDAP_CONNECTION_OPTIONS field if non-nil, zero value otherwise.

### GetAUTH_LDAP_CONNECTION_OPTIONSOk

`func (o *InlineObject60) GetAUTH_LDAP_CONNECTION_OPTIONSOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_CONNECTION_OPTIONSOk returns a tuple with the AUTH_LDAP_CONNECTION_OPTIONS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_CONNECTION_OPTIONS

`func (o *InlineObject60) SetAUTH_LDAP_CONNECTION_OPTIONS(v map[string]interface{})`

SetAUTH_LDAP_CONNECTION_OPTIONS sets AUTH_LDAP_CONNECTION_OPTIONS field to given value.

### HasAUTH_LDAP_CONNECTION_OPTIONS

`func (o *InlineObject60) HasAUTH_LDAP_CONNECTION_OPTIONS() bool`

HasAUTH_LDAP_CONNECTION_OPTIONS returns a boolean if a field has been set.

### GetAUTH_LDAP_DENY_GROUP

`func (o *InlineObject60) GetAUTH_LDAP_DENY_GROUP() string`

GetAUTH_LDAP_DENY_GROUP returns the AUTH_LDAP_DENY_GROUP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_DENY_GROUPOk

`func (o *InlineObject60) GetAUTH_LDAP_DENY_GROUPOk() (*string, bool)`

GetAUTH_LDAP_DENY_GROUPOk returns a tuple with the AUTH_LDAP_DENY_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_DENY_GROUP

`func (o *InlineObject60) SetAUTH_LDAP_DENY_GROUP(v string)`

SetAUTH_LDAP_DENY_GROUP sets AUTH_LDAP_DENY_GROUP field to given value.

### HasAUTH_LDAP_DENY_GROUP

`func (o *InlineObject60) HasAUTH_LDAP_DENY_GROUP() bool`

HasAUTH_LDAP_DENY_GROUP returns a boolean if a field has been set.

### GetAUTH_LDAP_GROUP_SEARCH

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_SEARCH() []string`

GetAUTH_LDAP_GROUP_SEARCH returns the AUTH_LDAP_GROUP_SEARCH field if non-nil, zero value otherwise.

### GetAUTH_LDAP_GROUP_SEARCHOk

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_SEARCHOk() (*[]string, bool)`

GetAUTH_LDAP_GROUP_SEARCHOk returns a tuple with the AUTH_LDAP_GROUP_SEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_GROUP_SEARCH

`func (o *InlineObject60) SetAUTH_LDAP_GROUP_SEARCH(v []string)`

SetAUTH_LDAP_GROUP_SEARCH sets AUTH_LDAP_GROUP_SEARCH field to given value.

### HasAUTH_LDAP_GROUP_SEARCH

`func (o *InlineObject60) HasAUTH_LDAP_GROUP_SEARCH() bool`

HasAUTH_LDAP_GROUP_SEARCH returns a boolean if a field has been set.

### GetAUTH_LDAP_GROUP_TYPE

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_TYPE() string`

GetAUTH_LDAP_GROUP_TYPE returns the AUTH_LDAP_GROUP_TYPE field if non-nil, zero value otherwise.

### GetAUTH_LDAP_GROUP_TYPEOk

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_TYPEOk() (*string, bool)`

GetAUTH_LDAP_GROUP_TYPEOk returns a tuple with the AUTH_LDAP_GROUP_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_GROUP_TYPE

`func (o *InlineObject60) SetAUTH_LDAP_GROUP_TYPE(v string)`

SetAUTH_LDAP_GROUP_TYPE sets AUTH_LDAP_GROUP_TYPE field to given value.

### HasAUTH_LDAP_GROUP_TYPE

`func (o *InlineObject60) HasAUTH_LDAP_GROUP_TYPE() bool`

HasAUTH_LDAP_GROUP_TYPE returns a boolean if a field has been set.

### GetAUTH_LDAP_GROUP_TYPE_PARAMS

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_TYPE_PARAMS() map[string]interface{}`

GetAUTH_LDAP_GROUP_TYPE_PARAMS returns the AUTH_LDAP_GROUP_TYPE_PARAMS field if non-nil, zero value otherwise.

### GetAUTH_LDAP_GROUP_TYPE_PARAMSOk

`func (o *InlineObject60) GetAUTH_LDAP_GROUP_TYPE_PARAMSOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_GROUP_TYPE_PARAMSOk returns a tuple with the AUTH_LDAP_GROUP_TYPE_PARAMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_GROUP_TYPE_PARAMS

`func (o *InlineObject60) SetAUTH_LDAP_GROUP_TYPE_PARAMS(v map[string]interface{})`

SetAUTH_LDAP_GROUP_TYPE_PARAMS sets AUTH_LDAP_GROUP_TYPE_PARAMS field to given value.

### HasAUTH_LDAP_GROUP_TYPE_PARAMS

`func (o *InlineObject60) HasAUTH_LDAP_GROUP_TYPE_PARAMS() bool`

HasAUTH_LDAP_GROUP_TYPE_PARAMS returns a boolean if a field has been set.

### GetAUTH_LDAP_ORGANIZATION_MAP

`func (o *InlineObject60) GetAUTH_LDAP_ORGANIZATION_MAP() map[string]interface{}`

GetAUTH_LDAP_ORGANIZATION_MAP returns the AUTH_LDAP_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetAUTH_LDAP_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_ORGANIZATION_MAPOk returns a tuple with the AUTH_LDAP_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_ORGANIZATION_MAP

`func (o *InlineObject60) SetAUTH_LDAP_ORGANIZATION_MAP(v map[string]interface{})`

SetAUTH_LDAP_ORGANIZATION_MAP sets AUTH_LDAP_ORGANIZATION_MAP field to given value.

### HasAUTH_LDAP_ORGANIZATION_MAP

`func (o *InlineObject60) HasAUTH_LDAP_ORGANIZATION_MAP() bool`

HasAUTH_LDAP_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetAUTH_LDAP_REQUIRE_GROUP

`func (o *InlineObject60) GetAUTH_LDAP_REQUIRE_GROUP() string`

GetAUTH_LDAP_REQUIRE_GROUP returns the AUTH_LDAP_REQUIRE_GROUP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_REQUIRE_GROUPOk

`func (o *InlineObject60) GetAUTH_LDAP_REQUIRE_GROUPOk() (*string, bool)`

GetAUTH_LDAP_REQUIRE_GROUPOk returns a tuple with the AUTH_LDAP_REQUIRE_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_REQUIRE_GROUP

`func (o *InlineObject60) SetAUTH_LDAP_REQUIRE_GROUP(v string)`

SetAUTH_LDAP_REQUIRE_GROUP sets AUTH_LDAP_REQUIRE_GROUP field to given value.

### HasAUTH_LDAP_REQUIRE_GROUP

`func (o *InlineObject60) HasAUTH_LDAP_REQUIRE_GROUP() bool`

HasAUTH_LDAP_REQUIRE_GROUP returns a boolean if a field has been set.

### GetAUTH_LDAP_SERVER_URI

`func (o *InlineObject60) GetAUTH_LDAP_SERVER_URI() string`

GetAUTH_LDAP_SERVER_URI returns the AUTH_LDAP_SERVER_URI field if non-nil, zero value otherwise.

### GetAUTH_LDAP_SERVER_URIOk

`func (o *InlineObject60) GetAUTH_LDAP_SERVER_URIOk() (*string, bool)`

GetAUTH_LDAP_SERVER_URIOk returns a tuple with the AUTH_LDAP_SERVER_URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_SERVER_URI

`func (o *InlineObject60) SetAUTH_LDAP_SERVER_URI(v string)`

SetAUTH_LDAP_SERVER_URI sets AUTH_LDAP_SERVER_URI field to given value.

### HasAUTH_LDAP_SERVER_URI

`func (o *InlineObject60) HasAUTH_LDAP_SERVER_URI() bool`

HasAUTH_LDAP_SERVER_URI returns a boolean if a field has been set.

### GetAUTH_LDAP_START_TLS

`func (o *InlineObject60) GetAUTH_LDAP_START_TLS() bool`

GetAUTH_LDAP_START_TLS returns the AUTH_LDAP_START_TLS field if non-nil, zero value otherwise.

### GetAUTH_LDAP_START_TLSOk

`func (o *InlineObject60) GetAUTH_LDAP_START_TLSOk() (*bool, bool)`

GetAUTH_LDAP_START_TLSOk returns a tuple with the AUTH_LDAP_START_TLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_START_TLS

`func (o *InlineObject60) SetAUTH_LDAP_START_TLS(v bool)`

SetAUTH_LDAP_START_TLS sets AUTH_LDAP_START_TLS field to given value.

### HasAUTH_LDAP_START_TLS

`func (o *InlineObject60) HasAUTH_LDAP_START_TLS() bool`

HasAUTH_LDAP_START_TLS returns a boolean if a field has been set.

### GetAUTH_LDAP_TEAM_MAP

`func (o *InlineObject60) GetAUTH_LDAP_TEAM_MAP() map[string]interface{}`

GetAUTH_LDAP_TEAM_MAP returns the AUTH_LDAP_TEAM_MAP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_TEAM_MAPOk

`func (o *InlineObject60) GetAUTH_LDAP_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_TEAM_MAPOk returns a tuple with the AUTH_LDAP_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_TEAM_MAP

`func (o *InlineObject60) SetAUTH_LDAP_TEAM_MAP(v map[string]interface{})`

SetAUTH_LDAP_TEAM_MAP sets AUTH_LDAP_TEAM_MAP field to given value.

### HasAUTH_LDAP_TEAM_MAP

`func (o *InlineObject60) HasAUTH_LDAP_TEAM_MAP() bool`

HasAUTH_LDAP_TEAM_MAP returns a boolean if a field has been set.

### GetAUTH_LDAP_USER_ATTR_MAP

`func (o *InlineObject60) GetAUTH_LDAP_USER_ATTR_MAP() map[string]interface{}`

GetAUTH_LDAP_USER_ATTR_MAP returns the AUTH_LDAP_USER_ATTR_MAP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_USER_ATTR_MAPOk

`func (o *InlineObject60) GetAUTH_LDAP_USER_ATTR_MAPOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_USER_ATTR_MAPOk returns a tuple with the AUTH_LDAP_USER_ATTR_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_USER_ATTR_MAP

`func (o *InlineObject60) SetAUTH_LDAP_USER_ATTR_MAP(v map[string]interface{})`

SetAUTH_LDAP_USER_ATTR_MAP sets AUTH_LDAP_USER_ATTR_MAP field to given value.

### HasAUTH_LDAP_USER_ATTR_MAP

`func (o *InlineObject60) HasAUTH_LDAP_USER_ATTR_MAP() bool`

HasAUTH_LDAP_USER_ATTR_MAP returns a boolean if a field has been set.

### GetAUTH_LDAP_USER_DN_TEMPLATE

`func (o *InlineObject60) GetAUTH_LDAP_USER_DN_TEMPLATE() string`

GetAUTH_LDAP_USER_DN_TEMPLATE returns the AUTH_LDAP_USER_DN_TEMPLATE field if non-nil, zero value otherwise.

### GetAUTH_LDAP_USER_DN_TEMPLATEOk

`func (o *InlineObject60) GetAUTH_LDAP_USER_DN_TEMPLATEOk() (*string, bool)`

GetAUTH_LDAP_USER_DN_TEMPLATEOk returns a tuple with the AUTH_LDAP_USER_DN_TEMPLATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_USER_DN_TEMPLATE

`func (o *InlineObject60) SetAUTH_LDAP_USER_DN_TEMPLATE(v string)`

SetAUTH_LDAP_USER_DN_TEMPLATE sets AUTH_LDAP_USER_DN_TEMPLATE field to given value.

### HasAUTH_LDAP_USER_DN_TEMPLATE

`func (o *InlineObject60) HasAUTH_LDAP_USER_DN_TEMPLATE() bool`

HasAUTH_LDAP_USER_DN_TEMPLATE returns a boolean if a field has been set.

### GetAUTH_LDAP_USER_FLAGS_BY_GROUP

`func (o *InlineObject60) GetAUTH_LDAP_USER_FLAGS_BY_GROUP() map[string]interface{}`

GetAUTH_LDAP_USER_FLAGS_BY_GROUP returns the AUTH_LDAP_USER_FLAGS_BY_GROUP field if non-nil, zero value otherwise.

### GetAUTH_LDAP_USER_FLAGS_BY_GROUPOk

`func (o *InlineObject60) GetAUTH_LDAP_USER_FLAGS_BY_GROUPOk() (*map[string]interface{}, bool)`

GetAUTH_LDAP_USER_FLAGS_BY_GROUPOk returns a tuple with the AUTH_LDAP_USER_FLAGS_BY_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_USER_FLAGS_BY_GROUP

`func (o *InlineObject60) SetAUTH_LDAP_USER_FLAGS_BY_GROUP(v map[string]interface{})`

SetAUTH_LDAP_USER_FLAGS_BY_GROUP sets AUTH_LDAP_USER_FLAGS_BY_GROUP field to given value.

### HasAUTH_LDAP_USER_FLAGS_BY_GROUP

`func (o *InlineObject60) HasAUTH_LDAP_USER_FLAGS_BY_GROUP() bool`

HasAUTH_LDAP_USER_FLAGS_BY_GROUP returns a boolean if a field has been set.

### GetAUTH_LDAP_USER_SEARCH

`func (o *InlineObject60) GetAUTH_LDAP_USER_SEARCH() []string`

GetAUTH_LDAP_USER_SEARCH returns the AUTH_LDAP_USER_SEARCH field if non-nil, zero value otherwise.

### GetAUTH_LDAP_USER_SEARCHOk

`func (o *InlineObject60) GetAUTH_LDAP_USER_SEARCHOk() (*[]string, bool)`

GetAUTH_LDAP_USER_SEARCHOk returns a tuple with the AUTH_LDAP_USER_SEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_LDAP_USER_SEARCH

`func (o *InlineObject60) SetAUTH_LDAP_USER_SEARCH(v []string)`

SetAUTH_LDAP_USER_SEARCH sets AUTH_LDAP_USER_SEARCH field to given value.

### HasAUTH_LDAP_USER_SEARCH

`func (o *InlineObject60) HasAUTH_LDAP_USER_SEARCH() bool`

HasAUTH_LDAP_USER_SEARCH returns a boolean if a field has been set.

### GetAUTOMATION_ANALYTICS_GATHER_INTERVAL

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_GATHER_INTERVAL() int32`

GetAUTOMATION_ANALYTICS_GATHER_INTERVAL returns the AUTOMATION_ANALYTICS_GATHER_INTERVAL field if non-nil, zero value otherwise.

### GetAUTOMATION_ANALYTICS_GATHER_INTERVALOk

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_GATHER_INTERVALOk() (*int32, bool)`

GetAUTOMATION_ANALYTICS_GATHER_INTERVALOk returns a tuple with the AUTOMATION_ANALYTICS_GATHER_INTERVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATION_ANALYTICS_GATHER_INTERVAL

`func (o *InlineObject60) SetAUTOMATION_ANALYTICS_GATHER_INTERVAL(v int32)`

SetAUTOMATION_ANALYTICS_GATHER_INTERVAL sets AUTOMATION_ANALYTICS_GATHER_INTERVAL field to given value.

### HasAUTOMATION_ANALYTICS_GATHER_INTERVAL

`func (o *InlineObject60) HasAUTOMATION_ANALYTICS_GATHER_INTERVAL() bool`

HasAUTOMATION_ANALYTICS_GATHER_INTERVAL returns a boolean if a field has been set.

### GetAUTOMATION_ANALYTICS_LAST_GATHER

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_LAST_GATHER() string`

GetAUTOMATION_ANALYTICS_LAST_GATHER returns the AUTOMATION_ANALYTICS_LAST_GATHER field if non-nil, zero value otherwise.

### GetAUTOMATION_ANALYTICS_LAST_GATHEROk

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_LAST_GATHEROk() (*string, bool)`

GetAUTOMATION_ANALYTICS_LAST_GATHEROk returns a tuple with the AUTOMATION_ANALYTICS_LAST_GATHER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATION_ANALYTICS_LAST_GATHER

`func (o *InlineObject60) SetAUTOMATION_ANALYTICS_LAST_GATHER(v string)`

SetAUTOMATION_ANALYTICS_LAST_GATHER sets AUTOMATION_ANALYTICS_LAST_GATHER field to given value.


### GetAUTOMATION_ANALYTICS_URL

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_URL() string`

GetAUTOMATION_ANALYTICS_URL returns the AUTOMATION_ANALYTICS_URL field if non-nil, zero value otherwise.

### GetAUTOMATION_ANALYTICS_URLOk

`func (o *InlineObject60) GetAUTOMATION_ANALYTICS_URLOk() (*string, bool)`

GetAUTOMATION_ANALYTICS_URLOk returns a tuple with the AUTOMATION_ANALYTICS_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTOMATION_ANALYTICS_URL

`func (o *InlineObject60) SetAUTOMATION_ANALYTICS_URL(v string)`

SetAUTOMATION_ANALYTICS_URL sets AUTOMATION_ANALYTICS_URL field to given value.

### HasAUTOMATION_ANALYTICS_URL

`func (o *InlineObject60) HasAUTOMATION_ANALYTICS_URL() bool`

HasAUTOMATION_ANALYTICS_URL returns a boolean if a field has been set.

### GetAWX_ANSIBLE_CALLBACK_PLUGINS

`func (o *InlineObject60) GetAWX_ANSIBLE_CALLBACK_PLUGINS() []string`

GetAWX_ANSIBLE_CALLBACK_PLUGINS returns the AWX_ANSIBLE_CALLBACK_PLUGINS field if non-nil, zero value otherwise.

### GetAWX_ANSIBLE_CALLBACK_PLUGINSOk

`func (o *InlineObject60) GetAWX_ANSIBLE_CALLBACK_PLUGINSOk() (*[]string, bool)`

GetAWX_ANSIBLE_CALLBACK_PLUGINSOk returns a tuple with the AWX_ANSIBLE_CALLBACK_PLUGINS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ANSIBLE_CALLBACK_PLUGINS

`func (o *InlineObject60) SetAWX_ANSIBLE_CALLBACK_PLUGINS(v []string)`

SetAWX_ANSIBLE_CALLBACK_PLUGINS sets AWX_ANSIBLE_CALLBACK_PLUGINS field to given value.

### HasAWX_ANSIBLE_CALLBACK_PLUGINS

`func (o *InlineObject60) HasAWX_ANSIBLE_CALLBACK_PLUGINS() bool`

HasAWX_ANSIBLE_CALLBACK_PLUGINS returns a boolean if a field has been set.

### GetAWX_COLLECTIONS_ENABLED

`func (o *InlineObject60) GetAWX_COLLECTIONS_ENABLED() bool`

GetAWX_COLLECTIONS_ENABLED returns the AWX_COLLECTIONS_ENABLED field if non-nil, zero value otherwise.

### GetAWX_COLLECTIONS_ENABLEDOk

`func (o *InlineObject60) GetAWX_COLLECTIONS_ENABLEDOk() (*bool, bool)`

GetAWX_COLLECTIONS_ENABLEDOk returns a tuple with the AWX_COLLECTIONS_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_COLLECTIONS_ENABLED

`func (o *InlineObject60) SetAWX_COLLECTIONS_ENABLED(v bool)`

SetAWX_COLLECTIONS_ENABLED sets AWX_COLLECTIONS_ENABLED field to given value.

### HasAWX_COLLECTIONS_ENABLED

`func (o *InlineObject60) HasAWX_COLLECTIONS_ENABLED() bool`

HasAWX_COLLECTIONS_ENABLED returns a boolean if a field has been set.

### GetAWX_ISOLATED_CHECK_INTERVAL

`func (o *InlineObject60) GetAWX_ISOLATED_CHECK_INTERVAL() int32`

GetAWX_ISOLATED_CHECK_INTERVAL returns the AWX_ISOLATED_CHECK_INTERVAL field if non-nil, zero value otherwise.

### GetAWX_ISOLATED_CHECK_INTERVALOk

`func (o *InlineObject60) GetAWX_ISOLATED_CHECK_INTERVALOk() (*int32, bool)`

GetAWX_ISOLATED_CHECK_INTERVALOk returns a tuple with the AWX_ISOLATED_CHECK_INTERVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ISOLATED_CHECK_INTERVAL

`func (o *InlineObject60) SetAWX_ISOLATED_CHECK_INTERVAL(v int32)`

SetAWX_ISOLATED_CHECK_INTERVAL sets AWX_ISOLATED_CHECK_INTERVAL field to given value.


### GetAWX_ISOLATED_CONNECTION_TIMEOUT

`func (o *InlineObject60) GetAWX_ISOLATED_CONNECTION_TIMEOUT() int32`

GetAWX_ISOLATED_CONNECTION_TIMEOUT returns the AWX_ISOLATED_CONNECTION_TIMEOUT field if non-nil, zero value otherwise.

### GetAWX_ISOLATED_CONNECTION_TIMEOUTOk

`func (o *InlineObject60) GetAWX_ISOLATED_CONNECTION_TIMEOUTOk() (*int32, bool)`

GetAWX_ISOLATED_CONNECTION_TIMEOUTOk returns a tuple with the AWX_ISOLATED_CONNECTION_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ISOLATED_CONNECTION_TIMEOUT

`func (o *InlineObject60) SetAWX_ISOLATED_CONNECTION_TIMEOUT(v int32)`

SetAWX_ISOLATED_CONNECTION_TIMEOUT sets AWX_ISOLATED_CONNECTION_TIMEOUT field to given value.

### HasAWX_ISOLATED_CONNECTION_TIMEOUT

`func (o *InlineObject60) HasAWX_ISOLATED_CONNECTION_TIMEOUT() bool`

HasAWX_ISOLATED_CONNECTION_TIMEOUT returns a boolean if a field has been set.

### GetAWX_ISOLATED_HOST_KEY_CHECKING

`func (o *InlineObject60) GetAWX_ISOLATED_HOST_KEY_CHECKING() bool`

GetAWX_ISOLATED_HOST_KEY_CHECKING returns the AWX_ISOLATED_HOST_KEY_CHECKING field if non-nil, zero value otherwise.

### GetAWX_ISOLATED_HOST_KEY_CHECKINGOk

`func (o *InlineObject60) GetAWX_ISOLATED_HOST_KEY_CHECKINGOk() (*bool, bool)`

GetAWX_ISOLATED_HOST_KEY_CHECKINGOk returns a tuple with the AWX_ISOLATED_HOST_KEY_CHECKING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ISOLATED_HOST_KEY_CHECKING

`func (o *InlineObject60) SetAWX_ISOLATED_HOST_KEY_CHECKING(v bool)`

SetAWX_ISOLATED_HOST_KEY_CHECKING sets AWX_ISOLATED_HOST_KEY_CHECKING field to given value.

### HasAWX_ISOLATED_HOST_KEY_CHECKING

`func (o *InlineObject60) HasAWX_ISOLATED_HOST_KEY_CHECKING() bool`

HasAWX_ISOLATED_HOST_KEY_CHECKING returns a boolean if a field has been set.

### GetAWX_ISOLATED_LAUNCH_TIMEOUT

`func (o *InlineObject60) GetAWX_ISOLATED_LAUNCH_TIMEOUT() int32`

GetAWX_ISOLATED_LAUNCH_TIMEOUT returns the AWX_ISOLATED_LAUNCH_TIMEOUT field if non-nil, zero value otherwise.

### GetAWX_ISOLATED_LAUNCH_TIMEOUTOk

`func (o *InlineObject60) GetAWX_ISOLATED_LAUNCH_TIMEOUTOk() (*int32, bool)`

GetAWX_ISOLATED_LAUNCH_TIMEOUTOk returns a tuple with the AWX_ISOLATED_LAUNCH_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ISOLATED_LAUNCH_TIMEOUT

`func (o *InlineObject60) SetAWX_ISOLATED_LAUNCH_TIMEOUT(v int32)`

SetAWX_ISOLATED_LAUNCH_TIMEOUT sets AWX_ISOLATED_LAUNCH_TIMEOUT field to given value.


### GetAWX_PROOT_BASE_PATH

`func (o *InlineObject60) GetAWX_PROOT_BASE_PATH() string`

GetAWX_PROOT_BASE_PATH returns the AWX_PROOT_BASE_PATH field if non-nil, zero value otherwise.

### GetAWX_PROOT_BASE_PATHOk

`func (o *InlineObject60) GetAWX_PROOT_BASE_PATHOk() (*string, bool)`

GetAWX_PROOT_BASE_PATHOk returns a tuple with the AWX_PROOT_BASE_PATH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_PROOT_BASE_PATH

`func (o *InlineObject60) SetAWX_PROOT_BASE_PATH(v string)`

SetAWX_PROOT_BASE_PATH sets AWX_PROOT_BASE_PATH field to given value.


### GetAWX_PROOT_ENABLED

`func (o *InlineObject60) GetAWX_PROOT_ENABLED() bool`

GetAWX_PROOT_ENABLED returns the AWX_PROOT_ENABLED field if non-nil, zero value otherwise.

### GetAWX_PROOT_ENABLEDOk

`func (o *InlineObject60) GetAWX_PROOT_ENABLEDOk() (*bool, bool)`

GetAWX_PROOT_ENABLEDOk returns a tuple with the AWX_PROOT_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_PROOT_ENABLED

`func (o *InlineObject60) SetAWX_PROOT_ENABLED(v bool)`

SetAWX_PROOT_ENABLED sets AWX_PROOT_ENABLED field to given value.


### GetAWX_PROOT_HIDE_PATHS

`func (o *InlineObject60) GetAWX_PROOT_HIDE_PATHS() []string`

GetAWX_PROOT_HIDE_PATHS returns the AWX_PROOT_HIDE_PATHS field if non-nil, zero value otherwise.

### GetAWX_PROOT_HIDE_PATHSOk

`func (o *InlineObject60) GetAWX_PROOT_HIDE_PATHSOk() (*[]string, bool)`

GetAWX_PROOT_HIDE_PATHSOk returns a tuple with the AWX_PROOT_HIDE_PATHS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_PROOT_HIDE_PATHS

`func (o *InlineObject60) SetAWX_PROOT_HIDE_PATHS(v []string)`

SetAWX_PROOT_HIDE_PATHS sets AWX_PROOT_HIDE_PATHS field to given value.

### HasAWX_PROOT_HIDE_PATHS

`func (o *InlineObject60) HasAWX_PROOT_HIDE_PATHS() bool`

HasAWX_PROOT_HIDE_PATHS returns a boolean if a field has been set.

### GetAWX_PROOT_SHOW_PATHS

`func (o *InlineObject60) GetAWX_PROOT_SHOW_PATHS() []string`

GetAWX_PROOT_SHOW_PATHS returns the AWX_PROOT_SHOW_PATHS field if non-nil, zero value otherwise.

### GetAWX_PROOT_SHOW_PATHSOk

`func (o *InlineObject60) GetAWX_PROOT_SHOW_PATHSOk() (*[]string, bool)`

GetAWX_PROOT_SHOW_PATHSOk returns a tuple with the AWX_PROOT_SHOW_PATHS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_PROOT_SHOW_PATHS

`func (o *InlineObject60) SetAWX_PROOT_SHOW_PATHS(v []string)`

SetAWX_PROOT_SHOW_PATHS sets AWX_PROOT_SHOW_PATHS field to given value.

### HasAWX_PROOT_SHOW_PATHS

`func (o *InlineObject60) HasAWX_PROOT_SHOW_PATHS() bool`

HasAWX_PROOT_SHOW_PATHS returns a boolean if a field has been set.

### GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL() float32`

GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL returns the AWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL field if non-nil, zero value otherwise.

### GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVALOk

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVALOk() (*float32, bool)`

GetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVALOk returns a tuple with the AWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL

`func (o *InlineObject60) SetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL(v float32)`

SetAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL sets AWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL field to given value.

### HasAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL

`func (o *InlineObject60) HasAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL() bool`

HasAWX_RESOURCE_PROFILING_CPU_POLL_INTERVAL returns a boolean if a field has been set.

### GetAWX_RESOURCE_PROFILING_ENABLED

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_ENABLED() bool`

GetAWX_RESOURCE_PROFILING_ENABLED returns the AWX_RESOURCE_PROFILING_ENABLED field if non-nil, zero value otherwise.

### GetAWX_RESOURCE_PROFILING_ENABLEDOk

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_ENABLEDOk() (*bool, bool)`

GetAWX_RESOURCE_PROFILING_ENABLEDOk returns a tuple with the AWX_RESOURCE_PROFILING_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_RESOURCE_PROFILING_ENABLED

`func (o *InlineObject60) SetAWX_RESOURCE_PROFILING_ENABLED(v bool)`

SetAWX_RESOURCE_PROFILING_ENABLED sets AWX_RESOURCE_PROFILING_ENABLED field to given value.

### HasAWX_RESOURCE_PROFILING_ENABLED

`func (o *InlineObject60) HasAWX_RESOURCE_PROFILING_ENABLED() bool`

HasAWX_RESOURCE_PROFILING_ENABLED returns a boolean if a field has been set.

### GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL() float32`

GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL returns the AWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL field if non-nil, zero value otherwise.

### GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVALOk

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVALOk() (*float32, bool)`

GetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVALOk returns a tuple with the AWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL

`func (o *InlineObject60) SetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL(v float32)`

SetAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL sets AWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL field to given value.

### HasAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL

`func (o *InlineObject60) HasAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL() bool`

HasAWX_RESOURCE_PROFILING_MEMORY_POLL_INTERVAL returns a boolean if a field has been set.

### GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL() float32`

GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL returns the AWX_RESOURCE_PROFILING_PID_POLL_INTERVAL field if non-nil, zero value otherwise.

### GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVALOk

`func (o *InlineObject60) GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVALOk() (*float32, bool)`

GetAWX_RESOURCE_PROFILING_PID_POLL_INTERVALOk returns a tuple with the AWX_RESOURCE_PROFILING_PID_POLL_INTERVAL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL

`func (o *InlineObject60) SetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL(v float32)`

SetAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL sets AWX_RESOURCE_PROFILING_PID_POLL_INTERVAL field to given value.

### HasAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL

`func (o *InlineObject60) HasAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL() bool`

HasAWX_RESOURCE_PROFILING_PID_POLL_INTERVAL returns a boolean if a field has been set.

### GetAWX_ROLES_ENABLED

`func (o *InlineObject60) GetAWX_ROLES_ENABLED() bool`

GetAWX_ROLES_ENABLED returns the AWX_ROLES_ENABLED field if non-nil, zero value otherwise.

### GetAWX_ROLES_ENABLEDOk

`func (o *InlineObject60) GetAWX_ROLES_ENABLEDOk() (*bool, bool)`

GetAWX_ROLES_ENABLEDOk returns a tuple with the AWX_ROLES_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_ROLES_ENABLED

`func (o *InlineObject60) SetAWX_ROLES_ENABLED(v bool)`

SetAWX_ROLES_ENABLED sets AWX_ROLES_ENABLED field to given value.

### HasAWX_ROLES_ENABLED

`func (o *InlineObject60) HasAWX_ROLES_ENABLED() bool`

HasAWX_ROLES_ENABLED returns a boolean if a field has been set.

### GetAWX_SHOW_PLAYBOOK_LINKS

`func (o *InlineObject60) GetAWX_SHOW_PLAYBOOK_LINKS() bool`

GetAWX_SHOW_PLAYBOOK_LINKS returns the AWX_SHOW_PLAYBOOK_LINKS field if non-nil, zero value otherwise.

### GetAWX_SHOW_PLAYBOOK_LINKSOk

`func (o *InlineObject60) GetAWX_SHOW_PLAYBOOK_LINKSOk() (*bool, bool)`

GetAWX_SHOW_PLAYBOOK_LINKSOk returns a tuple with the AWX_SHOW_PLAYBOOK_LINKS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_SHOW_PLAYBOOK_LINKS

`func (o *InlineObject60) SetAWX_SHOW_PLAYBOOK_LINKS(v bool)`

SetAWX_SHOW_PLAYBOOK_LINKS sets AWX_SHOW_PLAYBOOK_LINKS field to given value.

### HasAWX_SHOW_PLAYBOOK_LINKS

`func (o *InlineObject60) HasAWX_SHOW_PLAYBOOK_LINKS() bool`

HasAWX_SHOW_PLAYBOOK_LINKS returns a boolean if a field has been set.

### GetAWX_TASK_ENV

`func (o *InlineObject60) GetAWX_TASK_ENV() map[string]interface{}`

GetAWX_TASK_ENV returns the AWX_TASK_ENV field if non-nil, zero value otherwise.

### GetAWX_TASK_ENVOk

`func (o *InlineObject60) GetAWX_TASK_ENVOk() (*map[string]interface{}, bool)`

GetAWX_TASK_ENVOk returns a tuple with the AWX_TASK_ENV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAWX_TASK_ENV

`func (o *InlineObject60) SetAWX_TASK_ENV(v map[string]interface{})`

SetAWX_TASK_ENV sets AWX_TASK_ENV field to given value.

### HasAWX_TASK_ENV

`func (o *InlineObject60) HasAWX_TASK_ENV() bool`

HasAWX_TASK_ENV returns a boolean if a field has been set.

### GetCUSTOM_LOGIN_INFO

`func (o *InlineObject60) GetCUSTOM_LOGIN_INFO() string`

GetCUSTOM_LOGIN_INFO returns the CUSTOM_LOGIN_INFO field if non-nil, zero value otherwise.

### GetCUSTOM_LOGIN_INFOOk

`func (o *InlineObject60) GetCUSTOM_LOGIN_INFOOk() (*string, bool)`

GetCUSTOM_LOGIN_INFOOk returns a tuple with the CUSTOM_LOGIN_INFO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_LOGIN_INFO

`func (o *InlineObject60) SetCUSTOM_LOGIN_INFO(v string)`

SetCUSTOM_LOGIN_INFO sets CUSTOM_LOGIN_INFO field to given value.

### HasCUSTOM_LOGIN_INFO

`func (o *InlineObject60) HasCUSTOM_LOGIN_INFO() bool`

HasCUSTOM_LOGIN_INFO returns a boolean if a field has been set.

### GetCUSTOM_LOGO

`func (o *InlineObject60) GetCUSTOM_LOGO() string`

GetCUSTOM_LOGO returns the CUSTOM_LOGO field if non-nil, zero value otherwise.

### GetCUSTOM_LOGOOk

`func (o *InlineObject60) GetCUSTOM_LOGOOk() (*string, bool)`

GetCUSTOM_LOGOOk returns a tuple with the CUSTOM_LOGO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_LOGO

`func (o *InlineObject60) SetCUSTOM_LOGO(v string)`

SetCUSTOM_LOGO sets CUSTOM_LOGO field to given value.

### HasCUSTOM_LOGO

`func (o *InlineObject60) HasCUSTOM_LOGO() bool`

HasCUSTOM_LOGO returns a boolean if a field has been set.

### GetCUSTOM_VENV_PATHS

`func (o *InlineObject60) GetCUSTOM_VENV_PATHS() []string`

GetCUSTOM_VENV_PATHS returns the CUSTOM_VENV_PATHS field if non-nil, zero value otherwise.

### GetCUSTOM_VENV_PATHSOk

`func (o *InlineObject60) GetCUSTOM_VENV_PATHSOk() (*[]string, bool)`

GetCUSTOM_VENV_PATHSOk returns a tuple with the CUSTOM_VENV_PATHS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_VENV_PATHS

`func (o *InlineObject60) SetCUSTOM_VENV_PATHS(v []string)`

SetCUSTOM_VENV_PATHS sets CUSTOM_VENV_PATHS field to given value.

### HasCUSTOM_VENV_PATHS

`func (o *InlineObject60) HasCUSTOM_VENV_PATHS() bool`

HasCUSTOM_VENV_PATHS returns a boolean if a field has been set.

### GetDEFAULT_INVENTORY_UPDATE_TIMEOUT

`func (o *InlineObject60) GetDEFAULT_INVENTORY_UPDATE_TIMEOUT() int32`

GetDEFAULT_INVENTORY_UPDATE_TIMEOUT returns the DEFAULT_INVENTORY_UPDATE_TIMEOUT field if non-nil, zero value otherwise.

### GetDEFAULT_INVENTORY_UPDATE_TIMEOUTOk

`func (o *InlineObject60) GetDEFAULT_INVENTORY_UPDATE_TIMEOUTOk() (*int32, bool)`

GetDEFAULT_INVENTORY_UPDATE_TIMEOUTOk returns a tuple with the DEFAULT_INVENTORY_UPDATE_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_INVENTORY_UPDATE_TIMEOUT

`func (o *InlineObject60) SetDEFAULT_INVENTORY_UPDATE_TIMEOUT(v int32)`

SetDEFAULT_INVENTORY_UPDATE_TIMEOUT sets DEFAULT_INVENTORY_UPDATE_TIMEOUT field to given value.

### HasDEFAULT_INVENTORY_UPDATE_TIMEOUT

`func (o *InlineObject60) HasDEFAULT_INVENTORY_UPDATE_TIMEOUT() bool`

HasDEFAULT_INVENTORY_UPDATE_TIMEOUT returns a boolean if a field has been set.

### GetDEFAULT_JOB_TIMEOUT

`func (o *InlineObject60) GetDEFAULT_JOB_TIMEOUT() int32`

GetDEFAULT_JOB_TIMEOUT returns the DEFAULT_JOB_TIMEOUT field if non-nil, zero value otherwise.

### GetDEFAULT_JOB_TIMEOUTOk

`func (o *InlineObject60) GetDEFAULT_JOB_TIMEOUTOk() (*int32, bool)`

GetDEFAULT_JOB_TIMEOUTOk returns a tuple with the DEFAULT_JOB_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_JOB_TIMEOUT

`func (o *InlineObject60) SetDEFAULT_JOB_TIMEOUT(v int32)`

SetDEFAULT_JOB_TIMEOUT sets DEFAULT_JOB_TIMEOUT field to given value.

### HasDEFAULT_JOB_TIMEOUT

`func (o *InlineObject60) HasDEFAULT_JOB_TIMEOUT() bool`

HasDEFAULT_JOB_TIMEOUT returns a boolean if a field has been set.

### GetDEFAULT_PROJECT_UPDATE_TIMEOUT

`func (o *InlineObject60) GetDEFAULT_PROJECT_UPDATE_TIMEOUT() int32`

GetDEFAULT_PROJECT_UPDATE_TIMEOUT returns the DEFAULT_PROJECT_UPDATE_TIMEOUT field if non-nil, zero value otherwise.

### GetDEFAULT_PROJECT_UPDATE_TIMEOUTOk

`func (o *InlineObject60) GetDEFAULT_PROJECT_UPDATE_TIMEOUTOk() (*int32, bool)`

GetDEFAULT_PROJECT_UPDATE_TIMEOUTOk returns a tuple with the DEFAULT_PROJECT_UPDATE_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_PROJECT_UPDATE_TIMEOUT

`func (o *InlineObject60) SetDEFAULT_PROJECT_UPDATE_TIMEOUT(v int32)`

SetDEFAULT_PROJECT_UPDATE_TIMEOUT sets DEFAULT_PROJECT_UPDATE_TIMEOUT field to given value.

### HasDEFAULT_PROJECT_UPDATE_TIMEOUT

`func (o *InlineObject60) HasDEFAULT_PROJECT_UPDATE_TIMEOUT() bool`

HasDEFAULT_PROJECT_UPDATE_TIMEOUT returns a boolean if a field has been set.

### GetEVENT_STDOUT_MAX_BYTES_DISPLAY

`func (o *InlineObject60) GetEVENT_STDOUT_MAX_BYTES_DISPLAY() int32`

GetEVENT_STDOUT_MAX_BYTES_DISPLAY returns the EVENT_STDOUT_MAX_BYTES_DISPLAY field if non-nil, zero value otherwise.

### GetEVENT_STDOUT_MAX_BYTES_DISPLAYOk

`func (o *InlineObject60) GetEVENT_STDOUT_MAX_BYTES_DISPLAYOk() (*int32, bool)`

GetEVENT_STDOUT_MAX_BYTES_DISPLAYOk returns a tuple with the EVENT_STDOUT_MAX_BYTES_DISPLAY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEVENT_STDOUT_MAX_BYTES_DISPLAY

`func (o *InlineObject60) SetEVENT_STDOUT_MAX_BYTES_DISPLAY(v int32)`

SetEVENT_STDOUT_MAX_BYTES_DISPLAY sets EVENT_STDOUT_MAX_BYTES_DISPLAY field to given value.


### GetGALAXY_IGNORE_CERTS

`func (o *InlineObject60) GetGALAXY_IGNORE_CERTS() bool`

GetGALAXY_IGNORE_CERTS returns the GALAXY_IGNORE_CERTS field if non-nil, zero value otherwise.

### GetGALAXY_IGNORE_CERTSOk

`func (o *InlineObject60) GetGALAXY_IGNORE_CERTSOk() (*bool, bool)`

GetGALAXY_IGNORE_CERTSOk returns a tuple with the GALAXY_IGNORE_CERTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGALAXY_IGNORE_CERTS

`func (o *InlineObject60) SetGALAXY_IGNORE_CERTS(v bool)`

SetGALAXY_IGNORE_CERTS sets GALAXY_IGNORE_CERTS field to given value.

### HasGALAXY_IGNORE_CERTS

`func (o *InlineObject60) HasGALAXY_IGNORE_CERTS() bool`

HasGALAXY_IGNORE_CERTS returns a boolean if a field has been set.

### GetINSIGHTS_TRACKING_STATE

`func (o *InlineObject60) GetINSIGHTS_TRACKING_STATE() bool`

GetINSIGHTS_TRACKING_STATE returns the INSIGHTS_TRACKING_STATE field if non-nil, zero value otherwise.

### GetINSIGHTS_TRACKING_STATEOk

`func (o *InlineObject60) GetINSIGHTS_TRACKING_STATEOk() (*bool, bool)`

GetINSIGHTS_TRACKING_STATEOk returns a tuple with the INSIGHTS_TRACKING_STATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINSIGHTS_TRACKING_STATE

`func (o *InlineObject60) SetINSIGHTS_TRACKING_STATE(v bool)`

SetINSIGHTS_TRACKING_STATE sets INSIGHTS_TRACKING_STATE field to given value.

### HasINSIGHTS_TRACKING_STATE

`func (o *InlineObject60) HasINSIGHTS_TRACKING_STATE() bool`

HasINSIGHTS_TRACKING_STATE returns a boolean if a field has been set.

### GetLOGIN_REDIRECT_OVERRIDE

`func (o *InlineObject60) GetLOGIN_REDIRECT_OVERRIDE() string`

GetLOGIN_REDIRECT_OVERRIDE returns the LOGIN_REDIRECT_OVERRIDE field if non-nil, zero value otherwise.

### GetLOGIN_REDIRECT_OVERRIDEOk

`func (o *InlineObject60) GetLOGIN_REDIRECT_OVERRIDEOk() (*string, bool)`

GetLOGIN_REDIRECT_OVERRIDEOk returns a tuple with the LOGIN_REDIRECT_OVERRIDE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGIN_REDIRECT_OVERRIDE

`func (o *InlineObject60) SetLOGIN_REDIRECT_OVERRIDE(v string)`

SetLOGIN_REDIRECT_OVERRIDE sets LOGIN_REDIRECT_OVERRIDE field to given value.

### HasLOGIN_REDIRECT_OVERRIDE

`func (o *InlineObject60) HasLOGIN_REDIRECT_OVERRIDE() bool`

HasLOGIN_REDIRECT_OVERRIDE returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_ENABLED

`func (o *InlineObject60) GetLOG_AGGREGATOR_ENABLED() bool`

GetLOG_AGGREGATOR_ENABLED returns the LOG_AGGREGATOR_ENABLED field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_ENABLEDOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_ENABLEDOk() (*bool, bool)`

GetLOG_AGGREGATOR_ENABLEDOk returns a tuple with the LOG_AGGREGATOR_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_ENABLED

`func (o *InlineObject60) SetLOG_AGGREGATOR_ENABLED(v bool)`

SetLOG_AGGREGATOR_ENABLED sets LOG_AGGREGATOR_ENABLED field to given value.

### HasLOG_AGGREGATOR_ENABLED

`func (o *InlineObject60) HasLOG_AGGREGATOR_ENABLED() bool`

HasLOG_AGGREGATOR_ENABLED returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_HOST

`func (o *InlineObject60) GetLOG_AGGREGATOR_HOST() string`

GetLOG_AGGREGATOR_HOST returns the LOG_AGGREGATOR_HOST field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_HOSTOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_HOSTOk() (*string, bool)`

GetLOG_AGGREGATOR_HOSTOk returns a tuple with the LOG_AGGREGATOR_HOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_HOST

`func (o *InlineObject60) SetLOG_AGGREGATOR_HOST(v string)`

SetLOG_AGGREGATOR_HOST sets LOG_AGGREGATOR_HOST field to given value.

### HasLOG_AGGREGATOR_HOST

`func (o *InlineObject60) HasLOG_AGGREGATOR_HOST() bool`

HasLOG_AGGREGATOR_HOST returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_INDIVIDUAL_FACTS

`func (o *InlineObject60) GetLOG_AGGREGATOR_INDIVIDUAL_FACTS() bool`

GetLOG_AGGREGATOR_INDIVIDUAL_FACTS returns the LOG_AGGREGATOR_INDIVIDUAL_FACTS field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_INDIVIDUAL_FACTSOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_INDIVIDUAL_FACTSOk() (*bool, bool)`

GetLOG_AGGREGATOR_INDIVIDUAL_FACTSOk returns a tuple with the LOG_AGGREGATOR_INDIVIDUAL_FACTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_INDIVIDUAL_FACTS

`func (o *InlineObject60) SetLOG_AGGREGATOR_INDIVIDUAL_FACTS(v bool)`

SetLOG_AGGREGATOR_INDIVIDUAL_FACTS sets LOG_AGGREGATOR_INDIVIDUAL_FACTS field to given value.

### HasLOG_AGGREGATOR_INDIVIDUAL_FACTS

`func (o *InlineObject60) HasLOG_AGGREGATOR_INDIVIDUAL_FACTS() bool`

HasLOG_AGGREGATOR_INDIVIDUAL_FACTS returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_LEVEL

`func (o *InlineObject60) GetLOG_AGGREGATOR_LEVEL() string`

GetLOG_AGGREGATOR_LEVEL returns the LOG_AGGREGATOR_LEVEL field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_LEVELOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_LEVELOk() (*string, bool)`

GetLOG_AGGREGATOR_LEVELOk returns a tuple with the LOG_AGGREGATOR_LEVEL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_LEVEL

`func (o *InlineObject60) SetLOG_AGGREGATOR_LEVEL(v string)`

SetLOG_AGGREGATOR_LEVEL sets LOG_AGGREGATOR_LEVEL field to given value.

### HasLOG_AGGREGATOR_LEVEL

`func (o *InlineObject60) HasLOG_AGGREGATOR_LEVEL() bool`

HasLOG_AGGREGATOR_LEVEL returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_LOGGERS

`func (o *InlineObject60) GetLOG_AGGREGATOR_LOGGERS() []string`

GetLOG_AGGREGATOR_LOGGERS returns the LOG_AGGREGATOR_LOGGERS field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_LOGGERSOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_LOGGERSOk() (*[]string, bool)`

GetLOG_AGGREGATOR_LOGGERSOk returns a tuple with the LOG_AGGREGATOR_LOGGERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_LOGGERS

`func (o *InlineObject60) SetLOG_AGGREGATOR_LOGGERS(v []string)`

SetLOG_AGGREGATOR_LOGGERS sets LOG_AGGREGATOR_LOGGERS field to given value.

### HasLOG_AGGREGATOR_LOGGERS

`func (o *InlineObject60) HasLOG_AGGREGATOR_LOGGERS() bool`

HasLOG_AGGREGATOR_LOGGERS returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_MAX_DISK_USAGE_GB

`func (o *InlineObject60) GetLOG_AGGREGATOR_MAX_DISK_USAGE_GB() int32`

GetLOG_AGGREGATOR_MAX_DISK_USAGE_GB returns the LOG_AGGREGATOR_MAX_DISK_USAGE_GB field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_MAX_DISK_USAGE_GBOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_MAX_DISK_USAGE_GBOk() (*int32, bool)`

GetLOG_AGGREGATOR_MAX_DISK_USAGE_GBOk returns a tuple with the LOG_AGGREGATOR_MAX_DISK_USAGE_GB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_MAX_DISK_USAGE_GB

`func (o *InlineObject60) SetLOG_AGGREGATOR_MAX_DISK_USAGE_GB(v int32)`

SetLOG_AGGREGATOR_MAX_DISK_USAGE_GB sets LOG_AGGREGATOR_MAX_DISK_USAGE_GB field to given value.

### HasLOG_AGGREGATOR_MAX_DISK_USAGE_GB

`func (o *InlineObject60) HasLOG_AGGREGATOR_MAX_DISK_USAGE_GB() bool`

HasLOG_AGGREGATOR_MAX_DISK_USAGE_GB returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH

`func (o *InlineObject60) GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH() string`

GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH returns the LOG_AGGREGATOR_MAX_DISK_USAGE_PATH field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATHOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATHOk() (*string, bool)`

GetLOG_AGGREGATOR_MAX_DISK_USAGE_PATHOk returns a tuple with the LOG_AGGREGATOR_MAX_DISK_USAGE_PATH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH

`func (o *InlineObject60) SetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH(v string)`

SetLOG_AGGREGATOR_MAX_DISK_USAGE_PATH sets LOG_AGGREGATOR_MAX_DISK_USAGE_PATH field to given value.

### HasLOG_AGGREGATOR_MAX_DISK_USAGE_PATH

`func (o *InlineObject60) HasLOG_AGGREGATOR_MAX_DISK_USAGE_PATH() bool`

HasLOG_AGGREGATOR_MAX_DISK_USAGE_PATH returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_PASSWORD

`func (o *InlineObject60) GetLOG_AGGREGATOR_PASSWORD() string`

GetLOG_AGGREGATOR_PASSWORD returns the LOG_AGGREGATOR_PASSWORD field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_PASSWORDOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_PASSWORDOk() (*string, bool)`

GetLOG_AGGREGATOR_PASSWORDOk returns a tuple with the LOG_AGGREGATOR_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_PASSWORD

`func (o *InlineObject60) SetLOG_AGGREGATOR_PASSWORD(v string)`

SetLOG_AGGREGATOR_PASSWORD sets LOG_AGGREGATOR_PASSWORD field to given value.

### HasLOG_AGGREGATOR_PASSWORD

`func (o *InlineObject60) HasLOG_AGGREGATOR_PASSWORD() bool`

HasLOG_AGGREGATOR_PASSWORD returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_PORT

`func (o *InlineObject60) GetLOG_AGGREGATOR_PORT() int32`

GetLOG_AGGREGATOR_PORT returns the LOG_AGGREGATOR_PORT field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_PORTOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_PORTOk() (*int32, bool)`

GetLOG_AGGREGATOR_PORTOk returns a tuple with the LOG_AGGREGATOR_PORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_PORT

`func (o *InlineObject60) SetLOG_AGGREGATOR_PORT(v int32)`

SetLOG_AGGREGATOR_PORT sets LOG_AGGREGATOR_PORT field to given value.

### HasLOG_AGGREGATOR_PORT

`func (o *InlineObject60) HasLOG_AGGREGATOR_PORT() bool`

HasLOG_AGGREGATOR_PORT returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_PROTOCOL

`func (o *InlineObject60) GetLOG_AGGREGATOR_PROTOCOL() string`

GetLOG_AGGREGATOR_PROTOCOL returns the LOG_AGGREGATOR_PROTOCOL field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_PROTOCOLOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_PROTOCOLOk() (*string, bool)`

GetLOG_AGGREGATOR_PROTOCOLOk returns a tuple with the LOG_AGGREGATOR_PROTOCOL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_PROTOCOL

`func (o *InlineObject60) SetLOG_AGGREGATOR_PROTOCOL(v string)`

SetLOG_AGGREGATOR_PROTOCOL sets LOG_AGGREGATOR_PROTOCOL field to given value.

### HasLOG_AGGREGATOR_PROTOCOL

`func (o *InlineObject60) HasLOG_AGGREGATOR_PROTOCOL() bool`

HasLOG_AGGREGATOR_PROTOCOL returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_RSYSLOGD_DEBUG

`func (o *InlineObject60) GetLOG_AGGREGATOR_RSYSLOGD_DEBUG() bool`

GetLOG_AGGREGATOR_RSYSLOGD_DEBUG returns the LOG_AGGREGATOR_RSYSLOGD_DEBUG field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_RSYSLOGD_DEBUGOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_RSYSLOGD_DEBUGOk() (*bool, bool)`

GetLOG_AGGREGATOR_RSYSLOGD_DEBUGOk returns a tuple with the LOG_AGGREGATOR_RSYSLOGD_DEBUG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_RSYSLOGD_DEBUG

`func (o *InlineObject60) SetLOG_AGGREGATOR_RSYSLOGD_DEBUG(v bool)`

SetLOG_AGGREGATOR_RSYSLOGD_DEBUG sets LOG_AGGREGATOR_RSYSLOGD_DEBUG field to given value.

### HasLOG_AGGREGATOR_RSYSLOGD_DEBUG

`func (o *InlineObject60) HasLOG_AGGREGATOR_RSYSLOGD_DEBUG() bool`

HasLOG_AGGREGATOR_RSYSLOGD_DEBUG returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_TCP_TIMEOUT

`func (o *InlineObject60) GetLOG_AGGREGATOR_TCP_TIMEOUT() int32`

GetLOG_AGGREGATOR_TCP_TIMEOUT returns the LOG_AGGREGATOR_TCP_TIMEOUT field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_TCP_TIMEOUTOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_TCP_TIMEOUTOk() (*int32, bool)`

GetLOG_AGGREGATOR_TCP_TIMEOUTOk returns a tuple with the LOG_AGGREGATOR_TCP_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_TCP_TIMEOUT

`func (o *InlineObject60) SetLOG_AGGREGATOR_TCP_TIMEOUT(v int32)`

SetLOG_AGGREGATOR_TCP_TIMEOUT sets LOG_AGGREGATOR_TCP_TIMEOUT field to given value.

### HasLOG_AGGREGATOR_TCP_TIMEOUT

`func (o *InlineObject60) HasLOG_AGGREGATOR_TCP_TIMEOUT() bool`

HasLOG_AGGREGATOR_TCP_TIMEOUT returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_TOWER_UUID

`func (o *InlineObject60) GetLOG_AGGREGATOR_TOWER_UUID() string`

GetLOG_AGGREGATOR_TOWER_UUID returns the LOG_AGGREGATOR_TOWER_UUID field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_TOWER_UUIDOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_TOWER_UUIDOk() (*string, bool)`

GetLOG_AGGREGATOR_TOWER_UUIDOk returns a tuple with the LOG_AGGREGATOR_TOWER_UUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_TOWER_UUID

`func (o *InlineObject60) SetLOG_AGGREGATOR_TOWER_UUID(v string)`

SetLOG_AGGREGATOR_TOWER_UUID sets LOG_AGGREGATOR_TOWER_UUID field to given value.

### HasLOG_AGGREGATOR_TOWER_UUID

`func (o *InlineObject60) HasLOG_AGGREGATOR_TOWER_UUID() bool`

HasLOG_AGGREGATOR_TOWER_UUID returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_TYPE

`func (o *InlineObject60) GetLOG_AGGREGATOR_TYPE() string`

GetLOG_AGGREGATOR_TYPE returns the LOG_AGGREGATOR_TYPE field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_TYPEOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_TYPEOk() (*string, bool)`

GetLOG_AGGREGATOR_TYPEOk returns a tuple with the LOG_AGGREGATOR_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_TYPE

`func (o *InlineObject60) SetLOG_AGGREGATOR_TYPE(v string)`

SetLOG_AGGREGATOR_TYPE sets LOG_AGGREGATOR_TYPE field to given value.

### HasLOG_AGGREGATOR_TYPE

`func (o *InlineObject60) HasLOG_AGGREGATOR_TYPE() bool`

HasLOG_AGGREGATOR_TYPE returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_USERNAME

`func (o *InlineObject60) GetLOG_AGGREGATOR_USERNAME() string`

GetLOG_AGGREGATOR_USERNAME returns the LOG_AGGREGATOR_USERNAME field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_USERNAMEOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_USERNAMEOk() (*string, bool)`

GetLOG_AGGREGATOR_USERNAMEOk returns a tuple with the LOG_AGGREGATOR_USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_USERNAME

`func (o *InlineObject60) SetLOG_AGGREGATOR_USERNAME(v string)`

SetLOG_AGGREGATOR_USERNAME sets LOG_AGGREGATOR_USERNAME field to given value.

### HasLOG_AGGREGATOR_USERNAME

`func (o *InlineObject60) HasLOG_AGGREGATOR_USERNAME() bool`

HasLOG_AGGREGATOR_USERNAME returns a boolean if a field has been set.

### GetLOG_AGGREGATOR_VERIFY_CERT

`func (o *InlineObject60) GetLOG_AGGREGATOR_VERIFY_CERT() bool`

GetLOG_AGGREGATOR_VERIFY_CERT returns the LOG_AGGREGATOR_VERIFY_CERT field if non-nil, zero value otherwise.

### GetLOG_AGGREGATOR_VERIFY_CERTOk

`func (o *InlineObject60) GetLOG_AGGREGATOR_VERIFY_CERTOk() (*bool, bool)`

GetLOG_AGGREGATOR_VERIFY_CERTOk returns a tuple with the LOG_AGGREGATOR_VERIFY_CERT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOG_AGGREGATOR_VERIFY_CERT

`func (o *InlineObject60) SetLOG_AGGREGATOR_VERIFY_CERT(v bool)`

SetLOG_AGGREGATOR_VERIFY_CERT sets LOG_AGGREGATOR_VERIFY_CERT field to given value.

### HasLOG_AGGREGATOR_VERIFY_CERT

`func (o *InlineObject60) HasLOG_AGGREGATOR_VERIFY_CERT() bool`

HasLOG_AGGREGATOR_VERIFY_CERT returns a boolean if a field has been set.

### GetMANAGE_ORGANIZATION_AUTH

`func (o *InlineObject60) GetMANAGE_ORGANIZATION_AUTH() bool`

GetMANAGE_ORGANIZATION_AUTH returns the MANAGE_ORGANIZATION_AUTH field if non-nil, zero value otherwise.

### GetMANAGE_ORGANIZATION_AUTHOk

`func (o *InlineObject60) GetMANAGE_ORGANIZATION_AUTHOk() (*bool, bool)`

GetMANAGE_ORGANIZATION_AUTHOk returns a tuple with the MANAGE_ORGANIZATION_AUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMANAGE_ORGANIZATION_AUTH

`func (o *InlineObject60) SetMANAGE_ORGANIZATION_AUTH(v bool)`

SetMANAGE_ORGANIZATION_AUTH sets MANAGE_ORGANIZATION_AUTH field to given value.


### GetMAX_FORKS

`func (o *InlineObject60) GetMAX_FORKS() int32`

GetMAX_FORKS returns the MAX_FORKS field if non-nil, zero value otherwise.

### GetMAX_FORKSOk

`func (o *InlineObject60) GetMAX_FORKSOk() (*int32, bool)`

GetMAX_FORKSOk returns a tuple with the MAX_FORKS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_FORKS

`func (o *InlineObject60) SetMAX_FORKS(v int32)`

SetMAX_FORKS sets MAX_FORKS field to given value.

### HasMAX_FORKS

`func (o *InlineObject60) HasMAX_FORKS() bool`

HasMAX_FORKS returns a boolean if a field has been set.

### GetMAX_UI_JOB_EVENTS

`func (o *InlineObject60) GetMAX_UI_JOB_EVENTS() int32`

GetMAX_UI_JOB_EVENTS returns the MAX_UI_JOB_EVENTS field if non-nil, zero value otherwise.

### GetMAX_UI_JOB_EVENTSOk

`func (o *InlineObject60) GetMAX_UI_JOB_EVENTSOk() (*int32, bool)`

GetMAX_UI_JOB_EVENTSOk returns a tuple with the MAX_UI_JOB_EVENTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_UI_JOB_EVENTS

`func (o *InlineObject60) SetMAX_UI_JOB_EVENTS(v int32)`

SetMAX_UI_JOB_EVENTS sets MAX_UI_JOB_EVENTS field to given value.


### GetOAUTH2PROVIDER

`func (o *InlineObject60) GetOAUTH2PROVIDER() map[string]interface{}`

GetOAUTH2PROVIDER returns the OAUTH2PROVIDER field if non-nil, zero value otherwise.

### GetOAUTH2PROVIDEROk

`func (o *InlineObject60) GetOAUTH2PROVIDEROk() (*map[string]interface{}, bool)`

GetOAUTH2PROVIDEROk returns a tuple with the OAUTH2PROVIDER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH2PROVIDER

`func (o *InlineObject60) SetOAUTH2PROVIDER(v map[string]interface{})`

SetOAUTH2PROVIDER sets OAUTH2PROVIDER field to given value.

### HasOAUTH2PROVIDER

`func (o *InlineObject60) HasOAUTH2PROVIDER() bool`

HasOAUTH2PROVIDER returns a boolean if a field has been set.

### GetORG_ADMINS_CAN_SEE_ALL_USERS

`func (o *InlineObject60) GetORG_ADMINS_CAN_SEE_ALL_USERS() bool`

GetORG_ADMINS_CAN_SEE_ALL_USERS returns the ORG_ADMINS_CAN_SEE_ALL_USERS field if non-nil, zero value otherwise.

### GetORG_ADMINS_CAN_SEE_ALL_USERSOk

`func (o *InlineObject60) GetORG_ADMINS_CAN_SEE_ALL_USERSOk() (*bool, bool)`

GetORG_ADMINS_CAN_SEE_ALL_USERSOk returns a tuple with the ORG_ADMINS_CAN_SEE_ALL_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORG_ADMINS_CAN_SEE_ALL_USERS

`func (o *InlineObject60) SetORG_ADMINS_CAN_SEE_ALL_USERS(v bool)`

SetORG_ADMINS_CAN_SEE_ALL_USERS sets ORG_ADMINS_CAN_SEE_ALL_USERS field to given value.


### GetPROJECT_UPDATE_VVV

`func (o *InlineObject60) GetPROJECT_UPDATE_VVV() bool`

GetPROJECT_UPDATE_VVV returns the PROJECT_UPDATE_VVV field if non-nil, zero value otherwise.

### GetPROJECT_UPDATE_VVVOk

`func (o *InlineObject60) GetPROJECT_UPDATE_VVVOk() (*bool, bool)`

GetPROJECT_UPDATE_VVVOk returns a tuple with the PROJECT_UPDATE_VVV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROJECT_UPDATE_VVV

`func (o *InlineObject60) SetPROJECT_UPDATE_VVV(v bool)`

SetPROJECT_UPDATE_VVV sets PROJECT_UPDATE_VVV field to given value.


### GetPROXY_IP_ALLOWED_LIST

`func (o *InlineObject60) GetPROXY_IP_ALLOWED_LIST() []string`

GetPROXY_IP_ALLOWED_LIST returns the PROXY_IP_ALLOWED_LIST field if non-nil, zero value otherwise.

### GetPROXY_IP_ALLOWED_LISTOk

`func (o *InlineObject60) GetPROXY_IP_ALLOWED_LISTOk() (*[]string, bool)`

GetPROXY_IP_ALLOWED_LISTOk returns a tuple with the PROXY_IP_ALLOWED_LIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROXY_IP_ALLOWED_LIST

`func (o *InlineObject60) SetPROXY_IP_ALLOWED_LIST(v []string)`

SetPROXY_IP_ALLOWED_LIST sets PROXY_IP_ALLOWED_LIST field to given value.


### GetRADIUS_PORT

`func (o *InlineObject60) GetRADIUS_PORT() int32`

GetRADIUS_PORT returns the RADIUS_PORT field if non-nil, zero value otherwise.

### GetRADIUS_PORTOk

`func (o *InlineObject60) GetRADIUS_PORTOk() (*int32, bool)`

GetRADIUS_PORTOk returns a tuple with the RADIUS_PORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRADIUS_PORT

`func (o *InlineObject60) SetRADIUS_PORT(v int32)`

SetRADIUS_PORT sets RADIUS_PORT field to given value.

### HasRADIUS_PORT

`func (o *InlineObject60) HasRADIUS_PORT() bool`

HasRADIUS_PORT returns a boolean if a field has been set.

### GetRADIUS_SECRET

`func (o *InlineObject60) GetRADIUS_SECRET() string`

GetRADIUS_SECRET returns the RADIUS_SECRET field if non-nil, zero value otherwise.

### GetRADIUS_SECRETOk

`func (o *InlineObject60) GetRADIUS_SECRETOk() (*string, bool)`

GetRADIUS_SECRETOk returns a tuple with the RADIUS_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRADIUS_SECRET

`func (o *InlineObject60) SetRADIUS_SECRET(v string)`

SetRADIUS_SECRET sets RADIUS_SECRET field to given value.

### HasRADIUS_SECRET

`func (o *InlineObject60) HasRADIUS_SECRET() bool`

HasRADIUS_SECRET returns a boolean if a field has been set.

### GetRADIUS_SERVER

`func (o *InlineObject60) GetRADIUS_SERVER() string`

GetRADIUS_SERVER returns the RADIUS_SERVER field if non-nil, zero value otherwise.

### GetRADIUS_SERVEROk

`func (o *InlineObject60) GetRADIUS_SERVEROk() (*string, bool)`

GetRADIUS_SERVEROk returns a tuple with the RADIUS_SERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRADIUS_SERVER

`func (o *InlineObject60) SetRADIUS_SERVER(v string)`

SetRADIUS_SERVER sets RADIUS_SERVER field to given value.

### HasRADIUS_SERVER

`func (o *InlineObject60) HasRADIUS_SERVER() bool`

HasRADIUS_SERVER returns a boolean if a field has been set.

### GetREDHAT_PASSWORD

`func (o *InlineObject60) GetREDHAT_PASSWORD() string`

GetREDHAT_PASSWORD returns the REDHAT_PASSWORD field if non-nil, zero value otherwise.

### GetREDHAT_PASSWORDOk

`func (o *InlineObject60) GetREDHAT_PASSWORDOk() (*string, bool)`

GetREDHAT_PASSWORDOk returns a tuple with the REDHAT_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREDHAT_PASSWORD

`func (o *InlineObject60) SetREDHAT_PASSWORD(v string)`

SetREDHAT_PASSWORD sets REDHAT_PASSWORD field to given value.

### HasREDHAT_PASSWORD

`func (o *InlineObject60) HasREDHAT_PASSWORD() bool`

HasREDHAT_PASSWORD returns a boolean if a field has been set.

### GetREDHAT_USERNAME

`func (o *InlineObject60) GetREDHAT_USERNAME() string`

GetREDHAT_USERNAME returns the REDHAT_USERNAME field if non-nil, zero value otherwise.

### GetREDHAT_USERNAMEOk

`func (o *InlineObject60) GetREDHAT_USERNAMEOk() (*string, bool)`

GetREDHAT_USERNAMEOk returns a tuple with the REDHAT_USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREDHAT_USERNAME

`func (o *InlineObject60) SetREDHAT_USERNAME(v string)`

SetREDHAT_USERNAME sets REDHAT_USERNAME field to given value.

### HasREDHAT_USERNAME

`func (o *InlineObject60) HasREDHAT_USERNAME() bool`

HasREDHAT_USERNAME returns a boolean if a field has been set.

### GetREMOTE_HOST_HEADERS

`func (o *InlineObject60) GetREMOTE_HOST_HEADERS() []string`

GetREMOTE_HOST_HEADERS returns the REMOTE_HOST_HEADERS field if non-nil, zero value otherwise.

### GetREMOTE_HOST_HEADERSOk

`func (o *InlineObject60) GetREMOTE_HOST_HEADERSOk() (*[]string, bool)`

GetREMOTE_HOST_HEADERSOk returns a tuple with the REMOTE_HOST_HEADERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOTE_HOST_HEADERS

`func (o *InlineObject60) SetREMOTE_HOST_HEADERS(v []string)`

SetREMOTE_HOST_HEADERS sets REMOTE_HOST_HEADERS field to given value.


### GetSAML_AUTO_CREATE_OBJECTS

`func (o *InlineObject60) GetSAML_AUTO_CREATE_OBJECTS() bool`

GetSAML_AUTO_CREATE_OBJECTS returns the SAML_AUTO_CREATE_OBJECTS field if non-nil, zero value otherwise.

### GetSAML_AUTO_CREATE_OBJECTSOk

`func (o *InlineObject60) GetSAML_AUTO_CREATE_OBJECTSOk() (*bool, bool)`

GetSAML_AUTO_CREATE_OBJECTSOk returns a tuple with the SAML_AUTO_CREATE_OBJECTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAML_AUTO_CREATE_OBJECTS

`func (o *InlineObject60) SetSAML_AUTO_CREATE_OBJECTS(v bool)`

SetSAML_AUTO_CREATE_OBJECTS sets SAML_AUTO_CREATE_OBJECTS field to given value.

### HasSAML_AUTO_CREATE_OBJECTS

`func (o *InlineObject60) HasSAML_AUTO_CREATE_OBJECTS() bool`

HasSAML_AUTO_CREATE_OBJECTS returns a boolean if a field has been set.

### GetSCHEDULE_MAX_JOBS

`func (o *InlineObject60) GetSCHEDULE_MAX_JOBS() int32`

GetSCHEDULE_MAX_JOBS returns the SCHEDULE_MAX_JOBS field if non-nil, zero value otherwise.

### GetSCHEDULE_MAX_JOBSOk

`func (o *InlineObject60) GetSCHEDULE_MAX_JOBSOk() (*int32, bool)`

GetSCHEDULE_MAX_JOBSOk returns a tuple with the SCHEDULE_MAX_JOBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCHEDULE_MAX_JOBS

`func (o *InlineObject60) SetSCHEDULE_MAX_JOBS(v int32)`

SetSCHEDULE_MAX_JOBS sets SCHEDULE_MAX_JOBS field to given value.


### GetSESSIONS_PER_USER

`func (o *InlineObject60) GetSESSIONS_PER_USER() int32`

GetSESSIONS_PER_USER returns the SESSIONS_PER_USER field if non-nil, zero value otherwise.

### GetSESSIONS_PER_USEROk

`func (o *InlineObject60) GetSESSIONS_PER_USEROk() (*int32, bool)`

GetSESSIONS_PER_USEROk returns a tuple with the SESSIONS_PER_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSESSIONS_PER_USER

`func (o *InlineObject60) SetSESSIONS_PER_USER(v int32)`

SetSESSIONS_PER_USER sets SESSIONS_PER_USER field to given value.


### GetSESSION_COOKIE_AGE

`func (o *InlineObject60) GetSESSION_COOKIE_AGE() int32`

GetSESSION_COOKIE_AGE returns the SESSION_COOKIE_AGE field if non-nil, zero value otherwise.

### GetSESSION_COOKIE_AGEOk

`func (o *InlineObject60) GetSESSION_COOKIE_AGEOk() (*int32, bool)`

GetSESSION_COOKIE_AGEOk returns a tuple with the SESSION_COOKIE_AGE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSESSION_COOKIE_AGE

`func (o *InlineObject60) SetSESSION_COOKIE_AGE(v int32)`

SetSESSION_COOKIE_AGE sets SESSION_COOKIE_AGE field to given value.


### GetSOCIALAUTHAZUREADOAUTH2KEY

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2KEY() string`

GetSOCIALAUTHAZUREADOAUTH2KEY returns the SOCIALAUTHAZUREADOAUTH2KEY field if non-nil, zero value otherwise.

### GetSOCIALAUTHAZUREADOAUTH2KEYOk

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2KEYOk() (*string, bool)`

GetSOCIALAUTHAZUREADOAUTH2KEYOk returns a tuple with the SOCIALAUTHAZUREADOAUTH2KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHAZUREADOAUTH2KEY

`func (o *InlineObject60) SetSOCIALAUTHAZUREADOAUTH2KEY(v string)`

SetSOCIALAUTHAZUREADOAUTH2KEY sets SOCIALAUTHAZUREADOAUTH2KEY field to given value.

### HasSOCIALAUTHAZUREADOAUTH2KEY

`func (o *InlineObject60) HasSOCIALAUTHAZUREADOAUTH2KEY() bool`

HasSOCIALAUTHAZUREADOAUTH2KEY returns a boolean if a field has been set.

### GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP() map[string]interface{}`

GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP returns the SOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAPOk

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAPOk returns a tuple with the SOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) SetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP(v map[string]interface{})`

SetSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP sets SOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP field to given value.

### HasSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) HasSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP() bool`

HasSOCIALAUTHAZUREADOAUTH2ORGANIZATIONMAP returns a boolean if a field has been set.

### GetSOCIALAUTHAZUREADOAUTH2SECRET

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2SECRET() string`

GetSOCIALAUTHAZUREADOAUTH2SECRET returns the SOCIALAUTHAZUREADOAUTH2SECRET field if non-nil, zero value otherwise.

### GetSOCIALAUTHAZUREADOAUTH2SECRETOk

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2SECRETOk() (*string, bool)`

GetSOCIALAUTHAZUREADOAUTH2SECRETOk returns a tuple with the SOCIALAUTHAZUREADOAUTH2SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHAZUREADOAUTH2SECRET

`func (o *InlineObject60) SetSOCIALAUTHAZUREADOAUTH2SECRET(v string)`

SetSOCIALAUTHAZUREADOAUTH2SECRET sets SOCIALAUTHAZUREADOAUTH2SECRET field to given value.

### HasSOCIALAUTHAZUREADOAUTH2SECRET

`func (o *InlineObject60) HasSOCIALAUTHAZUREADOAUTH2SECRET() bool`

HasSOCIALAUTHAZUREADOAUTH2SECRET returns a boolean if a field has been set.

### GetSOCIALAUTHAZUREADOAUTH2TEAMMAP

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2TEAMMAP() map[string]interface{}`

GetSOCIALAUTHAZUREADOAUTH2TEAMMAP returns the SOCIALAUTHAZUREADOAUTH2TEAMMAP field if non-nil, zero value otherwise.

### GetSOCIALAUTHAZUREADOAUTH2TEAMMAPOk

`func (o *InlineObject60) GetSOCIALAUTHAZUREADOAUTH2TEAMMAPOk() (*map[string]interface{}, bool)`

GetSOCIALAUTHAZUREADOAUTH2TEAMMAPOk returns a tuple with the SOCIALAUTHAZUREADOAUTH2TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHAZUREADOAUTH2TEAMMAP

`func (o *InlineObject60) SetSOCIALAUTHAZUREADOAUTH2TEAMMAP(v map[string]interface{})`

SetSOCIALAUTHAZUREADOAUTH2TEAMMAP sets SOCIALAUTHAZUREADOAUTH2TEAMMAP field to given value.

### HasSOCIALAUTHAZUREADOAUTH2TEAMMAP

`func (o *InlineObject60) HasSOCIALAUTHAZUREADOAUTH2TEAMMAP() bool`

HasSOCIALAUTHAZUREADOAUTH2TEAMMAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_KEY

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_KEY() string`

GetSOCIAL_AUTH_GITHUB_KEY returns the SOCIAL_AUTH_GITHUB_KEY field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_KEYOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_KEYOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_KEYOk returns a tuple with the SOCIAL_AUTH_GITHUB_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_KEY

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_KEY(v string)`

SetSOCIAL_AUTH_GITHUB_KEY sets SOCIAL_AUTH_GITHUB_KEY field to given value.

### HasSOCIAL_AUTH_GITHUB_KEY

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_KEY() bool`

HasSOCIAL_AUTH_GITHUB_KEY returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP returns the SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP sets SOCIAL_AUTH_GITHUB_ORGANIZATION_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP() bool`

HasSOCIAL_AUTH_GITHUB_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORG_KEY

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_KEY() string`

GetSOCIAL_AUTH_GITHUB_ORG_KEY returns the SOCIAL_AUTH_GITHUB_ORG_KEY field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORG_KEYOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_KEYOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_ORG_KEYOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORG_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORG_KEY

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORG_KEY(v string)`

SetSOCIAL_AUTH_GITHUB_ORG_KEY sets SOCIAL_AUTH_GITHUB_ORG_KEY field to given value.

### HasSOCIAL_AUTH_GITHUB_ORG_KEY

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORG_KEY() bool`

HasSOCIAL_AUTH_GITHUB_ORG_KEY returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORG_NAME

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_NAME() string`

GetSOCIAL_AUTH_GITHUB_ORG_NAME returns the SOCIAL_AUTH_GITHUB_ORG_NAME field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORG_NAMEOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_NAMEOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_ORG_NAMEOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORG_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORG_NAME

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORG_NAME(v string)`

SetSOCIAL_AUTH_GITHUB_ORG_NAME sets SOCIAL_AUTH_GITHUB_ORG_NAME field to given value.

### HasSOCIAL_AUTH_GITHUB_ORG_NAME

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORG_NAME() bool`

HasSOCIAL_AUTH_GITHUB_ORG_NAME returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP returns the SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP sets SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP() bool`

HasSOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORG_SECRET

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_SECRET() string`

GetSOCIAL_AUTH_GITHUB_ORG_SECRET returns the SOCIAL_AUTH_GITHUB_ORG_SECRET field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORG_SECRETOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_SECRETOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_ORG_SECRETOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORG_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORG_SECRET

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORG_SECRET(v string)`

SetSOCIAL_AUTH_GITHUB_ORG_SECRET sets SOCIAL_AUTH_GITHUB_ORG_SECRET field to given value.

### HasSOCIAL_AUTH_GITHUB_ORG_SECRET

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORG_SECRET() bool`

HasSOCIAL_AUTH_GITHUB_ORG_SECRET returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP returns the SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP sets SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP() bool`

HasSOCIAL_AUTH_GITHUB_ORG_TEAM_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_SECRET

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_SECRET() string`

GetSOCIAL_AUTH_GITHUB_SECRET returns the SOCIAL_AUTH_GITHUB_SECRET field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_SECRETOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_SECRETOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_SECRETOk returns a tuple with the SOCIAL_AUTH_GITHUB_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_SECRET

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_SECRET(v string)`

SetSOCIAL_AUTH_GITHUB_SECRET sets SOCIAL_AUTH_GITHUB_SECRET field to given value.

### HasSOCIAL_AUTH_GITHUB_SECRET

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_SECRET() bool`

HasSOCIAL_AUTH_GITHUB_SECRET returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_ID

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_ID() string`

GetSOCIAL_AUTH_GITHUB_TEAM_ID returns the SOCIAL_AUTH_GITHUB_TEAM_ID field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_IDOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_IDOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_IDOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_ID

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_ID(v string)`

SetSOCIAL_AUTH_GITHUB_TEAM_ID sets SOCIAL_AUTH_GITHUB_TEAM_ID field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_ID

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_ID() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_ID returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_KEY

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_KEY() string`

GetSOCIAL_AUTH_GITHUB_TEAM_KEY returns the SOCIAL_AUTH_GITHUB_TEAM_KEY field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_KEYOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_KEYOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_KEYOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_KEY

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_KEY(v string)`

SetSOCIAL_AUTH_GITHUB_TEAM_KEY sets SOCIAL_AUTH_GITHUB_TEAM_KEY field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_KEY

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_KEY() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_KEY returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_TEAM_MAP returns the SOCIAL_AUTH_GITHUB_TEAM_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_TEAM_MAP sets SOCIAL_AUTH_GITHUB_TEAM_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_MAP() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP returns the SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP sets SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_SECRET

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_SECRET() string`

GetSOCIAL_AUTH_GITHUB_TEAM_SECRET returns the SOCIAL_AUTH_GITHUB_TEAM_SECRET field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_SECRETOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_SECRETOk() (*string, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_SECRETOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_SECRET

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_SECRET(v string)`

SetSOCIAL_AUTH_GITHUB_TEAM_SECRET sets SOCIAL_AUTH_GITHUB_TEAM_SECRET field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_SECRET

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_SECRET() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_SECRET returns a boolean if a field has been set.

### GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP() map[string]interface{}`

GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP returns the SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAPOk returns a tuple with the SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP sets SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP field to given value.

### HasSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP() bool`

HasSOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS() map[string]interface{}`

GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS returns the SOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTSOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTSOk() (*map[string]interface{}, bool)`

GetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTSOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS(v map[string]interface{})`

SetSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS sets SOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS() bool`

HasSOCIALAUTHGOOGLEOAUTH2AUTHEXTRAARGUMENTS returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2KEY

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2KEY() string`

GetSOCIALAUTHGOOGLEOAUTH2KEY returns the SOCIALAUTHGOOGLEOAUTH2KEY field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2KEYOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2KEYOk() (*string, bool)`

GetSOCIALAUTHGOOGLEOAUTH2KEYOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2KEY

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2KEY(v string)`

SetSOCIALAUTHGOOGLEOAUTH2KEY sets SOCIALAUTHGOOGLEOAUTH2KEY field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2KEY

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2KEY() bool`

HasSOCIALAUTHGOOGLEOAUTH2KEY returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP() map[string]interface{}`

GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP returns the SOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAPOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAPOk() (*map[string]interface{}, bool)`

GetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAPOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP(v map[string]interface{})`

SetSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP sets SOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP() bool`

HasSOCIALAUTHGOOGLEOAUTH2ORGANIZATIONMAP returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2SECRET

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2SECRET() string`

GetSOCIALAUTHGOOGLEOAUTH2SECRET returns the SOCIALAUTHGOOGLEOAUTH2SECRET field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2SECRETOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2SECRETOk() (*string, bool)`

GetSOCIALAUTHGOOGLEOAUTH2SECRETOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2SECRET

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2SECRET(v string)`

SetSOCIALAUTHGOOGLEOAUTH2SECRET sets SOCIALAUTHGOOGLEOAUTH2SECRET field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2SECRET

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2SECRET() bool`

HasSOCIALAUTHGOOGLEOAUTH2SECRET returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2TEAMMAP

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2TEAMMAP() map[string]interface{}`

GetSOCIALAUTHGOOGLEOAUTH2TEAMMAP returns the SOCIALAUTHGOOGLEOAUTH2TEAMMAP field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2TEAMMAPOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2TEAMMAPOk() (*map[string]interface{}, bool)`

GetSOCIALAUTHGOOGLEOAUTH2TEAMMAPOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2TEAMMAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2TEAMMAP

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2TEAMMAP(v map[string]interface{})`

SetSOCIALAUTHGOOGLEOAUTH2TEAMMAP sets SOCIALAUTHGOOGLEOAUTH2TEAMMAP field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2TEAMMAP

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2TEAMMAP() bool`

HasSOCIALAUTHGOOGLEOAUTH2TEAMMAP returns a boolean if a field has been set.

### GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS() []string`

GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS returns the SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS field if non-nil, zero value otherwise.

### GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINSOk

`func (o *InlineObject60) GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINSOk() (*[]string, bool)`

GetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINSOk returns a tuple with the SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS

`func (o *InlineObject60) SetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS(v []string)`

SetSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS sets SOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS field to given value.

### HasSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS

`func (o *InlineObject60) HasSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS() bool`

HasSOCIALAUTHGOOGLEOAUTH2WHITELISTEDDOMAINS returns a boolean if a field has been set.

### GetSOCIAL_AUTH_ORGANIZATION_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_ORGANIZATION_MAP() map[string]interface{}`

GetSOCIAL_AUTH_ORGANIZATION_MAP returns the SOCIAL_AUTH_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_ORGANIZATION_MAPOk returns a tuple with the SOCIAL_AUTH_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_ORGANIZATION_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_ORGANIZATION_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_ORGANIZATION_MAP sets SOCIAL_AUTH_ORGANIZATION_MAP field to given value.

### HasSOCIAL_AUTH_ORGANIZATION_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_ORGANIZATION_MAP() bool`

HasSOCIAL_AUTH_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_ENABLED_IDPS

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ENABLED_IDPS() map[string]interface{}`

GetSOCIAL_AUTH_SAML_ENABLED_IDPS returns the SOCIAL_AUTH_SAML_ENABLED_IDPS field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_ENABLED_IDPSOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ENABLED_IDPSOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_ENABLED_IDPSOk returns a tuple with the SOCIAL_AUTH_SAML_ENABLED_IDPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_ENABLED_IDPS

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_ENABLED_IDPS(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_ENABLED_IDPS sets SOCIAL_AUTH_SAML_ENABLED_IDPS field to given value.

### HasSOCIAL_AUTH_SAML_ENABLED_IDPS

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_ENABLED_IDPS() bool`

HasSOCIAL_AUTH_SAML_ENABLED_IDPS returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_EXTRA_DATA

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_EXTRA_DATA() []string`

GetSOCIAL_AUTH_SAML_EXTRA_DATA returns the SOCIAL_AUTH_SAML_EXTRA_DATA field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_EXTRA_DATAOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_EXTRA_DATAOk() (*[]string, bool)`

GetSOCIAL_AUTH_SAML_EXTRA_DATAOk returns a tuple with the SOCIAL_AUTH_SAML_EXTRA_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_EXTRA_DATA

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_EXTRA_DATA(v []string)`

SetSOCIAL_AUTH_SAML_EXTRA_DATA sets SOCIAL_AUTH_SAML_EXTRA_DATA field to given value.

### HasSOCIAL_AUTH_SAML_EXTRA_DATA

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_EXTRA_DATA() bool`

HasSOCIAL_AUTH_SAML_EXTRA_DATA returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR() map[string]interface{}`

GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR returns the SOCIAL_AUTH_SAML_ORGANIZATION_ATTR field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTROk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTROk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_ORGANIZATION_ATTROk returns a tuple with the SOCIAL_AUTH_SAML_ORGANIZATION_ATTR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_ORGANIZATION_ATTR sets SOCIAL_AUTH_SAML_ORGANIZATION_ATTR field to given value.

### HasSOCIAL_AUTH_SAML_ORGANIZATION_ATTR

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_ORGANIZATION_ATTR() bool`

HasSOCIAL_AUTH_SAML_ORGANIZATION_ATTR returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_ORGANIZATION_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORGANIZATION_MAP() map[string]interface{}`

GetSOCIAL_AUTH_SAML_ORGANIZATION_MAP returns the SOCIAL_AUTH_SAML_ORGANIZATION_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_ORGANIZATION_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORGANIZATION_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_ORGANIZATION_MAPOk returns a tuple with the SOCIAL_AUTH_SAML_ORGANIZATION_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_ORGANIZATION_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_ORGANIZATION_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_ORGANIZATION_MAP sets SOCIAL_AUTH_SAML_ORGANIZATION_MAP field to given value.

### HasSOCIAL_AUTH_SAML_ORGANIZATION_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_ORGANIZATION_MAP() bool`

HasSOCIAL_AUTH_SAML_ORGANIZATION_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_ORG_INFO

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORG_INFO() map[string]interface{}`

GetSOCIAL_AUTH_SAML_ORG_INFO returns the SOCIAL_AUTH_SAML_ORG_INFO field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_ORG_INFOOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_ORG_INFOOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_ORG_INFOOk returns a tuple with the SOCIAL_AUTH_SAML_ORG_INFO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_ORG_INFO

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_ORG_INFO(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_ORG_INFO sets SOCIAL_AUTH_SAML_ORG_INFO field to given value.


### GetSOCIAL_AUTH_SAML_SECURITY_CONFIG

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SECURITY_CONFIG() map[string]interface{}`

GetSOCIAL_AUTH_SAML_SECURITY_CONFIG returns the SOCIAL_AUTH_SAML_SECURITY_CONFIG field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SECURITY_CONFIGOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SECURITY_CONFIGOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_SECURITY_CONFIGOk returns a tuple with the SOCIAL_AUTH_SAML_SECURITY_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SECURITY_CONFIG

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SECURITY_CONFIG(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_SECURITY_CONFIG sets SOCIAL_AUTH_SAML_SECURITY_CONFIG field to given value.

### HasSOCIAL_AUTH_SAML_SECURITY_CONFIG

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_SECURITY_CONFIG() bool`

HasSOCIAL_AUTH_SAML_SECURITY_CONFIG returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_SP_ENTITY_ID

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_ENTITY_ID() string`

GetSOCIAL_AUTH_SAML_SP_ENTITY_ID returns the SOCIAL_AUTH_SAML_SP_ENTITY_ID field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SP_ENTITY_IDOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_ENTITY_IDOk() (*string, bool)`

GetSOCIAL_AUTH_SAML_SP_ENTITY_IDOk returns a tuple with the SOCIAL_AUTH_SAML_SP_ENTITY_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SP_ENTITY_ID

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SP_ENTITY_ID(v string)`

SetSOCIAL_AUTH_SAML_SP_ENTITY_ID sets SOCIAL_AUTH_SAML_SP_ENTITY_ID field to given value.

### HasSOCIAL_AUTH_SAML_SP_ENTITY_ID

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_SP_ENTITY_ID() bool`

HasSOCIAL_AUTH_SAML_SP_ENTITY_ID returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_SP_EXTRA

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_EXTRA() map[string]interface{}`

GetSOCIAL_AUTH_SAML_SP_EXTRA returns the SOCIAL_AUTH_SAML_SP_EXTRA field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SP_EXTRAOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_EXTRAOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_SP_EXTRAOk returns a tuple with the SOCIAL_AUTH_SAML_SP_EXTRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SP_EXTRA

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SP_EXTRA(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_SP_EXTRA sets SOCIAL_AUTH_SAML_SP_EXTRA field to given value.

### HasSOCIAL_AUTH_SAML_SP_EXTRA

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_SP_EXTRA() bool`

HasSOCIAL_AUTH_SAML_SP_EXTRA returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY() string`

GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY returns the SOCIAL_AUTH_SAML_SP_PRIVATE_KEY field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEYOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEYOk() (*string, bool)`

GetSOCIAL_AUTH_SAML_SP_PRIVATE_KEYOk returns a tuple with the SOCIAL_AUTH_SAML_SP_PRIVATE_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY(v string)`

SetSOCIAL_AUTH_SAML_SP_PRIVATE_KEY sets SOCIAL_AUTH_SAML_SP_PRIVATE_KEY field to given value.


### GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT() string`

GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT returns the SOCIAL_AUTH_SAML_SP_PUBLIC_CERT field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERTOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERTOk() (*string, bool)`

GetSOCIAL_AUTH_SAML_SP_PUBLIC_CERTOk returns a tuple with the SOCIAL_AUTH_SAML_SP_PUBLIC_CERT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT(v string)`

SetSOCIAL_AUTH_SAML_SP_PUBLIC_CERT sets SOCIAL_AUTH_SAML_SP_PUBLIC_CERT field to given value.


### GetSOCIAL_AUTH_SAML_SUPPORT_CONTACT

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SUPPORT_CONTACT() map[string]interface{}`

GetSOCIAL_AUTH_SAML_SUPPORT_CONTACT returns the SOCIAL_AUTH_SAML_SUPPORT_CONTACT field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_SUPPORT_CONTACTOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_SUPPORT_CONTACTOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_SUPPORT_CONTACTOk returns a tuple with the SOCIAL_AUTH_SAML_SUPPORT_CONTACT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_SUPPORT_CONTACT

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_SUPPORT_CONTACT(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_SUPPORT_CONTACT sets SOCIAL_AUTH_SAML_SUPPORT_CONTACT field to given value.


### GetSOCIAL_AUTH_SAML_TEAM_ATTR

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TEAM_ATTR() map[string]interface{}`

GetSOCIAL_AUTH_SAML_TEAM_ATTR returns the SOCIAL_AUTH_SAML_TEAM_ATTR field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_TEAM_ATTROk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TEAM_ATTROk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_TEAM_ATTROk returns a tuple with the SOCIAL_AUTH_SAML_TEAM_ATTR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_TEAM_ATTR

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_TEAM_ATTR(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_TEAM_ATTR sets SOCIAL_AUTH_SAML_TEAM_ATTR field to given value.

### HasSOCIAL_AUTH_SAML_TEAM_ATTR

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_TEAM_ATTR() bool`

HasSOCIAL_AUTH_SAML_TEAM_ATTR returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_TEAM_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TEAM_MAP() map[string]interface{}`

GetSOCIAL_AUTH_SAML_TEAM_MAP returns the SOCIAL_AUTH_SAML_TEAM_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_TEAM_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_TEAM_MAPOk returns a tuple with the SOCIAL_AUTH_SAML_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_TEAM_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_TEAM_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_TEAM_MAP sets SOCIAL_AUTH_SAML_TEAM_MAP field to given value.

### HasSOCIAL_AUTH_SAML_TEAM_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_SAML_TEAM_MAP() bool`

HasSOCIAL_AUTH_SAML_TEAM_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT() map[string]interface{}`

GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT returns the SOCIAL_AUTH_SAML_TECHNICAL_CONTACT field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACTOk

`func (o *InlineObject60) GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACTOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_SAML_TECHNICAL_CONTACTOk returns a tuple with the SOCIAL_AUTH_SAML_TECHNICAL_CONTACT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT

`func (o *InlineObject60) SetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT(v map[string]interface{})`

SetSOCIAL_AUTH_SAML_TECHNICAL_CONTACT sets SOCIAL_AUTH_SAML_TECHNICAL_CONTACT field to given value.


### GetSOCIAL_AUTH_TEAM_MAP

`func (o *InlineObject60) GetSOCIAL_AUTH_TEAM_MAP() map[string]interface{}`

GetSOCIAL_AUTH_TEAM_MAP returns the SOCIAL_AUTH_TEAM_MAP field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_TEAM_MAPOk

`func (o *InlineObject60) GetSOCIAL_AUTH_TEAM_MAPOk() (*map[string]interface{}, bool)`

GetSOCIAL_AUTH_TEAM_MAPOk returns a tuple with the SOCIAL_AUTH_TEAM_MAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_TEAM_MAP

`func (o *InlineObject60) SetSOCIAL_AUTH_TEAM_MAP(v map[string]interface{})`

SetSOCIAL_AUTH_TEAM_MAP sets SOCIAL_AUTH_TEAM_MAP field to given value.

### HasSOCIAL_AUTH_TEAM_MAP

`func (o *InlineObject60) HasSOCIAL_AUTH_TEAM_MAP() bool`

HasSOCIAL_AUTH_TEAM_MAP returns a boolean if a field has been set.

### GetSOCIAL_AUTH_USER_FIELDS

`func (o *InlineObject60) GetSOCIAL_AUTH_USER_FIELDS() []string`

GetSOCIAL_AUTH_USER_FIELDS returns the SOCIAL_AUTH_USER_FIELDS field if non-nil, zero value otherwise.

### GetSOCIAL_AUTH_USER_FIELDSOk

`func (o *InlineObject60) GetSOCIAL_AUTH_USER_FIELDSOk() (*[]string, bool)`

GetSOCIAL_AUTH_USER_FIELDSOk returns a tuple with the SOCIAL_AUTH_USER_FIELDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOCIAL_AUTH_USER_FIELDS

`func (o *InlineObject60) SetSOCIAL_AUTH_USER_FIELDS(v []string)`

SetSOCIAL_AUTH_USER_FIELDS sets SOCIAL_AUTH_USER_FIELDS field to given value.

### HasSOCIAL_AUTH_USER_FIELDS

`func (o *InlineObject60) HasSOCIAL_AUTH_USER_FIELDS() bool`

HasSOCIAL_AUTH_USER_FIELDS returns a boolean if a field has been set.

### GetSTDOUT_MAX_BYTES_DISPLAY

`func (o *InlineObject60) GetSTDOUT_MAX_BYTES_DISPLAY() int32`

GetSTDOUT_MAX_BYTES_DISPLAY returns the STDOUT_MAX_BYTES_DISPLAY field if non-nil, zero value otherwise.

### GetSTDOUT_MAX_BYTES_DISPLAYOk

`func (o *InlineObject60) GetSTDOUT_MAX_BYTES_DISPLAYOk() (*int32, bool)`

GetSTDOUT_MAX_BYTES_DISPLAYOk returns a tuple with the STDOUT_MAX_BYTES_DISPLAY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTDOUT_MAX_BYTES_DISPLAY

`func (o *InlineObject60) SetSTDOUT_MAX_BYTES_DISPLAY(v int32)`

SetSTDOUT_MAX_BYTES_DISPLAY sets STDOUT_MAX_BYTES_DISPLAY field to given value.


### GetTACACSPLUS_AUTH_PROTOCOL

`func (o *InlineObject60) GetTACACSPLUS_AUTH_PROTOCOL() string`

GetTACACSPLUS_AUTH_PROTOCOL returns the TACACSPLUS_AUTH_PROTOCOL field if non-nil, zero value otherwise.

### GetTACACSPLUS_AUTH_PROTOCOLOk

`func (o *InlineObject60) GetTACACSPLUS_AUTH_PROTOCOLOk() (*string, bool)`

GetTACACSPLUS_AUTH_PROTOCOLOk returns a tuple with the TACACSPLUS_AUTH_PROTOCOL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTACACSPLUS_AUTH_PROTOCOL

`func (o *InlineObject60) SetTACACSPLUS_AUTH_PROTOCOL(v string)`

SetTACACSPLUS_AUTH_PROTOCOL sets TACACSPLUS_AUTH_PROTOCOL field to given value.

### HasTACACSPLUS_AUTH_PROTOCOL

`func (o *InlineObject60) HasTACACSPLUS_AUTH_PROTOCOL() bool`

HasTACACSPLUS_AUTH_PROTOCOL returns a boolean if a field has been set.

### GetTACACSPLUS_HOST

`func (o *InlineObject60) GetTACACSPLUS_HOST() string`

GetTACACSPLUS_HOST returns the TACACSPLUS_HOST field if non-nil, zero value otherwise.

### GetTACACSPLUS_HOSTOk

`func (o *InlineObject60) GetTACACSPLUS_HOSTOk() (*string, bool)`

GetTACACSPLUS_HOSTOk returns a tuple with the TACACSPLUS_HOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTACACSPLUS_HOST

`func (o *InlineObject60) SetTACACSPLUS_HOST(v string)`

SetTACACSPLUS_HOST sets TACACSPLUS_HOST field to given value.

### HasTACACSPLUS_HOST

`func (o *InlineObject60) HasTACACSPLUS_HOST() bool`

HasTACACSPLUS_HOST returns a boolean if a field has been set.

### GetTACACSPLUS_PORT

`func (o *InlineObject60) GetTACACSPLUS_PORT() int32`

GetTACACSPLUS_PORT returns the TACACSPLUS_PORT field if non-nil, zero value otherwise.

### GetTACACSPLUS_PORTOk

`func (o *InlineObject60) GetTACACSPLUS_PORTOk() (*int32, bool)`

GetTACACSPLUS_PORTOk returns a tuple with the TACACSPLUS_PORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTACACSPLUS_PORT

`func (o *InlineObject60) SetTACACSPLUS_PORT(v int32)`

SetTACACSPLUS_PORT sets TACACSPLUS_PORT field to given value.

### HasTACACSPLUS_PORT

`func (o *InlineObject60) HasTACACSPLUS_PORT() bool`

HasTACACSPLUS_PORT returns a boolean if a field has been set.

### GetTACACSPLUS_SECRET

`func (o *InlineObject60) GetTACACSPLUS_SECRET() string`

GetTACACSPLUS_SECRET returns the TACACSPLUS_SECRET field if non-nil, zero value otherwise.

### GetTACACSPLUS_SECRETOk

`func (o *InlineObject60) GetTACACSPLUS_SECRETOk() (*string, bool)`

GetTACACSPLUS_SECRETOk returns a tuple with the TACACSPLUS_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTACACSPLUS_SECRET

`func (o *InlineObject60) SetTACACSPLUS_SECRET(v string)`

SetTACACSPLUS_SECRET sets TACACSPLUS_SECRET field to given value.

### HasTACACSPLUS_SECRET

`func (o *InlineObject60) HasTACACSPLUS_SECRET() bool`

HasTACACSPLUS_SECRET returns a boolean if a field has been set.

### GetTACACSPLUS_SESSION_TIMEOUT

`func (o *InlineObject60) GetTACACSPLUS_SESSION_TIMEOUT() int32`

GetTACACSPLUS_SESSION_TIMEOUT returns the TACACSPLUS_SESSION_TIMEOUT field if non-nil, zero value otherwise.

### GetTACACSPLUS_SESSION_TIMEOUTOk

`func (o *InlineObject60) GetTACACSPLUS_SESSION_TIMEOUTOk() (*int32, bool)`

GetTACACSPLUS_SESSION_TIMEOUTOk returns a tuple with the TACACSPLUS_SESSION_TIMEOUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTACACSPLUS_SESSION_TIMEOUT

`func (o *InlineObject60) SetTACACSPLUS_SESSION_TIMEOUT(v int32)`

SetTACACSPLUS_SESSION_TIMEOUT sets TACACSPLUS_SESSION_TIMEOUT field to given value.

### HasTACACSPLUS_SESSION_TIMEOUT

`func (o *InlineObject60) HasTACACSPLUS_SESSION_TIMEOUT() bool`

HasTACACSPLUS_SESSION_TIMEOUT returns a boolean if a field has been set.

### GetTOWER_URL_BASE

`func (o *InlineObject60) GetTOWER_URL_BASE() string`

GetTOWER_URL_BASE returns the TOWER_URL_BASE field if non-nil, zero value otherwise.

### GetTOWER_URL_BASEOk

`func (o *InlineObject60) GetTOWER_URL_BASEOk() (*string, bool)`

GetTOWER_URL_BASEOk returns a tuple with the TOWER_URL_BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOWER_URL_BASE

`func (o *InlineObject60) SetTOWER_URL_BASE(v string)`

SetTOWER_URL_BASE sets TOWER_URL_BASE field to given value.


### GetUI_LIVE_UPDATES_ENABLED

`func (o *InlineObject60) GetUI_LIVE_UPDATES_ENABLED() bool`

GetUI_LIVE_UPDATES_ENABLED returns the UI_LIVE_UPDATES_ENABLED field if non-nil, zero value otherwise.

### GetUI_LIVE_UPDATES_ENABLEDOk

`func (o *InlineObject60) GetUI_LIVE_UPDATES_ENABLEDOk() (*bool, bool)`

GetUI_LIVE_UPDATES_ENABLEDOk returns a tuple with the UI_LIVE_UPDATES_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUI_LIVE_UPDATES_ENABLED

`func (o *InlineObject60) SetUI_LIVE_UPDATES_ENABLED(v bool)`

SetUI_LIVE_UPDATES_ENABLED sets UI_LIVE_UPDATES_ENABLED field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



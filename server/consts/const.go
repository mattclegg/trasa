package consts

import "time"

//TODO change this

//TRASA_VERSION is current version of TRASA
const TRASA_VERSION = "2020.7.2"

const (
	ChangePassword = "changepassword"
)

type EntityConst string

const (

	//Entities

	ENTITY_USER                   EntityConst = "ENTITY_USER"
	ENTITY_APP                    EntityConst = "ENTITY_APP"
	ENTITY_APP_USER_MAP           EntityConst = "ENTITY_APP_USER_MAP"
	ENTITY_GROUP                  EntityConst = "ENTITY_GROUP"
	ENTITY_APP_GROUP_MAP          EntityConst = "ENTITY_APP_GROUP_MAP"
	ENTITY_USER_GROUP_MAP         EntityConst = "ENTITY_USER_GROUP_MAP"
	ENTITY_USERGROUP_APP          EntityConst = "ENTITY_USERGROUP_APP"
	ENTITY_USERGROUP_APPGROUP_MAP EntityConst = "ENTITY_USERGROUP_APPGROUP_MAP"
	ENTITY_USER_DEVICE            EntityConst = "ENTITY_USER_DEVICE"

	// Scopes
	ALL_USER        = "ALL_USER"
	ALL_ORGADMINS   = "ALL_ORGADMINS"
	ALL_NORMALUSERS = "ALL_NORMALUSERS"
	ALL_SERVICES    = "ALL_SERVICES"

	ALL_GROUPS           = "ALL_GROUPS"
	ALL_USERGROUPS       = "ALL_USERGROUPS"
	ALL_SERVICEGROUPS    = "ALL_SERVICEGROUPS"
	ALL_SSHSERVICES      = "ALL_SSHSERVICES"
	ALL_RDPSERVICES      = "ALL_RDPSERVICES"
	ALL_HTTPSERVICES     = "ALL_HTTPSERVICES"
	ALL_DATABASESERVICES = "ALL_DATABASESERVICES"

	ORG    = "ORG"
	DOMAIN = "DOMAIN"

	// Passwords
	ADMIN_FORGOT_PASSWORD = "ADMIN_FORGOT_PASSWORD"

	// Security
	SUSPICIOUS_LOGIN = "SUSPICIOUS_LOGIN"

	// System
	SYSTEM              = "SYSTEM"
	LOW_SYSTEM_RESOURCE = "LOW_SYSTEM_RESOURCE"
)

const (
	SCIM_USER_SCHEMA            = "urn:ietf:params:scim:schemas:core:2.0:User"
	SCIM_ENTERPRISE_USER_SCHEMA = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"

	SCIM_GROUP_SCHEMA = "urn:ietf:params:scim:schemas:core:2.0:Group"

	SCIM_ERR      = "urn:ietf:params:scim:api:messages:2.0:Error"
	SCIM_LISTRESP = "urn:ietf:params:scim:api:messages:2.0:ListResponse"
)

const (
	KEY_DOAPI = "KEY_DOAPI"
	KEY_LDAP  = "KEY_LDAP"
	KEY_SCIM  = "KEY_SCIM"
	KEY_SMTP  = "KEY_SMTP"
)

type EmailType string

const (
	EMAIL_ADHOC          EmailType = "EMAIL_ADHOC"
	EMAIL_DYNAMIC_ACCESS EmailType = "EMAIL_DYNAMIC_ACCESS"
	EMAIL_SECURITY_ALERT EmailType = "EMAIL_SECURITY_ALERT"
	EMAIL_USER_CRUD      EmailType = "EMAIL_USER_CRUD"
	EMAIL_ERR_REPORT     EmailType = "EMAIL_ERR_REPORT"
)

const (
	EMAIL_SMTP    EmailType = "EMAIL_SMTP"
	EMAIL_MAILGUN EmailType = "EMAIL_MAILGUN"
)

type VerifyTokenIntent string

const (
	VERIFY_TOKEN_CHANGEPASS VerifyTokenIntent = "VERIFY_TOKEN_CHANGEPASS"
	VERIFY_TOKEN_TFA        VerifyTokenIntent = "VERIFY_TOKEN_TFA"
)

type TokenExpiryTime int

const (
	TOKEN_EXPIRY_CHANGEPASS time.Duration = time.Second * 900   // 15 MINUTES
	TOKEN_EXPIRY_SIGNUP     time.Duration = time.Second * 90000 // 25 HOURS
)

const (
	AUTH_RESP_NOTIF_LICENSE   = "AUTH_RESP_NOTIF_LICENSE"
	AUTH_RESP_ENROL_DEVICE    = "AUTH_RESP_ENROL_DEVICE"
	AUTH_RESP_TFA_REQUIRED    = "AUTH_RESP_TFA_REQUIRED"
	AUTH_RESP_TFA_DH_REQUIRED = "AUTH_RESP_TFA_DH_REQUIRED"
	AUTH_RESP_SELECT_ORG      = "AUTH_RESP_SELECT_ORG"
	AUTH_RESP_CHANGE_PASS     = "AUTH_RESP_CHANGE_PASS"
	AUTH_RESP_RESET_PASS      = "AUTH_RESP_RESET_PASS"
	AUTH_RESP_FORGOT_PASS     = "AUTH_RESP_FORGOT_PASS"
)

const (
	AUTH_REQ_DASH_LOGIN    = "AUTH_REQ_DASH_LOGIN"
	AUTH_REQ_CHANGE_PASS   = "AUTH_REQ_CHANGE_PASS"
	AUTH_REQ_ENROL_DEVICE  = "AUTH_REQ_ENROL_DEVICE"
	AUTH_REQ_FORGOT_PASS   = "AUTH_REQ_FORGOT_PASS"
	AUTH_HTTP_ACCESS_PROXY = "AUTH_HTTP_ACCESS_PROXY"
	AUTH_REQ_TFA_DH        = "AUTH_REQ_TFA_DH"
)

const (
	IDP_TRASA   = "IDP_TRASA"
	IDP_FREEIPA = "IDP_FREEIPA"
)

const DEFAULT_ROOT_PASSWORD = "changeme"

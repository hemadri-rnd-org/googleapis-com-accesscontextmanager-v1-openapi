package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetIamPolicyRequest represents the GetIamPolicyRequest schema from the OpenAPI specification
type GetIamPolicyRequest struct {
	Options GetPolicyOptions `json:"options,omitempty"` // Encapsulates settings provided to GetIamPolicy.
}

// IngressTo represents the IngressTo schema from the OpenAPI specification
type IngressTo struct {
	Operations []ApiOperation `json:"operations,omitempty"` // A list of ApiOperations allowed to be performed by the sources specified in corresponding IngressFrom in this ServicePerimeter.
	Resources []string `json:"resources,omitempty"` // A list of resources, currently only projects in the form `projects/`, protected by this ServicePerimeter that are allowed to be accessed by sources defined in the corresponding IngressFrom. If a single `*` is specified, then access to all resources inside the perimeter are allowed.
}

// AuditLogConfig represents the AuditLogConfig schema from the OpenAPI specification
type AuditLogConfig struct {
	Logtype string `json:"logType,omitempty"` // The log type that this config enables.
	Exemptedmembers []string `json:"exemptedMembers,omitempty"` // Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
}

// OsConstraint represents the OsConstraint schema from the OpenAPI specification
type OsConstraint struct {
	Requireverifiedchromeos bool `json:"requireVerifiedChromeOs,omitempty"` // Only allows requests from devices with a verified Chrome OS. Verifications includes requirements that the device is enterprise-managed, conformant to domain policies, and the caller has permission to call the API targeted by the request.
	Minimumversion string `json:"minimumVersion,omitempty"` // The minimum allowed OS version. If not set, any version of this OS satisfies the constraint. Format: `"major.minor.patch"`. Examples: `"10.5.301"`, `"9.2.1"`.
	Ostype string `json:"osType,omitempty"` // Required. The allowed OS type.
}

// SetIamPolicyRequest represents the SetIamPolicyRequest schema from the OpenAPI specification
type SetIamPolicyRequest struct {
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
	Updatemask string `json:"updateMask,omitempty"` // OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
}

// VpcNetworkSource represents the VpcNetworkSource schema from the OpenAPI specification
type VpcNetworkSource struct {
	Vpcsubnetwork VpcSubNetwork `json:"vpcSubnetwork,omitempty"` // Sub-segment ranges inside of a VPC Network.
}

// EgressFrom represents the EgressFrom schema from the OpenAPI specification
type EgressFrom struct {
	Identities []string `json:"identities,omitempty"` // A list of identities that are allowed access through this [EgressPolicy], in the format of `user:{email_id}` or `serviceAccount:{email_id}`.
	Identitytype string `json:"identityType,omitempty"` // Specifies the type of identities that are allowed access to outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
	Sourcerestriction string `json:"sourceRestriction,omitempty"` // Whether to enforce traffic restrictions based on `sources` field. If the `sources` fields is non-empty, then this field must be set to `SOURCE_RESTRICTION_ENABLED`.
	Sources []EgressSource `json:"sources,omitempty"` // Sources that this EgressPolicy authorizes access from. If this field is not empty, then `source_restriction` must be set to `SOURCE_RESTRICTION_ENABLED`.
}

// VpcAccessibleServices represents the VpcAccessibleServices schema from the OpenAPI specification
type VpcAccessibleServices struct {
	Allowedservices []string `json:"allowedServices,omitempty"` // The list of APIs usable within the Service Perimeter. Must be empty unless 'enable_restriction' is True. You can specify a list of individual services, as well as include the 'RESTRICTED-SERVICES' value, which automatically includes all of the services protected by the perimeter.
	Enablerestriction bool `json:"enableRestriction,omitempty"` // Whether to restrict API calls within the Service Perimeter to the list of APIs specified in 'allowed_services'.
}

// ApiOperation represents the ApiOperation schema from the OpenAPI specification
type ApiOperation struct {
	Methodselectors []MethodSelector `json:"methodSelectors,omitempty"` // API methods or permissions to allow. Method or permission must belong to the service specified by `service_name` field. A single MethodSelector entry with `*` specified for the `method` field will allow all methods AND permissions for the service specified in `service_name`.
	Servicename string `json:"serviceName,omitempty"` // The name of the API whose methods or permissions the IngressPolicy or EgressPolicy want to allow. A single ApiOperation with `service_name` field set to `*` will allow all methods AND permissions for all services.
}

// IngressSource represents the IngressSource schema from the OpenAPI specification
type IngressSource struct {
	Resource string `json:"resource,omitempty"` // A Google Cloud resource that is allowed to ingress the perimeter. Requests from these resources will be allowed to access perimeter data. Currently only projects and VPCs are allowed. Project format: `projects/{project_number}` VPC network format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`. The project may be in any Google Cloud organization, not just the organization that the perimeter is defined in. `*` is not allowed, the case of allowing all Google Cloud resources only is not supported.
	Accesslevel string `json:"accessLevel,omitempty"` // An AccessLevel resource name that allow resources within the ServicePerimeters to be accessed from the internet. AccessLevels listed must be in the same policy as this ServicePerimeter. Referencing a nonexistent AccessLevel will cause an error. If no AccessLevel names are listed, resources within the perimeter can only be accessed via Google Cloud calls with request origins within the perimeter. Example: `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is specified for `access_level`, then all IngressSources will be allowed.
}

// Binding represents the Binding schema from the OpenAPI specification
type Binding struct {
	Role string `json:"role,omitempty"` // Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. For an overview of the IAM roles and permissions, see the [IAM documentation](https://cloud.google.com/iam/docs/roles-overview). For a list of the available pre-defined roles, see [here](https://cloud.google.com/iam/docs/understanding-roles).
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
	Members []string `json:"members,omitempty"` // Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
}

// BasicLevel represents the BasicLevel schema from the OpenAPI specification
type BasicLevel struct {
	Combiningfunction string `json:"combiningFunction,omitempty"` // How the `conditions` list should be combined to determine if a request is granted this `AccessLevel`. If AND is used, each `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. If OR is used, at least one `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. Default behavior is AND.
	Conditions []Condition `json:"conditions,omitempty"` // Required. A list of requirements for the `AccessLevel` to be granted.
}

// ListAccessPoliciesResponse represents the ListAccessPoliciesResponse schema from the OpenAPI specification
type ListAccessPoliciesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The pagination token to retrieve the next page of results. If the value is empty, no further results remain.
	Accesspolicies []AccessPolicy `json:"accessPolicies,omitempty"` // List of the AccessPolicy instances.
}

// GcpUserAccessBinding represents the GcpUserAccessBinding schema from the OpenAPI specification
type GcpUserAccessBinding struct {
	Accesslevels []string `json:"accessLevels,omitempty"` // Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	Dryrunaccesslevels []string `json:"dryRunAccessLevels,omitempty"` // Optional. Dry run access level that will be evaluated but will not be enforced. The access denial based on dry run policy will be logged. Only one access level is supported, not multiple. This list must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	Groupkey string `json:"groupKey,omitempty"` // Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	Name string `json:"name,omitempty"` // Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
}

// ListAuthorizedOrgsDescsResponse represents the ListAuthorizedOrgsDescsResponse schema from the OpenAPI specification
type ListAuthorizedOrgsDescsResponse struct {
	Authorizedorgsdescs []AuthorizedOrgsDesc `json:"authorizedOrgsDescs,omitempty"` // List of all the Authorized Orgs Desc instances.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The pagination token to retrieve the next page of results. If the value is empty, no further results remain.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// ListSupportedServicesResponse represents the ListSupportedServicesResponse schema from the OpenAPI specification
type ListSupportedServicesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The pagination token to retrieve the next page of results. If the value is empty, no further results remain.
	Supportedservices []SupportedService `json:"supportedServices,omitempty"` // List of services supported by VPC Service Controls instances.
}

// AccessContextManagerOperationMetadata represents the AccessContextManagerOperationMetadata schema from the OpenAPI specification
type AccessContextManagerOperationMetadata struct {
}

// Expr represents the Expr schema from the OpenAPI specification
type Expr struct {
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
}

// MethodSelector represents the MethodSelector schema from the OpenAPI specification
type MethodSelector struct {
	Method string `json:"method,omitempty"` // A valid method name for the corresponding `service_name` in ApiOperation. If `*` is used as the value for the `method`, then ALL methods and permissions are allowed.
	Permission string `json:"permission,omitempty"` // A valid Cloud IAM permission for the corresponding `service_name` in ApiOperation.
}

// IngressPolicy represents the IngressPolicy schema from the OpenAPI specification
type IngressPolicy struct {
	Ingressfrom IngressFrom `json:"ingressFrom,omitempty"` // Defines the conditions under which an IngressPolicy matches a request. Conditions are based on information about the source of the request. The request must satisfy what is defined in `sources` AND identity related fields in order to match.
	Ingressto IngressTo `json:"ingressTo,omitempty"` // Defines the conditions under which an IngressPolicy matches a request. Conditions are based on information about the ApiOperation intended to be performed on the target resource of the request. The request must satisfy what is defined in `operations` AND `resources` in order to match.
}

// ListServicePerimetersResponse represents the ListServicePerimetersResponse schema from the OpenAPI specification
type ListServicePerimetersResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The pagination token to retrieve the next page of results. If the value is empty, no further results remain.
	Serviceperimeters []ServicePerimeter `json:"servicePerimeters,omitempty"` // List of the Service Perimeter instances.
}

// ServicePerimeterConfig represents the ServicePerimeterConfig schema from the OpenAPI specification
type ServicePerimeterConfig struct {
	Egresspolicies []EgressPolicy `json:"egressPolicies,omitempty"` // List of EgressPolicies to apply to the perimeter. A perimeter may have multiple EgressPolicies, each of which is evaluated separately. Access is granted if any EgressPolicy grants it. Must be empty for a perimeter bridge.
	Ingresspolicies []IngressPolicy `json:"ingressPolicies,omitempty"` // List of IngressPolicies to apply to the perimeter. A perimeter may have multiple IngressPolicies, each of which is evaluated separately. Access is granted if any Ingress Policy grants it. Must be empty for a perimeter bridge.
	Resources []string `json:"resources,omitempty"` // A list of Google Cloud resources that are inside of the service perimeter. Currently only projects and VPCs are allowed. Project format: `projects/{project_number}` VPC network format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`.
	Restrictedservices []string `json:"restrictedServices,omitempty"` // Google Cloud services that are subject to the Service Perimeter restrictions. For example, if `storage.googleapis.com` is specified, access to the storage buckets inside the perimeter must meet the perimeter's access restrictions.
	Vpcaccessibleservices VpcAccessibleServices `json:"vpcAccessibleServices,omitempty"` // Specifies how APIs are allowed to communicate within the Service Perimeter.
	Accesslevels []string `json:"accessLevels,omitempty"` // A list of `AccessLevel` resource names that allow resources within the `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are listed, resources within the perimeter can only be accessed via Google Cloud calls with request origins within the perimeter. Example: `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`. For Service Perimeter Bridge, must be empty.
}

// CommitServicePerimetersResponse represents the CommitServicePerimetersResponse schema from the OpenAPI specification
type CommitServicePerimetersResponse struct {
	Serviceperimeters []ServicePerimeter `json:"servicePerimeters,omitempty"` // List of all the Service Perimeter instances in the Access Policy.
}

// Condition represents the Condition schema from the OpenAPI specification
type Condition struct {
	Requiredaccesslevels []string `json:"requiredAccessLevels,omitempty"` // A list of other access levels defined in the same `Policy`, referenced by resource name. Referencing an `AccessLevel` which does not exist is an error. All access levels listed must be granted for the Condition to be true. Example: "`accessPolicies/MY_POLICY/accessLevels/LEVEL_NAME"`
	Vpcnetworksources []VpcNetworkSource `json:"vpcNetworkSources,omitempty"` // The request must originate from one of the provided VPC networks in Google Cloud. Cannot specify this field together with `ip_subnetworks`.
	Devicepolicy DevicePolicy `json:"devicePolicy,omitempty"` // `DevicePolicy` specifies device specific restrictions necessary to acquire a given access level. A `DevicePolicy` specifies requirements for requests from devices to be granted access levels, it does not do any enforcement on the device. `DevicePolicy` acts as an AND over all specified fields, and each repeated field is an OR over its elements. Any unset fields are ignored. For example, if the proto is { os_type : DESKTOP_WINDOWS, os_type : DESKTOP_LINUX, encryption_status: ENCRYPTED}, then the DevicePolicy will be true for requests originating from encrypted Linux desktops and encrypted Windows desktops.
	Ipsubnetworks []string `json:"ipSubnetworks,omitempty"` // CIDR block IP subnetwork specification. May be IPv4 or IPv6. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly, for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32" is not. The originating IP of a request must be in one of the listed subnets in order for this Condition to be true. If empty, all IP addresses are allowed.
	Members []string `json:"members,omitempty"` // The request must be made by one of the provided user or service accounts. Groups are not supported. Syntax: `user:{emailid}` `serviceAccount:{emailid}` If not specified, a request may come from any user.
	Negate bool `json:"negate,omitempty"` // Whether to negate the Condition. If true, the Condition becomes a NAND over its non-empty fields. Any non-empty field criteria evaluating to false will result in the Condition to be satisfied. Defaults to false.
	Regions []string `json:"regions,omitempty"` // The request must originate from one of the provided countries/regions. Must be valid ISO 3166-1 alpha-2 codes.
}

// GetPolicyOptions represents the GetPolicyOptions schema from the OpenAPI specification
type GetPolicyOptions struct {
	Requestedpolicyversion int `json:"requestedPolicyVersion,omitempty"` // Optional. The maximum policy version that will be used to format the policy. Valid values are 0, 1, and 3. Requests specifying an invalid value will be rejected. Requests for policies with any conditional role bindings must specify version 3. Policies with no conditional role bindings may specify any valid value or leave the field unset. The policy in the response might use the policy version that you specified, or it might use a lower policy version. For example, if you specify version 3, but the policy has no conditional role bindings, the response uses version 1. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
}

// VpcSubNetwork represents the VpcSubNetwork schema from the OpenAPI specification
type VpcSubNetwork struct {
	Network string `json:"network,omitempty"` // Required. Network name. If the network is not part of the organization, the `compute.network.get` permission must be granted to the caller. Format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NETWORK_NAME}` Example: `//compute.googleapis.com/projects/my-project/global/networks/network-1`
	Vpcipsubnetworks []string `json:"vpcIpSubnetworks,omitempty"` // CIDR block IP subnetwork specification. The IP address must be an IPv4 address and can be a public or private IP address. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. If empty, all IP addresses are allowed.
}

// ReplaceServicePerimetersRequest represents the ReplaceServicePerimetersRequest schema from the OpenAPI specification
type ReplaceServicePerimetersRequest struct {
	Etag string `json:"etag,omitempty"` // Optional. The etag for the version of the Access Policy that this replace operation is to be performed on. If, at the time of replace, the etag for the Access Policy stored in Access Context Manager is different from the specified etag, then the replace operation will not be performed and the call will fail. This field is not required. If etag is not provided, the operation will be performed as if a valid etag is provided.
	Serviceperimeters []ServicePerimeter `json:"servicePerimeters,omitempty"` // Required. The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
}

// ReplaceServicePerimetersResponse represents the ReplaceServicePerimetersResponse schema from the OpenAPI specification
type ReplaceServicePerimetersResponse struct {
	Serviceperimeters []ServicePerimeter `json:"servicePerimeters,omitempty"` // List of the Service Perimeter instances.
}

// AccessLevel represents the AccessLevel schema from the OpenAPI specification
type AccessLevel struct {
	Description string `json:"description,omitempty"` // Description of the `AccessLevel` and its use. Does not affect behavior.
	Name string `json:"name,omitempty"` // Resource name for the `AccessLevel`. Format: `accessPolicies/{access_policy}/accessLevels/{access_level}`. The `access_level` component must begin with a letter, followed by alphanumeric characters or `_`. Its maximum length is 50 characters. After you create an `AccessLevel`, you cannot change its `name`.
	Title string `json:"title,omitempty"` // Human readable title. Must be unique within the Policy.
	Basic BasicLevel `json:"basic,omitempty"` // `BasicLevel` is an `AccessLevel` using a set of recommended features.
	Custom CustomLevel `json:"custom,omitempty"` // `CustomLevel` is an `AccessLevel` using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request. See CEL spec at: https://github.com/google/cel-spec
}

// EgressTo represents the EgressTo schema from the OpenAPI specification
type EgressTo struct {
	Resources []string `json:"resources,omitempty"` // A list of resources, currently only projects in the form `projects/`, that are allowed to be accessed by sources defined in the corresponding EgressFrom. A request matches if it contains a resource in this list. If `*` is specified for `resources`, then this EgressTo rule will authorize access to all resources outside the perimeter.
	Externalresources []string `json:"externalResources,omitempty"` // A list of external resources that are allowed to be accessed. Only AWS and Azure resources are supported. For Amazon S3, the supported format is s3://BUCKET_NAME. For Azure Storage, the supported format is azure://myaccount.blob.core.windows.net/CONTAINER_NAME. A request matches if it contains an external resource in this list (Example: s3://bucket/path). Currently '*' is not allowed.
	Operations []ApiOperation `json:"operations,omitempty"` // A list of ApiOperations allowed to be performed by the sources specified in the corresponding EgressFrom. A request matches if it uses an operation/service in this list.
}

// ReplaceAccessLevelsResponse represents the ReplaceAccessLevelsResponse schema from the OpenAPI specification
type ReplaceAccessLevelsResponse struct {
	Accesslevels []AccessLevel `json:"accessLevels,omitempty"` // List of the Access Level instances.
}

// SupportedService represents the SupportedService schema from the OpenAPI specification
type SupportedService struct {
	Title string `json:"title,omitempty"` // The name of the supported product, such as 'Cloud Product API'.
	Availableonrestrictedvip bool `json:"availableOnRestrictedVip,omitempty"` // True if the service is available on the restricted VIP. Services on the restricted VIP typically either support VPC Service Controls or are core infrastructure services required for the functioning of Google Cloud.
	Knownlimitations bool `json:"knownLimitations,omitempty"` // True if the service is supported with some limitations. Check [documentation](https://cloud.google.com/vpc-service-controls/docs/supported-products) for details.
	Name string `json:"name,omitempty"` // The service name or address of the supported service, such as `service.googleapis.com`.
	Supportstage string `json:"supportStage,omitempty"` // The support stage of the service.
	Supportedmethods []MethodSelector `json:"supportedMethods,omitempty"` // The list of the supported methods. This field exists only in response to GetSupportedService
}

// EgressPolicy represents the EgressPolicy schema from the OpenAPI specification
type EgressPolicy struct {
	Egressto EgressTo `json:"egressTo,omitempty"` // Defines the conditions under which an EgressPolicy matches a request. Conditions are based on information about the ApiOperation intended to be performed on the `resources` specified. Note that if the destination of the request is also protected by a ServicePerimeter, then that ServicePerimeter must have an IngressPolicy which allows access in order for this request to succeed. The request must match `operations` AND `resources` fields in order to be allowed egress out of the perimeter.
	Egressfrom EgressFrom `json:"egressFrom,omitempty"` // Defines the conditions under which an EgressPolicy matches a request. Conditions based on information about the source of the request. Note that if the destination of the request is also protected by a ServicePerimeter, then that ServicePerimeter must have an IngressPolicy which allows access in order for this request to succeed.
}

// EgressSource represents the EgressSource schema from the OpenAPI specification
type EgressSource struct {
	Accesslevel string `json:"accessLevel,omitempty"` // An AccessLevel resource name that allows protected resources inside the ServicePerimeters to access outside the ServicePerimeter boundaries. AccessLevels listed must be in the same policy as this ServicePerimeter. Referencing a nonexistent AccessLevel will cause an error. If an AccessLevel name is not specified, only resources within the perimeter can be accessed through Google Cloud calls with request origins within the perimeter. Example: `accessPolicies/MY_POLICY/accessLevels/MY_LEVEL`. If a single `*` is specified for `access_level`, then all EgressSources will be allowed.
}

// ServicePerimeter represents the ServicePerimeter schema from the OpenAPI specification
type ServicePerimeter struct {
	Title string `json:"title,omitempty"` // Human readable title. Must be unique within the Policy.
	Useexplicitdryrunspec bool `json:"useExplicitDryRunSpec,omitempty"` // Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists for all Service Perimeters, and that spec is identical to the status for those Service Perimeters. When this flag is set, it inhibits the generation of the implicit spec, thereby allowing the user to explicitly provide a configuration ("spec") to use in a dry-run version of the Service Perimeter. This allows the user to test changes to the enforced config ("status") without actually enforcing them. This testing is done through analyzing the differences between currently enforced and suggested restrictions. use_explicit_dry_run_spec must bet set to True if any of the fields in the spec are set to non-default values.
	Description string `json:"description,omitempty"` // Description of the `ServicePerimeter` and its use. Does not affect behavior.
	Name string `json:"name,omitempty"` // Resource name for the `ServicePerimeter`. Format: `accessPolicies/{access_policy}/servicePerimeters/{service_perimeter}`. The `service_perimeter` component must begin with a letter, followed by alphanumeric characters or `_`. After you create a `ServicePerimeter`, you cannot change its `name`.
	Perimetertype string `json:"perimeterType,omitempty"` // Perimeter type indicator. A single project or VPC network is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, the restricted service list as well as access level lists must be empty.
	Spec ServicePerimeterConfig `json:"spec,omitempty"` // `ServicePerimeterConfig` specifies a set of Google Cloud resources that describe specific Service Perimeter configuration.
	Status ServicePerimeterConfig `json:"status,omitempty"` // `ServicePerimeterConfig` specifies a set of Google Cloud resources that describe specific Service Perimeter configuration.
}

// CommitServicePerimetersRequest represents the CommitServicePerimetersRequest schema from the OpenAPI specification
type CommitServicePerimetersRequest struct {
	Etag string `json:"etag,omitempty"` // Optional. The etag for the version of the Access Policy that this commit operation is to be performed on. If, at the time of commit, the etag for the Access Policy stored in Access Context Manager is different from the specified etag, then the commit operation will not be performed and the call will fail. This field is not required. If etag is not provided, the operation will be performed as if a valid etag is provided.
}

// AuditConfig represents the AuditConfig schema from the OpenAPI specification
type AuditConfig struct {
	Auditlogconfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"` // The configuration for logging of each type of permission.
	Service string `json:"service,omitempty"` // Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
}

// AccessPolicy represents the AccessPolicy schema from the OpenAPI specification
type AccessPolicy struct {
	Parent string `json:"parent,omitempty"` // Required. The parent of this `AccessPolicy` in the Cloud Resource Hierarchy. Currently immutable once created. Format: `organizations/{organization_id}`
	Scopes []string `json:"scopes,omitempty"` // The scopes of the AccessPolicy. Scopes define which resources a policy can restrict and where its resources can be referenced. For example, policy A with `scopes=["folders/123"]` has the following behavior: - ServicePerimeter can only restrict projects within `folders/123`. - ServicePerimeter within policy A can only reference access levels defined within policy A. - Only one policy can include a given scope; thus, attempting to create a second policy which includes `folders/123` will result in an error. If no scopes are provided, then any resource within the organization can be restricted. Scopes cannot be modified after a policy is created. Policies can only have a single scope. Format: list of `folders/{folder_number}` or `projects/{project_number}`
	Title string `json:"title,omitempty"` // Required. Human readable title. Does not affect behavior.
	Etag string `json:"etag,omitempty"` // Output only. An opaque identifier for the current version of the `AccessPolicy`. This will always be a strongly validated etag, meaning that two Access Polices will be identical if and only if their etags are identical. Clients should not expect this to be in any specific format.
	Name string `json:"name,omitempty"` // Output only. Resource name of the `AccessPolicy`. Format: `accessPolicies/{access_policy}`
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
}

// TestIamPermissionsResponse represents the TestIamPermissionsResponse schema from the OpenAPI specification
type TestIamPermissionsResponse struct {
	Permissions []string `json:"permissions,omitempty"` // A subset of `TestPermissionsRequest.permissions` that the caller is allowed.
}

// ListGcpUserAccessBindingsResponse represents the ListGcpUserAccessBindingsResponse schema from the OpenAPI specification
type ListGcpUserAccessBindingsResponse struct {
	Gcpuseraccessbindings []GcpUserAccessBinding `json:"gcpUserAccessBindings,omitempty"` // GcpUserAccessBinding
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to get the next page of items. If blank, there are no more items.
}

// ReplaceAccessLevelsRequest represents the ReplaceAccessLevelsRequest schema from the OpenAPI specification
type ReplaceAccessLevelsRequest struct {
	Accesslevels []AccessLevel `json:"accessLevels,omitempty"` // Required. The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	Etag string `json:"etag,omitempty"` // Optional. The etag for the version of the Access Policy that this replace operation is to be performed on. If, at the time of replace, the etag for the Access Policy stored in Access Context Manager is different from the specified etag, then the replace operation will not be performed and the call will fail. This field is not required. If etag is not provided, the operation will be performed as if a valid etag is provided.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
}

// TestIamPermissionsRequest represents the TestIamPermissionsRequest schema from the OpenAPI specification
type TestIamPermissionsRequest struct {
	Permissions []string `json:"permissions,omitempty"` // The set of permissions to check for the `resource`. Permissions with wildcards (such as `*` or `storage.*`) are not allowed. For more information see [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
}

// CancelOperationRequest represents the CancelOperationRequest schema from the OpenAPI specification
type CancelOperationRequest struct {
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
	Operations []Operation `json:"operations,omitempty"` // A list of operations that matches the specified filter in the request.
}

// CustomLevel represents the CustomLevel schema from the OpenAPI specification
type CustomLevel struct {
	Expr Expr `json:"expr,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
}

// ListAccessLevelsResponse represents the ListAccessLevelsResponse schema from the OpenAPI specification
type ListAccessLevelsResponse struct {
	Accesslevels []AccessLevel `json:"accessLevels,omitempty"` // List of the Access Level instances.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The pagination token to retrieve the next page of results. If the value is empty, no further results remain.
}

// GcpUserAccessBindingOperationMetadata represents the GcpUserAccessBindingOperationMetadata schema from the OpenAPI specification
type GcpUserAccessBindingOperationMetadata struct {
}

// DevicePolicy represents the DevicePolicy schema from the OpenAPI specification
type DevicePolicy struct {
	Requirecorpowned bool `json:"requireCorpOwned,omitempty"` // Whether the device needs to be corp owned.
	Requirescreenlock bool `json:"requireScreenlock,omitempty"` // Whether or not screenlock is required for the DevicePolicy to be true. Defaults to `false`.
	Alloweddevicemanagementlevels []string `json:"allowedDeviceManagementLevels,omitempty"` // Allowed device management levels, an empty list allows all management levels.
	Allowedencryptionstatuses []string `json:"allowedEncryptionStatuses,omitempty"` // Allowed encryptions statuses, an empty list allows all statuses.
	Osconstraints []OsConstraint `json:"osConstraints,omitempty"` // Allowed OS versions, an empty list allows all types and all versions.
	Requireadminapproval bool `json:"requireAdminApproval,omitempty"` // Whether the device needs to be approved by the customer admin.
}

// IngressFrom represents the IngressFrom schema from the OpenAPI specification
type IngressFrom struct {
	Identitytype string `json:"identityType,omitempty"` // Specifies the type of identities that are allowed access from outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
	Sources []IngressSource `json:"sources,omitempty"` // Sources that this IngressPolicy authorizes access from.
	Identities []string `json:"identities,omitempty"` // A list of identities that are allowed access through this ingress policy, in the format of `user:{email_id}` or `serviceAccount:{email_id}`.
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Auditconfigs []AuditConfig `json:"auditConfigs,omitempty"` // Specifies cloud audit logging configuration for this policy.
	Bindings []Binding `json:"bindings,omitempty"` // Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Etag string `json:"etag,omitempty"` // `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Version int `json:"version,omitempty"` // Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
}

// AuthorizedOrgsDesc represents the AuthorizedOrgsDesc schema from the OpenAPI specification
type AuthorizedOrgsDesc struct {
	Authorizationtype string `json:"authorizationType,omitempty"` // A granular control type for authorization levels. Valid value is `AUTHORIZATION_TYPE_TRUST`.
	Name string `json:"name,omitempty"` // Resource name for the `AuthorizedOrgsDesc`. Format: `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`. The `authorized_orgs_desc` component must begin with a letter, followed by alphanumeric characters or `_`. After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.
	Orgs []string `json:"orgs,omitempty"` // The list of organization ids in this AuthorizedOrgsDesc. Format: `organizations/` Example: `organizations/123456`
	Assettype string `json:"assetType,omitempty"` // The asset type of this authorized orgs desc. Valid values are `ASSET_TYPE_DEVICE`, and `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	Authorizationdirection string `json:"authorizationDirection,omitempty"` // The direction of the authorization relationship between this organization and the organizations listed in the `orgs` field. The valid values for this field include the following: `AUTHORIZATION_DIRECTION_FROM`: Allows this organization to evaluate traffic in the organizations listed in the `orgs` field. `AUTHORIZATION_DIRECTION_TO`: Allows the organizations listed in the `orgs` field to evaluate the traffic in this organization. For the authorization relationship to take effect, all of the organizations must authorize and specify the appropriate relationship direction. For example, if organization A authorized organization B and C to evaluate its traffic, by specifying `AUTHORIZATION_DIRECTION_TO` as the authorization direction, organizations B and C must specify `AUTHORIZATION_DIRECTION_FROM` as the authorization direction in their `AuthorizedOrgsDesc` resource.
}

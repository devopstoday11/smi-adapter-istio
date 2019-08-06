// Code generated by protoc-gen-go.
// source: https://github.com/istio/api/blob/ece3d93c51b41adf415b38366ea642eb7c1bea93/rbac/v1alpha1/rbac.proto

// Istio RBAC (Role Based Access Control) defines ServiceRole and ServiceRoleBinding
// objects.
//
// A ServiceRole specification includes a list of rules (permissions). Each rule has
// the following standard fields:
//
//   * services: a list of services.
//   * methods: A list of HTTP methods. You can set the value to `*` to include all HTTP methods.
//              This field should not be set for TCP services. The policy will be ignored.
//              For gRPC services, only `POST` is allowed; other methods will result in denying services.
//   * paths: HTTP paths or gRPC methods. Note that gRPC methods should be
//     presented in the form of "/packageName.serviceName/methodName" and are case sensitive.
//
// In addition to the standard fields, operators can also use custom keys in the `constraints` field,
// the supported keys are listed in the "constraints and properties" page.
//
// Below is an example of ServiceRole object "product-viewer", which has "read" ("GET" and "HEAD")
// access to "products.svc.cluster.local" service at versions "v1" and "v2". "path" is not specified,
// so it applies to any path in the service.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ServiceRole
// metadata:
//   name: products-viewer
//   namespace: default
// spec:
//   rules:
//   - services: ["products.svc.cluster.local"]
//     methods: ["GET", "HEAD"]
//     constraints:
//     - key: "destination.labels[version]"
//       values: ["v1", "v2"]
// ```
//
// A ServiceRoleBinding specification includes two parts:
//
//  * The `roleRef` field that refers to a ServiceRole object in the same namespace.
//  * A list of `subjects` that are assigned the roles.
//
// In addition to a simple `user` field, operators can also use custom keys in the `properties` field,
// the supported keys are listed in the "constraints and properties" page.
//
// Below is an example of ServiceRoleBinding object "test-binding-products", which binds two subjects
// to ServiceRole "product-viewer":
//
//   * User "alice@yahoo.com"
//   * Services in "abc" namespace.
//
// ```yaml
// apiVersion: "rbac.istio.io/v1alpha1"
// kind: ServiceRoleBinding
// metadata:
//   name: test-binding-products
//   namespace: default
// spec:
//   subjects:
//   - user: alice@yahoo.com
//   - properties:
//       source.namespace: "abc"
//   roleRef:
//     kind: ServiceRole
//     name: "products-viewer"
// ```

package v1alpha1

// AccessRule defines a permission to access a list of services.
type AccessRule struct {
	// Required. A list of service names.
	// Exact match, prefix match, and suffix match are supported for service names.
	// For example, the service name "bookstore.mtv.cluster.local" matches
	// "bookstore.mtv.cluster.local" (exact match), or "bookstore*" (prefix match),
	// or "*.mtv.cluster.local" (suffix match).
	// If set to ["*"], it refers to all services in the namespace.
	Services []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	// Optional. A list of HTTP hosts. This is matched against the HOST header in
	// a HTTP request. Exact match, prefix match and suffix match are supported.
	// For example, the host "test.abc.com" matches "test.abc.com" (exact match),
	// or "*.abc.com" (prefix match), or "test.abc.*" (suffix match).
	// If not specified, it matches to any host.
	// This field should not be set for TCP services. The policy will be ignored.
	Hosts []string `protobuf:"bytes,5,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Optional. A list of HTTP hosts that must not be matched.
	NotHosts []string `protobuf:"bytes,6,rep,name=not_hosts,json=notHosts,proto3" json:"not_hosts,omitempty"`
	// Optional. A list of HTTP paths or gRPC methods.
	// gRPC methods must be presented as fully-qualified name in the form of
	// "/packageName.serviceName/methodName" and are case sensitive.
	// Exact match, prefix match, and suffix match are supported. For example,
	// the path "/books/review" matches "/books/review" (exact match),
	// or "/books/*" (prefix match), or "*/review" (suffix match).
	// If not specified, it matches to any path.
	// This field should not be set for TCP services. The policy will be ignored.
	Paths []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	// Optional. A list of HTTP paths or gRPC methods that must not be matched.
	NotPaths []string `protobuf:"bytes,7,rep,name=not_paths,json=notPaths,proto3" json:"not_paths,omitempty"`
	// Optional. A list of HTTP methods (e.g., "GET", "POST").
	// If not specified or specified as "*", it matches to any methods.
	// This field should not be set for TCP services. The policy will be ignored.
	// For gRPC services, only `POST` is allowed; other methods will result in denying services.
	Methods []string `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	// Optional. A list of HTTP methods that must not be matched.
	// Note: It's an error to set methods and not_methods at the same time.
	NotMethods []string `protobuf:"bytes,8,rep,name=not_methods,json=notMethods,proto3" json:"not_methods,omitempty"`
	// Optional. A list of port numbers of the request. If not specified, it matches
	// to any port number.
	// Note: It's an error to set ports and not_ports at the same time.
	Ports []int32 `protobuf:"varint,9,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	// Optional.  A list of port numbers that must not be matched.
	// Note: It's an error to set ports and not_ports at the same time.
	NotPorts []int32 `protobuf:"varint,10,rep,packed,name=not_ports,json=notPorts,proto3" json:"not_ports,omitempty"`
	// Optional. Extra constraints in the ServiceRole specification.
	Constraints []*AccessRule_Constraint `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints,omitempty"`
}

// Definition of a custom constraint. The supported keys are listed in the "constraint and properties" page.
type AccessRule_Constraint struct {
	// Key of the constraint.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// List of valid values for the constraint.
	// Exact match, prefix match, and suffix match are supported.
	// For example, the value "v1alpha2" matches "v1alpha2" (exact match),
	// or "v1*" (prefix match), or "*alpha2" (suffix match).
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

// Subject defines an identity. The identity is either a user or identified by a set of `properties`.
// The supported keys in `properties` are listed in "constraint and properties" page.
type Subject struct {
	// Optional. The user name/ID that the subject represents.
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Optional. A list of subject names. This is matched to the
	// `source.principal` attribute. If one of subject names is "*", it matches to a subject with any name.
	// Prefix and suffix matches are supported.
	Names []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
	// Optional. A list of subject names that must not be matched.
	NotNames []string `protobuf:"bytes,5,rep,name=not_names,json=notNames,proto3" json:"not_names,omitempty"`
	// Optional. The group that the subject belongs to.
	// Deprecated. Use groups and not_groups instead.
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"` // Deprecated: Do not use.
	// Optional. A list of groups that the subject represents. This is matched to the
	// `request.auth.claims[groups]` attribute. If not specified, it applies to any groups.
	Groups []string `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty"`
	// Optional. A list of groups that must not be matched.
	NotGroups []string `protobuf:"bytes,7,rep,name=not_groups,json=notGroups,proto3" json:"not_groups,omitempty"`
	// Optional. A list of namespaces that the subject represents. This is matched to
	// the `source.namespace` attribute. If not specified, it applies to any namespaces.
	Namespaces []string `protobuf:"bytes,8,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// Optional. A list of namespaces that must not be matched.
	NotNamespaces []string `protobuf:"bytes,9,rep,name=not_namespaces,json=notNamespaces,proto3" json:"not_namespaces,omitempty"`
	// Optional. A list of IP address or CIDR ranges that the subject represents.
	// E.g. 192.168.100.2 or 10.1.0.0/16. If not specified, it applies to any IP addresses.
	Ips []string `protobuf:"bytes,10,rep,name=ips,proto3" json:"ips,omitempty"`
	// Optional. A list of IP addresses or CIDR ranges that must not be matched.
	NotIps []string `protobuf:"bytes,11,rep,name=not_ips,json=notIps,proto3" json:"not_ips,omitempty"`
	// Optional. The set of properties that identify the subject.
	Properties map[string]string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

// RoleRef refers to a role object.
type RoleRef struct {
	// Required. The type of the role being referenced.
	// Currently, "ServiceRole" is the only supported value for "kind".
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Required. The name of the ServiceRole object being referenced.
	// The ServiceRole object must be in the same namespace as the ServiceRoleBinding object.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}
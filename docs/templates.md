## K8sAllowedRepos

> Requires container images to begin with a string from the specified list.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sAllowedRepos
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # repos <array>: The list of prefixes a container image is allowed to have.
    repos:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### repo-is-openpolicyagent
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sAllowedRepos
metadata:
  name: repo-is-openpolicyagent
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
    namespaces:
    - default
  parameters:
    repos:
    - openpolicyagent/
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: nginx-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: nginx-disallowed
spec:
  containers: []
  initContainers:
  - image: nginx
    name: nginx
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: nginx-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
  initContainers:
  - image: nginx
    name: nginx
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
```
  </details>

---

## K8sBlockEndpointEditDefaultRole

> Many Kubernetes installations by default have a system:aggregate-to-edit ClusterRole which does not properly restrict access to editing Endpoints. This ConstraintTemplate forbids the system:aggregate-to-edit ClusterRole from granting permission to create/patch/update Endpoints.
ClusterRole/system:aggregate-to-edit should not allow Endpoint edit permissions due to CVE-2021-25740, Endpoint & EndpointSlice permissions allow cross-Namespace forwarding, https://github.com/kubernetes/kubernetes/issues/103675

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockEndpointEditDefaultRole
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### block-endpoint-edit-default-role
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockEndpointEditDefaultRole
metadata:
  name: block-endpoint-edit-default-role
spec:
  match:
    kinds:
    - apiGroups:
      - rbac.authorization.k8s.io
      kinds:
      - ClusterRole
```
  #### Allowed
  ```yaml
  apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  creationTimestamp: null
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: system:aggregate-to-edit
rules:
- apiGroups:
  - ""
  resources:
  - pods/attach
  - pods/exec
  - pods/portforward
  - pods/proxy
  - secrets
  - services/proxy
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - impersonate
- apiGroups:
  - ""
  resources:
  - pods
  - pods/attach
  - pods/exec
  - pods/portforward
  - pods/proxy
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - ""
  resources:
  - configmaps
  - persistentvolumeclaims
  - replicationcontrollers
  - replicationcontrollers/scale
  - secrets
  - serviceaccounts
  - services
  - services/proxy
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - apps
  resources:
  - daemonsets
  - deployments
  - deployments/rollback
  - deployments/scale
  - replicasets
  - replicasets/scale
  - statefulsets
  - statefulsets/scale
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - autoscaling
  resources:
  - horizontalpodautoscalers
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - batch
  resources:
  - cronjobs
  - jobs
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - extensions
  resources:
  - daemonsets
  - deployments
  - deployments/rollback
  - deployments/scale
  - ingresses
  - networkpolicies
  - replicasets
  - replicasets/scale
  - replicationcontrollers/scale
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - policy
  resources:
  - poddisruptionbudgets
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - networking.k8s.io
  resources:
  - ingresses
  - networkpolicies
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
```
  #### Disallowed
  ```yaml
  apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  creationTimestamp: null
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
  name: system:aggregate-to-edit
rules:
- apiGroups:
  - ""
  resources:
  - pods/attach
  - pods/exec
  - pods/portforward
  - pods/proxy
  - secrets
  - services/proxy
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - impersonate
- apiGroups:
  - ""
  resources:
  - pods
  - pods/attach
  - pods/exec
  - pods/portforward
  - pods/proxy
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - ""
  resources:
  - configmaps
  - persistentvolumeclaims
  - replicationcontrollers
  - replicationcontrollers/scale
  - secrets
  - serviceaccounts
  - services
  - services/proxy
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
- apiGroups:
  - apps
  resources:
  - daemonsets
  - deployments
  - deployments/rollback
  - deployments/scale
  - endpoints
  - replicasets
  - replicasets/scale
  - statefulsets
  - statefulsets/scale
  verbs:
  - create
  - delete
  - deletecollection
  - patch
  - update
```
  </details>

---

## K8sBlockNodePort

> Disallows all Services with type NodePort.
https://kubernetes.io/docs/concepts/services-networking/service/#nodeport

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockNodePort
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### block-node-port
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockNodePort
metadata:
  name: block-node-port
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Service
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: my-service-disallowed
spec:
  ports:
  - nodePort: 30007
    port: 80
    targetPort: 80
  type: NodePort
```
  </details>

---

## K8sBlockWildcardIngress

> Users should not be able to create Ingresses with a blank or wildcard (*) hostname since that would enable them to intercept traffic for other services in the cluster, even if they don't have access to those services.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockWildcardIngress
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### block-wildcard-ingress
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sBlockWildcardIngress
metadata:
  name: block-wildcard-ingress
spec:
  match:
    kinds:
    - apiGroups:
      - extensions
      - networking.k8s.io
      kinds:
      - Ingress
```
  #### Allowed
  ```yaml
  apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: non-wildcard-ingress
spec:
  rules:
  - host: myservice.example.com
    http:
      paths:
      - backend:
          service:
            name: example
            port:
              number: 80
        path: /
        pathType: Prefix
```
  #### Disallowed
  ```yaml
  apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: wildcard-ingress
spec:
  rules:
  - host: ""
    http:
      paths:
      - backend:
          service:
            name: example
            port:
              number: 80
        path: /
        pathType: Prefix
```
  
  ```yaml
  apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: wildcard-ingress
spec:
  rules:
  - http:
      paths:
      - backend:
          service:
            name: example
            port:
              number: 80
        path: /
        pathType: Prefix
```
  
  ```yaml
  apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: wildcard-ingress
spec:
  rules:
  - host: '*.example.com'
    http:
      paths:
      - backend:
          service:
            name: example
            port:
              number: 80
        path: /
        pathType: Prefix
  - host: valid.example.com
    http:
      paths:
      - backend:
          service:
            name: example
            port:
              number: 80
        path: /
        pathType: Prefix
```
  </details>

---

## K8sContainerLimits

> Requires containers to have memory and CPU limits set and constrains limits to be within the specified maximum values.
https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerLimits
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # cpu <string>: The maximum allowed cpu limit on a Pod, exclusive.
    cpu: <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # memory <string>: The maximum allowed memory limit on a Pod, exclusive.
    memory: <string>
```

<details>
  <summary>Examples</summary>

  #### container-must-have-limits
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerLimits
metadata:
  name: container-must-have-limits
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    cpu: 200m
    memory: 1Gi
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 100m
        memory: 1Gi
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 100m
        memory: 2Gi
```
  </details>

---

## K8sContainerRatios

> Sets a maximum ratio for container resource limits to requests.
https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerRatios
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # cpuRatio <string>: The maximum allowed ratio of `resources.limits.cpu` to
    # `resources.requests.cpu` on a container. If not specified, equal to
    # `ratio`.
    cpuRatio: <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # ratio <string>: The maximum allowed ratio of `resources.limits` to
    # `resources.requests` on a container.
    ratio: <string>
```

<details>
  <summary>Examples</summary>

  #### container-must-meet-ratio
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerRatios
metadata:
  name: container-must-meet-ratio
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    ratio: "2"
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 200m
        memory: 200Mi
      requests:
        cpu: 100m
        memory: 100Mi
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 800m
        memory: 2Gi
      requests:
        cpu: 100m
        memory: 100Mi
```
  

  #### container-must-meet-memory-and-cpu-ratio
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerRatios
metadata:
  name: container-must-meet-memory-and-cpu-ratio
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    cpuRatio: "10"
    ratio: "1"
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: "4"
        memory: 2Gi
      requests:
        cpu: "1"
        memory: 2Gi
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: "4"
        memory: 2Gi
      requests:
        cpu: 100m
        memory: 2Gi
```
  </details>

---

## K8sContainerRequests

> Requires containers to have memory and CPU requests set and constrains requests to be within the specified maximum values.
https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerRequests
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # cpu <string>: The maximum allowed cpu request on a Pod, exclusive.
    cpu: <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # memory <string>: The maximum allowed memory request on a Pod, exclusive.
    memory: <string>
```

<details>
  <summary>Examples</summary>

  #### container-must-have-requests
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sContainerRequests
metadata:
  name: container-must-have-requests
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    cpu: 200m
    memory: 1Gi
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      requests:
        cpu: 100m
        memory: 1Gi
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      requests:
        cpu: 100m
        memory: 2Gi
```
  </details>

---

## K8sDisallowAnonymous

> Disallows associating ClusterRole and Role resources to the system:anonymous user and system:unauthenticated group.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sDisallowAnonymous
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedRoles <array>: The list of ClusterRoles and Roles that may be
    # associated with the `system:unauthenticated` group and `system:anonymous`
    # user.
    allowedRoles:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### no-anonymous
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sDisallowAnonymous
metadata:
  name: no-anonymous
spec:
  match:
    kinds:
    - apiGroups:
      - rbac.authorization.k8s.io
      kinds:
      - ClusterRoleBinding
    - apiGroups:
      - rbac.authorization.k8s.io
      kinds:
      - RoleBinding
  parameters:
    allowedRoles:
    - cluster-role-1
```
  #### Allowed
  ```yaml
  apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-role-binding-1
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-role-1
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:authenticated
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:unauthenticated
```
  #### Disallowed
  ```yaml
  apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cluster-role-binding-2
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-role-2
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:authenticated
- apiGroup: rbac.authorization.k8s.io
  kind: Group
  name: system:unauthenticated
```
  </details>

---

## K8sDisallowedTags

> Requires container images to have an image tag different from the ones in the specified list.
https://kubernetes.io/docs/concepts/containers/images/#image-names

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sDisallowedTags
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # tags <array>: Disallowed container image tags.
    tags:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### container-image-must-not-have-latest-tag
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sDisallowedTags
metadata:
  name: container-image-must-not-have-latest-tag
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
    namespaces:
    - default
  parameters:
    exemptImages:
    - openpolicyagent/opa-exp:latest
    - openpolicyagent/opa-exp2:latest
    tags:
    - latest
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-exempt-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa-exp:latest
    name: opa-exp
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/init:v1
    name: opa-init
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa-exp2:latest
    name: opa-exp2
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa
    name: opa
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-disallowed-2
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:latest
    name: opa
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-disallowed-3
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa-exp:latest
    name: opa
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/init:latest
    name: opa-init
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa-exp2:latest
    name: opa-exp2
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/monitor:latest
    name: opa-monitor
```
  </details>

---

## K8sExternalIPs

> Restricts Service externalIPs to an allowed list of IP addresses.
https://kubernetes.io/docs/concepts/services-networking/service/#external-ips

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sExternalIPs
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedIPs <array>: An allow-list of external IP addresses.
    allowedIPs:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### external-ips
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sExternalIPs
metadata:
  name: external-ips
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Service
  parameters:
    allowedIPs:
    - 203.0.113.0
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: allowed-external-ip
spec:
  externalIPs:
  - 203.0.113.0
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    app: MyApp
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: disallowed-external-ip
spec:
  externalIPs:
  - 1.1.1.1
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    app: MyApp
```
  </details>

---

## K8sHttpsOnly

> Requires Ingress resources to be HTTPS only.
Ingress resources must: - include a valid TLS configuration - include the `kubernetes.io/ingress.allow-http` annotation, set to
  `false`.

https://kubernetes.io/docs/concepts/services-networking/ingress/#tls

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sHttpsOnly
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### ingress-https-only
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sHttpsOnly
metadata:
  name: ingress-https-only
spec:
  match:
    kinds:
    - apiGroups:
      - extensions
      - networking.k8s.io
      kinds:
      - Ingress
```
  #### Allowed
  ```yaml
  apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.allow-http: "false"
  name: ingress-demo-disallowed
spec:
  rules:
  - host: example-host.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
  tls:
  - {}
```
  #### Disallowed
  ```yaml
  apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-demo-disallowed
spec:
  rules:
  - host: example-host.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
```
  </details>

---

## K8sImageDigests

> Requires container images to contain a digest.
https://kubernetes.io/docs/concepts/containers/images/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sImageDigests
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### container-image-must-have-digest
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sImageDigests
metadata:
  name: container-image-must-have-digest
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
    namespaces:
    - default
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2@sha256:04ff8fce2afd1a3bc26260348e5b290e8d945b1fad4b4c16d22834c2f3a1814a
    name: opa
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
```
  </details>

---

## K8sPSPAllowPrivilegeEscalationContainer

> Controls restricting escalation to root privileges. Corresponds to the `allowPrivilegeEscalation` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#privilege-escalation

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAllowPrivilegeEscalationContainer
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-allow-privilege-escalation-container
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAllowPrivilegeEscalationContainer
metadata:
  name: psp-allow-privilege-escalation-container
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-privilege-escalation
  name: nginx-privilege-escalation-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      allowPrivilegeEscalation: false
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-privilege-escalation
  name: nginx-privilege-escalation-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      allowPrivilegeEscalation: true
```
  </details>

---

## K8sPSPAllowedUsers

> Controls the user and group IDs of the container and some volumes. Corresponds to the `runAsUser`, `runAsGroup`, `supplementalGroups`, and `fsGroup` fields in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#users-and-groups

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAllowedUsers
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # fsGroup <object>: Controls the fsGroup values that are allowed in a Pod
    # or container-level SecurityContext.
    fsGroup:
      # ranges <array>: A list of group ID ranges affected by the rule.
      ranges:
        # <list item: object>: The range of group IDs affected by the rule.
        - # max <integer>: The maximum group ID in the range, inclusive.
          max: <integer>
          # min <integer>: The minimum group ID in the range, inclusive.
          min: <integer>
      # rule <string>: A strategy for applying the fsGroup restriction.
      # Allowed Values: MustRunAs, MayRunAs, RunAsAny
      rule: <string>
    # runAsGroup <object>: Controls which group ID values are allowed in a Pod
    # or container-level SecurityContext.
    runAsGroup:
      # ranges <array>: A list of group ID ranges affected by the rule.
      ranges:
        # <list item: object>: The range of group IDs affected by the rule.
        - # max <integer>: The maximum group ID in the range, inclusive.
          max: <integer>
          # min <integer>: The minimum group ID in the range, inclusive.
          min: <integer>
      # rule <string>: A strategy for applying the runAsGroup restriction.
      # Allowed Values: MustRunAs, MayRunAs, RunAsAny
      rule: <string>
    # runAsUser <object>: Controls which user ID values are allowed in a Pod or
    # container-level SecurityContext.
    runAsUser:
      # ranges <array>: A list of user ID ranges affected by the rule.
      ranges:
        # <list item: object>: The range of user IDs affected by the rule.
        - # max <integer>: The maximum user ID in the range, inclusive.
          max: <integer>
          # min <integer>: The minimum user ID in the range, inclusive.
          min: <integer>
      # rule <string>: A strategy for applying the runAsUser restriction.
      # Allowed Values: MustRunAs, MustRunAsNonRoot, RunAsAny
      rule: <string>
    # supplementalGroups <object>: Controls the supplementalGroups values that
    # are allowed in a Pod or container-level SecurityContext.
    supplementalGroups:
      # ranges <array>: A list of group ID ranges affected by the rule.
      ranges:
        # <list item: object>: The range of group IDs affected by the rule.
        - # max <integer>: The maximum group ID in the range, inclusive.
          max: <integer>
          # min <integer>: The minimum group ID in the range, inclusive.
          min: <integer>
      # rule <string>: A strategy for applying the supplementalGroups
      # restriction.
      # Allowed Values: MustRunAs, MayRunAs, RunAsAny
      rule: <string>
```

<details>
  <summary>Examples</summary>

  #### psp-pods-allowed-user-ranges
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAllowedUsers
metadata:
  name: psp-pods-allowed-user-ranges
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    fsGroup:
      ranges:
      - max: 200
        min: 100
      rule: MustRunAs
    runAsGroup:
      ranges:
      - max: 200
        min: 100
      rule: MustRunAs
    runAsUser:
      ranges:
      - max: 200
        min: 100
      rule: MustRunAs
    supplementalGroups:
      ranges:
      - max: 200
        min: 100
      rule: MustRunAs
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-users
  name: nginx-users-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      runAsGroup: 199
      runAsUser: 199
  securityContext:
    fsGroup: 199
    supplementalGroups:
    - 199
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-users
  name: nginx-users-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      runAsGroup: 250
      runAsUser: 250
  securityContext:
    fsGroup: 250
    supplementalGroups:
    - 250
```
  </details>

---

## K8sPSPAppArmor

> Configures an allow-list of AppArmor profiles for use by containers. This corresponds to specific annotations applied to a PodSecurityPolicy. For information on AppArmor, see https://kubernetes.io/docs/tutorials/clusters/apparmor/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAppArmor
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedProfiles <array>: An array of AppArmor profiles. Examples:
    # `runtime/default`, `unconfined`.
    allowedProfiles:
      - <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-apparmor
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAppArmor
metadata:
  name: psp-apparmor
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    allowedProfiles:
    - runtime/default
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    container.apparmor.security.beta.kubernetes.io/nginx: runtime/default
  labels:
    app: nginx-apparmor
  name: nginx-apparmor-allowed
spec:
  containers:
  - image: nginx
    name: nginx
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    container.apparmor.security.beta.kubernetes.io/nginx: unconfined
  labels:
    app: nginx-apparmor
  name: nginx-apparmor-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
```
  </details>

---

## K8sPSPAutomountServiceAccountTokenPod

> Controls the ability of any Pod to enable automountServiceAccountToken.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAutomountServiceAccountTokenPod
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    <object>
```

<details>
  <summary>Examples</summary>

  #### psp-automount-serviceaccount-token-pod
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPAutomountServiceAccountTokenPod
metadata:
  name: psp-automount-serviceaccount-token-pod
spec:
  match:
    excludedNamespaces:
    - kube-system
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-not-automountserviceaccounttoken
  name: nginx-automountserviceaccounttoken-allowed
spec:
  automountServiceAccountToken: false
  containers:
  - image: nginx
    name: nginx
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-automountserviceaccounttoken
  name: nginx-automountserviceaccounttoken-disallowed
spec:
  automountServiceAccountToken: true
  containers:
  - image: nginx
    name: nginx
```
  </details>

---

## K8sPSPCapabilities

> Controls Linux capabilities on containers. Corresponds to the `allowedCapabilities` and `requiredDropCapabilities` fields in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#capabilities

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPCapabilities
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedCapabilities <array>: A list of Linux capabilities that can be
    # added to a container.
    allowedCapabilities:
      - <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # requiredDropCapabilities <array>: A list of Linux capabilities that are
    # required to be dropped from a container.
    requiredDropCapabilities:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### capabilities-demo
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPCapabilities
metadata:
  name: capabilities-demo
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
    namespaces:
    - default
  parameters:
    allowedCapabilities:
    - something
    requiredDropCapabilities:
    - must_drop
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-allowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
    securityContext:
      capabilities:
        add:
        - something
        drop:
        - must_drop
        - another_one
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    owner: me.agilebank.demo
  name: opa-disallowed
spec:
  containers:
  - args:
    - run
    - --server
    - --addr=localhost:8080
    image: openpolicyagent/opa:0.9.2
    name: opa
    resources:
      limits:
        cpu: 100m
        memory: 30Mi
    securityContext:
      capabilities:
        add:
        - disallowedcapability
```
  </details>

---

## K8sPSPFSGroup

> Controls allocating an FSGroup that owns the Pod's volumes. Corresponds to the `fsGroup` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPFSGroup
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # ranges <array>: GID ranges affected by the rule.
    ranges:
      - # max <integer>: The maximum GID in the range, inclusive.
        max: <integer>
        # min <integer>: The minimum GID in the range, inclusive.
        min: <integer>
    # rule <string>: An FSGroup rule name.
    # Allowed Values: MayRunAs, MustRunAs, RunAsAny
    rule: <string>
```

<details>
  <summary>Examples</summary>

  #### psp-fsgroup
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPFSGroup
metadata:
  name: psp-fsgroup
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    ranges:
    - max: 1000
      min: 1
    rule: MayRunAs
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: fsgroup-disallowed
spec:
  containers:
  - command:
    - sh
    - -c
    - sleep 1h
    image: busybox
    name: fsgroup-demo
    volumeMounts:
    - mountPath: /data/demo
      name: fsgroup-demo-vol
  securityContext:
    fsGroup: 500
  volumes:
  - emptyDir: {}
    name: fsgroup-demo-vol
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: fsgroup-disallowed
spec:
  containers:
  - command:
    - sh
    - -c
    - sleep 1h
    image: busybox
    name: fsgroup-demo
    volumeMounts:
    - mountPath: /data/demo
      name: fsgroup-demo-vol
  securityContext:
    fsGroup: 2000
  volumes:
  - emptyDir: {}
    name: fsgroup-demo-vol
```
  </details>

---

## K8sPSPFlexVolumes

> Controls the allowlist of FlexVolume drivers. Corresponds to the `allowedFlexVolumes` field in PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#flexvolume-drivers

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPFlexVolumes
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedFlexVolumes <array>: An array of AllowedFlexVolume objects.
    allowedFlexVolumes:
      - # driver <string>: The name of the FlexVolume driver.
        driver: <string>
```

<details>
  <summary>Examples</summary>

  #### psp-flexvolume-drivers
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPFlexVolumes
metadata:
  name: psp-flexvolume-drivers
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    allowedFlexVolumes:
    - driver: example/lvm
    - driver: example/cifs
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-flexvolume-driver
  name: nginx-flexvolume-driver-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /test
      name: test-volume
      readOnly: true
  volumes:
  - flexVolume:
      driver: example/lvm
    name: test-volume
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-flexvolume-driver
  name: nginx-flexvolume-driver-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /test
      name: test-volume
      readOnly: true
  volumes:
  - flexVolume:
      driver: example/testdriver
    name: test-volume
```
  </details>

---

## K8sPSPForbiddenSysctls

> Controls the `sysctl` profile used by containers. Corresponds to the `forbiddenSysctls` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPForbiddenSysctls
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # forbiddenSysctls <array>: A disallow-list of sysctls. `*` forbids all
    # sysctls.
    forbiddenSysctls:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-forbidden-sysctls
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPForbiddenSysctls
metadata:
  name: psp-forbidden-sysctls
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    forbiddenSysctls:
    - kernel.*
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-forbidden-sysctls
  name: nginx-forbidden-sysctls-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
  securityContext:
    sysctls:
    - name: net.core.somaxconn
      value: "1024"
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-forbidden-sysctls
  name: nginx-forbidden-sysctls-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
  securityContext:
    sysctls:
    - name: kernel.msgmax
      value: "65536"
    - name: net.core.somaxconn
      value: "1024"
```
  </details>

---

## K8sPSPHostFilesystem

> Controls usage of the host filesystem. Corresponds to the `allowedHostPaths` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostFilesystem
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedHostPaths <array>: An array of hostpath objects, representing
    # paths and read/write configuration.
    allowedHostPaths:
      - # pathPrefix <string>: The path prefix that the host volume must
        # match.
        pathPrefix: <string>
        # readOnly <boolean>: when set to true, any container volumeMounts
        # matching the pathPrefix must include `readOnly: true`.
        readOnly: <boolean>
```

<details>
  <summary>Examples</summary>

  #### psp-host-filesystem
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostFilesystem
metadata:
  name: psp-host-filesystem
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    allowedHostPaths:
    - pathPrefix: /foo
      readOnly: true
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-filesystem-disallowed
  name: nginx-host-filesystem
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
      readOnly: true
  volumes:
  - hostPath:
      path: /foo/bar
    name: cache-volume
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-filesystem-disallowed
  name: nginx-host-filesystem
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
      readOnly: true
  volumes:
  - hostPath:
      path: /tmp
    name: cache-volume
```
  </details>

---

## K8sPSPHostNamespace

> Disallows sharing of host PID and IPC namespaces by pod containers. Corresponds to the `hostPID` and `hostIPC` fields in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#host-namespaces

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostNamespace
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    <object>
```

<details>
  <summary>Examples</summary>

  #### psp-host-namespace
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostNamespace
metadata:
  name: psp-host-namespace
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-namespace
  name: nginx-host-namespace-allowed
spec:
  containers:
  - image: nginx
    name: nginx
  hostIPC: false
  hostPID: false
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-namespace
  name: nginx-host-namespace-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
  hostIPC: true
  hostPID: true
```
  </details>

---

## K8sPSPHostNetworkingPorts

> Controls usage of host network namespace by pod containers. Specific ports must be specified. Corresponds to the `hostNetwork` and `hostPorts` fields in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#host-namespaces

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostNetworkingPorts
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # hostNetwork <boolean>: Determines if the policy allows the use of
    # HostNetwork in the pod spec.
    hostNetwork: <boolean>
    # max <integer>: The end of the allowed port range, inclusive.
    max: <integer>
    # min <integer>: The start of the allowed port range, inclusive.
    min: <integer>
```

<details>
  <summary>Examples</summary>

  #### psp-host-network-ports
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPHostNetworkingPorts
metadata:
  name: psp-host-network-ports
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    hostNetwork: true
    max: 9000
    min: 80
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-networking-ports
  name: nginx-host-networking-ports-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 9000
      hostPort: 80
  hostNetwork: false
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-host-networking-ports
  name: nginx-host-networking-ports-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 9001
      hostPort: 9001
  hostNetwork: true
```
  </details>

---

## K8sPSPPrivilegedContainer

> Controls the ability of any container to enable privileged mode. Corresponds to the `privileged` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#privileged

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPPrivilegedContainer
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-privileged-container
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPPrivilegedContainer
metadata:
  name: psp-privileged-container
spec:
  match:
    excludedNamespaces:
    - kube-system
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-privileged
  name: nginx-privileged-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      privileged: false
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-privileged
  name: nginx-privileged-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      privileged: true
```
  </details>

---

## K8sPSPProcMount

> Controls the allowed `procMount` types for the container. Corresponds to the `allowedProcMountTypes` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#allowedprocmounttypes

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPProcMount
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
    # procMount <string>: Defines the strategy for the security exposure of
    # certain paths in `/proc` by the container runtime. Setting to `Default`
    # uses the runtime defaults, where `Unmasked` bypasses the default
    # behavior.
    # Allowed Values: Default, Unmasked
    procMount: <string>
```

<details>
  <summary>Examples</summary>

  #### psp-proc-mount
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPProcMount
metadata:
  name: psp-proc-mount
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    procMount: Default
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-proc-mount
  name: nginx-proc-mount-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      procMount: Default
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-proc-mount
  name: nginx-proc-mount-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      procMount: Unmasked
```
  </details>

---

## K8sPSPReadOnlyRootFilesystem

> Requires the use of a read-only root file system by pod containers. Corresponds to the `readOnlyRootFilesystem` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPReadOnlyRootFilesystem
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-readonlyrootfilesystem
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPReadOnlyRootFilesystem
metadata:
  name: psp-readonlyrootfilesystem
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-readonlyrootfilesystem
  name: nginx-readonlyrootfilesystem-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      readOnlyRootFilesystem: true
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-readonlyrootfilesystem
  name: nginx-readonlyrootfilesystem-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      readOnlyRootFilesystem: false
```
  </details>

---

## K8sPSPSELinuxV2

> Defines an allow-list of seLinuxOptions configurations for pod containers. Corresponds to a PodSecurityPolicy requiring SELinux configs. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#selinux

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPSELinuxV2
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedSELinuxOptions <array>: An allow-list of SELinux options
    # configurations.
    allowedSELinuxOptions:
      # <list item: object>: An allowed configuration of SELinux options for a
      # pod container.
      - # level <string>: An SELinux level.
        level: <string>
        # role <string>: An SELinux role.
        role: <string>
        # type <string>: An SELinux type.
        type: <string>
        # user <string>: An SELinux user.
        user: <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-selinux-v2
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPSELinuxV2
metadata:
  name: psp-selinux-v2
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    allowedSELinuxOptions:
    - level: s0:c123,c456
      role: object_r
      type: svirt_sandbox_file_t
      user: system_u
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-selinux
  name: nginx-selinux-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      seLinuxOptions:
        level: s0:c123,c456
        role: object_r
        type: svirt_sandbox_file_t
        user: system_u
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-selinux
  name: nginx-selinux-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      seLinuxOptions:
        level: s1:c234,c567
        role: sysadm_r
        type: svirt_lxc_net_t
        user: sysadm_u
```
  </details>

---

## K8sPSPSeccomp

> Controls the seccomp profile used by containers. Corresponds to the `seccomp.security.alpha.kubernetes.io/allowedProfileNames` annotation on a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPSeccomp
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedLocalhostFiles <array>: When using securityContext naming scheme
    # for seccomp and including `Localhost` this array holds the allowed
    # profile JSON files. Putting a `*` in this array will allows all JSON
    # files to be used. This field is required to allow `Localhost` in
    # securityContext as with an empty list it will block.
    allowedLocalhostFiles:
      - <string>
    # allowedProfiles <array>: An array of allowed profile values for seccomp
    # on Pods/Containers. Can use the annotation naming scheme:
    # `runtime/default`, `docker/default`, `unconfined` and/or
    # `localhost/some-profile.json`. The item `localhost/*` will allow any
    # localhost based profile. Can also use the securityContext naming scheme:
    # `RuntimeDefault`, `Unconfined` and/or `Localhost`. For securityContext
    # `Localhost`, use the parameter `allowedLocalhostProfiles` to list the
    # allowed profile JSON files. The policy code will translate between the
    # two schemes so it is not necessary to use both. Putting a `*` in this
    # array allows all Profiles to be used. This field is required since with
    # an empty list this policy will block all workloads.
    allowedProfiles:
      - <string>
    # exemptImages <array>: Any container that uses an image that matches an
    # entry in this list will be excluded from enforcement. Prefix-matching can
    # be signified with `*`. For example: `my-image-*`. It is recommended that
    # users use the fully-qualified Docker image name (e.g. start with a domain
    # name) in order to avoid unexpectedly exempting images from an untrusted
    # repository.
    exemptImages:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-seccomp
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPSeccomp
metadata:
  name: psp-seccomp
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    allowedProfiles:
    - runtime/default
    - docker/default
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    container.seccomp.security.alpha.kubernetes.io/nginx: runtime/default
  labels:
    app: nginx-seccomp
  name: nginx-seccomp-allowed
spec:
  containers:
  - image: nginx
    name: nginx
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    seccomp.security.alpha.kubernetes.io/pod: runtime/default
  labels:
    app: nginx-seccomp
  name: nginx-seccomp-allowed2
spec:
  containers:
  - image: nginx
    name: nginx
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    seccomp.security.alpha.kubernetes.io/pod: unconfined
  labels:
    app: nginx-seccomp
  name: nginx-seccomp-disallowed2
spec:
  containers:
  - image: nginx
    name: nginx
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  annotations:
    container.seccomp.security.alpha.kubernetes.io/nginx: unconfined
  labels:
    app: nginx-seccomp
  name: nginx-seccomp-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
```
  </details>

---

## K8sPSPVolumeTypes

> Restricts mountable volume types to those specified by the user. Corresponds to the `volumes` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#volumes-and-file-systems

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPVolumeTypes
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # volumes <array>: `volumes` is an array of volume types. All volume types
    # can be enabled using `*`.
    volumes:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### psp-volume-types
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPVolumeTypes
metadata:
  name: psp-volume-types
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    volumes:
    - configMap
    - emptyDir
    - projected
    - secret
    - downwardAPI
    - persistentVolumeClaim
    - flexVolume
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-volume-types
  name: nginx-volume-types-allowed
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
  - image: nginx
    name: nginx2
    volumeMounts:
    - mountPath: /cache2
      name: demo-vol
  volumes:
  - emptyDir: {}
    name: cache-volume
  - emptyDir: {}
    name: demo-vol
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  labels:
    app: nginx-volume-types
  name: nginx-volume-types-disallowed
spec:
  containers:
  - image: nginx
    name: nginx
    volumeMounts:
    - mountPath: /cache
      name: cache-volume
  - image: nginx
    name: nginx2
    volumeMounts:
    - mountPath: /cache2
      name: demo-vol
  volumes:
  - hostPath:
      path: /tmp
    name: cache-volume
  - emptyDir: {}
    name: demo-vol
```
  </details>

---

## K8sPodDisruptionBudget

> Disallow the following scenarios when deploying PodDisruptionBudgets or resources that implement the replica subresource (e.g. Deployment, ReplicationController, ReplicaSet, StatefulSet): 1. Deployment of PodDisruptionBudgets with .spec.maxUnavailable == 0 2. Deployment of PodDisruptionBudgets with .spec.minAvailable == .spec.replicas of the resource with replica subresource This will prevent PodDisruptionBudgets from blocking voluntary disruptions such as node draining.
https://kubernetes.io/docs/concepts/workloads/pods/disruptions/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPodDisruptionBudget
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### pod-distruption-budget
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPodDisruptionBudget
metadata:
  name: pod-distruption-budget
spec:
  match:
    kinds:
    - apiGroups:
      - apps
      kinds:
      - Deployment
      - ReplicaSet
      - StatefulSet
    - apiGroups:
      - policy
      kinds:
      - PodDisruptionBudget
    - apiGroups:
      - ""
      kinds:
      - ReplicationController
```
  #### Allowed
  ```yaml
  apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: nginx-pdb-allowed
  namespace: default
spec:
  maxUnavailable: 1
  selector:
    matchLabels:
      foo: bar
```
  
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment-allowed-1
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
      example: allowed-deployment-1
  template:
    metadata:
      labels:
        app: nginx
        example: allowed-deployment-1
    spec:
      containers:
      - image: nginx:1.14.2
        name: nginx
        ports:
        - containerPort: 80
---
# Referential Data
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: inventory-nginx-pdb-allowed-1
  namespace: default
spec:
  minAvailable: 2
  selector:
    matchLabels:
      app: nginx
      example: allowed-deployment-1
```
  
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment-allowed-2
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
      example: allowed-deployment-2
  template:
    metadata:
      labels:
        app: nginx
        example: allowed-deployment-2
    spec:
      containers:
      - image: nginx:1.14.2
        name: nginx
        ports:
        - containerPort: 80
---
# Referential Data
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: inventory-nginx-pdb-allowed-2
  namespace: default
spec:
  maxUnavailable: 1
  selector:
    matchLabels:
      app: nginx
      example: allowed-deployment-2
```
  #### Disallowed
  ```yaml
  apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: nginx-pdb-disallowed
  namespace: default
spec:
  maxUnavailable: 0
  selector:
    matchLabels:
      foo: bar
```
  
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx-deployment-disallowed
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
      example: disallowed-deployment
  template:
    metadata:
      labels:
        app: nginx
        example: disallowed-deployment
    spec:
      containers:
      - image: nginx:1.14.2
        name: nginx
        ports:
        - containerPort: 80
---
# Referential Data
apiVersion: policy/v1
kind: PodDisruptionBudget
metadata:
  name: inventory-nginx-pdb-disallowed
  namespace: default
spec:
  minAvailable: 3
  selector:
    matchLabels:
      app: nginx
      example: disallowed-deployment
```
  </details>

---

## K8sReplicaLimits

> Requires that objects with the field `spec.replicas` (Deployments, ReplicaSets, etc.) specify a number of replicas within defined ranges.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sReplicaLimits
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # ranges <array>: Allowed ranges for numbers of replicas.  Values are
    # inclusive.
    ranges:
      # <list item: object>: A range of allowed replicas.  Values are
      # inclusive.
      - # max_replicas <integer>: The maximum number of replicas allowed,
        # inclusive.
        max_replicas: <integer>
        # min_replicas <integer>: The minimum number of replicas allowed,
        # inclusive.
        min_replicas: <integer>
```

<details>
  <summary>Examples</summary>

  #### replica-limits
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sReplicaLimits
metadata:
  name: replica-limits
spec:
  match:
    kinds:
    - apiGroups:
      - apps
      kinds:
      - Deployment
  parameters:
    ranges:
    - max_replicas: 50
      min_replicas: 3
```
  #### Allowed
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  name: allowed-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx:1.14.2
        name: nginx
        ports:
        - containerPort: 80
```
  #### Disallowed
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  name: disallowed-deployment
spec:
  replicas: 100
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx:1.14.2
        name: nginx
        ports:
        - containerPort: 80
```
  </details>

---

## K8sRequiredAnnotations

> Requires resources to contain specified annotations, with values matching provided regular expressions.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredAnnotations
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # annotations <array>: A list of annotations and values the object must
    # specify.
    annotations:
      - # allowedRegex <string>: If specified, a regular expression the
        # annotation's value must match. The value must contain at least one
        # match for the regular expression.
        allowedRegex: <string>
        # key <string>: The required annotation.
        key: <string>
    message: <string>
```

<details>
  <summary>Examples</summary>

  #### all-must-have-certain-set-of-annotations
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredAnnotations
metadata:
  name: all-must-have-certain-set-of-annotations
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Service
  parameters:
    annotations:
    - allowedRegex: ^([A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}|[a-z]{1,39})$
      key: a8r.io/owner
    - allowedRegex: ^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$
      key: a8r.io/runbook
    message: All services must have a `a8r.io/owner` and `a8r.io/runbook` annotations.
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  annotations:
    a8r.io/owner: dev-team-alfa@contoso.com
    a8r.io/runbook: https://confluence.contoso.com/dev-team-alfa/runbooks
  name: allowed-service
spec:
  ports:
  - name: http
    port: 80
    targetPort: 8080
  selector:
    app: foo
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: disallowed-service
spec:
  ports:
  - name: http
    port: 80
    targetPort: 8080
  selector:
    app: foo
```
  </details>

---

## K8sRequiredLabels

> Requires resources to contain specified labels, with values matching provided regular expressions.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredLabels
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # labels <array>: A list of labels and values the object must specify.
    labels:
      - # allowedRegex <string>: If specified, a regular expression the
        # annotation's value must match. The value must contain at least one
        # match for the regular expression.
        allowedRegex: <string>
        # key <string>: The required label.
        key: <string>
    message: <string>
```

<details>
  <summary>Examples</summary>

  #### all-must-have-owner
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredLabels
metadata:
  name: all-must-have-owner
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Namespace
  parameters:
    labels:
    - allowedRegex: ^[a-zA-Z]+.agilebank.demo$
      key: owner
    message: All namespaces must have an `owner` label that points to your company
      username
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Namespace
metadata:
  labels:
    owner: user.agilebank.demo
  name: allowed-namespace
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Namespace
metadata:
  name: disallowed-namespace
```
  </details>

---

## K8sRequiredProbes

> Requires Pods to have readiness and/or liveness probes.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredProbes
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # probeTypes <array>: The probe must define a field listed in `probeType`
    # in order to satisfy the constraint (ex. `tcpSocket` satisfies
    # `['tcpSocket', 'exec']`)
    probeTypes:
      - <string>
    # probes <array>: A list of probes that are required (ex: `readinessProbe`)
    probes:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### must-have-probes
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequiredProbes
metadata:
  name: must-have-probes
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - Pod
  parameters:
    probeTypes:
    - tcpSocket
    - httpGet
    - exec
    probes:
    - readinessProbe
    - livenessProbe
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: test-pod1
spec:
  containers:
  - image: tomcat
    livenessProbe:
      initialDelaySeconds: 5
      periodSeconds: 10
      tcpSocket:
        port: 80
    name: tomcat
    ports:
    - containerPort: 8080
    readinessProbe:
      initialDelaySeconds: 5
      periodSeconds: 10
      tcpSocket:
        port: 8080
  volumes:
  - emptyDir: {}
    name: cache-volume
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: test-pod1
spec:
  containers:
  - image: nginx:1.7.9
    livenessProbe: null
    name: nginx-1
    ports:
    - containerPort: 80
    volumeMounts:
    - mountPath: /tmp/cache
      name: cache-volume
  - image: tomcat
    name: tomcat
    ports:
    - containerPort: 8080
    readinessProbe:
      initialDelaySeconds: 5
      periodSeconds: 10
      tcpSocket:
        port: 8080
  volumes:
  - emptyDir: {}
    name: cache-volume
```
  
  ```yaml
  apiVersion: v1
kind: Pod
metadata:
  name: test-pod2
spec:
  containers:
  - image: nginx:1.7.9
    livenessProbe:
      initialDelaySeconds: 5
      periodSeconds: 10
      tcpSocket:
        port: 80
    name: nginx-1
    ports:
    - containerPort: 80
    readinessProbe: null
    volumeMounts:
    - mountPath: /tmp/cache
      name: cache-volume
  - image: tomcat
    name: tomcat
    ports:
    - containerPort: 8080
    readinessProbe:
      initialDelaySeconds: 5
      periodSeconds: 10
      tcpSocket:
        port: 8080
  volumes:
  - emptyDir: {}
    name: cache-volume
```
  </details>

---

## K8sUniqueIngressHost

> Requires all Ingress rule hosts to be unique.
Does not handle hostname wildcards: https://kubernetes.io/docs/concepts/services-networking/ingress/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sUniqueIngressHost
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### unique-ingress-host
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sUniqueIngressHost
metadata:
  name: unique-ingress-host
spec:
  match:
    kinds:
    - apiGroups:
      - extensions
      - networking.k8s.io
      kinds:
      - Ingress
```
  #### Allowed
  ```yaml
  apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-host-disallowed
  namespace: default
spec:
  rules:
  - host: example-host.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
  - host: example-host1.example.com
    http:
      paths:
      - backend:
          serviceName: nginx2
          servicePort: 80
```
  #### Disallowed
  ```yaml
  apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-host-disallowed
  namespace: default
spec:
  rules:
  - host: example-host.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
---
# Referential Data
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-host-example
  namespace: default
spec:
  rules:
  - host: example-host.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
```
  
  ```yaml
  apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: ingress-host-disallowed2
  namespace: default
spec:
  rules:
  - host: example-host2.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
  - host: example-host3.example.com
    http:
      paths:
      - backend:
          serviceName: nginx2
          servicePort: 80
---
# Referential Data
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: ingress-host-example2
  namespace: default
spec:
  rules:
  - host: example-host2.example.com
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
```
  </details>

---

## K8sUniqueServiceSelector

> Requires Services to have unique selectors within a namespace. Selectors are considered the same if they have identical keys and values. Selectors may share a key/value pair so long as there is at least one distinct key/value pair between them.
https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sUniqueServiceSelector
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
```

<details>
  <summary>Examples</summary>

  #### unique-service-selector
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sUniqueServiceSelector
metadata:
  labels:
    owner: admin.agilebank.demo
  name: unique-service-selector
```
  #### Allowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: gatekeeper-test-service-disallowed
  namespace: default
spec:
  ports:
  - port: 443
  selector:
    key: other-value
```
  #### Disallowed
  ```yaml
  apiVersion: v1
kind: Service
metadata:
  name: gatekeeper-test-service-disallowed
  namespace: default
spec:
  ports:
  - port: 443
  selector:
    key: value
---
# Referential Data
apiVersion: v1
kind: Service
metadata:
  name: gatekeeper-test-service-example
  namespace: default
spec:
  ports:
  - port: 443
  selector:
    key: value
```
  </details>

---

## NoUpdateServiceAccount

> Blocks updating the service account on resources that abstract over Pods. This policy is ignored in audit mode.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: NoUpdateServiceAccount
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint. For match schema details please refer to:
  # https://open-policy-agent.github.io/gatekeeper/website/docs/howto/#the-match-field
  match:
    [match schema]
  parameters:
    # allowedGroups <array>: Groups that should be allowed to bypass the
    # policy.
    allowedGroups:
      - <string>
    # allowedUsers <array>: Users that should be allowed to bypass the policy.
    allowedUsers:
      - <string>
```

<details>
  <summary>Examples</summary>

  #### no-update-kube-system-service-account
  #### Constraint 
  ```yaml
  apiVersion: constraints.gatekeeper.sh/v1beta1
kind: NoUpdateServiceAccount
metadata:
  name: no-update-kube-system-service-account
spec:
  match:
    kinds:
    - apiGroups:
      - ""
      kinds:
      - ReplicationController
    - apiGroups:
      - apps
      kinds:
      - ReplicaSet
      - Deployment
      - StatefulSet
      - DaemonSet
    - apiGroups:
      - batch
      kinds:
      - CronJob
    namespaces:
    - kube-system
  parameters:
    allowedGroups: []
    allowedUsers: []
```
  #### Allowed
  ```yaml
  apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: policy-test
  name: policy-test
  namespace: kube-system
spec:
  replicas: 1
  selector:
    matchLabels:
      app: policy-test-deploy
  template:
    metadata:
      labels:
        app: policy-test-deploy
    spec:
      containers:
      - command:
        - /bin/bash
        - -c
        - sleep 99999
        image: ubuntu
        name: policy-test
      serviceAccountName: policy-test-sa-1
```
  #### Disallowed</details>

---
## DisallowedAuthzPrefix

Requires that principals and namespaces in Istio `AuthorizationPolicy` rules not have a prefix from a specified list.
https://istio.io/latest/docs/reference/config/security/authorization-policy/

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: DisallowedAuthzPrefix
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint.  Please see the match criteria documentation for more information:
  # https://cloud.google.com/anthos-config-management/docs/reference/match
  match:
    [match schema]
  parameters:
    # disallowedprefixes <array>: Disallowed prefixes of principals and
    # namespaces.
    disallowedprefixes:
      - <string>
```

<div>
<devsite-expandable>
<h3 class="showalways">Examples</h3>
<h4>disallowed-authz-prefix-constraint</h4>
<h5>Constraint</h5>
<pre class="prettyprint lang-yaml">
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: DisallowedAuthzPrefix
metadata:
  name: disallowed-authz-prefix-constraint
spec:
  enforcementAction: dryrun
  match:
    kinds:
    - apiGroups:
      - security.istio.io
      kinds:
      - AuthorizationPolicy
  parameters:
    disallowedprefixes:
    - badprefix
    - reallybadprefix
</pre>
<h5>Disallowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: bad-source-namespace
  namespace: foo
spec:
  rules:
  - from:
    - source:
        principals:
        - cluster.local/ns/default/sa/sleep
    - source:
        namespaces:
        - badprefix-test
    to:
    - operation:
        methods:
        - GET
        paths:
        - /info*
    - operation:
        methods:
        - POST
        paths:
        - /data
    when:
    - key: request.auth.claims[iss]
      values:
      - https://accounts.google.com
  selector:
    matchLabels:
      app: httpbin
      version: v1
</pre>
</devsite-expandable>
</div>

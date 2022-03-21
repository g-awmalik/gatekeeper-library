## K8sPodsRequireSecurityContext

Requires all Pods and containers to have a SecurityContext defined at the Pod or container level.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPodsRequireSecurityContext
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint.  Please see the match criteria documentation for more information:
  # https://cloud.google.com/anthos-config-management/docs/reference/match
  match:
    [match schema]
```

<div>
<devsite-expandable>
<h3 class="showalways">Examples</h3>
<h4>pods-require-security-context</h4>
<h5>Constraint</h5>
<pre class="prettyprint lang-yaml">
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPodsRequireSecurityContext
metadata:
  name: pods-require-security-context
</pre>
<h5>Allowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Pod
metadata:
  name: allowed-example
spec:
  containers:
  - image: nginx
    name: nginx
    securityContext:
      runAsUser: 2000
</pre>
<h5>Disallowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Pod
metadata:
  name: disallowed-example
spec:
  containers:
  - image: nginx
    name: nginx
</pre>
</devsite-expandable>
</div>

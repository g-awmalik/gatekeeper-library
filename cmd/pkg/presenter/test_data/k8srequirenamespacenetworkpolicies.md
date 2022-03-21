## K8sRequireNamespaceNetworkPolicies

Requires that every namespace defined in the cluster has a NetworkPolicy.
Note: This constraint is referential. See https://cloud.google.com/anthos-config-management/docs/how-to/creating-constraints#referential for details.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequireNamespaceNetworkPolicies
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
<h4>require-namespace-network-policies</h4>
<h5>Constraint</h5>
<pre class="prettyprint lang-yaml">
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequireNamespaceNetworkPolicies
metadata:
  name: require-namespace-network-policies
spec:
  enforcementAction: dryrun
</pre>
<h5>Allowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Namespace
metadata:
  name: require-namespace-network-policies-example
---
# Referential Data
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: require-namespace-network-policies-example
</pre>
<h5>Disallowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Namespace
metadata:
  name: require-namespace-network-policies-example
</pre>
</devsite-expandable>
</div>

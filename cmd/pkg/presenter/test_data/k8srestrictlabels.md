## K8sRestrictLabels

Disallows resources with any of the specified `restrictedLabels`. Matches on label key names only.  Single object exceptions can be included, identified by their group, kind, namespace, and name.

### Constraint schema

```yaml
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRestrictLabels
metadata:
  name: example
spec:
  # match <object>: allows you to configure which resources fall in scope for
  # this constraint.  Please see the match criteria documentation for more information:
  # https://cloud.google.com/anthos-config-management/docs/reference/match
  match:
    [match schema]
  parameters:
    # exceptions <array>: A list of objects that are exempted from the label
    # restrictions.
    exceptions:
      # <list item: object>: A single object's identification, based on group,
      # kind, namespace, and name.
      - group: <string>
        kind: <string>
        name: <string>
        namespace: <string>
    # restrictedLabels <array>: A list of label keys strings.
    restrictedLabels:
      - <string>
```

<div>
<devsite-expandable>
<h3 class="showalways">Examples</h3>
<h4>restrict-label-example</h4>
<h5>Constraint</h5>
<pre class="prettyprint lang-yaml">
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRestrictLabels
metadata:
  name: restrict-label-example
spec:
  enforcementAction: dryrun
  parameters:
    exceptions:
    - group: ""
      kind: Pod
      name: allowed-example
      namespace: default
    restrictedLabels:
    - label-example
</pre>
<h5>Allowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Pod
metadata:
  labels:
    label-example: example
  name: allowed-example
  namespace: default
spec:
  containers:
  - image: nginx
    name: nginx
</pre>
<h5>Disallowed</h5>
<pre class="prettyprint lang-yaml">
apiVersion: v1
kind: Pod
metadata:
  labels:
    label-example: example
  name: disallowed-example
  namespace: default
spec:
  containers:
  - image: nginx
    name: nginx
</pre>
</devsite-expandable>
</div>

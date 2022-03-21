```yaml
# excludedNamespaces <array>: `excludedNamespaces` is a list of namespace
# names. If defined, a constraint only applies to resources not in a listed
# namespace. ExcludedNamespaces also supports a prefix-based glob.  For
# example, `excludedNamespaces: [kube-*]` matches both `kube-system` and
# `kube-public`.
excludedNamespaces:
  - <string>
kinds:
  # <list item: object>: The Group and Kind of objects that should be matched.
  # If multiple groups/kinds combinations are specified, an incoming resource
  # need only match one to be in scope.
  - apiGroups:
      - <string>
    kinds:
      - <string>
# labelSelector <object>: `labelSelector` is the combination of two optional
# fields: `matchLabels` and `matchExpressions`.  These two fields provide
# different methods of selecting or excluding k8s objects based on the label
# keys and values included in object metadata.  All selection expressions from
# both sections are ANDed to determine if an object meets the cumulative
# requirements of the selector.
labelSelector:
  # matchExpressions <array>: a list of label selection expressions. A selected
  # resource will match all of these expressions.
  matchExpressions:
    # <list item: object>: a selector that specifies a label key, a set of
    # label values, an operator that defines the relationship between the two
    # that will match the selector.
    - # key <string>: the label key that the selector applies to.
      key: <string>
      # operator <string>: the relationship between the label and value set
      # that defines a matching selection.
      # Allowed Values: In, NotIn, Exists, DoesNotExist
      operator: <string>
      # values <array>: a set of label values.
      values:
        - <string>
  # matchLabels <object>: A mapping of label keys to sets of allowed label
  # values for those keys.  A selected resource will match all of these
  # expressions.
  matchLabels:
    [key]: <string>
# name <string>: `name` is the name of an object.  If defined, it matches
# against objects with the specified name.  Name also supports a prefix-based
# glob.  For example, `name: pod-*` matches both `pod-a` and `pod-b`.
name: <string>
# namespaceSelector <object>: `namespaceSelector` is a label selector against
# an object's containing namespace or the object itself, if the object is a
# namespace.
namespaceSelector:
  # matchExpressions <array>: a list of label selection expressions. A selected
  # resource will match all of these expressions.
  matchExpressions:
    # <list item: object>: a selector that specifies a label key, a set of
    # label values, an operator that defines the relationship between the two
    # that will match the selector.
    - # key <string>: the label key that the selector applies to.
      key: <string>
      # operator <string>: the relationship between the label and value set
      # that defines a matching selection.
      # Allowed Values: In, NotIn, Exists, DoesNotExist
      operator: <string>
      # values <array>: a set of label values.
      values:
        - <string>
  # matchLabels <object>: A mapping of label keys to sets of allowed label
  # values for those keys.  A selected resource will match all of these
  # expressions.
  matchLabels:
    [key]: <string>
# namespaces <array>: `namespaces` is a list of namespace names. If defined, a
# constraint only applies to resources in a listed namespace.  Namespaces also
# supports a prefix-based glob.  For example, `namespaces: [kube-*]` matches
# both `kube-system` and `kube-public`.
namespaces:
  - <string>
# scope <string>: `scope` determines if cluster-scoped and/or namespaced-scoped
# resources are matched.  Accepts `*`, `Cluster`, or `Namespaced`. (defaults to
# `*`)
# Allowed Values: *, Cluster, Namespaced
scope: <string>
```

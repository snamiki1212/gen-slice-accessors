# go-gen-slice-accessors

Generate accessors for each field in the slice struct.

## Usage

### 1. Add go:generate directive.

```diff filename="user.go"
package main

+//go:generate go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go
type User struct {
  UserID    string
}

type Users []User
```

### 2. Run `go generate` and got generated code.

```diff filename="user_gen.go"
+// Code generated by go generate DO NOT EDIT.
+
+package main
+
+// UserIDs
+func (xs Users) UserIDs() []string {
+	sli := make([]string, 0, len(xs))
+	for i := range xs {
+		sli = append(sli, xs[i].UserID)
+	}
+	return sli
+}
```

## Help

```shell
$ go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors@latest --help

Generate accessors for each field in the slice struct.

Usage:
  gen-slice-accessors [flags]

Flags:
  -a, --accessor strings   accessor name for field / e.g. --accessor=Name:GetName
  -e, --entity string      target entity name
  -x, --exclude strings    field names to exclude
  -h, --help               help for gen-slice-accessors
  -i, --input string       input file name
  -o, --output string      output file name
  -s, --slice string       target slice name
```

## Examples

### Exclude

```diff filename="user.go"
package main

+//go:generate go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go --exclude=CreatedAt,UpdatedAt
type User struct {
	UserID    string
	CreatedAt int64
	UpdatedAt int64
}

type Users []User
```

```diff filename="user_gen.go"
+// Code generated by go generate DO NOT EDIT.
+
+package main
+
+// UserIDs
+func (xs Users) UserIDs() []string {
+	sli := make([]string, 0, len(xs))
+	for i := range xs {
+		sli = append(sli, xs[i].UserID)
+	}
+	return sli
+}
```

### Define accessor name

```diff filename="user.go"
package main

+//go:generate go run -mod=mod github.com/snamiki1212/go-gen-slice-accessors --entity User --slice Users --input user.go --output user_gen.go --accessor=UserID:GetUserIDs --accessor=Age:AgeList
type User struct {
  UserID    string
  Age       int64
}

type Users []User
```

```diff filename="user_gen.go"
+// Code generated by go generate DO NOT EDIT.
+
+package main
+
+// GetUserIDs
+func (xs Users) GetUserIDs() []string {
+	sli := make([]string, 0, len(xs))
+	for i := range xs {
+		sli = append(sli, xs[i].UserID)
+	}
+	return sli
+}
+
+// AgeList
+func (xs Users) AgeList() []int64 {
+	sli := make([]int64, 0, len(xs))
+	for i := range xs {
+		sli = append(sli, xs[i].Age)
+	}
+	return sli
+}
```

## E2E

TODO: Add E2E test on CI

```zsh
$ go generate ./example
$ go run ./example
```

## LICENSE

[MIT](./LICENSE)

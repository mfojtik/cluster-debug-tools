module github.com/openshift/cluster-debug-tools

go 1.12

require (
	github.com/google/go-github/v32 v32.1.0
	github.com/openshift/library-go v0.0.0-20190927184318-c355e2019bb3
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/apiserver v0.0.0-20190918160949-bfa5e2e684ad
	k8s.io/cli-runtime v0.0.0-20190918162238-f783a3654da8
	k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)

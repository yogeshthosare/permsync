Permsync is my practice repo to play out with k8s programming in general. Will write k8s controllers that keeps the the permissions on filesystem in-sync with k8s object i define in api-specs. 

Created the spec file and generated deepcopy, clientset, listeners and informers boiler plate code with k8s-code-generator. (so far)

```
ubox@ubox:~/go/src/permsync/github.com/yogeshthosare/permsync$ execDir=~/go/pkg/mod/k8s.io/code-generator@v0.20.2

ubox@ubox:~/go/src/permsync$ "${execDir}"/generate-groups.sh all github.com/yogeshthosare/permsync/pkg/client github.com/yogeshthosare/permsync/pkg/apis yogesh.dev:v1alpha1 --go-header-file "${execDir}"/hack/boilerplate.go.txt

Generating deepcopy funcs

Generating clientset for yogesh.dev:v1alpha1 at github.com/yogeshthosare/permsync/pkg/client/clientset

Generating listers for yogesh.dev:v1alpha1 at github.com/yogeshthosare/permsync/pkg/client/listers

Generating informers for yogesh.dev:v1alpha1 at github.com/yogeshthosare/permsync/pkg/client/informers

ubox@ubox:~/go/src/permsync$ tree

.

├── github.com

│   └── yogeshthosare

│       └── permsync

│           └── pkg

│               ├── apis

│               │   └── yogesh.dev

│               │       └── v1alpha1

│               │           └── zz_generated.deepcopy.go

│               └── client

│                   ├── clientset

│                   │   └── versioned

│                   │       ├── clientset.go

│                   │       ├── doc.go

│                   │       ├── fake

│                   │       │   ├── clientset_generated.go

│                   │       │   ├── doc.go

│                   │       │   └── register.go

│                   │       ├── scheme

│                   │       │   ├── doc.go

│                   │       │   └── register.go

│                   │       └── typed

│                   │           └── yogesh.dev

│                   │               └── v1alpha1

│                   │                   ├── doc.go

│                   │                   ├── fake

│                   │                   │   ├── doc.go

│                   │                   │   ├── fake_permsync.go

│                   │                   │   └── fake_yogesh.dev_client.go

│                   │                   ├── generated_expansion.go

│                   │                   ├── permsync.go

│                   │                   └── yogesh.dev_client.go

│                   ├── informers

│                   │   └── externalversions

│                   │       ├── factory.go

│                   │       ├── generic.go

│                   │       ├── internalinterfaces

│                   │       │   └── factory_interfaces.go

│                   │       └── yogesh.dev

│                   │           ├── interface.go

│                   │           └── v1alpha1

│                   │               ├── interface.go

│                   │               └── permsync.go

│                   └── listers

│                       └── yogesh.dev

│                           └── v1alpha1

│                               ├── expansion_generated.go

│                               └── permsync.go


├── go.mod

├── go.sum

├── main.go

└── pkg

    └── apis
    
        └── yogesh.dev
	
            └── v1alpha1
	    
                ├── doc.go
		
                ├── register.go
		
                └── types.go
```
This is a bit off, I had to move doc.go, register.go, types.go alongside with zz_generated.deepcopy.go. Which is something I am to trying to figure out. rest is all good. 
Need to update this doc when I will get it resolved. 

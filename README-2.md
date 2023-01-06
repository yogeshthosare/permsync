Key to make this thing work (atleast on my machine) - 1. Create abc directory 2. Add remote abc repository 3. Create a mod of that set git origin within abc 4. Run code generation script from abc

1)
cd go/src 
mkdir abc
cd abc
pwd
~/go/src/abc 

2)
~/go/src/abc$ echo "# abc" >> README.md
~/go/src/abc$ git init
~/go/src/abc$ git add README.md
~/go/src/abc$ git commit -m "first commit"
~/go/src/abc$ git remote add origin git@github.com:yogeshthosare/abc.git
~/go/src/abc$ git push  origin master
~/go/src/abc$ cp -R ../k8s-object-skeleton/pkg/ .

3)
~/go/src/abc$ go mod init github.com/yogeshthosare/abc
~/go/src/abc$ cp -R ../k8s-object-skeleton/main.go .
~/go/src/abc$ go mod tidy

4)
~/go/src/abc$ execDir=~/go/pkg/mod/k8s.io/code-generator@v0.20.2
~/go/src/abc$ "${execDir}"/generate-groups.sh all github.com/yogeshthosare/abc/pkg/client github.com/yogeshthosare/abc/pkg/apis yogesh.dev:v1alpha1 --go-header-file "${execDir}"/hack/boilerplate.go.txt

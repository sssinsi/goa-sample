# goa-sample
goa sample

- vendering tool is glide

for Mac

```
brew update
brew install glide
glide init
glide get github.com/goadesign/goa
glide create
cd vender/github.com/goadesign/goa/goagen
go build
cd -
./vendor/github.com/goadesign/goa/goagen/goagen bootstrap -d github.com/sssinsi/goa-sample/cellar/design
```
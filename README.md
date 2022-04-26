# substrait-go

Experimental Go bindings for [substrait](https://substrait.io)

## Generate from proto files

initialize the submodule if you haven't yet done so:

```bash
git submodule update --init
```

Update the submodule for any upstream changes 

```bash
git submodule update --remote substrait
```

Then generate the files and copy them out

```bash
pushd substrait
buf generate
cp -r gen/proto/go/substrait/* ../proto
popd
```

After this you can commit the updated `.pb.go` files.
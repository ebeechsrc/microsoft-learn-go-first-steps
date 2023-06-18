# Creating modules and workspaces

This creates a workspace whose root is in this directory.

From here you can create the module:
```bash
cd calculator
go build
```

.. and the application, to refer to the local version of the module use:
```golang
require github.com/ebeechsrc/microsoft-learn-go-first-steps-calculator v0.0.0

replace github.com/ebeechsrc/microsoft-learn-go-first-steps-calculator v0.0.0 => ../calculator
```
and to run:
```bash
cd helloworld
go build
```

To add additional modules:

```bash
go mod init $MODULE_NAME
cd $MODULE_NAME
# Create your go package files...
go work use .
```
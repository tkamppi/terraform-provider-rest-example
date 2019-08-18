# terraform-provider-rest-example
A basic terraform provider and a HTTP client integrated towards dummy endpoints from https://httpbin.org 

The provider makes POST/GET/DELETE/PUT requests towards httpbin for the matching CRUD operations of Terraform.

## Build:
Place directories httpbin_client & terraform-provider-rest-example in your $GOPATH/src directory.
```
cd $GOPATH/src/terraform-provider-rest-example
go build -o terraform-provider-example
```

To build this for windows a .exe ending is required, replace the last command above with:
```
go build -o terraform-provider-example.exe
```

## Usage
The provider is called `restexample`, and it has a single resource defined `restexample_server` which takes the argument `name`.

Example:
Define a file `main.tf`
```
provider "restexample" {
  hostname = "https://httpbin.org"
}

resource "restexample_server" "my-server" {
  name = "my-server-name"
}
```

Make sure your go build produced binary is located in your `%APPDATA%\terraform.d\plugins` if on Windows or in `~/.terraform.d/plugins` on other systems. Alternatively just place the binary in the same folder as the `main.tf` file.

The provider now works with the usual `terraform init`, `terraform plan`, `terraform apply`, `terraform destroy` etc.

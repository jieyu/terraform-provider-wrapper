# Terraform Provider Wrapper

This allows users to wrap a Terraform Provider plugin and inject custom logging and auditing on operations at per Resource level.

## Build and Test

Make sure you project is under `$GOPATH`.

```bash
$ go build -o <target> ./cmd/aws
```

The `<target>` here should follow Terraform plugin [naming convention](https://www.terraform.io/docs/configuration/providers.html#plugin-names-and-versions).
In other word, it should be `terraform-provider-aws_vX.Y.Z`.

Then, put the built plugin to the [third party plugins directory](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) (i.e., `.terraform/plugins/`).

```bash
$ mkdir -p ~/.terraform/plugins/
$ cp <target> ~/.terraform/plugins/
```

Then, run `terraform init` to pick up the new plugin.

```bash
$ cd <config_dir>
$ terraform init -upgrade
```

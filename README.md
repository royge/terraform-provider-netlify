Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Installation
------------

**NOTE:** This is not an official provider's repo. Installation would be different.

1. Download pre-built binary

On Mac

```sh
$ curl -o terraform-provider-netlify_v0.4.1_x4 https://github.com/royge/terraform-provider-netlify/releases/download/v0.4.1/terraform-provider-netlify_v0.4.1_x4.darwin-amd64
```

On Linux

```sh
$ curl -o terraform-provider-netlify_v0.4.1_x4 https://github.com/royge/terraform-provider-netlify/releases/download/v0.4.1/terraform-provider-netlify_v0.4.1_x4.linux-amd64
```

2. Make executable and move into plugins directory

```sh
$ chmod +x terraform-provider-netlify_v0.4.1_x4 && mv terraform-provider-netlify_v0.4.1_x4 ~/.terraform.d/plugins/
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-netlify`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-netlify.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-netlify
$ make build
# or if you're on a mac:
$ gnumake build
```

Using the provider
----------------------

> *NOTE*: This provider is best used when paired with a VCS system provider such as the [Github provider](https://www.terraform.io/docs/providers/github/index.html), which will be used for reference in these examples, since Netlify integrates directly with your VCS system in order to continuously deploy your website.

Using this provider requires an auth token from Netlify. You can generate a token here: https://app.netlify.com/account/applications. You will also likely need an auth token for your VCS system. In this example, we'll use Github, so you'll want to get a Github token as well. We'll start by configuring Github. In this example, we'll assume that we're using a repo at `github.com/username/repo`.

```js
// configure the github provider
provider "github" {
  organization = "<username>"
}

// Configure the repository with the dynamically created Netlify key.
resource "github_repository_deploy_key" "key" {
  title      = "Netlify"
  repository = "<repo>"
  key        = "${netlify_deploy_key.key.public_key}"
  read_only  = false
}

// Create a webhook that triggers Netlify builds on push.
resource "github_repository_webhook" "main" {
  repository = "<repo>"
  name       = "web"
  events     = ["delete", "push", "pull_request"]

  configuration {
    content_type = "json"
    url          = "https://api.netlify.com/hooks/github"
  }

  depends_on = ["netlify_site.main"]
}
```

This pairs closely with the Netlify provider instructions as you can see, example shown below:

```js
// A new, unique deploy key for this specific website
resource "netlify_deploy_key" "key" {}

resource "netlify_site" "main" {
  name = "<name of netlify site>"

  repo {
    repo_branch = "<github branch to deploy>"
    command = "<command used to build your website>"
    deploy_key_id = "${netlify_deploy_key.key.id}"
    dir = "<directory your website is built into, relative to root>"
    provider = "github"
    repo_path = "<username/repo>"
  }
}
```

With all the details filled in here, you should be able to run the script and have your site deploy. Of course, it's likely that you will want to configure some of these values as variables, and you can use `GITHUB_TOKEN` and `NETLIFY_TOKEN` environment variables as well to represent these API keys.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-netlify
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

# Meta Context Manager

# Overview

Most command line tools that I use on a daily basis determine the user’s desired destination for API calls to remote services based on environment variables. These variables are often referenced using the same key by multiple tools. For each of these CLI tools, users need to supply the required environment variables. Maintaining the same “context” across multiple tools with differing variable schemas, and switching between these contexts can quickly become tedious.

![env_vars](./docs/images/env_vars.png)

You may find yourself struggling to remember which account, org, project, region, cluster, or namespace you had exported in your previous terminal instance. You could implement some terminal hints to signal the selected context to yourself…

![env_vars](./docs/images/terminal_hints.png)

But this requires a lot of custom shell scripting to get working, and isn’t very easy to share across an organization.

Wouldn’t it be nice if all of your environment variables could be inferred from a simple, easy-to-read “context” alias that you only had to set up once?

## Goals

1. Export a set of environment variables from an easily configured context.
2. Enable use of multiple tools with the same context, despite

## Terminology

* **context** - an abstraction around a set of environments that define how you are able to interact with cloud resources through your command line
* **scope** - a single component of a context, for example an AWS REGION, which may or may not comprise the entire context.

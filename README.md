# Hive-based extendable CLI

A test to build an extendable CLI based on the [hive](https://docs.cilium.io/en/latest/contributing/development/hive/) framework

Build the original CLI with:

`make`

the available commands are:

```
cli
├── foo
└── bar
    └── one
```

Each with its specific flags.

Build the extended CLI with:

`make -C extended`

the additional `bar` subcommand named `two` is listed:

```
cli
├── foo
└── bar
    ├── one
    └── two
```
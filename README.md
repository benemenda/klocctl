# klocctl
Build:

make build

or 

./klocctl go build -o ./klocctl main.go

Usage:

./klocctl -h

Contributing:

klocctl's CLI utlises the open-source spf13 projects cobra and viper, so documentation on extending its API can be readily learned there.

Extensions should follow the format defined well by cobra, i.e.:

---
Cobra is built on a structure of commands, arguments & flags.

Commands represent actions, Args are things and Flags are modifiers for those actions.

The best applications will read like sentences when used. Users will know how to use the application because they will natively understand how to use it.

The pattern to follow is APPNAME VERB NOUN --ADJECTIVE. or APPNAME COMMAND ARG --FLAG

A few good real world examples may better illustrate this point.

In the following example, 'server' is a command, and 'port' is a flag:

hugo server --port=1313

In this command we are telling Git to clone the url bare.

git clone URL --bare

---

Example Repo where Cobra and Viper are used brilliantly: https://github.com/jgsqware/clairctl/



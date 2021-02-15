# instago
instago is a CLI tool that helps gather instagram analytics for a certain account.

## features

opposed to those sketchy third party apps, you can do the following from your secure local machine:

* track ghost followers
* see who you are not following back

## example

`instago track ghosts`</br>checks which users don't follow you back

`instago track fans`</br>checks which users follow you that you don't follow back

## local usage
To run this tool:

1. clone repo & `cd`
2. `go install instago`
3. `instago login`

See [usage.md](./docs/usage.md) for more about tool's functionality, or run `instago --help`
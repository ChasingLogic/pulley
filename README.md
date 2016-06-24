# Pulley
A suckless way to use SSH as a client in golang.

## Installation

Simply import it in your project and run `go get`

```go
import "github.com/chasinglogic/pulley"
```

Alternatively you can go get it directly:

```bash
go get github.com/chasinglogic/pulley
```

## Usage

For full usage you can check out the
[godoc](https://godoc.org/github.com/chasinglogic/pulley)

TODO: Add examples.

## Contributing

1. Fork :fork_and_knife: it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push :shipit: to the branch: `git push origin my-new-feature`
5. :fire: Submit a pull request :D :fire:

All pull(ey) requests should go to the develop branch not master. Thanks!

## History

This library was inspired by a combination of 
[paramiko](http://www.paramiko.org/) and [parallel
ssh](https://github.com/pkittenis/parallel-ssh) I was working on a tool for
remote management of Unix systems and was writing so much wrapper code for the
golang ssh package that I realized it should be a library of it's own and thus
Pulley was born.

## License

This code is distributed under the Apache 2.0 License.

```
Copyright 2016 Mathew Robinson

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

```

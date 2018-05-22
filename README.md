# issues2markdown

[![License][License-Image]][License-Url]
[![CircleCI](https://circleci.com/gh/issues2markdown/cli.svg?style=svg)](https://circleci.com/gh/issues2markdown/cli)
[![Coverage Status](https://coveralls.io/repos/github/issues2markdown/cli/badge.svg?branch=develop)](https://coveralls.io/github/issues2markdown/cli?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/issues2markdown/cli)](https://goreportcard.com/report/github.com/issues2markdown/cli)

Convert a list of issues to markdown.

## Installation

Once you have [installed Go](http://golang.org/doc/install.html#releases), run these commands to install `issues2markdown` tool:

```bash
go get github.com/issues2markdown/cli
```

## Documentation

Execute the following command to get the provided command line tool usage information:

```bash
$ issues2markdown --help
issues2markdown converts a list of github issues to markdown list format

Usage:
  issues2markdown [flags]

Flags:
      --github-token string   github token
  -h, --help                  help for issues2markdown
  -v, --verbose               enable verbose mode
      --version               version for issues2markdown
```

### Example output

An example of the output could be:

```markdown
- [ ] org/repo : [Issue Title 1](https://github.com/org/repo/issues/1)
- [x] org/repo : [Issue Title 2](https://github.com/org/repo/issues/2)
- [x] org/repo : [Issue Title 3](https://github.com/org/repo/issues/3)
- [ ] org/repo2 : [Issue Title 1](https://github.com/org/repo2/issues/1)
```

Which will render as:

- [ ] org/repo : [Issue Title 1](https://github.com/org/repo/issues/1)
- [x] org/repo : [Issue Title 2](https://github.com/org/repo/issues/2)
- [x] org/repo : [Issue Title 3](https://github.com/org/repo/issues/3)
- [ ] org/repo2 : [Issue Title 1](https://github.com/org/repo2/issues/1)

## License

Copyright (c) 2018 issues2markdown Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[License-Url]: http://opensource.org/licenses/Apache
[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg

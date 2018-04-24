# issues2markdown

Convert a list of issues to markdown

Take a look at the [ROADMAP](ROADMAP.md) to folow the current development 
status.

## Installation

Once you have [installed Go](http://golang.org/doc/install.html#releases), run these commands to install the gomock package and the mockgen tool:

	go get github.com/repejota/issues2markdown
	
## Documentation

Execute the following command to get the provided command line tool usage information:

```
$ issues2markdown --help
issues2markdown converts a list of github issues to markdown list format

Usage:
issues2markdown [flags]

Flags:
    --github-token string   github token
-h, --help                  help for issues2markdown
-o, --organization string   filter by organization
-r, --repository string     filter by repository
    --state string          filter by issue state (default "all")
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


## Badges

* License [![License][License-Image]][License-Url]

* Test Coverage Master [![Coverage Status](https://coveralls.io/repos/github/repejota/issues2markdown/badge.svg?branch=master)](https://coveralls.io/github/repejota/issues2markdown?branch=master)
* Test Coverage Develop [![Coverage Status](https://coveralls.io/repos/github/repejota/issues2markdown/badge.svg?branch=develop)](https://coveralls.io/github/repejota/issues2markdown?branch=develop)

* Test Status Master [![CircleCI](https://circleci.com/gh/repejota/issues2markdown/tree/master.svg?style=svg)](https://circleci.com/gh/repejota/issues2markdown/tree/master)
* Test Status Develop [![CircleCI](https://circleci.com/gh/repejota/issues2markdown/tree/develop.svg?style=svg)](https://circleci.com/gh/repejota/issues2markdown/tree/develop)

* Golang ReportCard [![Go Report Card](https://goreportcard.com/badge/github.com/repejota/issues2markdown)](https://goreportcard.com/report/github.com/repejota/issues2markdown)

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

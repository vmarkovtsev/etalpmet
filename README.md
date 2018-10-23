<p align="center">
    <img src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f914.png" alt="🤔">
</p>
<h1 align="center">etalpmet</h1>
<p align="center">
      Extract templates from strings in Go.<br><br>
      <a href="http://godoc.org/gopkg.in/vmarkovtsev/etalpmet.v1"><img src="https://godoc.org/gopkg.in/vmarkovtsev/etalpmet.v1?status.svg" alt="GoDoc"></a>
      <a href="https://travis-ci.org/vmarkovtsev/etalpmet"><img src="https://travis-ci.org/vmarkovtsev/etalpmet.svg?branch=master" alt="Travis build Status"></a>
      <a href="https://codecov.io/gh/vmarkovtsev/etalpmet"><img src="https://codecov.io/github/vmarkovtsev/etalpmet/coverage.svg" alt="Code coverage"></a>
      <a href="https://goreportcard.com/report/github.com/vmarkovtsev/etalpmet"><img src="https://goreportcard.com/badge/github.com/vmarkovtsev/etalpmet" alt="Go Report Card"></a>
      <a href="https://opensource.org/licenses/Apache-2.0"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache 2.0 license"></a>
</p>
<p align="center">
  <a href="#overview">Overview</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#installation">Installation</a> •
  <a href="#contributions">Contributions</a> •
  <a href="#license">License</a>
</p>

--------

## Overview

Given some byte strings, this library extracts the common template between them -- some people call
it "reverse templating". For example, given the following list of file names:

```
IMG_20180930_171704.jpg
IMG_20181001_150308.jpg
IMG_20181001_190338.jpg
IMG_20181021_122346.jpg
```

we can infer the common file name template: `IMG_2018*_1*.jpg`.

Idea credits: [turicas/templater](https://github.com/turicas/templater/).

## How To Use

There are two functions: "basic" and "advanced". Both return a slice of byte slices which correspond to
template constant blocks.
The former function is straightforward:

```go
import "gopkg.in/vmarkovtsev/etalpmet.v1"

template := etalpmet.ReverseTemplate(
	[]byte("<b> spam and eggs </b>"),
	[]byte("<b> ham and spam </b>"),
    []byte("<b> white and black </b>"))
// ["<b>", "and", "</b>"]
```

The "advanced" function allows to change some parameters of the template extraction.

```go

import "gopkg.in/vmarkovtsev/etalpmet.v1"


template := etalpmet.ReverseTemplateWithParameters(
	5,     // min block length
	false, // trim
	[]byte("<b> spam and eggs </b>"),
	[]byte("<b> ham and spam </b>"),
    []byte("<b> white and black </b>"))
// [nil, " and ", " </b>"]
```

Note `nil` which signals that there is text before the leftmost template block `" and "`.

## Installation

```
go get gopkg.in/vmarkovtsev/etalpmet.v1
```

The code supports building under Go >= 1.8.

## Contributions

...are pretty much welcome! See [contributing.md](contributing.md) and [code_of_conduct.md](code_of_conduct.md).

## License

Apache 2.0, see [LICENSE](LICENSE). It allows you to use this code in proprietary projects.
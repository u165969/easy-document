# easy-document

## What Is This?
This repository gives you a useful command.
The command enables you to create a beautiful HTML document from markdown.

### Details
This command receives a markdown file, then returns an HTML file.
The markdown file must satisfy some conditions for the beautiful HTML outputs.
You can see the conditions in "input/document.md".

## How To Use?
```
$ git clone https://github.com/u165969/easy-document.git
$ cd easy-document
$ go build
$ mv ./easy-document $GOPATH/bin/document
$ document
```

### Interface
You can see the interface of the document command as below:
```
$ document --help
```

## ToDos

* when you find only one h4 in input
* code-prettify

## References
* markdown parser: https://github.com/russross/blackfriday 
* sanitization: https://github.com/microcosm-cc/bluemonday

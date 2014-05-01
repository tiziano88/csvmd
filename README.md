# csvmd

CSV to Markdown converter command line utility.

## Installation

 1. You should have a correctly installed Go compiler environment and your personal workspace ($GOPATH). If you have no idea what **$GOPATH** is, take a look [here](http://golang.org/doc/code.html). Please make sure that your **$GOPATH/bin** is available in your **$PATH**. 

    Do these steps only if you understand why you need to do them:

    `export GOPATH=$HOME/goprojects`

    `export PATH=$PATH:$GOPATH/bin`

 2. Then you need to get the appropriate version of `csvmd`, for 6g/8g/5g compiler you can do this:

    `go get -u github.com/tiziano88/csvmd` (-u flag for "update")


## Usage

Just pipe your csv file throug `csvmd`:

```
 % echo "abc,def\n12345,67890" | csvmd
 |   abc |   def |
 | ----- | ----- |
 | 12345 | 67890 |
```

```
 % cat in.csv | csvmd > out.md
```

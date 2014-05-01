# csvmd

CSV to Markdown converter command line utility.

## Installation

Run `go get -u github.com/tiziano88/csvmd`.

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

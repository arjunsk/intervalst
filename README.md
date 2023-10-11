# Interval Search Tree

An interval tree is a data structure useful for storing values associated with intervals,
and efficiently search those values based on intervals that overlap with any given interval.
This generic implementation uses a self-balancing binary search tree algorithm, so searching
for any intersection has a worst-case time-complexity guarantee of <= 2 log N, where N is the number of items in the tree.

For more on interval trees, [see](https://en.wikipedia.org/wiki/Interval_tree). This library is extracted from [etcd](https://github.com/etcd-io/etcd/tree/main)

## Installing

```sh
$ go get github.com/rdleal/intervalst
```

## Usage

```go

```
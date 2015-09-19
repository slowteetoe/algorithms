# Algorithms
Some random algorithms implemented in Go, from a Coursera class on Algorithms

## Union-Find
*       find : check if 2 objects are in the same connected component
*       union : replace the component containing two objects with their union

#### quick-find
#### quick-union
#### weighted quick-union
#### TODO: weighted quick-union with path compression

## Search
#### binary search (don't use this btw, just use <a href="http://golang.org/pkg/sort/#Search" target="_blank">sort.Search()</a> )

## Calculator
#### twostack : use Djikstra's two-stack algorithm to handle fully parenthesised basic math expressions, e.g.  ( 1 + ( 10 * 15 ) ) - note: this is kinda hacky since golang doesn't have generics and I didn't feel like writing a bunch of boilerplate to handle a stack of strings vs. a stack of ints

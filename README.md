# geddit

> Note: the code here is what you get by following Andrew Gerrand's [Getting Started With Go](https://talks.golang.org/2012/tutorial.slide#1) tutorial.

A simple command-line program that fetches and displays the latest headlines from the golang page on Reddit.

The program will: 
* make an HTTP request to the Reddit API, 
* decode the JSON response into a Go data structure, and 
* print each link's title, URL, and number of comments.

Install with `go get github.com/joyrexus/reddit/geddit` and then try running
`geddit`.

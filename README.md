### Simple rating similarity calculator

#### Development

```go
go run main.go
```

Example of the input:

```go
testRatings := [][]int{
		{0, 4, 0, 4, 4, 0},
		{0, 1, 0, 0, 0, 4},
		{4, 0, 4, 0, 4, 2},
		{3, 5, 0, 1, 0, 0},
	}
```

Example of the result:

```
Top users (id + similarity score) with most `similar` ratings to user with id 3:
[{0 24} {2 12} {1 5}]

IDs of top 2 users with most `similar` ratings to user with id 3:
[0 2]
```

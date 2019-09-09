# frontmatter

Handle frontmatter in files.


## example
```go
data, rest, err := frontmatter.Parse(strings.NewReader(`---
name: frontmatter
version: 1.0.0
---

# hello
hello world!`), "---")

/*
data == `name: frontmatter
version: 1.0.0`
*/

/*
rest == `# hello
hello world!`
*/
```

## licence
MIT. More in LICENCE file

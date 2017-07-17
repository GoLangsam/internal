# internal

***Hide***, *what You don't want the audience to* ***see*** - **[Daryl](Daryl.md)**

In *[go](http://golang.org)* any package below/inside `internal` is **not** exported/published to any package above - (only to siblings and their descendants).

Thus: A great place to 'hide' stuff in plain sight :-)

Stuff that is either
- not good enough (yet) or
- not general enough (yet)
in order to be used in Your project.

Stuff such as:

## cmd/
- `glob` - simple CLI to play with `filepath/glob`

## container/

### ccsafe/
- `dotpath`

### oneway
no packages yet

## do/
- `dot`
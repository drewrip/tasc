# OpenTA

Compiler and Player for Interactive Stories


### Compiler

To build the compiler:

```
go build tasc-cli.go
```

Every different scenario/branch in your story gets its own .tam file in the format specificied in [syntax.tam](syntax.tam). For an example of how this works, check out the [samples folder](/samples).

To compile a story:

```
tasc-cli path/to/folder/with/tam/files/
```

* -o: Name of the project to generate (similar to the -o flag in gcc)

### Player

To build the player:

```
go build tap-cli.go
```

To play a story on the terminal:

```
tap-cli path/to/compiled/story/
```
___

**Feel free to submit a PR or contact me if you are interested in contributing. Writers also welcome for beta testing!**

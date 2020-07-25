# Pashua-Binding-Go

Binding code to use Pashua from Go programms

## Overview

This is the binding (glue code) for using Pashua from Go programs. Pashua is a macOS application for using native GUI dialog windows in various programming languages.

This code can be found in a GitHub repository at https://github.com/derLinkshaender/Pashua-Binding-Go. For examples in other programming languages, see https://github.com/BlueM/Pashua-Bindings.

Other related links:
* [Pashua homepage](https://www.bluem.net/jump/pashua)
* [Pashua repository on GitHub](https://github.com/BlueM/Pashua)

## Usage

The file `pashua.go` contains the binding code itself as a package

The folder `example` contains an example, which does not do much more 
than define how the dialog window should look like and use the functions in the `pashua.go` file.

You need to have [Pashua](https://www.bluem.net/jump/pashua) installed on your Mac 
to make use of this repository. 
The code expects Pashua.app in one of the “typical” locations, such as the global or 
the user’s “Applications” folder, or in the folder which contains `example.go`.

## Compatibility

This code 

It is compatible and has been tested with Pashua 0.11. 
It requires a version of Pashua that handles UTF-8 encoded input.


## Author

This code was written by Armin Hanisch. You can reach the author on [Twitter](https://twitter.com/derLinkshaender)


## License

MIT License, see the file [LICENSE](./LICENSE)

# Gompt: an enhanced prompt for your shell

[Gompt](https://github.com/RomainGiraud/gompt) is a [Powerline](https://github.com/Lokaltog/vim-powerline) like prompt,
 inspired by [Powerline-shell](https://github.com/b-ryan/powerline-shell).

*It is mainly a personal project to learn [Go Programming Language](https://golang.org/).*

I share it, feel free to use and to improve.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Setup](#setup)
  - [Bash](#bash)
  - [Fonts](#fonts)
- [Configuration](#configuration)
  - [Prompt](#prompt)
  - [Segments](#segments)
  - [Styles](#styles)
  - [Brushes](#brushes)
  - [Colors](#colors)
- [Miscellaneous](#miscellaneous)
  - [Name](#name)
- [Contributing](#contributing)
- [FAQ](#faq)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Setup

You have to [install](https://golang.org/doc/install) Go toolchain on your system.

Get source code with `go` tool:
```bash
go get github.com/RomainGiraud/gompt
```

Then you can modify your configuration in `main.go` file (see [Configuration](#configuration)).

Compile and install it:
```bash
go install
cp $GOPATH/bin/gompt ~/bin/
```

### Bash

Add these lines to your `~/.bashrc`:

```bash
function _update_ps1() {
    export PS1="$(gompt -s $?)"
}

export PROMPT_COMMAND="_update_ps1;"
```

### Fonts

You have to use a patched font, like [Powerline fonts](https://github.com/powerline/fonts)
 or [Nerd fonts](https://nerdfonts.com/).


## Configuration

If you do not know Go, I strongly encourage you to take the [Go Tour](https://golang.org/doc/#go_tour).

All configuration is done in `main.go` file in a hard-coded way.
If you want to change default prompt, you have to modify the source code.

In the future, I will try to create a tool to generate configuration (maybe another project!).

### Prompt

`prompt` is the main component. It is composed of multiple segments with different behaviours.

### Segments

In general, a segment has a style (can have more) and always output a string.

For now, following segments are available:
- `ComplexPath` - current path with different options
- `Hostname` - machine's hostname
- `Text` - preprocessed text: environment variable `${ENV}` and external command `${$cmd> ls}`.
- `Username` - current logged-in username

`Text` is one of the main segment, you can do lots of things just with this one:
- separator between segments
- replace `Hostname` and `Username` (with environment variable or command)

### Styles

A style has foreground and background brushes.
I will add more attribute in the future (bold, underline, blink...).

For now, there are two styles:
- `StyleStandard` - basic style
- `StyleChameleon` - useful to make a nice transition with certain utf-8 characters

### Brushes

A brush stores one or more colors.
- `UniBrush` - a monochromatic color
- `GradientBrush` - interpolate between two colors

### Colors

Colors can be specified in three formats:
- named colors: `Black`, `Red`, `Green`, `Yellow`, `Blue`, `Magenta`, `Cyan`, `White` and `Default`
- 256 color lookup table: integer between 0 and 255 (named colors are at the beginning of this table)
- true colors: in hexadecimal format `#f0f` or `#ff00ff`.

Interesting link: [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code).


## Miscellaneous

### Name

Gompt came from concatenation of `go` and `prompt`.

## Contributing

If you enjoying this tool and want to say something to me: hi <at> romaingiraud <dot> com

If you are a developer and want to contribute, submit a pull-request :)


## FAQ

**Another powerline-shell like?**

I did this tool on my free time to learn Go.
I encourage everybody to try to do the same type of project.
It is a real training, simple enough but with a lot of design thoughts.

*It is not a replacement for the real powerline-shell.*

**Why is there no configuration file?**

Initially, a configuration file existed.
But the maintenance and the level of imbrication (style > brush > color) were too heavy.

Then, I try to focus `gompt` on the rendering speed.
Reading a configuration file at each display is useless.
Generally, you do not change your configuration each seconds.

A possible solution is to cache the parsing of configuration.
But, for the moment, this is not the priority.

**Why does not GradientBrush work with named color?**

Named colors are limited to eight values.
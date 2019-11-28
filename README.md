# Round

Round is a command-line for rounding the images.

## Installation

```
go get github.com/mingrammer/round
```

## Usage

Round an image. (default option: **4** corners with **0.25** rounding rate)

```shell
$ round /path/to/image.png
```

Round multiple images.

```shell
$ round /path/to/image.png /path/to/image2.png
$ round *.png
```

Change the rounding rate. (1.0 means that make an image circular)

```shell
$ round -r 0.5 /path/to/image.png
$ round -r 1.0 /path/to/image.png
```

You can round only specific corners.

```shell
# tl (top left), tr (top right), bl (bottom left), br (bottom right)
$ round -c tl,br /path/to/image.png
```

Add prefix or suffix. (default suffix: **_rounded**)

```shell
$ round -p new -s _circle /path/to/square.png
```

##  Examples

|          | rect+jpg                                                 | rect+png                                                 | square+jpg                                                 | square+png                                                 |
| -------- | -------------------------------------------------------- | -------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| original | ![flower.jpg](examples/flower.jpg)                       | ![flower.png](examples/flower.png)                       | ![flower2.jpg](examples/flower2.jpg)                       | ![flower2.png](examples/flower2.png)                       |
| half     | ![flower_r0.5.jpg](examples/flower_r0.5.jpg)             | ![flower_r0.5.png](examples/flower_r0.5.png)             | ![flower2_r0.5.jpg](examples/flower2_r0.5.jpg)             | ![flower2_r0.5.png](examples/flower2_r0.5.png)             |
| leaf     | ![flower_leaf.jpg](examples/flower_leaf.jpg)             | ![flower_leaf.png](examples/flower_leaf.png)             | ![flower2_leaf.jpg](examples/flower2_leaf.jpg)             | ![flower2_leaf.png](examples/flower2_leaf.png)             |
| circle   | ![new_flower_ciecle.jpg](examples/new_flower_circle.jpg) | ![new_flower_ciecle.png](examples/new_flower_circle.png) | ![new_flower2_ciecle.jpg](examples/new_flower2_circle.jpg) | ![new_flower2_ciecle.png](examples/new_flower2_circle.png) |

| original   | half     | leaf                        | circle                      |
| ---------- | -------- | --------------------------- | --------------------------- |
| no options | `-r 0.5` | `-r 0.75 -c tl,br -s _leaf` | `-r 1.0 -p new_ -s _circle` |

Photo by Alex Blăjan on Unsplash

## License

MIT
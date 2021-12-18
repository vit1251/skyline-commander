# Skyline Commander

![Снимок экрана 2021-12-16 в 23 32 51](https://user-images.githubusercontent.com/110534/146445010-7ead6837-772f-4283-876b-fa9888aa6411.png)

## Requirements

  * Golang
  * libncurses

## Source code compile
  
### Prepare Debian bases distributions

    $ sudo apt-get install pkg-config
    $ go build
	
### Prepare MacOS X system

    $ brew install ncurses
    $ export PKG_CONFIG_PATH="/usr/local/opt/ncurses/lib/pkgconfig"
    $ go build
    

# Skyline Commander

![Снимок экрана 2021-12-16 в 23 05 24](https://user-images.githubusercontent.com/110534/146444786-b788c398-00ea-4000-a9e6-49d69c25d954.png)

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
    

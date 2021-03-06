# go-to-commons

Upload media files to Wikimedia Commons.

## Usage

Tip: [Create a bot password](https://commons.wikimedia.org/wiki/Special:BotPasswords?wpappId=go-to-commons) instead of using your main Wikimedia password!

```sh
$ go-to-commons -username $MEDIAWIKI_USERNAME -password $MEDIAWIKI_PASSWORD -file go-to-commons.png -filename go-to-commons.png -text "=={{int:filedesc}}==\n..."
```

```sh
$ ./go-to-commons -help
Usage of ./go-to-commons:
  -api string
        MediaWiki API URL (default "https://commons.wikimedia.org/w/api.php")
  -comment string
        Upload comment (default "Uploaded with go-to-commons")
  -file string
        Media file to upload
  -filename string
        Filename on Wikimedia Commons
  -overwrite
        Overwrite existing file on Wikimedia Commons
  -password string
        Wikimedia password (default "$MEDIAWIKI_PASSWORD")
  -text string
        Wikitext of media file on Wikimedia Commons or (if specified as @file.txt, the text is read from file.txt)
  -username string
        Wikimedia username (default "$MEDIAWIKI_USERNAME")
  -version
        Display version information
```

## Open Source Libraries

- https://github.com/cgt/go-mwclient (CC0 licensed)
- https://github.com/antonholmquist/jason (MIT licensed)
- https://github.com/mrjones/oauth (MIT licensed)

## Author and License

- Author: Simon Legner ([simon04](https://github.com/simon04))
- License: [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl.html)

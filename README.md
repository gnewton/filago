# filago
Monitor open files (via /proc/pid) of a process.
Lists time when it first sees the file open, and when it is no longer open.


## Building
```
$ git clone https://github.com/gnewton/filago.git
$ cd filago
$ go build
```

### Usage:
```
$ filago pid
```

### Output

Sample output (tab separated):

```
$ filago 1792
2017-04-06T00:08:57.959133-04:00	open	/home/bsmith/.xsession-errors
2017-04-06T00:09:00.519095-04:00	open	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:00.529033-04:00	close	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:14.839044-04:00	open	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:14.849037-04:00	close	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:35.089172-04:00	open	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:35.099067-04:00	close	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg

2017-04-06T00:09:42.029112-04:00	open	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:09:42.039114-04:00	close	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:10:18.379028-04:00	open	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:10:18.38908-04:00	close	/usr/share/emacs/24.5/etc/images/icons/hicolor/scalable/apps/emacs.svg
2017-04-06T00:10:33.629016-04:00	open	/home/bsmith/.emacs
2017-04-06T00:10:33.639025-04:00	close	/home/bsmith/.emacs
2017-04-06T00:10:40.349049-04:00	open	/home/bsmith/.emacs
2017-04-06T00:10:40.35902-04:00	close	/home/bsmith/.emacs
2017-04-06T00:14:28.219021-04:00	open	/home/bsmith/work/flogg/main.go
2017-04-06T00:14:28.229087-04:00	close	/home/bsmith/work/flogg/main.go

```

## Caveat
`filago` checks every 10 milliseconds, so can miss open/closes smaller than this granularity.
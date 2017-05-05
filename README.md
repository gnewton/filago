# filago
Monitor open files (via /proc/pid) of a process.
Lists time when it first sees the file open (which may be later than when the file actually opened, especially at startup of filago), and when it is no longer open.


## Building
```
$ git clone https://github.com/gnewton/filago.git
$ cd filago
$ go build
```

### Usage:
```
$ ./filago --help
Usage of ./filago:
  -d uint
    	Time granularity for checking files, in milliseconds (default 10)
  -r	Show only real files, i.e. no pipes, sockets, etc.
$
```

### Output

Sample output (tab separated):

```
$ filago 1792
2017-05-05T00:01:25.002325-04:00 open pipe:[25629]
2017-05-05T00:01:25.002325-04:00 open /home/guser/.xsession-errors
2017-05-05T00:01:25.002325-04:00 open socket:[28116]
2017-05-05T00:01:25.002325-04:00 open /home/guser/.local/share/gvfs-metadata/home-405fe1ec.log (deleted)
2017-05-05T00:01:25.002325-04:00 open socket:[28117]
2017-05-05T00:01:25.002325-04:00 open pipe:[28118]
2017-05-05T00:01:25.002325-04:00 open pipe:[96074]
2017-05-05T00:01:25.002325-04:00 open /home/guser/install/firefox/omni.ja
2017-05-05T00:01:25.002325-04:00 open socket:[96077]
2017-05-05T00:01:25.002325-04:00 open /home/guser/.mozilla/firefox/uqsr6u0q.default/features/{67ce503f-44fb-407b-86fc-eb27d1da49b7}/shield-recipe-client@mozilla.org.xpi
2017-05-05T00:01:25.002325-04:00 open /home/guser/install/firefox/browser/omni.ja
2017-05-05T00:01:25.002325-04:00 open pipe:[22385]
2017-05-05T00:01:25.002325-04:00 open /home/guser/.cache/event-sound-cache.tdb.9bf539dba0e34f7aaf456bd844b6826e.x86_64-redhat-linux-gnu
2017-05-05T00:01:25.002325-04:00 open pipe:[29953]
2017-05-05T00:01:25.002325-04:00 open socket:[20246]
2017-05-05T00:01:25.002325-04:00 open pipe:[29954]
2017-05-05T00:01:25.002325-04:00 open socket:[20407]
2017-05-05T00:01:25.002325-04:00 open socket:[27198]
2017-05-05T00:01:25.002325-04:00 open anon_inode:[eventfd]
2017-05-05T00:01:25.002325-04:00 open anon_inode:inotify
2017-05-05T00:01:25.002325-04:00 open socket:[26094]
```

Only real files:
```
$ ./filago   -r 1943
2017-05-05T00:02:29.39932-04:00 open /home/guser/.xsession-errors
2017-05-05T00:02:29.39932-04:00 open /home/guser/.local/share/gvfs-metadata/home-405fe1ec.log (deleted)
2017-05-05T00:02:29.39932-04:00 open /home/guser/install/firefox/omni.ja
2017-05-05T00:02:29.39932-04:00 open /home/guser/.mozilla/firefox/uqsr6u0q.default/features/{67ce503f-44fb-407b-86fc-eb27d1da49b7}/shield-recipe-client@mozilla.org.xpi
2017-05-05T00:02:29.39932-04:00 open /home/guser/install/firefox/browser/omni.ja
2017-05-05T00:02:29.39932-04:00 open /home/guser/.cache/event-sound-cache.tdb.9bf539dba0e34f7aaf456bd844b6826e.x86_64-redhat-linux-gnu
2017-05-05T00:02:29.39932-04:00 open /home/guser/.mozilla/firefox/uqsr6u0q.default/cert8.db
2017-05-05T00:02:29.39932-04:00 open /home/guser/.mozilla/firefox/uqsr6u0q.default/key3.db
2017-05-05T00:02:29.39932-04:00 open /home/guser/install/firefox/browser/features/aushelper@mozilla.org.xpi
2017-05-05T00:02:29.39932-04:00 open /home/guser/install/firefox/browser/features/e10srollout@mozilla.org.xpi
2017-05-05T00:02:29.39932-04:00 open /home/guser/install/firefox/browser/features/firefox@getpocket.com.xpi
2017-05-05T00:02:29.39932-04:00 open /home/guser/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-MnnxcxisBPnSXQ-eff@jetpack.xpi
2017-05-05T00:02:29.39932-04:00 open /home/guserx/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-uqbSKwXpf2K6yl@jetpack.xpi
```

## Caveat
`filago` checks every 10 milliseconds, so can miss open/closes smaller than this granularity.
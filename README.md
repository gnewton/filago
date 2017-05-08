# filago
Monitor open files (via /proc/pid) of a process.
Lists time when it first sees the file open (which may be later than when the file actually opened, especially at startup of filago), and when it is no longer open.

This includes tcp and unix sockets.

*Only works on Linux*


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
    	Time granularity for checking files, in milliseconds (default 100)
  -r	Show only real files, i.e. no pipes, sockets, etc.
$
```

### Output
#### File:
```
timestamp [open|close] path
```

#### Socket:
##### General:
```
timestamp [open|close] path
```

##### Unix socket:
```
timestamp [open|close] path unix [path|-]
```

##### tcp socket:
```
timestamp [open|close] path tcp ip:port fqhn
```

Hostnames are looked-up and cached.


Sample output on firefox:

```
$ ./filago 2262
2017-05-07T12:40:40.663016-04:00 open socket:[31037] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[31038] unix -
2017-05-07T12:40:40.663016-04:00 open pipe:[31039]
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/omni.ja
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/browser/omni.ja
2017-05-07T12:40:40.663016-04:00 open pipe:[31040]
2017-05-07T12:40:40.663016-04:00 open pipe:[30144]
2017-05-07T12:40:40.663016-04:00 open socket:[31041]
2017-05-07T12:40:40.663016-04:00 open pipe:[26138]
2017-05-07T12:40:40.663016-04:00 open socket:[35650] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[33988] unix -
2017-05-07T12:40:40.663016-04:00 open anon_inode:[eventfd]
2017-05-07T12:40:40.663016-04:00 open anon_inode:inotify
2017-05-07T12:40:40.663016-04:00 open socket:[27647] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[30153] unix -
2017-05-07T12:40:40.663016-04:00 open pipe:[29118]
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cert8.db
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/key3.db
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/browser/features/aushelper@mozilla.org.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/browser/features/e10srollout@mozilla.org.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/browser/features/firefox@getpocket.com.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-MnnxcxisBPnSXQ-eff@jetpack.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-uqbSKwXpf2K6yl@jetpack.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jsonview@brh.numbera.com.xpi
2017-05-07T12:40:40.663016-04:00 open socket:[30131] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/features/{67ce503f-44fb-407b-86fc-eb27d1da49b7}/shield-recipe-client@mozilla.org.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/uBlock0@raymondhill.net.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/install/firefox/browser/features/webcompat@mozilla.org.xpi
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/permissions.sqlite
2017-05-07T12:40:40.663016-04:00 open socket:[30151] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-wal
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-shm
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/content-prefs.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/formhistory.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extension-data/ublock0.sqlite
2017-05-07T12:40:40.663016-04:00 open socket:[30154] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-wal
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-shm
2017-05-07T12:40:40.663016-04:00 open socket:[30155] unix -
2017-05-07T12:40:40.663016-04:00 open pipe:[30156]
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.local/share/gvfs-metadata/home (deleted)
2017-05-07T12:40:40.663016-04:00 open socket:[31081] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[25560] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[1630334] tcp 52.84.139.212:80 server-52-84-139-212.yto50.r.cloudfront.net.
2017-05-07T12:40:40.663016-04:00 open socket:[30157] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[30163] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[30159] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[25564] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[30161] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-wal
2017-05-07T12:40:40.663016-04:00 open socket:[30140] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-shm
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.cache/mozilla/firefox/uqsr6u0q.default/OfflineCache/index.sqlite
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.cache/mozilla/firefox/uqsr6u0q.default/startupCache/startupCache.8.little
2017-05-07T12:40:40.663016-04:00 open socket:[1610905] tcp 107.161.13.237:443 cache.google.com.
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/.parentlock
2017-05-07T12:40:40.663016-04:00 open pipe:[511788]
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.local/share/gvfs-metadata/home-6bd0f44f.log (deleted)
2017-05-07T12:40:40.663016-04:00 open socket:[416916] tcp 198.252.206.25:443 stackoverflow.com.
2017-05-07T12:40:40.663016-04:00 open anon_inode:[eventpoll]
2017-05-07T12:40:40.663016-04:00 open socket:[511791] unix -
2017-05-07T12:40:40.663016-04:00 open /home/gnewton/.cache/event-sound-cache.tdb.9bf539dba0e34f7aaf456bd844b6826e.x86_64-redhat-linux-gnu
2017-05-07T12:40:40.663016-04:00 open socket:[1265977] unix -
2017-05-07T12:40:40.663016-04:00 open socket:[1574737] tcp 173.194.175.189:443 qs-in-f189.1e100.net.
2017-05-07T12:41:21.110231-04:00 close socket:[1630334] tcp 52.84.139.212:80 server-52-84-139-212.yto50.r.cloudfront.net.
2017-05-07T12:42:35.083678-04:00 open socket:[1635598]
2017-05-07T12:42:35.155006-04:00 close socket:[1635598]
2017-05-07T12:42:35.155006-04:00 open socket:[1630691] tcp 172.217.6.229:443 lga25s55-in-f5.1e100.net.
```

Only real files:
```
$ ./filago -r 2262
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/omni.ja
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/browser/omni.ja
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cert8.db
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/key3.db
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/browser/features/aushelper@mozilla.org.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/browser/features/e10srollout@mozilla.org.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/browser/features/firefox@getpocket.com.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-MnnxcxisBPnSXQ-eff@jetpack.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-uqbSKwXpf2K6yl@jetpack.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jsonview@brh.numbera.com.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/features/{67ce503f-44fb-407b-86fc-eb27d1da49b7}/shield-recipe-client@mozilla.org.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/uBlock0@raymondhill.net.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/install/firefox/browser/features/webcompat@mozilla.org.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/https-everywhere-eff@eff.org.xpi
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/permissions.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-wal
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-shm
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/content-prefs.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/formhistory.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/extension-data/ublock0.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-wal
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-shm
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.local/share/gvfs-metadata/home (deleted)
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-wal
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-shm
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.cache/mozilla/firefox/uqsr6u0q.default/OfflineCache/index.sqlite
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.cache/mozilla/firefox/uqsr6u0q.default/startupCache/startupCache.8.little
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.mozilla/firefox/uqsr6u0q.default/.parentlock
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.local/share/gvfs-metadata/home-6bd0f44f.log (deleted)
2017-05-07T12:44:41.405442-04:00 open /home/gnewton/.cache/event-sound-cache.tdb.9bf539dba0e34f7aaf456bd844b6826e.x86_64-redhat-linux-gnu
```

## Caveat
*`filago` checks every 100 milliseconds (modifiable by command line -d), so can miss open/closes smaller than this temporal granularity.*

## TODO
- Allow turning off of DNS hostname lookups (can be slow/expensive) & some servers not on DNS
- Unix and tcp socket info in /proc/net/[unix|tcp] read too often: need to change so not read so much...
- JSON output

# filago
Monitor open files (via /proc/pid) of a process.
Lists time when it first sees the file open (which may be later than when the file actually opened, especially at startup of filago), and when it is no longer open.
Stop running when target process ends.

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
  -j	Output json (complete json per line)
  -l	Turn on hostname lookup (default is a "-"
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


Sample default output on firefox:

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

json output with hostnames looked up:
```
$ ./filago -l -j 3737
{"filename":"socket:[95066]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803e6c99f80","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":95066}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[95067]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803e6c99c00","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":95067}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"pipe:[95068]","type":"pipe","status":"open","mod_time":"2017-06-02T23:20:10.616170585-04:00"}
{"filename":"/home/gnewton/install/firefox/omni.ja","type":"file","status":"open","mod_time":"2017-05-27T23:24:59.740556075-04:00"}
{"filename":"/home/gnewton/install/firefox/browser/omni.ja","type":"file","status":"open","mod_time":"2017-05-27T23:25:00.173556279-04:00"}
{"filename":"pipe:[95069]","type":"pipe","status":"open","mod_time":"2017-06-02T23:19:57.685144488-04:00"}
{"filename":"pipe:[110949]","type":"pipe","status":"open","mod_time":"2017-06-02T23:19:57.685144488-04:00"}
{"filename":"socket:[95070]","type":"other","status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"pipe:[102268]","type":"pipe","status":"open","mod_time":"2017-06-02T23:20:13.688176034-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extension-data/ublock0.sqlite","type":"file","status":"open","mod_time":"2017-06-02T23:19:09.638842227-04:00"}
{"filename":"socket:[110951]","type":"unix","socket_info":{"unix_socket":{"num":"ffff880403259c00","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":110951}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"anon_inode:[eventfd]","type":"anon_inode","status":"open","mod_time":"2017-06-02T22:20:43.183999937-04:00"}
{"filename":"anon_inode:inotify","type":"anon_inode","status":"open","mod_time":"2017-06-02T22:20:43.183999937-04:00"}
{"filename":"socket:[110956]","type":"unix","socket_info":{"unix_socket":{"num":"ffff880403258000","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":110956}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[106057]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803ecffb100","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":106057}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"pipe:[106054]","type":"pipe","status":"open","mod_time":"2017-06-02T23:20:13.696176048-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/cert8.db","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.630087523-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/key3.db","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.630087523-04:00"}
{"filename":"/home/gnewton/install/firefox/browser/features/aushelper@mozilla.org.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:25:00.177556281-04:00"}
{"filename":"/home/gnewton/install/firefox/browser/features/e10srollout@mozilla.org.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:25:00.177556281-04:00"}
{"filename":"/home/gnewton/install/firefox/browser/features/firefox@getpocket.com.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:25:00.177556281-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-MnnxcxisBPnSXQ-eff@jetpack.xpi","type":"file","status":"open","mod_time":"2017-05-09T21:22:57-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jid1-uqbSKwXpf2K6yl@jetpack.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:32:57-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/jsonview@brh.numbera.com.xpi","type":"file","status":"open","mod_time":"2017-01-26T20:21:02-05:00"}
{"filename":"socket:[102264]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803ece40e00","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":102264}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/features/{fc212d3b-cc8b-4e09-8225-838f8501fd5b}/shield-recipe-client@mozilla.org.xpi","type":"file","status":"open","mod_time":"2017-05-28T23:32:58-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/uBlock0@raymondhill.net.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:32:57-04:00"}
{"filename":"/home/gnewton/install/firefox/browser/features/webcompat@mozilla.org.xpi","type":"file","status":"open","mod_time":"2017-05-27T23:25:00.173556279-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/extensions/https-everywhere-eff@eff.org.xpi","type":"file","status":"open","mod_time":"2017-05-28T08:52:05-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/permissions.sqlite","type":"file","status":"open","mod_time":"2017-01-29T08:36:35.238948679-05:00"}
{"filename":"socket:[106055]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803ecffb480","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":106055}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.268085245-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-wal","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.261085201-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/places.sqlite-shm","type":"file","status":"open","mod_time":"2017-06-02T23:19:58.393148896-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/content-prefs.sqlite","type":"file","status":"open","mod_time":"2016-01-27T22:55:20.649349151-05:00"}
{"filename":"socket:[113056]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf247380","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0005","st":"03","inode":113056}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.625087491-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-wal","type":"file","status":"open","mod_time":"2017-06-02T23:19:59.225150374-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/cookies.sqlite-shm","type":"file","status":"open","mod_time":"2017-06-02T23:19:59.225150374-04:00"}
{"filename":"socket:[102313]","type":"tcp","socket_info":{"inode":102313,"tcp_socket":{"sl":3,"local_port":40862,"remote_port":443,"st":1,"uid":1000,"extra":"1","rem_address":"151.101.0.201","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113057]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf247700","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0005","st":"03","inode":113057}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"pipe:[113058]","type":"pipe","status":"open","mod_time":"2017-06-02T23:19:59.486150837-04:00"}
{"filename":"socket:[95144]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803e6c99880","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":95144}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113356]","type":"other","status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[110966]","type":"unix","socket_info":{"unix_socket":{"num":"ffff88040325aa00","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":110966}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113059]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf246c80","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":113059}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113065]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf245b00","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":113065}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113061]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf246900","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":113061}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[106058]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8803ecffc600","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":106058}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113063]","type":"unix","socket_info":{"unix_socket":{"num":"ffff8800bf246200","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":113063}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite","type":"file","status":"open","mod_time":"2017-06-02T23:19:48.669087768-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-wal","type":"file","status":"open","mod_time":"2017-06-02T23:19:59.84315147-04:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/webappsstore.sqlite-shm","type":"file","status":"open","mod_time":"2017-06-02T23:19:59.84315147-04:00"}
{"filename":"socket:[110944]","type":"unix","socket_info":{"unix_socket":{"num":"ffff88040325bb80","refcount":"00000003","protocol":"00000000","flags":"00000000","type":"0001","st":"03","inode":110944}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.mozilla/firefox/uqsr6u0q.default/.parentlock","type":"file","status":"open","mod_time":"2017-06-02T23:19:57.667144375-04:00"}
{"filename":"anon_inode:[eventpoll]","type":"anon_inode","status":"open","mod_time":"2017-06-02T22:20:43.183999937-04:00"}
{"filename":"socket:[113356]","type":"other","status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113400]","type":"tcp","socket_info":{"inode":113400,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58785,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":13,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113400]","type":"tcp","socket_info":{"inode":113400,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58785,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":13,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117372]","type":"tcp","socket_info":{"inode":117372,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58786,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117372]","type":"tcp","socket_info":{"inode":117372,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58786,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117397]","type":"tcp","socket_info":{"inode":117397,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58787,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":18,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117397]","type":"tcp","socket_info":{"inode":117397,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58787,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":18,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117402]","type":"tcp","socket_info":{"inode":117402,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58788,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117402]","type":"tcp","socket_info":{"inode":117402,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58788,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113445]","type":"tcp","socket_info":{"inode":113445,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58789,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":18,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113445]","type":"tcp","socket_info":{"inode":113445,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58789,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":18,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113446]","type":"tcp","socket_info":{"inode":113446,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58790,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":8,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113446]","type":"tcp","socket_info":{"inode":113446,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58790,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":8,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113447]","type":"tcp","socket_info":{"inode":113447,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58791,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":20,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113447]","type":"tcp","socket_info":{"inode":113447,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58791,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":20,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[114875]","type":"tcp","socket_info":{"inode":114875,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58792,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[114875]","type":"tcp","socket_info":{"inode":114875,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":6,"local_port":58792,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":19,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117508]","type":"tcp","socket_info":{"inode":117508,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58793,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":7,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[117508]","type":"tcp","socket_info":{"inode":117508,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58793,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":7,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[109318]","type":"tcp","socket_info":{"inode":109318,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58794,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":20,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[109318]","type":"tcp","socket_info":{"inode":109318,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":5,"local_port":58794,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":20,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[109319]","type":"tcp","socket_info":{"inode":109319,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58795,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":21,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[109319]","type":"tcp","socket_info":{"inode":109319,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":4,"local_port":58795,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":21,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[106227]","type":"tcp","socket_info":{"inode":106227,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":1,"local_port":58796,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":21,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[106227]","type":"tcp","socket_info":{"inode":106227,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":1,"local_port":58796,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":21,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[111430]","type":"tcp","socket_info":{"inode":111430,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":2,"local_port":58797,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":22,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[111430]","type":"tcp","socket_info":{"inode":111430,"tcp_socket":{"rem_hostname":"a23-63-227-177.deploy.static.akamaitechnologies.com.","sl":2,"local_port":58797,"remote_port":80,"st":1,"tr":1,"uid":1000,"tx_queue":296,"tm_when":22,"extra":"3","rem_address":"23.63.227.177","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[106273]","type":"tcp","socket_info":{"inode":106273,"tcp_socket":{"rem_hostname":"ec2-34-211-94-203.us-west-2.compute.amazonaws.com.","sl":4,"local_port":47111,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":193,"tm_when":21,"extra":"3","rem_address":"34.211.94.203","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[110114]","type":"tcp","socket_info":{"inode":110114,"tcp_socket":{"sl":4,"local_port":34260,"remote_port":80,"st":1,"tr":2,"uid":1000,"tm_when":987,"extra":"3","rem_address":"72.21.91.29","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"/home/gnewton/.cache/mozilla/firefox/uqsr6u0q.default/OfflineCache/index.sqlite","type":"file","status":"open","mod_time":"2015-05-27T06:40:57.729790739-04:00"}
{"filename":"socket:[106278]","type":"tcp","socket_info":{"inode":106278,"tcp_socket":{"sl":5,"local_port":42312,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":187,"tm_when":20,"extra":"3","rem_address":"151.101.193.140","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[110116]","type":"tcp","socket_info":{"inode":110116,"tcp_socket":{"sl":2,"local_port":48623,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":197,"tm_when":20,"extra":"3","rem_address":"151.101.1.140","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[110116]","type":"tcp","socket_info":{"inode":110116,"tcp_socket":{"sl":2,"local_port":48623,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":197,"tm_when":20,"extra":"3","rem_address":"151.101.1.140","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113560]","type":"tcp","socket_info":{"inode":113560,"tcp_socket":{"sl":6,"local_port":48628,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":180,"rx_queue":258,"tm_when":21,"extra":"3","rem_address":"151.101.1.140","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113560]","type":"tcp","socket_info":{"inode":113560,"tcp_socket":{"sl":6,"local_port":48628,"remote_port":443,"st":1,"tr":1,"uid":1000,"tx_queue":180,"rx_queue":258,"tm_when":21,"extra":"3","rem_address":"151.101.1.140","local_address":"192.168.0.101"}},"status":"close","mod_time":"1969-12-31T19:00:00-05:00"}
{"filename":"socket:[113617]","type":"tcp","socket_info":{"inode":113617,"tcp_socket":{"sl":5,"local_port":48631,"remote_port":443,"st":1,"tr":2,"uid":1000,"tm_when":999,"extra":"3","rem_address":"151.101.1.140","local_address":"192.168.0.101"}},"status":"open","mod_time":"1969-12-31T19:00:00-05:00"}
```

## Caveat
*`filago` checks every 100 milliseconds (modifiable by command line -d), so can miss open/closes smaller than this temporal granularity.*

## TODO
- Unix and tcp socket info in /proc/net/[unix|tcp] read too often: need to change so not read so much...


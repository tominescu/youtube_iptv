# youtube_iptv
Transfer youtube live streaming to http live stream for IPTV

# Usage:

## Server
Build this app and run it on a server within the same lan

## TV
Add http://server_addr:8080/index.m3u8?id=xxxxx&q=1 to iptv.txt

Param id, e.g.: for https://www.youtube.com/watch?v=sI8KEnewCx4 id is sI8KEnewCx4
Param q is the quality, 1 means the best, larger number means worse quality

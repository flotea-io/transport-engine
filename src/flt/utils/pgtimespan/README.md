1. plgo must be in /etc/environment PATH
2. plgo ./ under pgtimespan, that generate BUILD
3. cd BUILD and make install with_llvm=no
4. Plugin should be installed
5. CREATE EXTENSION pgtimespan;

HowTo Build&Test:
```
/app/src/flt/utils/pgtimespan# plgo
/app/src/flt/utils/pgtimespan# cd build
/app/src/flt/utils/pgtimespan/build# make install with_llvm=no
/app/src/flt/utils/pgtimespan/build# service postgresql restart

DROP EXTENSION pgtimespan; // if exists before
CREATE EXTENSION pgtimespan; // updated
```

Example of using:

```SELECT * FROM gtfs_routes WHERE IsOpen(schedule::TEXT, ARRAY[2019,10,12]);```


Code for installation (can be automated later in vagrant provision)
```
/bin/mkdir -p '/usr/share/postgresql/11/extension'
/bin/mkdir -p '/usr/share/postgresql/11/extension'
/bin/mkdir -p '/usr/lib/postgresql/11/lib'
/usr/bin/install -c -m 644 .//pgtimespan.control '/usr/share/postgresql/11/extension/'
/usr/bin/install -c -m 644 .//pgtimespan--0.1.sql  '/usr/share/postgresql/11/extension/'
/usr/bin/install -c -m 755  pgtimespan.so '/usr/lib/postgresql/11/lib/'
```

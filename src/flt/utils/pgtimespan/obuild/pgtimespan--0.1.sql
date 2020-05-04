-- complain if script is sourced in psql, rather than via CREATE EXTENSION
\echo Use "CREATE EXTENSION pgtimespan" to load this file. \quit
CREATE OR REPLACE FUNCTION IsOpen(json text,event bigint[])
RETURNS boolean AS
'$libdir/pgtimespan', 'IsOpen'
LANGUAGE c VOLATILE STRICT;

CREATE OR REPLACE FUNCTION StartInDay(schedule text,day bigint[])
RETURNS boolean AS
'$libdir/pgtimespan', 'StartInDay'
LANGUAGE c VOLATILE STRICT;


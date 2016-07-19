#!/usr/bin/perl

#use strict;
use warnings;
 my $ip1="10.1.0.127"; # 10.1.0.127
 my $p1=161;
 my $c1="public";
 my $ip2="10.1.0.123"; # 10.1.0.123
 my $p2=161;
 my $c2="public";
 my $OID_sysUpTime = '1.3.6.1.2.1.1.3.0';
 my $in1="1.3.6.1.2.1.2.2.1.10.1";
 my $out1="1.3.6.1.2.1.2.2.1.16.1";
 my $in8="1.3.6.1.2.1.2.2.1.10.8";
 my $out8="1.3.6.1.2.1.2.2.1.16.8";
 my $IN1="1.3.6.1.2.1.2.2.1.10.1";
 my $OUT1="1.3.6.1.2.1.2.2.1.16.1";
 my $IN2="1.3.6.1.2.1.2.2.1.10.2";
 my $OUT2="1.3.6.1.2.1.2.2.1.16.2";
 my @oid1 =($OID_sysUpTime,$in1,$out1,$in8,$out8);
 my @oid2 =($OID_sysUpTime,$IN1,$OUT1,$IN2,$OUT2);

open (FOO,"/mnt/LONTAS/ExpControl/dire15/info/details.txt") ||
    die "ERROR Unable to open file: $!\n";

my $last;
my $first = <FOO>;

while (<FOO>) { $last = $_ }
close FOO;
my @words = split / /,$last;
my @exp = split (':',$words[0]);
my @run = split (':',$words[1]);
my @sampling = split(':',$words[6])
my $runid = $run[1];
my $expid = $exp[1];
my $sampling_interval = $sampling[1];



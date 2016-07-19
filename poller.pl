#!/usr/bin/perl
use strict;
use warnings;
use Net::SNMP qw(snmp_dispatcher ticks_to_time);
use Time::HiRes qw(usleep gettimeofday);
use JSON;
my $script="/home/ats/dire15/thesis/config.pl";
do "$script";
#my ($second,$microsecond)=gettimeofday;
#open FILE1, ">>",""
print "$expid\n";

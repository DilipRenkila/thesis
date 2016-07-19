#!/usr/bin/perl
use strict;
use warnings;
use Net::SNMP qw(snmp_dispatcher ticks_to_time);
use Time::HiRes qw(usleep gettimeofday);
use JSON;
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
my @sampling = split(':',$words[6]);
my $runid = $run[1];
my $expid = $exp[1];
my $sampling_interval = $sampling[1];
my $foo = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/in-%d-%d.txt",$expid,$runid;
my $bar = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/out-%d-%d.txt",$expid,$runid;

#open FOO, ">>",$foo or die $!;
#open BAR, ">>",$bar or die $!;

my $i1=1;
my $i2=2;

my ($session1,$error1) = Net::SNMP->session(
    -hostname => $ip1,
    -community => $c1,
    -port => $p1,
    -nonblocking => 1,
    -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric

);

if (!defined $session1) {

    printf "ERROR: %s.\n", $error1;
    exit 1;
}

my ($session2,$error2) = Net::SNMP->session(
    -hostname => $ip2,
    -community => $c2,
    -port => $p2,
    -nonblocking => 1,
    -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric

);

if (!defined $session2) {

    printf "ERROR: %s.\n", $error2;
    exit 1;
}


while(1) {

    my $result1 = $session1->get_request(-varbindlist => \@oid1,-callback =>[ \&table_callback ,$i1]);
    if (!defined $result1) {
      printf "ERROR: %s.\n", $session1->error();
      $session1->close();
      exit 1;
   }

    my $result2 = $session2->get_request(-varbindlist => \@oid2,-callback =>[ \&table_callback ,$i2]);
    if (!defined $result2) {
      printf "ERROR: %s.\n", $session2->error();
      $session2->close();
      exit 1;
   }

    my ($seconds, $microseconds) = gettimeofday;
    my $time1 = $seconds * 1000000 + $microseconds ;
    snmp_dispatcher();
    my ($seconds1, $microseconds1) = gettimeofday;
    my $time2 = $seconds1 * 1000000 + $microseconds1 ;
    my $delta = ($time2-$time1); #in microsecond
    my $Delta = $sampling_interval*1000000 - $delta ;

    if ( $Delta == abs($Delta) ) {
        usleep   $Delta;
    } else {
        usleep 0;
    }
}

sub table_callback()
{
    my ($session,$host,$x) = @_;
    my @names = $session->var_bind_names();
    my $next  = undef;
    my $InOctet="1.3.6.1.2.1.2.2.1.10";
    my $list = $session->var_bind_list();
    my ($second, $microsecond) = gettimeofday;
    my $time = sprintf('%d.%d',$second,$microsecond);
    open FOO, ">>",$foo or die $!;
    open BAR, ">>",$bar or die $!;
    if (!defined $list) {
         printf "ERROR: %s\n", $session->error();
         return;
    }
    my $JSON = JSON->new->utf8;

    if ($host == $i1)
    {
        my $In1    = $list->{$in1};
        my $Out1    = $list->{$out1};
        my $In8    = $list->{$in8};
        my $Out8    = $list->{$out8};
        my $uptime = $list->{$OID_sysUpTime}; # in microseconds
           $uptime = $uptime*0.01;
        my $Output = {Switch_IP=>$ip1,Uptime=>$uptime,unixtime=>$time,IN_Interface_1=>$In1,OUT_Interface_1=>$Out1,IN_Interface_8=>$In8,OUT_Interface_8=>$Out8};
        my $json = $JSON->encode($Output) ;
        print FOO "$json\n";
    }

    if ($host == $i2)
    {
        my $In1     = $list->{$IN1};
        my $Out1    = $list->{$OUT1};
        my $In2     = $list->{$IN2};
        my $Out2    = $list->{$OUT2};
        my $uptime = $list->{$OID_sysUpTime}; # in microseconds
           $uptime = $uptime*0.01;
        my $Output = {Switch_IP=>$ip2,Uptime=>$uptime,unixtime=>$time,IN_Interface_1=>$In1,OUT_Interface_1=>$Out1,IN_Interface_2=>$In2,OUT_Interface_2=>$Out2};
        my $json = $JSON->encode($Output) ;
#       print "$json\n";
        print BAR "$json\n";
#        print FILE2 "$ip2,$uptime,$seconds,$In1,out1=$Out1,in2=$In2,$Out2\n";
   }

}


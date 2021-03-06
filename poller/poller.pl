#!/usr/bin/perl
use strict;
use warnings FATAL => 'all';
use YAML::Tiny;
use Net::SNMP qw(snmp_dispatcher ticks_to_time);
use Time::HiRes qw(usleep gettimeofday);
use JSON;

# Open the config yaml file
my $yaml = YAML::Tiny->read( '/mnt/LONTAS/ExpControl/dire15/thesis/poller/config.yml' );

# Storing the properties of yaml file into variables
my $ingress_switch  = $yaml->[0]->{ingress_switch_details}->{switch_management_ip};
my $egress_switch  = $yaml->[0]->{egress_switch_details}->{switch_management_ip};
my $ingress_interface = $yaml->[0]->{ingress_switch_details}->{interface_id};
my $egress_interface = $yaml->[0]->{egress_switch_details}->{interface_id};

my $sysUpTime = '1.3.6.1.2.1.1.3.0';
my @ingress_oid = ($sysUpTime,(sprintf "1.3.6.1.2.1.2.2.1.10.%s",$ingress_interface));
my @egress_oid =  ($sysUpTime,(sprintf "1.3.6.1.2.1.2.2.1.16.%s",$egress_interface));

open (FOO,"/mnt/LONTAS/ExpControl/dire15/info/details.txt") ||
    die "ERROR Unable to open file: $!\n";

my $last;
my $first = <FOO>;

while (<FOO>) { $last = $_ }
close FOO;
my @words = split / /,$last;
my @exp = split (':',$words[0]);
my @run = split (':',$words[1]);
my @sampling = split(':',$words[8]);
my $runid = $run[1];
my $expid = $exp[1];
my $sampling_interval = $sampling[1];
my $foo = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/in-%d-%d.txt",$expid,$runid;
my $bar = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/out-%d-%d.txt",$expid,$runid;

# initializing a session variable to probe ingress switch
my ($session1,$error1) = Net::SNMP->session(
    -hostname => $ingress_switch,
    -community => "public",
    -port => 161,
    -nonblocking => 1,
    -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric
);

if (!defined $session1) {
    printf "ERROR: %s.\n", $error1;
    exit 1;
}

# initializing a session variable to probe egress switch
my ($session2,$error2) = Net::SNMP->session(
    -hostname => $egress_switch,
    -community => "public",
    -port => 161,
    -nonblocking => 1,
    -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric
);

if (!defined $session2) {
    printf "ERROR: %s.\n", $error2;
    exit 1;
}

my $pointer=0;

while(1) {

    my $result1 = $session1->get_request(-varbindlist => \@ingress_oid,-callback =>[ \&table_callback ,$ingress_switch,$pointer]);
    if (!defined $result1) {
      printf "ERROR: %s.\n", $session1->error();
      $session1->close();
      exit 1;
   }

    my $result2 = $session2->get_request(-varbindlist => \@egress_oid,-callback =>[ \&table_callback ,$egress_switch,$pointer]);
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
    my $Delta = $sampling_interval*1000 - $delta ;

    if ( $Delta > 0 ) {
        usleep   $Delta;
    } else {
        usleep 0;
    }
    $pointer = $pointer +1 ;
}

sub table_callback()
{
    my ($session,$host,$x) = @_;
    my $list = $session->var_bind_list();
    if (!defined $list) {
         printf "ERROR: %s\n", $session->error();
         return;
    }

    my ($second, $microsecond) = gettimeofday;
    my $time = sprintf('%d.%d',$second,$microsecond);

    #opening files for printing the output
    open FOO, ">>",$foo or die $!;
    open BAR, ">>",$bar or die $!;

    my $JSON = JSON->new->utf8;

    if ($host eq $ingress_switch)
    {
        my $In1    = $list->{$ingress_oid[1]};
        my $uptime = $list->{$sysUpTime}; # in microseconds
           $uptime = $uptime*0.01; # in seconds
        my $Output = {switch_management_ip=>$ingress_switch,uptime=>$uptime,unixtime=>$time,in_1=>$In1,serial_id=>$x};
        my $json = $JSON->encode($Output) ;
        print FOO "$json\n";
    }

    else
    {
        my $Out8   = $list->{$egress_oid[1]};
        my $uptime = $list->{$sysUpTime}; # in microseconds
           $uptime = $uptime*0.01; # in seconds
        my $Output = {switch_management_ip=>$egress_switch,uptime=>$uptime,unixtime=>$time,out_8=>$Out8,serial_id=>$x};
        my $json = $JSON->encode($Output) ;
        print BAR "$json\n";
    }

}

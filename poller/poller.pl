#!/usr/bin/perl
use strict;
use warnings FATAL => 'all';
use YAML::Tiny;
use Net::SNMP qw(snmp_dispatcher ticks_to_time);
use Time::HiRes qw(usleep gettimeofday);

# Open the config yaml file
my $yaml = YAML::Tiny->read( '/mnt/LONTAS/ExpControl/dire15/thesis/config.yml' );

# Storing the properties of yaml file into variables
my $ingress_switch  = $yaml->[0]->{ingress_switch_details}->{switch_management_ip};
my $egress_switch  = $yaml->[0]->{egress_switch_details}->{switch_management_ip};
my $ingress_interface = $yaml->[0]->{ingress_switch_details}->{interface_id};
my $egress_interface = $yaml->[0]->{egress_switch_details}->{interface_id};
print "$egress_switch\n";
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
my @sampling = split(':',$words[6]);
my $runid = $run[1];
my $expid = $exp[1];
my $sampling_interval = $sampling[1];
my $foo = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/in-%d-%d.txt",$expid,$runid;
my $bar = sprintf "/mnt/LONTAS/ExpControl/dire15/logs/out-%d-%d.txt",$expid,$runid;

# initializing a session variable to probe ingress switch
my ($session1,$error1) = Net::SNMP->session(
    -hostname => $ingress_switch,
    -community => 'public',
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
    -community => 'public',
    -port => 161,
    -nonblocking => 1,
    -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric
);

if (!defined $session2) {
    printf "ERROR: %s.\n", $error2;
    exit 1;
}


while(1) {

    my $result1 = $session1->get_request(-varbindlist => \@ingress_oid,-callback =>[ \&table_callback ,$ingress_switch]);
    if (!defined $result1) {
      printf "ERROR: %s.\n", $session1->error();
      $session1->close();
      exit 1;
   }

    my $result2 = $session2->get_request(-varbindlist => \@egress_oid,-callback =>[ \&table_callback ,$egress_switch]);
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

    if ( $Delta > 0 ) {
        usleep   $Delta;
    } else {
        usleep 0;
    }
}

sub table_callback()
{
    my ($session,$host) = @_;
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

    if ($host == $ingress_switch)
    {
        my $In1    = $list->{$ingress_oid[1]};
        my $uptime = $list->{$sysUpTime}; # in microseconds
           $uptime = $uptime*0.01; # in seconds
        my $Output = {switch_management_ip=>$ingress_switch,uptime=>$uptime,unixtime=>$time,in_1=>$In1};
        my $json = $JSON->encode($Output) ;
        print FOO "$json\n";
    }

    if ($host == $egress_switch)
    {
        my $Out1   = $list->{$egress_oid[1]};
        my $uptime = $list->{$sysUpTime}; # in microseconds
           $uptime = $uptime*0.01; # in seconds
        my $Output = {switch_management_ip=>$egress_switch,uptime=>$uptime,unixtime=>$time,in_1=>$Out1};
        my $json = $JSON->encode($Output) ;
        print BAR "$json\n";
    }

}

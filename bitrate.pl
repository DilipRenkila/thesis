   #!/usr/bin/perl


   use Net::SNMP qw(snmp_dispatcher ticks_to_time);
   use Time::HiRes qw(usleep gettimeofday);
   use JSON;

   my $filename = "/home/ats/dire15/thesis/delay.txt"; 
   my $delay;
    open(my $fh, '<', $filename) or die "cannot open file $filename";
    {
        local $/;
        $delay = <$fh>;
    }
    close($fh);

   my $filename = "/home/ats/dire15/thesis/expid.txt";
   my $expid;
    open(my $fh, '<', $filename) or die "cannot open file $filename";
    {
        local $/;
        $expid = <$fh>;
    }
    close($fh);

   my $filename = "/home/ats/dire15/thesis/runid.txt";
   my $runid;
    open(my $fh, '<', $filename) or die "cannot open file $filename";
    {
        local $/;
        $runid = <$fh>;
    }
    close($fh);

  
   my $path ="/home/ats/dire15/thesis/config.pl";
   do "$path";
   ($second, $microsecond) = gettimeofday;
   open FILE1, ">>", "/home/ats/dire15/thesis/logs/$ip1-ingress-$expid-$runid-$delay.txt" or die $!;
   open FILE2, ">>", "/home/ats/dire15/thesis/logs/$ip2-egress-$expid-$runid-$delay.txt" or die $!;
   $i1=1;$i2=2;
   my ($session1, $error1) = Net::SNMP->session(
      -hostname     => $ip1,
      -community    => $c1,
      -port         => $p1,
      -nonblocking  => 1,
      -translate    => [-timeticks => 0x0 ]  # Turned off so  that the sysUpTime is numeric   
   );

   if (!defined $session1) {
      printf "ERROR: %s.\n", $error;
      exit 1;
   }

   my ($session2, $error2) = Net::SNMP->session(
      -hostname     => $ip2,
      -community    => $c2,
      -port         => $p2,
      -nonblocking  => 1,
      -translate    => [-timeticks => 0x0 ]  # Turn off so sysUpTime is numeric   
   );

   if (!defined $session2) {
      printf "ERROR: %s.\n", $error2;
      exit 1;
   }

 while(1)
 {  
   my $result1 = $session1->get_request(-varbindlist => \@oid1,-callback =>[ \&table_callback ,$i1,$second]);
     
     if (!defined $result1) {
      printf "ERROR: %s.\n", $session1->error();
      $session1->close();
      exit 1;
   }
   my $result2 = $session2->get_request(-varbindlist => \@oid2,-callback =>[ \&table_callback ,$i2,$second]);

   if (!defined $result2) {
      printf "ERROR: %s.\n", $session2->error();
      $session2->close();
      exit 1;
   }
 ($seconds, $microseconds) = gettimeofday;
  $t1 = $seconds * 1000000 + $microseconds ;
   snmp_dispatcher();
 ($seconds1, $microseconds1) = gettimeofday; 
  $t2 = $seconds1 * 1000000 + $microseconds1 ;
 $delta = ($t2-$t1); #in microseconds
 $Delta = 1000000 - $delta ; # probing for 100 milliseconds # 1 millisecond = 10000 microseconds
 
 
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
      ($seconds, $microseconds) = gettimeofday;
      open FILE1, ">>", "/home/ats/dire15/thesis/logs/$ip1-ingress-$expid-$runid-$delay.txt" or die $!;
      open FILE2, ">>", "/home/ats/dire15/thesis/logs/$ip2-egress-$expid-$runid-$delay.txt" or die $!;
      if (!defined $list) {
         printf "ERROR: %s\n", $session->error();
         return;
      }
      my $JSON = JSON->new->utf8;   
      if ($host == $i1)
      
      {  

       $In1    = $list->{$in1};
       $Out1    = $list->{$out1};
       $In8    = $list->{$in8};
       $Out8    = $list->{$out8};   
       $uptime = $list->{$OID_sysUpTime}; # in microseconds
       $uptime = $uptime*0.01;
       $Output = {Switch_IP=>$ip1,Uptime=>$uptime,unixtime=>$seconds,IN_Interface_1=>$In1,OUT_Interface_1=>$Out1,IN_Interface_8=>$In8,OUT_Interface_8=>$Out8}; 
       $json = $JSON->encode($Output) ;
#       print "$json\n";
       print FILE1 "$json\n";
#        print FILE1 "$ip1,$uptime,$seconds,$In1,out1=$Out1,in8=$In8,$Out8\n";
   }
      if ($host == $i2)
      {
		  
       $In1     = $list->{$IN1};
       $Out1    = $list->{$OUT1};
       $In2     = $list->{$IN2};
       $Out2    = $list->{$OUT2};
       $uptime = $list->{$OID_sysUpTime}; # in microseconds
       $uptime = $uptime*0.01;
       $Output = {Switch_IP=>$ip2,Uptime=>$uptime,unixtime=>$seconds,IN_Interface_1=>$In1,OUT_Interface_1=>$Out1,IN_Interface_2=>$In2,OUT_Interface_2=>$Out2}; 
       $json = $JSON->encode($Output) ;
#       print "$json\n";
       print FILE2 "$json\n";
#        print FILE2 "$ip2,$uptime,$seconds,$In1,out1=$Out1,in2=$In2,$Out2\n";
   }
      
}
   

#! /usr/bin/perl 
if ( @ARGV !=2)
{
        print "Usage:please enter sourefile + targetfile  \n";
        exit;
}
 #   open(OUT,">> $test1")  or die $!;
#    chomp($test=<STDIN>);
    $test=shift;
    $test1=shift;
    my $counter =1;
    open(OUT,">> $test1")  or die $!;
    open(IN, "< $test") or die $!;

             while(my $line = <IN> )
                {


                      my @huadan = split/\|/,"SerialNo|Version|TicketType|TimeStamp|PayTime|HostID|ServiceScenarious|ChargedParty|CallingParty|CalledParty|OrignialCallingParty|OrignialCalledParty|PayFlag|ServID|CustID|Brand|SessionId|SessionBeginTime|SessionTerminatedTime|TerminatedCause|OrignialHost|Balanceinfo|Accumlatorinfo|TariffInfo|MasterProductID|BearerCapability|DoCredit|ServiceKey|MSISDN|Account|BOSSSEQ|OCSSEQ|TradeTime|TradeType|AccountLeft|CallServiceStop|AccountStop|ErrorType|TradeSeq|AccountNumber|HomePLMN|Addedactivedays|OSSOperatorID|OverFirstCharge|MsgSrc|Controlflag|";

                        chomp($line);
                      my @array=split (/\|/,$line);
                      print OUT "record:$counter\n";
                      print OUT ".............................................\n";
                       for ($index=0;$index<@array-1;$index++){
                          printf OUT ("%2d|%-30s|%s\n",$index,$huadan[$index],$array[$index]);
                      }
                      $counter++
                }
                close IN or die $!;
        print "the end\n";
           close OUT or die $!;

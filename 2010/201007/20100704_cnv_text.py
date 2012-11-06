# -*- coding: utf-8 -*- 
for line in open('f:/mywb.dic'):
		wd=''.join(line).rstrip('\n').rstrip('\r').split(' ');
		print '%(key)s\t' % {'key':wd[0]},
		for ww in range(len(wd)):
		   if ww>0 :
		      print ' %(word)s' % {'word':wd[ww]},
		print '';
		#for ww in range(len(wd)):
		#    if ww>0 :
		#      print '%(key)s\t%(num)d\t%(word)s' % {'key':wd[ww],'num':0,'word':wd[0]};

//============================================================================
// Name        : bdb.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : debug info
// 新版的HP-UX 及 Linux都不支持 ucontext中的regs
//============================================================================

#include <cstdlib>
#include <cstring>
#include <csignal>
#include <cstdio>
#include <cerrno>
#include <ucontext.h>

static void seghandler (unsigned int sn, siginfo_t si, \
		                    struct ucontext *sc)
{
	unsigned int mnip;
	int i;
	mnip = *(unsigned int*)(((struct pt_regs *)
			((&(sc->uc_mcontext))->regs))->nip);
	printf("Signal number = %d, Signal errno = %d\n",
			si.si_signo,si.si_errno);
	switch(si.si_code)
	{
		case 1:
			printf(" SI code = %d (Address not mapped to object)\n",si.si_code);
		  break;
		case 2:
			printf(" SI code = %d (Invalid permissions for mapped object)\n",si.si_code);
			break;
		default:
			printf(" SI code = %d (Unknown SI Code)\n",si.si_code);
			break;
	}
	printf(" Intruction pointer = %x \n",mnip);
	printf(" Fault addr = 0x%x \n",si.si_addr);
	printf(" dar = 0x%x \n",
			(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->dar));
  printf(" trap = 0x%x \n",
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->trap));
  printf(" Op-Code [nip - 4] = 0x%x at address = 0x%x \n",
  		*(unsigned int*)
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip-4),
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip-4) );
  printf(" Failed Op-code = 0x%x at address = 0x%x \n",
  		*(unsigned int*)
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip),
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip) );
  printf(" Op-Code [nip + 1] = 0x%x at address = 0x%x \n",
  		*(unsigned int*)
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip+4),
  		(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip+4));
  printf("***GPR values are the time of fault***\n");
  for (i=0; i<11; ++i)
  	printf(" Gpr[%d] = 0x%x \n",i,
  			(((struct pt_regs*)((&(sc->uc_mcontext))->regs))->gpr[i]));

  (((struct pt_regs*)((&(sc->uc_mcontext))->regs))->nip)+=4;
}

int main()
{
	struct sigaction m;
	char *p,*q,arr[]="Ma";
	q=arr;
	m.sa_flags = SA_SIGINFO;
	m.sa_sigaction = (void*)seghandler;
	sigaction(SIGSEGV,&m,(struct sigaction*)NULL);
	*p++ = *q++;
	return 0;
}


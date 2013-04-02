#include <time.h>
#include <unistd.h>
#include <cstdio>

#define FEATURE_DEBUG_OPT
#define DebugOpt 1

int main(){
		time_t t1 = time(NULL);
		time_t t2;
		long dt;
		short rescan = 60;
		short sleep_time = 60;

		for (;;) {
			sleep((sleep_time + 1) - (short) (time(NULL) % sleep_time));

			t2 = time(NULL);
			dt = t2 - t1;

			/*
			 * The file 'cron.update' is checked to determine new cron
			 * jobs.  The directory is rescanned once an hour to deal
			 * with any screwups.
			 *
			 * check for disparity.  Disparities over an hour either way
			 * result in resynchronization.  A reverse-indexed disparity
			 * less then an hour causes us to effectively sleep until we
			 * match the original time (i.e. no re-execution of jobs that
			 * have just been run).  A forward-indexed disparity less then
			 * an hour causes intermediate jobs to be run, but only once
			 * in the worst case.
			 *
			 * when running jobs, the inequality used is greater but not
			 * equal to t1, and less then or equal to t2.
			 */

			if (--rescan == 0) {
				rescan = 60;
				//SynchronizeDir();
			}
			//CheckUpdates();
#ifdef FEATURE_DEBUG_OPT
			if (DebugOpt)
				printf("\005Wakeup dt=%d\n", dt);
#endif
			if (dt < -60 * 60 || dt > 60 * 60) {
				t1 = t2;
				printf("\111time disparity of %d minutes detected\n", dt / 60);
			} else if (dt > 0) {
				//TestJobs(t1, t2);
				printf("ok time\n");
				//RunJobs();
				sleep(5);
				//if (CheckJobs() > 0) { ///返回还有任务的数量
				if (bool has_next=false){
					sleep_time = 10;
				} else {
					sleep_time = 60;
				}
				t1 = t2;
			}
		}
	}
	
/**
 * @file 20100319_FileLock.cpp
 * @brief
 *
 * @author bonly
 * @date 2011-11-22 bonly created
 */
#include <fcntl.h>
#include <cerrno>
#include <cstdio>
#include <sys/types.h>
#include <unistd.h>
#include <sys/wait.h>
#include <cstdlib>
#include <string.h>

class FileLock
{
    public:
        enum Status{finish=0,inited=1};
        FileLock():file_(0),status_(finish)
        {
            pid_ = getpid();
        }
        ~FileLock()
        {
            if (file_ != 0)
            {
                file_unlock();
            }
        }

        /**
         *  @return 成功:文件ID 失败:-1
         */
        int get_lock()
        {
            int li_lck_st = -1;
            file_ = open("/tmp/dbinit", O_RDWR|O_CREAT, 0777);
            if( file_ <= 0 )
            {
                //perror("file open error");
                return -1;
            }
            if(-1 == (li_lck_st = file_lock()))
            {
                close(file_);
                file_ = 0;
                //perror("lock file fail");
                return -1;
            }
            if(-1 == lseek(file_, 0, SEEK_SET))
            {
                perror("lseek faild");
                file_unlock();
                return -1;
            }
            if(-1 == read(file_, &status_, sizeof(int)))
            {
                //perror("read faild");
            }

            printf("read status %d\n", status_);
            return li_lck_st;
        }
        int release_lock()
        {
            if (file_ != 0)
            {
                if(-1 == lseek(file_, 0, SEEK_SET))
                {
                    perror("lseek faild");
                }
                if(-1 == write(file_, &status_, sizeof(int)))
                {
                    perror("write faile");
                }
                return file_unlock();
            }
            return -1;
        }
        int status(){return (int)status_;}
        void status(int st)
        {
            status_ = (Status)st;
        }

    private:
        /**
         *  @ return 成功:文件ID 失败:-1
         */
        int file_lock()
        {
            struct flock s_flock;
            s_flock.l_type = F_WRLCK;
            s_flock.l_whence = SEEK_SET;
            s_flock.l_start = 0;
            s_flock.l_len = 0;
            s_flock.l_pid = pid_;

            //F_SETLKW对加锁操作进行阻塞，
            //F_SETLK不对加锁操作进行阻塞，立即返回
            if (-1 == fcntl(file_, F_SETLKW, &s_flock))
            {
                perror("lock file fail");
                return -1;
            }

            return file_;
        }

        /**
         *  @return 失败: -1
         */
        int file_unlock()
        {
            struct flock fl = {0};
            fl.l_type   = F_UNLCK;
            fl.l_whence = SEEK_SET;
            fl.l_start  = 0;
            fl.l_len    = 0;

            int res = fcntl(file_, F_SETLKW, &fl);;
            if (-1 != res)
            {
                close (file_);
                file_ = 0;
            }
            return res;
        }

    private:
        int file_;
        pid_t pid_;
        Status status_;
};

int main(int argc, char* argv[])
{
    int pid = -1;
    FileLock flk;

    int i = 5;
    while(i--)
    {
        if ((pid = fork()) < 0)
        { //fork出错
             puts("fork1 error");
        }
        else if(pid >0) //父进程
        {
            sleep(5);

            if (waitpid(pid, NULL, 0) < 0)
              puts("waitpid error");
        }
        else //子进程
        {
            sleep(1);
            int res = 0;
            if ((res = flk.get_lock()) < 0)
            {
                printf("lock fail pid=[%d]\n", getpid() );
            }
            else
            {
                printf("lock succ [%d] pid =%d\n", res, getpid() );
                printf("org status is: %d\n",flk.status());
                if (argc >= 2)
                {
                  flk.status(strtol(argv[1], 0, 10));
                }
                printf("chg status to: %d\n",flk.status());
                sleep(5);

                if(flk.release_lock()==-1)
                {
                    printf("unlock fail pid =%d\n", getpid() );
                }
                else
                {
                    printf("unlock succ [%d] pid =%d\n", res, getpid() );
                }
            }
        }
    }

    return 0;
}


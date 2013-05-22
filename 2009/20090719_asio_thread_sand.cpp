#include <iostream>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/date_time/posix_time/posix_time.hpp>
using namespace boost;
class printer
{
  public:
    printer(boost::asio::io_service& io) :
      strand_(io), timer1_(io, boost::posix_time::seconds(1)), timer2_(io,
          boost::posix_time::seconds(1)), count_(0)
    {
      timer1_.async_wait(strand_.wrap(boost::bind(&printer::print1, this)));
      timer2_.async_wait(strand_.wrap(boost::bind(&printer::print2, this)));
    }

    ~printer()
    {
      std::cout << "Final count is " << this_thread::get_id() <<"\t"<< count_ << "\n";
    }
    void print1()
    {
      if (count_ < 10)
      {
        std::cout << "Timer 1: " << this_thread::get_id() <<"\t"<< count_ << "\n";
        ++count_;

        timer1_.expires_at(timer1_.expires_at() + boost::posix_time::seconds(1));
        timer1_.async_wait(strand_.wrap(boost::bind(&printer::print1, this)));
      }
    }

    void print2()
    {
      if (count_ < 10)
      {
        std::cout << "Timer 2: " << this_thread::get_id() <<"\t"<< count_ << "\n";
        ++count_;

        timer2_.expires_at(timer2_.expires_at() + boost::posix_time::seconds(1));
        timer2_.async_wait(strand_.wrap(boost::bind(&printer::print2, this)));
      }
    }

  private:
    boost::asio::strand strand_;
    boost::asio::deadline_timer timer1_;
    boost::asio::deadline_timer timer2_;
    int count_;
};

int main()
{
  boost::asio::io_service io;
  printer p(io);
  boost::thread t(boost::bind(&boost::asio::io_service::run, &io));
  io.run();
  t.join();

  return 0;
}

/*
Timer 2: 0x8075780	0
Timer 1: 0x8075780	1
Timer 2: 0x8075590	2
Timer 1: 0x8075590	3
Timer 2: 0x8075590	4
Timer 1: 0x8075590	5
Timer 2: 0x8075590	6
Timer 1: 0x8075590	7
Timer 2: 0x8075590	8
Timer 1: 0x8075590	9
Final count is 0x8075780	10
*/

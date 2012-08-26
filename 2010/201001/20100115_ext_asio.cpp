//============================================================================
// Name        : ext_asio.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/scoped_ptr.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/weak_ptr.hpp>
#include <boost/system/error_code.hpp>
template<typename Service>
class basic_timer: public boost::asio::basic_io_object<Service>
//这是一个封装类，它提供了对外的接口函数，
//同时它拥有了两个关键的变量，
//一个是IO操作实现类implementation （用boost::shared_ptr来管理）
//一个是IO服务类service
//可以看到它的接口只是简单的运用上面两个变量
//我暂且叫它IO接口类好了。
//接口第一个参数都是this->implementation，看起来好像只是把IO操作实现类对象传递给IO服务类service对象
//但我想这样的接口可以灵活地更换IO操作实现类来实现不同的操作。
{
   public:
      explicit basic_timer(boost::asio::io_service &io_service) :
         boost::asio::basic_io_object<Service>(io_service)
      {
      }
      void wait(std::size_t seconds)
      {
         return this->service.wait(this->implementation, seconds);
      }
      template<typename Handler>
      void async_wait(std::size_t seconds, Handler handler)
      {
         this->service.async_wait(this->implementation, seconds, handler);
      }
};

class timer_impl;

template<typename TimerImplementation = timer_impl>
class basic_timer_service: public boost::asio::io_service::service
//这就是上文的IO服务类service
{
   public:
      static boost::asio::io_service::id id;
      explicit basic_timer_service(boost::asio::io_service &io_service) :
                  boost::asio::io_service::service(io_service),
                  async_work_(
                           new boost::asio::io_service::work(async_io_service_)),
                  async_thread_(
                           boost::bind(&boost::asio::io_service::run,
                                    &async_io_service_))
      //为了方便下文描述，在这里把io_service叫外部io_service，而async_io_service_叫内部io_service_
      //可以看出它在构造函数起了一个线程，这个线程只执行了成员变量内部io_service_的run函数，
      //async_work_是用来控制线程生命周期的
      {
      }
      ~basic_timer_service()
      //async_work_在类析构函数中释放，表示这个线程的生命周期正好是这个类对象的生命周期。
      {
         async_work_.reset();
         async_io_service_.stop();
         async_thread_.join();
      }
      typedef boost::shared_ptr<TimerImplementation> implementation_type;
      void construct(implementation_type &impl)
      //接口类的变量IO操作实现类implementation在这里被创建
      {
         impl.reset(new TimerImplementation());
      }
      void destroy(implementation_type &impl)
      //接口类的变量IO操作实现类implementation在这里被销毁
      //在接口类的构造函数中我们可以看到显示地调用了service.construct(implementation);
      //在接口类的析构函数中我们可以看到显示地调用了service.destroy(implementation);
      //这样设计的目的，我想也许是为了能在接口类中能管理多个IO操作实现类的对象的生成与释放
      //而服务类并不关心实际使用的是哪一个IO操作实现类，也不用去理会IO操作实现类的生命周期。
      //这样分工明确，扩展灵活。
      {
         impl->destroy();
         impl.reset();
      }
      void wait(implementation_type &impl, std::size_t seconds)
      //这看起来一层又一层地封装确实会让程序效率低了，
      //本来boost::asio就是用来异步操作的，像wait这种阻塞的东西就不要加进来，
      //当然我想这个例子只是为了让wait和下面的async_wait作一个比较
      {
         boost::system::error_code ec;
         impl->wait(seconds, ec);
         boost::asio::detail::throw_error(ec);
      }
      template<typename Handler>
      class wait_operation
      //我们可以通过
      //template <typename CompletionHandler>
      //void boost::asio::io_service::post(CompletionHandler handler);
      //来为io_service添加一个函数任务，CompletionHandler这个类必须重载void operator()() const或者void operator()()
      //io_service会在调用run(), run_one(), poll() or poll_one()的线程中运行这个函数。
      //我们通常用
      //template <typename Handler, typename Arg1>
      //binder1<Handler, Arg1> boost::asio::detail::bind_handler(const Handler& handler,
      //  const Arg1& arg1)
      //{
      //  return binder1<Handler, Arg1>(handler, arg1);
      //}
      //来绑定一个一元函数，返回的binder1就是一个符合CompletionHandler标准的类。
      //但是必须是一个一元函数，这很让人恼火。
      //所以我们构造一个类似于wait_operation的类来给io_service添加一个函数任务。
      //wait_operation同样重载void operator()() const
      //如果我们需要io_service调用函数的参数很多或者像wait_operation一样在处理一些数据后还要转发给其他io_service(既涉及单位很多复杂的函数)，
      //这个时候可以建立一个符合CompletionHandler标准的类
      //因为io_service中所有的任务都是一个类对象，它不需要理会是什么样的类，只是安排一个时间在某个线程运行操作().
      {
         public:
            wait_operation(implementation_type &impl,
                     boost::asio::io_service &io_service, std::size_t seconds,
                     Handler handler) :
               impl_(impl), io_service_(io_service), work_(io_service),
                        seconds_(seconds), handler_(handler)
            {
            }
            void operator()() const
            {
               implementation_type impl = impl_.lock();
               if (impl)
               {
                  boost::system::error_code ec;
                  impl->wait(seconds_, ec);
                  this->io_service_.post(
                           boost::asio::detail::bind_handler(handler_, ec));
                  //内部io_service在调用完实际IO对象impl的wait之后，会给外部io_service添加一个任务。
               }
               else
               {
                  this->io_service_.post(
                           boost::asio::detail::bind_handler(handler_,
                                    boost::asio::error::operation_aborted));
                  //由于内部io_service::run是在内部线程中运行，而IO操作实现类impl是接口类的一个成员变量，
                  //虽然IO服务类有construct和destroy这样的接口操作O操作实现类的生命，但前面说过了这不是它的职责，
                  //接口类才会调用这两个接口，所以IO操作实现类impl说不定已经被释放了。
                  //也许你会想到接口类不是在析构的时候才会释放IO操作实现类impl，但前面我说过接口类可以管理多个IO操作实现类。
                  //所以这里用到了boost::weak_ptr。
                  //如果IO操作实现类impl已经被释放了，我们会给外部io_service添加任务的参数中显示的写入错误ID。
               }
            }
         private:
            boost::weak_ptr<TimerImplementation> impl_;
            boost::asio::io_service &io_service_;
            boost::asio::io_service::work work_;
            std::size_t seconds_;
            Handler handler_;
      };
      template<typename Handler>
      void async_wait(implementation_type &impl, std::size_t seconds,
               Handler handler)
      {
         this->async_io_service_.post(
                  wait_operation<Handler> (impl, this->get_io_service(),
                           seconds, handler));
         //为内部io_service添加一个任务
         //这里可以看到内部io_service（内部线程）会调用IO操作实现类的阻塞函数wait
         //而外部io_service会完成Handler的调用，在这个例子中是在主线程调用的。
      }
   private:
      void shutdown_service()
      {
      }
      boost::asio::io_service async_io_service_;
      boost::scoped_ptr<boost::asio::io_service::work> async_work_;
      boost::thread async_thread_;
};

template<typename TimerImplementation>
boost::asio::io_service::id basic_timer_service<TimerImplementation>::id;
class timer_impl
//IO操作实现类,这个一目了然，但注意它的destroy接口，
//我们在IO服务类的void destroy(implementation_type &impl) 接口显示调用过了
//在调用过之后释放了内存，那么destroy接口必须满足一个条件，或许就是让自己所有的操作都中断，
//然后静静地等待消失。
{
   public:
      timer_impl() :
         handle_(CreateEvent(NULL, FALSE, FALSE, NULL))
      {
      }
      ~timer_impl()
      {
         CloseHandle(handle_);
      }
      void destroy()
      {
         SetEvent(handle_);
      }
      void wait(std::size_t seconds, boost::system::error_code &ec)
      {
         DWORD res = WaitForSingleObject(handle_, seconds * 1000);
         if (res == WAIT_OBJECT_0)
            ec = boost::asio::error::operation_aborted;
         else
            ec = boost::system::error_code();
      }
   private:
      HANDLE handle_;
};

void wait_handler(const boost::system::error_code &ec)
//外部io_service的任务函数。
{
   std::cout << "5 s." << std::endl;
}
typedef basic_timer<basic_timer_service<> > timer;

#include <cstddef>
int main(int argC, char* argv[])
{

   boost::asio::io_service io_service;
   timer t(io_service);
   t.async_wait(5, wait_handler);
   io_service.run();

   return 0;
}

/*
http://zh.highscore.de/cpp/boost/

虽然 Boost.Asio 主要是支持网络功能的，但是加入其它 I/O 对象以执行其它的异步操作也非常容易。 本节将介绍 Boost.Asio 扩展的一个总体布局。 虽然这不是必须的，但它为其它扩展提供了一个可行的框架作为起点。

要向 Boost.Asio 中增加新的异步操作，需要实现以下三个类：

一个派生自 boost::asio::basic_io_object 的类，以表示新的 I/O 对象。使用这个新的 Boost.Asio 扩展的开发者将只会看到这个 I/O 对象。

一个派生自 boost::asio::io_service::service 的类，表示一个服务，它被注册为 I/O 服务，可以从 I/O 对象访问它。 服务与 I/O 对象之间的区别是很重要的，因为在任意给定的时间点，每个 I/O 服务只能有一个服务实例，而一个服务可以被多个 I/O 对象访问。

一个不派生自任何其它类的类，表示该服务的具体实现。 由于在任意给定的时间点每个 I/O 服务只能有一个服务实例，所以服务会为每个 I/O 对象创建一个其具体实现的实例。 该实例管理与相应 I/O 对象有关的内部数据。

本节中开发的 Boost.Asio 扩展并不仅仅提供一个框架，而是模拟一个可用的 boost::asio::deadline_timer 对象。 它与原来的 boost::asio::deadline_timer 的区别在于，计时器的时长是作为参数传递给 wait() 或 async_wait() 方法的，而不是传给构造函数。

#include <boost/asio.hpp>
#include <cstddef>

template <typename Service>
class basic_timer
  : public boost::asio::basic_io_object<Service>
{
  public:
    explicit basic_timer(boost::asio::io_service &io_service)
      : boost::asio::basic_io_object<Service>(io_service)
    {
    }

    void wait(std::size_t seconds)
    {
      return this->service.wait(this->implementation, seconds);
    }

    template <typename Handler>
    void async_wait(std::size_t seconds, Handler handler)
    {
      this->service.async_wait(this->implementation, seconds, handler);
    }
};
下载源代码
每个 I/O 对象通常被实现为一个模板类，要求以一个服务来实例化 - 通常就是那个特定为此 I/O 对象开发的服务。 当一个 I/O 对象被实例化时，该服务会通过父类 boost::asio::basic_io_object 自动注册为 I/O 服务，除非它之前已经注册。 这样可确保任何 I/O 对象所使用的服务只会每个 I/O 服务只注册一次。

在 I/O 对象的内部，可以通过 service 引用来访问相应的服务，通常的访问就是将方法调用前转至该服务。 由于服务需要为每一个 I/O 对象保存数据，所以要为每一个使用该服务的 I/O 对象自动创建一个实例。 这还是在父类 boost::asio::basic_io_object 的帮助下实现的。 实际的服务实现被作为一个参数传递给任一方法调用，使得服务可以知道是哪个 I/O 对象启动了这次调用。 服务的具体实现是通过 implementation 属性来访问的。

一般一上谕，I/O 对象是相对简单的：服务的安装以及服务实现的创建都是由父类 boost::asio::basic_io_object 来完成的，方法调用则只是前转至相应的服务；以 I/O 对象的实际服务实现作为参数即可。

#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <boost/scoped_ptr.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/weak_ptr.hpp>
#include <boost/system/error_code.hpp>

template <typename TimerImplementation = timer_impl>
class basic_timer_service
  : public boost::asio::io_service::service
{
  public:
    static boost::asio::io_service::id id;

    explicit basic_timer_service(boost::asio::io_service &io_service)
      : boost::asio::io_service::service(io_service),
      async_work_(new boost::asio::io_service::work(async_io_service_)),
      async_thread_(boost::bind(&boost::asio::io_service::run, &async_io_service_))
    {
    }

    ~basic_timer_service()
    {
      async_work_.reset();
      async_io_service_.stop();
      async_thread_.join();
    }

    typedef boost::shared_ptr<TimerImplementation> implementation_type;

    void construct(implementation_type &impl)
    {
      impl.reset(new TimerImplementation());
    }

    void destroy(implementation_type &impl)
    {
      impl->destroy();
      impl.reset();
    }

    void wait(implementation_type &impl, std::size_t seconds)
    {
      boost::system::error_code ec;
      impl->wait(seconds, ec);
      boost::asio::detail::throw_error(ec);
    }

    template <typename Handler>
    class wait_operation
    {
      public:
        wait_operation(implementation_type &impl, boost::asio::io_service &io_service, std::size_t seconds, Handler handler)
          : impl_(impl),
          io_service_(io_service),
          work_(io_service),
          seconds_(seconds),
          handler_(handler)
        {
        }

        void operator()() const
        {
          implementation_type impl = impl_.lock();
          if (impl)
          {
              boost::system::error_code ec;
              impl->wait(seconds_, ec);
              this->io_service_.post(boost::asio::detail::bind_handler(handler_, ec));
          }
          else
          {
              this->io_service_.post(boost::asio::detail::bind_handler(handler_, boost::asio::error::operation_aborted));
          }
      }

      private:
        boost::weak_ptr<TimerImplementation> impl_;
        boost::asio::io_service &io_service_;
        boost::asio::io_service::work work_;
        std::size_t seconds_;
        Handler handler_;
    };

    template <typename Handler>
    void async_wait(implementation_type &impl, std::size_t seconds, Handler handler)
    {
      this->async_io_service_.post(wait_operation<Handler>(impl, this->get_io_service(), seconds, handler));
    }

  private:
    void shutdown_service()
    {
    }

    boost::asio::io_service async_io_service_;
    boost::scoped_ptr<boost::asio::io_service::work> async_work_;
    boost::thread async_thread_;
};

template <typename TimerImplementation>
boost::asio::io_service::id basic_timer_service<TimerImplementation>::id;
下载源代码
为了与 Boost.Asio 集成，一个服务必须符合几个要求：

它必须派生自 boost::asio::io_service::service。 构造函数必须接受一个指向 I/O 服务的引用，该 I/O 服务会被相应地传给 boost::asio::io_service::service 的构造函数。

任何服务都必须包含一个类型为 boost::asio::io_service::id 的静态公有属性 id。在 I/O 服务的内部是用该属性来识别服务的。

必须定义两个名为 construct() 和 destruct() 的公有方法，均要求一个类型为 implementation_type 的参数。 implementation_type 通常是该服务的具体实现的类型定义。 正如上面例子所示，在 construct() 中可以很容易地使用一个 boost::shared_ptr 对象来初始化一个服务实现，以及在 destruct() 中相应地析构它。 由于这两个方法都会在一个 I/O 对象被创建或销毁时自动被调用，所以一个服务可以分别使用 construct() 和 destruct() 为每个 I/O 对象创建和销毁服务实现。

必须定义一个名为 shutdown_service() 的方法；不过它可以是私有的。 对于一般的 Boost.Asio 扩展来说，它通常是一个空方法。 只有与 Boost.Asio 集成得非常紧密的服务才会使用它。 但是这个方法必须要有，这样扩展才能编译成功。

为了将方法调用前转至相应的服务，必须为相应的 I/O 对象定义要前转的方法。 这些方法通常具有与 I/O 对象中的方法相似的名字，如上例中的 wait() 和 async_wait()。 同步方法，如 wait()，只是访问该服务的具体实现去调用一个阻塞式的方法，而异步方法，如 async_wait()，则是在一个线程中调用这个阻塞式方法。

在线程的协助下使用异步操作，通常是通过访问一个新的 I/O 服务来完成的。 上述例子中包含了一个名为 async_io_service_ 的属性，其类型为 boost::asio::io_service。 这个 I/O 服务的 run() 方法是在它自己的线程中启动的，而它的线程是在该服务的构造函数内部由类型为 boost::thread 的 async_thread_ 创建的。 第三个属性 async_work_ 的类型为 boost::scoped_ptr<boost::asio::io_service::work>，用于避免 run() 方法立即返回。 否则，这可能会发生，因为已没有其它的异步操作在创建。 创建一个类型为 boost::asio::io_service::work 的对象并将它绑定至该 I/O 服务，这个动作也是发生在该服务的构造函数中，可以防止 run() 方法立即返回。

一个服务也可以无需访问它自身的 I/O 服务来实现 - 单线程就足够的。 为新增的线程使用一个新的 I/O 服务的原因是，这样更简单： 线程间可以用 I/O 服务来非常容易地相互通信。 在这个例子中，async_wait() 创建了一个类型为 wait_operation 的函数对象，并通过 post() 方法将它传递给内部的 I/O 服务。 然后，在用于执行这个内部 I/O 服务的 run() 方法的线程内，调用该函数对象的重载 operator()()。 post() 提供了一个简单的方法，在另一个线程中执行一个函数对象。

wait_operation 的重载 operator()() 操作符基本上就是执行了和 wait() 方法相同的工作：调用服务实现中的阻塞式 wait() 方法。 但是，有可能这个 I/O 对象以及它的服务实现在这个线程执行 operator()() 操作符期间被销毁。 如果服务实现是在 destruct() 中销毁的，则 operator()() 操作符将不能再访问它。 这种情形是通过使用一个弱指针来防止的，从第一章中我们知道：如果在调用 lock() 时服务实现仍然存在，则弱指针 impl_ 返回它的一个共享指针，否则它将返回0。 在这种情况下，operator()() 不会访问这个服务实现，而是以一个 boost::asio::error::operation_aborted 错误来调用句柄。

#include <boost/system/error_code.hpp>
#include <cstddef>
#include <windows.h>

class timer_impl
{
  public:
    timer_impl()
      : handle_(CreateEvent(NULL, FALSE, FALSE, NULL))
    {
    }

    ~timer_impl()
    {
      CloseHandle(handle_);
    }

    void destroy()
    {
      SetEvent(handle_);
    }

    void wait(std::size_t seconds, boost::system::error_code &ec)
    {
      DWORD res = WaitForSingleObject(handle_, seconds * 1000);
      if (res == WAIT_OBJECT_0)
        ec = boost::asio::error::operation_aborted;
      else
        ec = boost::system::error_code();
    }

private:
    HANDLE handle_;
};
下载源代码
服务实现 timer_impl 使用了 Windows API 函数，只能在 Windows 中编译和使用。 这个例子的目的只是为了说明一种潜在的实现。

timer_impl 提供两个基本方法：wait() 用于等待数秒。 destroy() 则用于取消一个等待操作，这是必须要有的，因为对于异步操作来说，wait() 方法是在其自身的线程中调用的。 如果 I/O 对象及其服务实现被销毁，那么阻塞式的 wait() 方法就要尽使用 destroy() 来取消。

这个 Boost.Asio 扩展可以如下使用。

#include <boost/asio.hpp>
#include <iostream>
#include "basic_timer.hpp"
#include "timer_impl.hpp"
#include "basic_timer_service.hpp"

void wait_handler(const boost::system::error_code &ec)
{
  std::cout << "5 s." << std::endl;
}

typedef basic_timer<basic_timer_service<> > timer;

int main()
{
  boost::asio::io_service io_service;
  timer t(io_service);
  t.async_wait(5, wait_handler);
  io_service.run();
}
下载源代码
与本章开始的例子相比，这个 Boost.Asio 扩展的用法类似于 boost::asio::deadline_timer。 在实践上，应该优先使用 boost::asio::deadline_timer，因为它已经集成在 Boost.Asio 中了。 这个扩展的唯一目的就是示范一下 Boost.Asio 是如何扩展新的异步操作的。

目录监视器(Directory Monitor) 是现实中的一个 Boost.Asio 扩展，它提供了一个可以监视目录的 I/O 对象。 如果被监视目录中的某个文件被创建、修改或是删除，就会相应地调用一个句柄。 当前的版本支持 Windows 和 Linux (内核版本 2.6.13 或以上)。
*/


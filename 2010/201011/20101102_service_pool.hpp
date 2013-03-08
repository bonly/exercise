#ifndef __SERVICE_POOL_HPP_
#define __SERVICE_POOL_HPP_
#include <boost/shared_ptr.hpp>
#include "20101101_gameobj.hpp"

namespace bus{
class ServicePool{
	public:
	  boost::shared_ptr<bus::srv::ObjSrv> _obj_srv;
	
};
extern boost::shared_ptr<bus::ServicePool> g_spool;
}
#endif

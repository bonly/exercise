#include "20101101_gameobj.hpp"
#include "20101102_service_pool.hpp"
#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <boost/thread.hpp>

void push_lvl(boost::shared_ptr<bus::ServicePool> sp){
	for (int i=0; i<10; ++i){
		boost::shared_ptr<bus::srv::Level_attr> lvl = boost::make_shared<bus::srv::Level_attr>();
		lvl->level = i;
		lvl->lvl_minap = i*10;
		
		sp->_obj_srv->_lvl.insert(std::make_pair(i, lvl));
	}
}

int main(){
	bus::g_spool = boost::make_shared<bus::ServicePool>();
	bus::g_spool->_obj_srv = boost::make_shared<bus::srv::ObjSrv>();
	push_lvl(bus::g_spool);

	//boost::thread th(boost::bind(&bus::srv::CalPaladinPower,123,1,1,1,1));
	//th.join();
	
	bus::srv::TFightPower tp = bus::srv::CalPaladinPower(123,1,1,1,1);
	std::clog << "power: " << tp.power << std::endl;
	return 0;
}

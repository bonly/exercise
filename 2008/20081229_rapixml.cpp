
void
main ()
    while (true)
    {
BONLY_DEBUG              
    CLife_Pre_Modify_Status_V01 sms;
    char buf[255];
BONLY_DEBUG    
    
    for (int i=1; i<=3; ++i)
    {
	    boost::shared_ptr<User> user(new User);
BONLY_DEBUG    	
	    ParamList lstparam;
	    lstparam.insert (make_pair(
	    		std::string("123"),std::string("current life status")));
	    lstparam.insert (make_pair(
	    		std::string("124"),std::string("obj life status")));
	    lstparam.insert (make_pair(
	    		std::string("125"),std::string("left days"))); 
	    lstparam.insert (make_pair(
	    		std::string("125"),std::string("last day")));      			   			    			
BONLY_DEBUG
      strcpy(user->MSISDN,boost::lexical_cast<std::string>(13719360000+i).c_str());
	    user->UserParam.push_back(lstparam);
BONLY_DEBUG    
	    sms.add_user(user);
	   }
    sms.get_xml(buf,255);
BONLY_DEBUG    
    std::cout << buf << std::endl;
BONLY_DEBUG    	
    }
}

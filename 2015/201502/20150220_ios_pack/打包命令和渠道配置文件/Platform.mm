
extern "C" {
    
    const char* NodeName = "UIMainScence_1";

    void backToGame(){
        
    }
    
    void initPlatform()
    {
//        NSString* str =  [NSString stringWithFormat:@"%@", [[NSBundle mainBundle] objectForInfoDictionaryKey:(NSString*)kCFBundleVersionKey]];
//        
//        UnitySendMessage(NodeName, "SetVer", str.UTF8String );
//        UnitySendMessage(NodeName, "SetPlaformIP", "test.game83.com" );
        
        UnitySendMessage(NodeName, "SetPlaformIP", "thero.game83.com" );
        UnitySendMessage(NodeName, "SetPlaformName", "google" );
        UnitySendMessage(NodeName, "SetPlaformPort", "8675" );
    }
    
    void loginIos()
    {
        UnitySendMessage(NodeName, "SetUserId", "" );
        UnitySendMessage(NodeName,  "Login", "test key string" );
    }
    
    
    void initBilling()
    {
    }
   
    
    void billingIos(const char* userId, int id)
    {
        
    }
    
    void openURL(const char* URL)
    {
        NSURL * URLStr = [NSURL URLWithString:@"http://itunes.apple.com/cn/app/cai-bao-lian-meng/id844957874"];
        [[UIApplication sharedApplication]openURL:URLStr];
    }
    
}
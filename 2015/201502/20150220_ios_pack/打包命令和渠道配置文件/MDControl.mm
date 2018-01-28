//
//  MDControl.m
//  Unity-iPhone
//
//  Created by Bonly on 14-9-16.
//
//
#ifdef PLATFORM_PP

#import "MDControl.h"
#include "Platform.h"

extern "C"{
    const char* Platform = "1008";
    const char* Notice = "13";
}
NSString *UserID;
NSString *Token;

@implementation MDControl

static MDControl *sharedHelper = nil;
+ (MDControl *) sharedInstance {
    D();
    if (!sharedHelper) {
        sharedHelper = [[MDControl alloc] init];
        [sharedHelper Init];
    }
    return sharedHelper;
}

-(void)Back2Game{
    D();
}

-(void) UserCenter{
    D();
    [[PPAppPlatformKit sharedInstance] showCenter];
}

-(void) Init{
    D();
    [[PPAppPlatformKit sharedInstance] setDelegate:self];

    [[PPAppPlatformKit sharedInstance] setAppId:4415 AppKey:@"2f8346b4a44ee100120afdf73f9619ba"];
    [[PPAppPlatformKit sharedInstance] setIsNSlogData:YES];
    [[PPAppPlatformKit sharedInstance] setRechargeAmount:10];
    [[PPAppPlatformKit sharedInstance] setIsLongComet:YES];
    [[PPAppPlatformKit sharedInstance] setIsLogOutPushLoginView:YES];
    [[PPAppPlatformKit sharedInstance] setIsOpenRecharge:YES];
    [[PPAppPlatformKit sharedInstance] setCloseRechargeAlertMessage:@"关闭充值提示语"];

    [[PPUIKit sharedInstance] checkGameUpdate];
    [PPUIKit setIsDeviceOrientationLandscapeLeft:YES];
    [PPUIKit setIsDeviceOrientationLandscapeRight:YES];
//    [PPUIKit setIsDeviceOrientationPortrait:YES];
//    [PPUIKit setIsDeviceOrientationPortraitUpsideDown:YES];
}

-(void)ShowLogin{
    D();
    [[PPAppPlatformKit sharedInstance] showLogin];
//    UnitySendMessage(NodeName, "SetUserId", UserID.UTF8String);
//    UnitySendMessage(NodeName,  "Login", Token.UTF8String );

}
#pragma mark - 消息回调处理

-(void)ppVerifyingUpdatePassCallBack
{
    D();
    NSLog(@"验证游戏版本完毕回调");

//    [[PPAppPlatformKit sharedInstance] showLogin];
}

- (void)ppLoginStrCallBack:(NSString *)paramStrToKenKey
{
    D();

    [[PPAppPlatformKit sharedInstance] getUserInfoSecurity];
    NSLog([NSString stringWithFormat:@"登录后tokenKey 验证通过 == tokenKey:%@ ",paramStrToKenKey],nil);

    UserID = @"";
    Token = paramStrToKenKey;
    
    UnitySendMessage(NodeName, "SetUserId", "" );
    UnitySendMessage(NodeName,  "Login", Token.UTF8String );
}

-(void) Pay:(int64_t)userId :(const char*)name :(int)idKey :(const char*)idName :(float)money{
    D();

    NSDate *currentTime = [NSDate date];
    NSDateFormatter *dateFormatter = [[NSDateFormatter alloc] init];
    [dateFormatter setDateFormat:@"yyymmddhhmmss"];
    NSString *strTm = [dateFormatter stringFromDate: currentTime];

    [[PPAppPlatformKit sharedInstance] exchangeGoods:(int)money
                                              BillNo:[NSString stringWithFormat:@"%lld%@", userId, strTm]
                    BillTitle:[NSString stringWithFormat:@"%@(%.2f元)", [NSString stringWithUTF8String:idName], money]                                             RoleId:[NSString stringWithFormat:@"%lld", userId]
                                              ZoneId:0];
}

- (void)ppCloseWebViewCallBack:(PPWebViewCode)paramPPWebViewCode{
    D();
    //可根据关闭的WEB页面做你需要的业务处理
    NSLog(@"当前关闭的WEB页面回调是%d", paramPPWebViewCode);
//    _showLabelForPropView.text = [NSString stringWithFormat:@"当前关闭的WEB页面回调是%d", paramPPWebViewCode];
//    _showLabelForBgloginImageView.text = [NSString stringWithFormat:@"当前关闭的WEB页面回调是%d", paramPPWebViewCode];
}

-(void)ppClosePageViewCallBack:(PPPageCode)paramPPPageCode{
    D();
    //可根据关闭的VIEW页面做你需要的业务处理
    NSLog(@"当前关闭的VIEW页面回调是%d", paramPPPageCode);
    UnitySendMessage(NodeName,  "Login", "" );

//    _showLabelForPropView.text = [NSString stringWithFormat:@"当前关闭的VIEW页面回调是%d", paramPPPageCode];
//    _showLabelForBgloginImageView.text = [NSString stringWithFormat:@"当前关闭的VIEW页面回调是%d", paramPPPageCode];

}

- (void)ppPayResultCallBack:(PPPayResultCode)paramPPPayResultCode{
    D();
}

- (void)ppLogOffCallBack{
    D();
}

- (BOOL)application:(UIApplication *)application handleOpenURL:(NSURL *)url {
    D();
    [[PPAppPlatformKit sharedInstance] alixPayResult:url];
    return YES; }
@end

#endif

//
//  MDControl.h
//  Unity-iPhone
//
//  Created by Bonly on 14-9-16.
//
//
#ifdef PLATFORM_PP

#ifndef __PP_MDCONTROL_H__
#define __PP_MDCONTROL_H__

#import <PPAppPlatformKit/PPAppPlatformKit.h>

@interface MDControl : UIViewController<PPAppPlatformKitDelegate>
+(MDControl*) sharedInstance;
-(void)Init;
-(void)Pay:(int64_t)userId :(const char*)iname :(int)idKey :(const char*)idName :(float)money;
-(void)ShowLogin;
-(void)Back2Game;
-(void)UserCenter;
@end
#endif

#endif

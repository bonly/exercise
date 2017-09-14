from com.android.monkeyrunner import MonkeyRunner as MR    
from com.android.monkeyrunner import MonkeyDevice as MD    
from com.android.monkeyrunner import MonkeyImage as MI    
from com.android.monkeyrunner.easy import EasyMonkeyDevice,By    
  
device=MR.waitForConnection(10)    
if device:    
      print("Connect device successful!")    
else:    
      print("Connect device failed!")    
device=EasyMonkeyDevice(device)    
device.installPackage("D:\\MonkeyRunnerDemo\\Apps\\estore.apk")    
device.startActivity(component="com.eshore.ezone/.StartActivity")    
MR.sleep(3)    
device.touch(By.id("id/btn_disagree"),device.DOWN_AND_UP)    

#EasyMonkeyDevice类里面还有很多方法，包括exists(By)、getText(By)、type(By,String)、visible(By)
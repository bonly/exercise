#!/home/opt/adt/sdk/tools/monkeyrunner
from com.android.monkeyrunner import MonkeyDevice, MonkeyRunner
from com.android.monkeyrunner.easy import EasyMonkeyDevice, By

device = MonkeyRunner.waitForConnection()
if device:
	print("connect device success")
else:
	print("connect device failed!")

easy_device = EasyMonkeyDevice(device)

#print("get")
#print(By.id('id/btn_east'))
#print("end")

view = device.getHierarchyViewer()
print(view);
btn = view.findViewById('id/setStepValue')
print(btn)
text = view.getText(btn)
print text.encode('utf-8')
#device.startActivity(component="com.axndx.prithvee.pokemongocontrols/.MainActivity")
#easy_device.touch(By.id('id/btn_east'), device.DOWN_AND_UP) #By.id('id/btn_east') #0x187548e

# device.getViewIdList()


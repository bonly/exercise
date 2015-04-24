#!/usr/bin/python
from com.android.monkeyrunner import MonkeyRunner, MonkeyDevice
from com.android.monkeyrunner import MonkeyImage
from com.android.monkeyrunner.easy import EasyMonkeyDevice
#from com.android.monkeyrunner.easy import HierarchyViewer
from com.android.monkeyrunner.easy import By
#from com.android.hierarchyviewerlib.device import ViewNode

device = MonkeyRunner.waitForConnection(0,"SH13TPL09945")
#device = MonkeyRunner.waitForConnection(0, "19761202")
easy_device = EasyMonkeyDevice(device)
print 'start Contacts'
#device.installPackage('./es.apk')
#device.shell('am start -a android.intent.action.MAIN -n com.android.contacts/.DialtactsContactsEntryActivity')
device.startActivity(component='com.polyvi.cupmp/.ui.user.UserRegisterActivity')
#device.startActivity(component='com.polyvi.cupmp/.ui.ImppActivity')
MonkeyRunner.sleep(3)

#text=easy_device.visible(By.id('id/phoneNumber'))
#print text
#hierarchy_viewer = device.getHierarchyViewer()
#view_node=hierarchy_viewer.findViewById('id/phoneNumber')
#view_node=Byid('id/phoneNumber')
#print view_node

device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
device.type('15360534225')
#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
#device.type('mypasswd')
#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
#device.type('mypasswd')
#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
#device.type('3231')
#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')
#!/usr/bin/python
from com.android.monkeyrunner import MonkeyRunner, MonkeyDevice
from com.android.monkeyrunner import MonkeyImage
from com.android.monkeyrunner.easy import EasyMonkeyDevice
#from com.android.monkeyrunner.easy import HierarchyViewer
from com.android.monkeyrunner.easy import By
#from com.android.hierarchyviewerlib.device import ViewNode

device = MonkeyRunner.waitForConnection(0,"SH13TPL09945")
easy_device = EasyMonkeyDevice(device)
print 'start Contacts'
#device.installPackage('./es.apk')
#device.shell('am start -a android.intent.action.MAIN -n com.android.contacts/.DialtactsContactsEntryActivity')
device.startActivity(component='netease.pomelo.chat/.MainActivity')
MonkeyRunner.sleep(3)

#text=easy_device.visible(By.id('id/phoneNumber'))
#print text
hierarchy_viewer = device.getHierarchyViewer()
view_node=hierarchy_viewer.findViewById('id/channel')
print view_node
easy_device.touch(By.id('id/channel'),device.DOWN_AND_UP) 
device.type('mychn\n')

view_node=hierarchy_viewer.findViewById('id/name')
print view_node
easy_device.touch(By.id('id/name'),device.DOWN_AND_UP) 
device.type('myname\n')

#device.press('KEYCODE_DPAD_DOWN', 'DOWN_AND_UP')


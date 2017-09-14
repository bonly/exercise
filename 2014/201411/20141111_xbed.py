#!/home/opt/adt/sdk/tools/monkeyrunner
# -*- coding: UTF-8 -*-
#auth: bonly
#create: 2015.11.1

from com.android.monkeyrunner import MonkeyDevice, MonkeyRunner
from com.android.monkeyrunner.easy import EasyMonkeyDevice, By

def search_city(viewid, str):
	view = device.getHierarchyViewer()
	gv = view.findViewById(viewid).children #找到GridView的子列表
	children_count = len(gv)                #list的大小
	for idx in range(0, children_count):    #遍历
		viewLayout = gv[idx]                #子layout
		btn = viewLayout.children[0]  #第一个子节点就是要找的对象
		text = view.getText(btn)
		if text.encode('utf-8') == str:
			return btn
	return None


import sys
reload(sys)
sys.setdefaultencoding('utf-8')

device = MonkeyRunner.waitForConnection()
if device:
	print("连接成功")
else:
	print("连接失败")

device.startActivity(component="com.xbed.xbed/.ui.MainActivity")
device.startActivity(component="com.xbed.xbed/.ui.SearchActivity")

easy_device = EasyMonkeyDevice(device)

MonkeyRunner.sleep(1)

# easy_device.touch(By.id('id/edtTxt_search'), MonkeyDevice.DOWN_AND_UP);
# MonkeyRunner.sleep(1)
# device.press('KEYCODE_BACK', MonkeyDevice.DOWN_AND_UP);  
# device.type('aaaa')

# easy_device.type(By.id('id/edtTxt_search'), 'yyyt ');

# gz = device.getViewsByText('深圳市')

sr = search_city('id/gv_hot_city', '广州市d')
if sr is not None:
	print "ok"
else:
	print "no"

# btn = view.findViewById('id/tv_hot_city')
# text = view.getText(btn)
# print text.encode('utf-8')


# view = device.getHierarchyViewer()
# print(view);
# btn = view.findViewById('id/setStepValue')
# print(btn)
# text = view.getText(btn)
# print text.encode('utf-8')
#device.startActivity(component="com.axndx.prithvee.pokemongocontrols/.MainActivity")
#easy_device.touch(By.id('id/btn_east'), device.DOWN_AND_UP) #By.id('id/btn_east') #0x187548e

# device.getViewIdList()


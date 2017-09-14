#!/home/opt/adt/sdk/tools/monkeyrunner
# -*- coding: UTF-8 -*-
#auth: bonly
#create: 2015.11.1

from com.android.monkeyrunner import MonkeyDevice, MonkeyRunner
from com.android.monkeyrunner.easy import EasyMonkeyDevice, By

import sys
reload(sys)
sys.setdefaultencoding('utf-8')

device = MonkeyRunner.waitForConnection()
if device:
	print("连接成功")
else:
	print("连接失败")

def search_city(viewid, str):
	view = device.getHierarchyViewer()
	gv = view.findViewById(viewid).children #找到GridView的子列表
	children_count = len(gv)                #list的大小
	for idx in range(0, children_count):    #遍历
		viewLayout = gv[idx]                #子layout
		btn = viewLayout.children[0]        #第一个子节点就是要找的对象
		text = view.getText(btn)
		if text.encode('utf-8') == str:
			return btn
	return None


def search_room(viewid, str):
	view = device.getHierarchyViewer()
	fl = view.findViewById(viewid).children   #lv_room_list:LinearLayout  --> FrameLayout
	lv = fl[0].children_count
	# children_count = len(gv)                  #
	# for idx in range(0, children_count):
	# 	viewLayout = gv[idx]
	# 	btn = viewLayout.children[0]
	# 	text = view.getText(btn)
	# 	print(text.encode('utf-8'))


device.startActivity(component="com.xbed.xbed/.ui.MainActivity")
device.startActivity(component="com.xbed.xbed/.ui.SearchActivity")

easy_device = EasyMonkeyDevice(device)

MonkeyRunner.sleep(1)

sr = search_city('id/gv_hot_city', '广州市')
if sr is not None:
	print("找到广州市，点击")
else:
	print("找不到广州市")
	exit()

view = device.getHierarchyViewer()
pt = view.getAbsoluteCenterOfView(sr)
device.touch(pt.x, pt.y, device.DOWN_AND_UP)  #点击广州

MonkeyRunner.sleep(5)

search_room('id/lv_room_list', '574新name_琶洲会展中心邦泰国际大床房')


# device.getViewIdList()


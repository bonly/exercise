#!/usr/bin/python
#-*-coding:utf-8-*-
'''
Created on 2011-3-10

@author: bonly
'''

from GChartWrapper import GChart,Line

def simple():
   G = GChart()
   # Set the chart type, either Google API type or regular name
   G.type('pie')
   # Update the chart's dataset, can be two dimensional or contain string data
   G.dataset('helloworld')
   # Set the size of the chart, default is 300x150
   G.size(250, 100) 
   print G
   
def margins_test():
   G = Line([[8,12,11,16,22],[7,8,10,12,19,23]], encoding='simple')   #实际数据
   G.size(1000, 200) #整个图大小
   G.label(1,5,10,20,40,60,100,160,260,420,68,1100,1780) #下标量
   G.fill('bg', 's', 'e0e0e0') #背景
   G.color('black', 'blue')  #线颜色
   G.margin(20,20,20,30,80,20)  #图位置
   G.legend('Merge', 'BigTable')  #线标签
   print G
   
def line_test():
   G = Line([[0,10,20], [0,20,25]], encoding='text')  #数据  simple模式和text模式
   G.size(1000, 200)  #整个图的大小
   G.axes.type('xy')  #XY两轴,可用xyz
   G.axes.label(0, 30) #设置0(x)轴标签
   #G.axes.label(1, '5', '10', '15', '20', '25', '30')  #设置1(Y)轴标签
   #G.axes.range(0, 0, 5, 1)  #设置0(x)轴范围
   G.axes.range(1, 0, 50)   #设置1(Y)轴范围
   G.scale(0, 50)  #设置放大比率,最小最大
   G.fill('bg', 's', 'e0e0e0') #背景色
   G.color('black', 'blue')  #两条线的颜色
   G.margin(20, 20, 20, 30, 80, 20)  #表在图中的位置
   G.legend('Merge', 'BigTable')  #两标签名字
   G.show()  #网页显示
   #print G   #打印网址
   #G.save('/tmp/my.png') #保存文件
   
def mylabel(index, *args):
   if isinstance(args, tuple):
      print 'True'
   print str('%s:|%s' % (index, '|'.join(map(str, args)))).replace('None', '')
   print str('%s:|%s' % (index, '|'.join(map(str, args)))).replace(' ', '').replace('[', '').replace(']', '').replace('[', '').replace(']', '').replace(',', '|')
  
     
if __name__ == '__main__':
   #simple()
   #margins_test()
   line_test()
   

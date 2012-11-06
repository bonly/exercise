#!/usr/bin/python
import sys
from PySide.QtCore import *
from PySide.QtGui import *
 
if(__name__ == '__main__'):
    #app = QApplication(sys.argv)
    app = QApplication("myname")
    hellobt = QPushButton('Say Hello');
    hellobt.show()
    sys.exit(app.exec_())

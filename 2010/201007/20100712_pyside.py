#!/usr/bin/python
import sys
from PySide.QtCore import *
from PySide.QtGui import *
import emb
 
if(__name__ == '__main__'):
    #app = QApplication(sys.argv)
    app = QApplication("myname");
    hellobt = QPushButton('Say Hello');
    hellobt.clicked.connect(emb.hello);
    hellobt.show();
    sys.exit(app.exec_());


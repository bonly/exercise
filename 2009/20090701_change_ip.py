# -*- coding: cp936 -*-
import wmi
import pythoncom;

con =wmi.WMI()
ip = "10.1.10.20"
subnet = "255.255.255.0"

wql = "SELECT * FROM Win32_NetworkAdapterConfiguration WHERE IPEnabled = TRUE"

for adapter in con.query(wql):
    print adapter
    #ReturnValue = adapter.EnableStatic(IPAddress=ip, SubnetMask=subnet)
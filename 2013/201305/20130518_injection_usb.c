#include <usb.h>
#include <dlfcn.h>
#include <stdio.h>

usb_dev_handle *usb_open(struct usb_device *dev)
{
  static usb_dev_handle *(*libusb_open)
             (struct usb_device *dev) = NULL;
  void *handle;
  usb_dev_handle *usb_handle;
  char *error;

  if (!libusb_open) {
    handle = dlopen("/usr/lib/libusb.so",
                    RTLD_LAZY);
    if (!handle) {
      fputs(dlerror(), stderr);
      exit(1);
    }
    libusb_open = dlsym(handle, "usb_open");
    if ((error = dlerror()) != NULL) {
      fprintf(stderr, "%s\n", error);
      exit(1);
    }
  }

  printf("calling usb_open(%s)\n", dev->filename);
  usb_handle = libusb_open(dev);
  printf("usb_open() returned %p\n", usb_handle);
  return usb_handle;
}
// user for LD_PRELOAD 
/*
 gcc -Wall -O2 -fpic -shared -ldl -o shim.so 20130518_injection_usb.c
*/
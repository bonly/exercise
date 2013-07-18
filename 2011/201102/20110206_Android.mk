LOCAL_PATH:= $(call my-dir)

thread_SRC_FILES := \
        libs/thread/src/pthread/thread.cpp \
        libs/thread/src/pthread/once.cpp \

filesystem_SRC_FILES := \
        libs/filesystem/src/codecvt_error_category.cpp \
        libs/filesystem/src/operations.cpp \
        libs/filesystem/src/path.cpp \
        libs/filesystem/src/path_traits.cpp \
        libs/filesystem/src/portability.cpp \
        libs/filesystem/src/utf8_codecvt_facet.cpp \
        libs/system/src/error_code.cpp \

regex_SRC_FILES := \
        libs/regex/src/cpp_regex_traits.cpp \
        libs/regex/src/cregex.cpp \
        libs/regex/src/c_regex_traits.cpp \
        libs/regex/src/fileiter.cpp \
        libs/regex/src/icu.cpp \
        libs/regex/src/instances.cpp \
        libs/regex/src/posix_api.cpp \
        libs/regex/src/regex.cpp \
        libs/regex/src/regex_debug.cpp \
        libs/regex/src/regex_raw_buffer.cpp \
        libs/regex/src/regex_traits_defaults.cpp \
        libs/regex/src/static_mutex.cpp \
        libs/regex/src/usinstances.cpp \
        libs/regex/src/w32_regex_traits.cpp \
        libs/regex/src/wc_regex_traits.cpp \
        libs/regex/src/wide_posix_api.cpp \
        libs/regex/src/winstances.cpp \

common_SRC_FILES := $(thread_SRC_FILES) $(filesystem_SRC_FILES)

include $(CLEAR_VARS)
LOCAL_MODULE:= boost
LOCAL_SRC_FILES := $(common_SRC_FILES)

#prebuilt_stdcxx_PATH := prebuilts/ndk/current/sources/cxx-stl/gnu-libstdc++

LOCAL_C_INCLUDES := \
        $(LOCAL_PATH)/boost \
        $(prebuilt_stdcxx_PATH)/include \
       $(prebuilt_stdcxx_PATH)/libs/$(TARGET_CPU_ABI)/include/ \

LOCAL_C_INCLUDES := \
        $(prebuilt_stdcxx_PATH)/include \
        $(prebuilt_stdcxx_PATH)/libs/$(TARGET_CPU_ABI)/include/ \
        $(prebuilt_supccxx_PATH)/include

LOCAL_CFLAGS += -fvisibility=hidden -lpthread
LOCAL_CPPFLAGS += -fexceptions -frtti

LOCAL_SHARED_LIBRARIES := libc libstdc++ libstlport

#用CrystaX版本编译器的可能需要 LOCAL_LDFLAGS += -L$(prebuilt_stdcxx_PATH)/libs/$(TARGET_CPU_ABI) -lgnustl_static -lsupc++
LOCAL_LDFLAGS += -L$(prebuilt_stdcxx_PATH)/libs/$(TARGET_CPU_ABI) 

#LOCAL_MODULE_TAGS := optional
include $(BUILD_SHARED_LIBRARY)


include $(CLEAR_VARS)
LOCAL_MODULE:= filesystem
LOCAL_SRC_FILES := $(filesystem_SRC_FILES)
include $(BUILD_SHARED_LIBRARY)

include $(CLEAR_VARS)
LOCAL_MODULE:= system
LOCAL_SRC_FILES := libs/system/src/error_code.cpp 
include $(BUILD_SHARED_LIBRARY)

include $(CLEAR_VARS)
LOCAL_MODULE:= regex
LOCAL_SRC_FILES := $(regex_SRC_FILES)
include $(BUILD_SHARED_LIBRARY)

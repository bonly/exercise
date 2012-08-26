class _20091114_jni
{

  static ///要动态加载库,必须写此静态无名方法在类中 
  {
    System.loadLibrary("hello");
  }

  public native void displayHelloWorld();/// 如果函数要成为本地方法,必须声明为native, 并且不能实现,hello是动态库名
  public static void main(String[] args)
  {
    _20091114_jni JN = new _20091114_jni();
    JN.displayHelloWorld();
  }
}

/**
编译 javac _20091114_jni.java
*/

class Hello {
        private native void print();
        public static void main(String[] args) {
                new Hello().print();
                System.err.println("Hello from Java");
        }
        static {
                System.loadLibrary("hello");
        }
}`
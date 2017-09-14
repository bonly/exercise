
class Qunar{
	public native String Encode(String data, String key);
	public static void main(String[] args){
		// for (int idx=0; idx<1; idx++){
			Qunar obj = new Qunar();
			String str;
			str = obj.Encode("timeStamp=2016-11-02 10:50:09&caller=Xbed&traceId=201611021050092&groupId=Xbed&roomIds=xbedtestroom101","gvgmzItX3W_K1XDdFd009aoxk3AsQ0Dbt0OUw9RnA6cQhakxo39P8t-y6Qu62HvEENyo5jfAdSziaOR2Y5CawQ");
			System.out.println(str);
		// }
	}
	static{
		System.loadLibrary("qunar");
	}
}
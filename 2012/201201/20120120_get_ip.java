public static String GetNetIp()
{
  URL infoUrl = null;
  InputStream inStream = null;
  try
  {
      infoUrl = new URL("http://city.ip138.com/city0.asp");
      URLConnection connection = infoUrl.openConnection();
      HttpURLConnection httpConnection = (HttpURLConnection)connection;
      int responseCode = httpConnection.getResponseCode();
      if(responseCode == HttpURLConnection.HTTP_OK)
      { 
          inStream = httpConnection.getInputStream(); 
          BufferedReader reader = new BufferedReader(new InputStreamReader(inStream,"utf-8"));
          StringBuilder strber = new StringBuilder();
          String line = null;
          while ((line = reader.readLine()) != null) 
              strber.append(line + "\n");
          inStream.close();

          //从反馈的结果中提取出IP地址

          int start = strber.indexOf("[");
          int end = strber.indexOf("]", start + 1);
          line = strber.substring(start + 1, end);
          return line; 
      }
  }

  catch(MalformedURLException e) {
      e.printStackTrace();
  }

  catch (IOException e) {
      e.printStackTrace();
  }

  return null;

}
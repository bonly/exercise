public class CommandResult {
    public String result = "";
    public String error = "";

    public static runCommand(String command, boolean root) {
        Log.e("runCommand", command);

        Process process = null;
        DataOutputStream os = null;
        DataInputStream stdout = null;
        DataInputStream stderr = null;
        CommandResult ret = new CommandResult();
        try {
            StringBuffer output = new StringBuffer();
            StringBuffer error = new StringBuffer();
            if (root) {
                process = Runtime.getRuntime().exec("su");
                os = new DataOutputStream(process.getOutputStream());
                os.writeBytes(command + "\n");
                os.writeBytes("exit\n");
                os.flush();
            } else {
                process = Runtime.getRuntime().exec(command);
            }

            stdout = new DataInputStream(process.getInputStream());
            String line;
            while ((line = stdout.readLine()) != null) {
                output.append(line).append('\n');
            }
            stderr = new DataInputStream(process.getErrorStream());
            while ((line = stderr.readLine()) != null) {
                error.append(line).append('\n');
            }
            process.waitFor();
            ret.result = output.toString().trim();
            ret.error = error.toString().trim();
        } catch (Exception e) {
            ret.result = "";
            ret.error = e.getMessage();
        } finally {
            try {
                if (os != null) {
                    os.close();
                }
                if (stdout != null) {
                    stdout.close();
                }
                if (stderr != null) {
                    stderr.close();
                }
                process.destroy();
            } catch (Exception e) {
                ret.result = "";
                ret.error = e.getMessage();
            }
        }

        return ret;
    }
}
//http://paradigmx.net/blog/2012/02/05/go-for-android/
//http://gimite.net/en/index.php?Run%20native%20executable%20in%20Android%20App
/*
Include the binary go-exec in the assets folder.
Use getAssets().open("go-exec") to get an InputStream.
Write it to /data/data/app-package-name/, where the app has access to write files and make it executable.
Make it executable using the code above, i.e. CommandResult.runCommand("/system/bin/chmod 744 /data/data/app-package-name/go-exec", 0)
Run /data/data/app-package-name/go-exec using the code above.
*/

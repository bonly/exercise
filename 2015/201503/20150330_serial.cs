[Serializable]
public class Student : ISerializable
{
    private string _name;

    public string Name
    {
        get { return _name; }
        set { _name = value; }
    }
    [SecurityPermission(SecurityAction.Demand, SerializationFormatter = true)]
    public void GetObjectData(SerializationInfo info, StreamingContext context)
    {
        info.SetType(typeof(SerializationHelper));
    }
}

[Serializable]
public class SerializationHelper : IObjectReference
{
    public object GetRealObject(StreamingContext context)
    {
        return "新的类型哦";
    }
}

static void Main(string[] args)
{
    Student student = new Student { Name = "马里奥" };
    using (var stream = new MemoryStream())
    {
        BinaryFormatter formatter = new BinaryFormatter();
        formatter.Serialize(stream, student);
        stream.Position = 0;

        var deserializeValue = formatter.Deserialize(stream);
        Console.Write(deserializeValue.ToString());
        Console.Read();
    }
}
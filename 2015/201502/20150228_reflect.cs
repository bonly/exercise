using System;
using System.Reflection;


public class Rf{
  public int Abc;
  public string str;

  static void Main(string[] arg){
    Rf rf = new Rf();
    rf.Abc = 10;
    rf.str = "ok";
    rf.Write(rf);
  }

  public void Write(Rf rf){
      Type typ = typeof(Rf);
      FieldInfo[]  fields = typ.GetFields(BindingFlags.c | BindingFlags.Instance);
      foreach( var field in fields){
          Console.Write(field.Name + "\t");
          Console.WriteLine(field.GetValue(rf));
      }
  }
};


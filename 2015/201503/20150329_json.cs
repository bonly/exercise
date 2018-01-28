using System;
using System.IO;
using System.Runtime.Serialization;
using System.Runtime.Serialization.Formatters.Binary;
using System.Runtime.Serialization.Formatters.Soap;
using System.Text;
using System.Windows.Forms;

namespace WinFormsFormatter
{
    public partial class Form1 : Form
    {
        /// <summary>
        /// 序列化檔案名稱
        /// </summary>
        string FileName = string.Format(@"{0}\\{1}", Application.StartupPath, "demo.txt");

        public Form1()
        {
            InitializeComponent();
        }

        #region 按鈕事件
        private void btnSerialize_Click(object sender, EventArgs e)
        {
            if (rbBinary.Checked == true)
            {
                SerializeBinary();
            }
            else
            {
                SerializeSoap();
            }
        }

        private void btnDeserialize_Click(object sender, EventArgs e)
        {
            ClsSerializable clsSerializable = null;

            if (rbBinary.Checked)
            {
                clsSerializable = DeserializeBinary();
            }
            else
            {
                clsSerializable = DeserializeSoap();
            }

            // 顯示還原序列化後的類別物件於畫面中
            string strContect = string.Format("Number: {0}\\nName: {1}\\nCmt: {2}", clsSerializable.Number, clsSerializable.Name, clsSerializable.Cmt);
            this.rtbContent.Text = strContect;
        }

        #endregion
        #region BinaryFormatter

        /// <summary>
        /// 使用 BinaryFormatter 進行序列化
        /// </summary>
        private void SerializeBinary()
        {
            // 建立 ClsSerializable 類別物件
            ClsSerializable clsSerializable = new ClsSerializable();

            // 建立檔案資料流物件
            using (FileStream fileStream = new FileStream(FileName, FileMode.Create, FileAccess.Write))
            {
                // 建立 BinaryFormatter 物件
                BinaryFormatter binaryFormatter = new BinaryFormatter();

                // 將物件進行二進位序列化，並且儲存檔案
                binaryFormatter.Serialize(fileStream, clsSerializable);
            }

            // 將序列化後的檔案內容呈現到表單畫面
            StringBuilder sbContent = new StringBuilder();

            foreach (var byteData in File.ReadAllBytes(FileName))
            {
                sbContent.Append(byteData);
                sbContent.Append(" ");
            }
            this.rtbContent.Text = sbContent.ToString();
        }

        /// <summary>
        /// 使用 BinaryFormatter 進行還原序列化
        /// </summary>
        /// <returns></returns>
        private ClsSerializable DeserializeBinary()
        {
            // 建立 ClsSerializable 類別物件
            ClsSerializable clsSerializable = null;

            // 建立檔案資料流物件
            using (FileStream fileStream = new FileStream(FileName, FileMode.Open))
            {
                // 建立 BinaryFormatter 物件
                BinaryFormatter binaryFormatter = new BinaryFormatter();

                // 將檔案內容還原序列化成 Object 物件，並且進一步轉型成正確的型別 ClsSerializable
                clsSerializable = (ClsSerializable)binaryFormatter.Deserialize(fileStream);
            }

            return clsSerializable;
        }

        #endregion
        #region SoapFormatter

        /// <summary>
        /// 使用 SoapFormatter 進行序列化
        /// </summary>
        private void SerializeSoap()
        {
            // 建立 ClsSerializable 類別物件
            ClsSerializable clsSerializable = new ClsSerializable();

            // 建立檔案資料流物件
            using (FileStream fileStream = new FileStream(FileName, FileMode.Create, FileAccess.Write))
            {
                // 建立 SoapFormatter 物件
                SoapFormatter soapFormatter = new SoapFormatter();

                // 將物件進行 SOAP 序列化，並且儲存檔案
                soapFormatter.Serialize(fileStream, clsSerializable);
            }

            // 將序列化後的檔案內容呈現到表單畫面
            rtbContent.Text = File.ReadAllText(FileName);
        }

        /// <summary>
        /// 使用 SoapFormatter 進行還原序列化
        /// </summary>
        /// <returns></returns>
        private ClsSerializable DeserializeSoap()
        {
            // 建立 ClsSerializable 類別物件
            ClsSerializable clsSerializable = null;

            // 建立檔案資料流物件
            using (FileStream fileStream = new FileStream(FileName, FileMode.Open))
            {
                // 建立 SoapFormatter 物件
                SoapFormatter soapFormatter = new SoapFormatter();

                // 將檔案內容還原序列化成 Object 物件，並且進一步轉型成正確的型別 ClsSerializable
                clsSerializable = (ClsSerializable)soapFormatter.Deserialize(fileStream);
            }

            return clsSerializable;
        }

        #endregion
        [Serializable]

        public class ClsSerializable : ISerializable
        {
            private int _Number = 0;
            private string _Name = "John";
            private string _Cmt = "沒有";

            public int Number
            {
                get 
                { 
                    return this._Number; 
                }
            }

            public string Name
            {
                get 
                { 
                    return this._Name; 
                }
            }

            public string Cmt
            {
                get 
                { 
                    return this._Cmt; 
                }
            }

            public ClsSerializable()
            {
                ModifyMemberValue();
            }

            public void GetObjectData(SerializationInfo info, StreamingContext context)
            {
                info.AddValue("_Number", _Number);
                info.AddValue("_Name", _Name);
                info.AddValue("_Cmt", _Cmt);
            }

            public ClsSerializable(SerializationInfo info, StreamingContext context)
            {
                _Number = (int)info.GetValue("_Number", _Number.GetType());
                _Name = (string)info.GetValue("_Name", _Name.GetType());
                _Cmt = (string)info.GetValue("_Cmt", _Cmt.GetType());
            }

            public void ModifyMemberValue()
            {
                this._Number = 7;
                this._Name = "Ou";
                this._Cmt = "喜歡音樂";
            }
        }
    }
}
#include <fstream>
#include <iostream>
#include <cstring>
#include <cerrno>
#include <cstdlib>
#include <iomanip>
using namespace std;

struct Addition
{
        char change_part[1024];
};


struct CBE
{
        char EventFormatType[2];
        char roll_flag[1];
        char roll_count[2];
        char file_id[10];
        char exc_id[4];
        char FileType[2];
        char subno[24];
        char IMSI[15];
        char IMEI[15];
        char start_time[14];
        char special_flag[2];
        char proc_time[14];
        char event_id[20];
        char Switch_flag[2];
        char District[2];
        char Brand[1];
        char User_Type[2];
        char Visit_Area[8];
        char B_subno[24];
        char Bill_type[2];
        char ACCT_Mob[14];
        char ACCT_Toll[14];
        char ACCT_Inf[14];
        char Mob_fee[8];
        char Toll_fee[8];
        char Inf_fee[8];
        char Pay_mode[1];
        char dis_id[64];
        char reserve[36];
        char cbe_flag[1];
        char period_flag[1];
        char SubsID[14];
        char A_pay_type[1];
        char A_pay_subno[24];
        char A_pay_switch_flag[2];
        char A_pay_district[2];
        char A_pay_brand[1];
        char A_pay_user_type[2];
        char A_AcctID[14];
        char A_deducted[1];
        char A_ACCT_BALANCE[12];
        char A_ACCT_BALANCE_ID1[18];
        char A_ACCT_BALANCE_AMT1[8];
        char A_ACCT_BALANCE_ID2[18];
        char A_ACCT_BALANCE_AMT2[8];
        char A_ACCT_BALANCE_ID3[18];
        char A_ACCT_BALANCE_AMT3[8];
        char A_ACCT_BALANCE_ID4[18];
        char A_ACCT_BALANCE_AMT4[8];
        char B_pay_type[1];
        char B_pay_subno[24];
        char B_pay_switch_flag[2];
        char B_pay_district[2];
        char B_pay_brand[1];
        char B_pay_user_type[2];
        char B_AcctID[14];
        char B_deducted[1];
        char B_ACCT_BALANCE[12];
        char B_ACCT_BALANCE_ID1[18];
        char B_ACCT_BALANCE_AMT1[8];
        char B_ACCT_BALANCE_ID2[18];
        char B_ACCT_BALANCE_AMT2[8];
        char B_ACCT_BALANCE_ID3[18];
        char B_ACCT_BALANCE_AMT3[8];
};

class opCBE
{
    public:
        opCBE(CBE &c,Addition &add):cbe(c),add(add){};
        char* sql_insert(char* buf, size_t &len)
        {
            if (!buf || !len)
                return 0;

            string sql;
            sql.reserve(len);
            sql.append("insert into BILL_201112_01_0000_001 values ");
            sql.append("(");
#define SQLAP(N,L) { \
    if(!L) sql.append("'").append(N,sizeof(N)).append("',\n"); \
    else  sql.append("'").append(N,sizeof(N)).append("'\n"); \
}
            SQLAP(cbe.EventFormatType,false);
            SQLAP(cbe.roll_flag,false);
            SQLAP(cbe.roll_count,false);
            SQLAP(cbe.file_id,false);
            SQLAP(cbe.exc_id,false);
            SQLAP(cbe.FileType,false);
            SQLAP(cbe.subno,false);
            SQLAP(cbe.IMSI,false);
            SQLAP(cbe.IMEI,false);
            SQLAP(cbe.start_time,false);
            SQLAP(cbe.special_flag,false);
            SQLAP(cbe.proc_time,false);
            SQLAP(cbe.event_id,false);
            SQLAP(cbe.Switch_flag,false);
            SQLAP(cbe.District,false);
            SQLAP(cbe.Brand,false);
            SQLAP(cbe.User_Type,false);
            SQLAP(cbe.Visit_Area,false);
            SQLAP(cbe.B_subno,false);
            SQLAP(cbe.Bill_type,false);
            SQLAP(cbe.ACCT_Mob,false);
            SQLAP(cbe.ACCT_Toll,false);
            SQLAP(cbe.ACCT_Inf,false);
            SQLAP(cbe.Mob_fee,false);
            SQLAP(cbe.Toll_fee,false);
            SQLAP(cbe.Inf_fee,false);
            SQLAP(cbe.Pay_mode,false);
            SQLAP(cbe.dis_id,false);
            SQLAP(cbe.reserve,false);
            SQLAP(cbe.cbe_flag,false);
            SQLAP(cbe.period_flag,false);
            SQLAP(cbe.SubsID,false);
            SQLAP(cbe.A_pay_type,false);
            SQLAP(cbe.A_pay_subno,false);
            SQLAP(cbe.A_pay_switch_flag,false);
            SQLAP(cbe.A_pay_district,false);
            SQLAP(cbe.A_pay_brand,false);
            SQLAP(cbe.A_pay_user_type,false);
            SQLAP(cbe.A_AcctID,false);
            SQLAP(cbe.A_deducted,false);
            SQLAP(cbe.A_ACCT_BALANCE,false);
            SQLAP(cbe.A_ACCT_BALANCE_ID1,false);
            SQLAP(cbe.A_ACCT_BALANCE_AMT1,false);
            SQLAP(cbe.A_ACCT_BALANCE_ID2,false);
            SQLAP(cbe.A_ACCT_BALANCE_AMT2,false);
            SQLAP(cbe.A_ACCT_BALANCE_ID3,false);
            SQLAP(cbe.A_ACCT_BALANCE_AMT3,false);
            SQLAP(cbe.A_ACCT_BALANCE_ID4,false);
            SQLAP(cbe.A_ACCT_BALANCE_AMT4,false);
            SQLAP(cbe.B_pay_type,false);
            SQLAP(cbe.B_pay_subno,false);
            SQLAP(cbe.B_pay_switch_flag,false);
            SQLAP(cbe.B_pay_district,false);
            SQLAP(cbe.B_pay_brand,false);
            SQLAP(cbe.B_pay_user_type,false);
            SQLAP(cbe.B_AcctID,false);
            SQLAP(cbe.B_deducted,false);
            SQLAP(cbe.B_ACCT_BALANCE,false);
            SQLAP(cbe.B_ACCT_BALANCE_ID1,false);
            SQLAP(cbe.B_ACCT_BALANCE_AMT1,false);
            SQLAP(cbe.B_ACCT_BALANCE_ID2,false);
            SQLAP(cbe.B_ACCT_BALANCE_AMT2,false);
            SQLAP(cbe.B_ACCT_BALANCE_ID3,false);
            SQLAP(cbe.B_ACCT_BALANCE_AMT3,false);
            sql.append("'").append(add.change_part,strlen(add.change_part)).append("'\n");
            sql.append(");");
            size_t sqlen = sql.length();
            if (sqlen >= len)
            {
                cerr << "len too small to set the sql string, need " << sqlen << endl;
                return 0;
            }
            memcpy(buf,sql.c_str(),sqlen);
            len = sqlen;
            return buf;
        }
    public:
        CBE &cbe;
        Addition &add;
};

#define PV(V) printv(#V,V,sizeof(V))
void printv(const char* name, const char* value, size_t len)
{
    //cout << name << ": " << setw(len) << value << endl; ///即使设置了输出长度,还是不会截断
    //printf("%s: %2s\n", name, value); ///同样不会截断
    char tmp[len+1];
    memcpy(tmp, value, len);
    tmp[len+1]='\0';
    clog << name << ": " << tmp << endl;
}

int loadCBE(CBE &cbe, char* mem)
{
   return 0;
}

int main(int argc, char* argv[])
{
    fstream fs(argv[1]);
    string outfile(argv[1]);
    outfile.append(".sql");
    fstream out(outfile.c_str(),ios_base::out);

    if (fs.is_open()==false||out.is_open()==false)
    {
        clog << strerror(errno);
        exit(-1);
    }

    for(;;)
    {
        CBE cbe;

        ///getline会把最后一位补为0结束
        //fs.getline((char*)&cbe, sizeof(CBE), '\n');
        //fs.getline((char*)&cbe, sizeof(CBE));
        fs.read((char*)&cbe,sizeof(CBE));

        Addition ad;
        fs.get((char*)&ad.change_part, (int)sizeof(ad.change_part),'\r');

        if(!fs.good())
            break;
        ///跳过\r\n
        fs.seekg(+2,ios::cur);
        //PV(cbe.EventFormatType);

        char buf[2048]="";

        opCBE p (cbe,ad);
        size_t len = sizeof(buf);
        if(!p.sql_insert(buf, len))
        {
            exit(-1);
        }
        printf("len:%d\n%s\n", len, buf);
        out.write(buf, len);
        out.write("\n",1);
        // clog << buf << endl;

    }
    fs.close();
    out.close();
    return 0;
}

/*
CREATE TABLE `indb_newtable_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tablename` char(32) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `ins_count` bigint(20) DEFAULT NULL,
  `table_num` int(11) DEFAULT NULL,
  `business_type` char(2) DEFAULT NULL,
  `day` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39404 DEFAULT CHARSET=utf8;

insert into indb_newtable_record values(0,'BILL_201112_01',201112,0,1000,0,'01',27);

create table BILL_201112_01_4727_000
(
   EventFormatType char(2),
   roll_flag char(1),
   roll_count char(2),
   file_id char(10),
   exc_id char(4),
   FileType char(2),
   subno char(24),
   IMSI char(15),
   IMEI char(15),
   start_time char(14),
   special_flag char(2),
   proc_time char(14),
   event_id char(20),
   Switch_flag char(2),
   District char(2),
   Brand char(1),
   User_Type char(2),
   Visit_Area char(8),
   B_subno char(24),
   Bill_type char(2),
   ACCT_Mob char(14),
   ACCT_Toll char(14),
   ACCT_Inf char(14),
   Mob_fee char(8),
   Toll_fee char(8),
   Inf_fee char(8),
   Pay_mode char(1),
   dis_id char(64),
   reserve char(36),
   cbe_flag char(1),
   period_flag char(1),
   SubsID char(14),
   A_pay_type char(1),
   A_pay_subno char(24),
   A_pay_switch_flag char(2),
   A_pay_district char(2),
   A_pay_brand char(1),
   A_pay_user_type char(2),
   A_AcctID char(14),
   A_deducted char(1),
   A_ACCT_BALANCE char(12),
   A_ACCT_BALANCE_ID1 char(18),
   A_ACCT_BALANCE_AMT1 char(8),
   A_ACCT_BALANCE_ID2 char(18),
   A_ACCT_BALANCE_AMT2 char(8),
   A_ACCT_BALANCE_ID3 char(18),
   A_ACCT_BALANCE_AMT3 char(8),
   A_ACCT_BALANCE_ID4 char(18),
   A_ACCT_BALANCE_AMT4 char(8),
   B_pay_type char(1),
   B_pay_subno char(24),
   B_pay_switch_flag char(2),
   B_pay_district char(2),
   B_pay_brand char(1),
   B_pay_user_type char(2),
   B_AcctID char(14),
   B_deducted char(1),
   B_ACCT_BALANCE char(12),
   B_ACCT_BALANCE_ID1 char(18),
   B_ACCT_BALANCE_AMT1 char(8),
   B_ACCT_BALANCE_ID2 char(18),
   B_ACCT_BALANCE_AMT2 char(8),
   B_ACCT_BALANCE_ID3 char(18),
   B_ACCT_BALANCE_AMT3 char(8),
   change_part varchar(1024),
   KEY idx_subno (subno)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


*/

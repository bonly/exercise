//============================================================================
// Name        : mysql.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <cmdline.h>
//#include "printdata.h"
#include <mysql++.h>
#include <ssqls.h>

/*
sql_create_4(stock,
    1, 4, // The meaning of these values is covered in the user manual
    mysqlpp::sql_char, item,
    mysqlpp::sql_bigint, num,
    mysqlpp::sql_bigint, weight,
    mysqlpp::sql_bigint, price
    )
//*/

//#define sql4(date) sql_create_4(stock##date, 1, 4, mysqlpp::sql_char, item,mysqlpp::sql_bigint, num,mysqlpp::sql_bigint, weight,mysqlpp::sql_bigint, price)

#include <fstream>

using namespace std;

#define tbn(date) table_##date

sql_create_4(tbn(1024),
    1, 4, // The meaning of these values is covered in the user manual
    mysqlpp::sql_char, item,
    mysqlpp::sql_bigint, num,
    mysqlpp::sql_bigint, weight,
    mysqlpp::sql_bigint, price
    )

int
main(int argc, char *argv[])
{
    vector<tbn(1024)> stock_vector;
    for(int i=0; i<10; ++i)
    {
        tbn(1024) st;
        st.item = "item";
        st.num = i;
        st.weight = i * 10;
        st.price = i * 1.5;
        stock_vector.push_back(st);
    }

    try {
        // Establish the connection to the database server.
        mysqlpp::Connection con("test",
                "127.0.0.1", "bonly", "");

        mysqlpp::Query query = con.query();
        //query.exec("DELETE FROM stock");


        //mysqlpp::Query::MaxPacketInsertPolicy<> insert_policy(1000);
        query.insert(stock_vector.begin(), stock_vector.end()
                ).execute();

        // Retrieve and print out the new table contents.
        //print_stock_table(query);
        clog << "insert finish\n";
    }
    catch (const mysqlpp::BadQuery& er) {
        // Handle any query errors
        cerr << "Query error: " << er.what() << endl;
        return -1;
    }
    catch (const mysqlpp::BadConversion& er) {
        // Handle bad conversions
        cerr << "Conversion error: " << er.what() << endl <<
                "\tretrieved data size: " << er.retrieved <<
                ", actual size: " << er.actual_size << endl;
        return -1;
    }
    catch (const mysqlpp::BadInsertPolicy& er) {
        // Handle bad conversions
        cerr << "InsertPolicy error: " << er.what() << endl;
        return -1;
    }
    catch (const mysqlpp::Exception& er) {
        // Catch-all for any other MySQL++ exceptions
        cerr << "Error: " << er.what() << endl;
        return -1;
    }

    return 0;
}

/*
create table stock(item char(10),num int, weight bigint, price bigint);
 */

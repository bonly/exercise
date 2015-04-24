import QtQuick 2.0
import QtQuick.LocalStorage 2.0
import "qrc:/Lib/Db.js" as Db

Item {
          function getQml(){
//              console.log("begin");
//              var db = Db.getDatabase();
//              Db.initAppScript();
//              db.transaction(
//                          function(tx){
//                              var rs = tx.executeSql('select script from AppScript limit 1');
//                              for (var i=0; i<rs.rows.length; i++){
//                                  my_rec.mycontent  = rs.rows.item(i).script;
//                              }
//                          }
//              );

              var QuanZi = Qt.createComponent("qrc:/Page/quanzi.qml" );
//              var newObject = Qt.createQmlObject(my_rec.mycontent, my_rec, 'testObj');
//              aTxt.text="u click" + newObject.myvar
          }

           Component.onCompleted: getQml()

}

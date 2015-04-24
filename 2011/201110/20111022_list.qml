import QtQuick 2.2

/*
Item {
    width: 200; height: 250

    ListModel {
        id: myModel
        ListElement { type: "Dog"; age: 8 }
        ListElement { type: "Cat"; age: 5 }
    }

    Component {
        id: myDelegate
        Text { text: type + ", " + age }
    }

    ListView {
        anchors.fill: parent
        model: myModel
        delegate: myDelegate
    }
}
*/

//*
Item{
   width:200
   height:250
    ListModel{
        id :furitM
        ListElement{
            name: "Apple";
            cost: 2.45;
        }
        ListElement{
            name: "Orange";
            cost: 3.5;
        }

    }

    Component{   //must be Component
        id : acomp
        Column{ //must be column if u need multi line
          Text{
              text: name  + ":" + cost
          }
          Text{text: "=============" } //or u can not has two line
        }
    }

    ListView{
        model: furitM
        anchors.fill: parent
        delegate: acomp
    }
}
//*/


import QtQuick 2.2

Rectangle{
    height: parent.height

    ListModel {
        id: fruitModel

        ListElement {
            name: "Apple"
            cost: 2.45
        }
        ListElement {
            name: "Orange"
            cost: 3.25
        }
        ListElement {
            name: "Banana"
            cost: 1.95
        }
    }

    Component{
        id: highlight
        Rectangle{
            color: "lightsteelblue"
            radius: 5
        }
    }

    ListView {
        anchors.fill: parent
        model: fruitModel
        delegate: Row {
//            x:5; y:5;
            Text { text: "Fruit: " + name }
            Text { text: "Cost: $" + cost }
        }
        highlight: highlight
        focus:true
    }

    MouseArea {
        anchors.fill: parent
        onClicked: fruitModel.append({"cost": 5.95, "name":"Pizza"})
    }
}


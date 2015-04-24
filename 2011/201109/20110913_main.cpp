#include <QtGui/QGuiApplication>
#include "qtquick2applicationviewer.h"


/*
 * viewer.rootContext()->setContextProperty("AuthName", "Benny.Zou");
console.log(AuthName);
*/

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);

    QtQuick2ApplicationViewer viewer;

    viewer.setDebug(1);
//    viewer.setMainQmlFile(QStringLiteral("qrc:/Page/quanzi.qml"));
    viewer.setMainQmlFile(QStringLiteral("/App/main.qml"));
    viewer.showExpanded();

    return app.exec();
}

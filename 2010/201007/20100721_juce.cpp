#define JUCE_MODULE_AVAILABLE_juce_core                 1
#define JUCE_MODULE_AVAILABLE_juce_cryptography         1
#define JUCE_MODULE_AVAILABLE_juce_data_structures      1
#define JUCE_MODULE_AVAILABLE_juce_events               1
#define JUCE_MODULE_AVAILABLE_juce_graphics             1
#define JUCE_MODULE_AVAILABLE_juce_gui_basics           1
#define JUCE_MODULE_AVAILABLE_juce_gui_extra            1
#define JUCE_MODULE_AVAILABLE_juce_opengl               1
#define JUCE_MODULE_AVAILABLE_juce_video                1
#include "modules/juce_core/juce_core.h"
#include "modules/juce_cryptography/juce_cryptography.h"
#include "modules/juce_data_structures/juce_data_structures.h"
#include "modules/juce_events/juce_events.h"
#include "modules/juce_graphics/juce_graphics.h"
#include "modules/juce_gui_basics/juce_gui_basics.h"
#include "modules/juce_gui_extra/juce_gui_extra.h"
#include "modules/juce_opengl/juce_opengl.h"
#include "modules/juce_video/juce_video.h"

using namespace juce;

class MainWindow    : public DocumentWindow
{
   public:
        MainWindow()  : DocumentWindow ("MainWindow",
                                        Colours::lightgrey,
                                        DocumentWindow::allButtons)
        {
            setContentOwned (new MainContentComponent(), true);

            centreWithSize (getWidth(), getHeight());
            setVisible (true);
        }
};

class hello_juceApplication  : public JUCEApplication
{
  public:
    void initialise (const String& commandLine)
    {   
        // This method is where you should put your application's initialisation code..

        mainWindow = new MainWindow();
    }   
    MainWindow *mainWindow;
};

int main(int argc, char* argv[])
{
  hello_juceApplication app;
  app.initialiseApp();
  MessageManager::getInstance()->runDispatchLoop();
  return app.shutdownApp();
}


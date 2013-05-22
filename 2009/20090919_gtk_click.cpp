#include <gtk/gtk.h> 
#include <cstdio>

void 
on_window_destroy (GtkObject *object, gpointer user_data)
{
    gtk_main_quit ();
    printf("quit\n");
}

//G_MODULE_EXPORT 
void on_button1_clicked(GtkObject *object, gpointer user_data)
{
    printf("click\n");
}

int
main (int argc, char *argv[])
{
    GtkBuilder      *builder; 
    GtkWidget       *window;

    gtk_init (&argc, &argv);

    builder = gtk_builder_new ();
    gtk_builder_add_from_file (builder, "tutorial.xml", NULL);
    window = GTK_WIDGET (gtk_builder_get_object (builder, "dialog1"));
    GtkWidget *button = GTK_WIDGET(gtk_builder_get_object(builder,"button1"));
    g_signal_connect(window,"destroy",G_CALLBACK(on_window_destroy),NULL);
    g_signal_connect(button,"clicked",G_CALLBACK(on_button1_clicked),NULL);

    //gtk_builder_connect_signals (builder, NULL);
    g_object_unref (G_OBJECT (builder));
        
    gtk_widget_show (window);                
    gtk_main ();

    return 0;
}
//g++ `pkg-config --cflags --libs gtk+-2.0` 20090918_gtk.cpp 

/*
 
<?xml version="1.0"?>
<interface>
  <requires lib="gtk+" version="2.16"/>
  <!-- interface-naming-policy project-wide -->
  <object class="GtkDialog" id="dialog1">
    <property name="border_width">5</property>
    <property name="type_hint">normal</property>
    <property name="has_separator">False</property>
    <child internal-child="vbox">
      <object class="GtkVBox" id="dialog-vbox1">
        <property name="visible">True</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child>
          <placeholder/>
        </child>
        <child internal-child="action_area">
          <object class="GtkHButtonBox" id="dialog-action_area1">
            <property name="visible">True</property>
            <property name="layout_style">end</property>
            <child>
              <object class="GtkButton" id="button1">
                <property name="label" translatable="yes">button</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <signal name="clicked" handler="on_button1_clicked"/>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">False</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <placeholder/>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="pack_type">end</property>
            <property name="position">0</property>
          </packing>
        </child>
      </object>
    </child>
    <action-widgets>
      <action-widget response="0">button1</action-widget>
    </action-widgets>
  </object>
</interface>
*/

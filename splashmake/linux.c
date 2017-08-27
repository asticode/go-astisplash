#include <gtk/gtk.h>
#include "shared.c"

// Signal handler
void signal_handler(int signum)
{
   exit(signum);
}

// GTK activate callback
static void activate (GtkApplication* app, gpointer user_data)
{
    // Build window
    GtkWidget *window = gtk_application_window_new (app);

    // Window is frameless
    gtk_window_set_decorated(GTK_WINDOW (window), 0);

    // Center window
    gtk_window_set_position (GTK_WINDOW (window), GTK_WIN_POS_CENTER);

    // Resize
    gtk_window_set_default_size (GTK_WINDOW (window), width, height);

    // Add title
    gtk_window_set_title (GTK_WINDOW (window), "Splash");

    // Build box
    GtkWidget *box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 0);
    gtk_container_add(GTK_CONTAINER (window), box);

    // Add background to the box
    GtkWidget *image = gtk_image_new_from_file(background);
    gtk_box_pack_start(box, image, TRUE, FALSE, 0);

    // Show window
    gtk_widget_show_all (window);
}

// Main
int main (int argc, char **argv)
{
    // Register signal handler
    signal(SIGINT, signal_handler);

    // Parse flags
    argc = parseFlags(argc, argv);

    // Build application
    GtkApplication *app;
    int status;
    app = gtk_application_new ("org.asticode.astisplash", G_APPLICATION_FLAGS_NONE);
    g_signal_connect (app, "activate", G_CALLBACK (activate), NULL);
    status = g_application_run (G_APPLICATION (app), argc, argv);
    g_object_unref (app);
    return status;
}
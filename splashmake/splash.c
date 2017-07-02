#include <gtk/gtk.h>
#include <stdlib.h>

// Global vars
int width, height;
char *icon, *background, *title;

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

    // Add title
    if (title == NULL) {
        title = "Window";
    }
    gtk_window_set_title (GTK_WINDOW (window), title);

    // Resize
    if (width > 0 && height > 0) {
        gtk_window_set_default_size (GTK_WINDOW (window), width, height);
    }

    // Build layout
    GtkWidget *layout = gtk_layout_new(NULL, NULL);
    gtk_container_add(GTK_CONTAINER (window), layout);
    gtk_widget_show(layout);

    // Add background to the layout
    GtkWidget *image = gtk_image_new_from_file(background);
    gtk_layout_put(GTK_LAYOUT(layout), image, 0, 0);

    // Add spinner
    GtkWidget *spinner = gtk_spinner_new ();
    gtk_spinner_start (GTK_SPINNER (spinner));
    gtk_layout_put(GTK_LAYOUT(layout), spinner,  0, 0);

    // Show window
    gtk_widget_show_all (window);

    // Move spinner
    GtkAllocation* alloc = g_new(GtkAllocation, 1);
    gtk_widget_get_allocation(spinner, alloc);
    gtk_layout_move(layout, spinner, width/2-alloc->width/2, height/2-alloc->height/2+30);
    g_free(alloc);
}

// Main
int main (int argc, char **argv)
{
    // Register signal handler
    signal(SIGINT, signal_handler);

    // Parse flags
    while ((argc > 1) && (argv[1][0] == '-'))
    {
        switch (argv[1][1])
        {
            case 'b':
                background = &argv[1][2];
                break;
            case 'h':
                height = atoi(&argv[1][2]);
                break;
            case 'i':
                icon = &argv[1][2];
                break;
            case 't':
                title = &argv[1][2];
                break;
            case 'w':
                width = atoi(&argv[1][2]);
                break;
        }
        ++argv;
        --argc;
    }

    // Build application
    GtkApplication *app;
    int status;
    app = gtk_application_new ("org.asticode.astisplash", G_APPLICATION_FLAGS_NONE);
    g_signal_connect (app, "activate", G_CALLBACK (activate), NULL);
    status = g_application_run (G_APPLICATION (app), argc, argv);
    g_object_unref (app);
    return status;
}
#include <gtk/gtk.h>
#include <stdlib.h>

int width, height, x, y, isCentered;
char *imagePath;
GtkApplication *app;
GtkWidget *window;

void signal_handler(int signum)
{
   exit(signum);
}

static void activate (GtkApplication* app, gpointer user_data)
{
    window = gtk_application_window_new (app);
    gtk_window_set_title (GTK_WINDOW (window), "Window");
    gtk_window_set_decorated(GTK_WINDOW (window), 0); // Frameless window
    if (isCentered == 1) {
        gtk_window_set_position (GTK_WINDOW (window), GTK_WIN_POS_CENTER);
    } else {
        gtk_window_move(GTK_WINDOW (window), x, y);
    }
    GtkWidget *image = gtk_image_new_from_file(imagePath);
    if (width > 0 && height > 0) {
        gtk_window_set_default_size (GTK_WINDOW (window), width, height);
        GdkPixbuf *pixbuf =	gtk_image_get_pixbuf(GTK_IMAGE(image));
        pixbuf = gdk_pixbuf_scale_simple(pixbuf, width, height, GDK_INTERP_BILINEAR);
        gtk_image_set_from_pixbuf(GTK_IMAGE(image), pixbuf);
    }
    gtk_container_add (GTK_CONTAINER (window), image);
    gtk_widget_show_all (window);
}

int main (int argc, char **argv)
{
    // Register signal and signal handler
    signal(SIGINT, signal_handler);

    // Parse flags
    while ((argc > 1) && (argv[1][0] == '-'))
    {
        switch (argv[1][1])
        {
            case 'c':
                isCentered = 1;
                break;
            case 'h':
                height = atoi(&argv[1][2]);
                break;
            case 'i':
                imagePath = &argv[1][2];
                break;
            case 'w':
                width = atoi(&argv[1][2]);
                break;
            case 'x':
                x = atoi(&argv[1][2]);
                break;
            case 'y':
                y = atoi(&argv[1][2]);
                break;
        }
        ++argv;
        --argc;
    }

    // Build application
    int status;

    app = gtk_application_new ("org.gtk.example", G_APPLICATION_FLAGS_NONE);
    g_signal_connect (app, "activate", G_CALLBACK (activate), NULL);
    status = g_application_run (G_APPLICATION (app), argc, argv);
    g_object_unref (app);

    return status;
}
// Global vars
int width, height;
char *background;

// Parses flags
void parseFlags(int argc, char **argv) {
    while ((argc > 1) && (argv[1][0] == '-')) {
        switch (argv[1][1]) {
            case 'b':
                background = &argv[1][2];
                break;
            case 'h':
                height = atoi(&argv[1][2]);
                break;
            case 'w':
                width = atoi(&argv[1][2]);
                break;
        }
        ++argv;
        --argc;
    }
}
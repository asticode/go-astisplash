#import <Cocoa/Cocoa.h>
#include "shared.c"

// Build a new application
void* NewApplication(void) {
    return [NSApplication sharedApplication];
}

// Run an application
void Run(void *app) {
    @autoreleasepool {
        NSApplication* a = (NSApplication*)app;
        [a setActivationPolicy:NSApplicationActivationPolicyRegular];
        [a run];
    }
}

// Build a new window
void NewWindow() {
    // Build window
    NSWindow* w = [[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, width, height)
        styleMask:NSWindowStyleMaskBorderless
        backing:NSBackingStoreBuffered
        defer:NO
    ];

    // Make sure the window is displayed
    [w makeKeyAndOrderFront:nil];

    // Center window
    [w center];

    // Add background image
    NSString *p = [NSString stringWithUTF8String:background];
    NSImage *i =  [[NSImage alloc] initWithContentsOfFile:p];
    [[w contentView] setWantsLayer:YES];
    [[w contentView] layer].contents = i;
}

// Main
int main(int argc, char **argv) {
    // Parse flags
    parseFlags(argc, argv);

    // New application
    NSApplication *a = NewApplication();

    // Build window
    NewWindow();

    // Run app
    Run(a);
}
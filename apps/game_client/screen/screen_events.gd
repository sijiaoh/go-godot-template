extends Node


@warning_ignore("unused_signal")
signal replace_screen(screen_entry: ScreenEntry)
@warning_ignore("unused_signal")
signal append_screen(screen_entry: ScreenEntry)

@warning_ignore("unused_signal")
signal exit_screen()
@warning_ignore("unused_signal")
signal exited_screen(screen_entry: ScreenEntry)

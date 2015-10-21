Go library to interface with LDTP Servers
=========================================

This library allows you to pilot LDTP implementations for Windows, Mac
and Linux through the same endpoints used by
http://ldtp.freedesktop.org/user-doc/index.html, allowing you to
click, focus on windows, type some text in editboxes, push some
buttons, take screenshots and other desktop automation
functionalities.

See http://ldtp.freedesktop.org/user-doc/d5/db1/a00140.html and the
links to the Mac, Windows and Linux source code therein for a better
understanding of what is offered.

Example usage
-------------

```
	client := New("localhost:4118")

	guiExists, err := client.GUIExist("*Chrome")
    if err != nil || !guiExists {
        log.Fatalln("Chrome doesn't exist")
    }

	client.CaptureScreen("/tmp/boo1.jpg")
```


License
-------

GPLv3

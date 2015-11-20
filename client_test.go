package ldtp

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestSomeMouseMovement(t *testing.T) {
	client := New("localhost:4118")

	client.GenerateMouseEvent(0, 0, "b1p")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(5, 5, "rel")
	client.GenerateMouseEvent(25, 25, "rel")
	client.GenerateMouseEvent(255, 255, "rel")

}
func TestScreenCapture(t *testing.T) {
	client := New("localhost:4118")

	if client.CaptureScreen("/tmp/boo1.jpg") != nil {
		t.Fatal("error capturing")
	}
	if client.CaptureWindow("/tmp/boo2.jpg", "dlgHud") != nil {
		t.Fatal("error capturing")
	}
	if client.CaptureSized("/tmp/boo3.jpg", Size{0, 0, 25, 25}) != nil {
		t.Fatal("error capturing")
	}
}

func TestBunchOfThings(t *testing.T) {
	client := New("localhost:4118")

	_, err := client.GUITimeout(1)
	if err != nil {
		log.Fatalln("guitimeout:", err)
	}

	//windows, err := client.GetAppList()

	windows, err := client.GetWindowList()
	if err != nil {
		log.Fatalln("Couldn't get window list", err)
	}

	fmt.Printf("Window list: %s\n", windows)

	for _, window := range windows {
		fmt.Println("Dealing with window", window)

		guiExists, err := client.GUIExist(window)
		if err != nil {
			fmt.Println("    error:", err)
		}
		fmt.Println("    guiExists: ", guiExists)

		// continue

		// childs, err := client.GetChild(window, "", "label")
		// if err != nil {
		// 	fmt.Println("    error:", err)
		// 	continue
		// }
		// fmt.Printf("    childs Rechercher push_buttons: %#v\n", childs)

		objectList, err := client.GetObjectList(window)
		if err != nil {
			fmt.Println("    error fetching objects:", err)
			continue
		}

		for _, object := range objectList {
			fmt.Println("    for object", object)

			states, err := client.GetAllStates(window, object)
			if err != nil {
				fmt.Println("      error:", err)
				continue
			}
			fmt.Printf("      States: %#v\n", states)

			continue

			guiExists, err := client.GUIObjectExist(window, object)
			if err != nil {
				fmt.Println("      error:", err)
			}
			fmt.Println("      guiExists: ", guiExists)

			size, err := client.GetObjectSize(window, object)
			if err != nil {
				//fmt.Println("      error:", err)
				continue
			}
			fmt.Printf("    %#v\n", size)

			props, err := client.GetObjectProperties(window, object)
			if err != nil {
				fmt.Println("      error:", err)
				continue
			}
			fmt.Printf("      Props: %#v\n", props)
		}
	}
}

func TestAddExtensionInChrome(t *testing.T) {
	client := New("localhost:4118")

	//client.GetWindowList()

	windowName := "*Google*Chrome"

	if client.Click(windowName, "btnAddextension") != nil {
		log.Fatalln("Couldn't click btnAddextension")
	}

	objects, _ := client.GetObjectList(windowName)
	fmt.Println("  objects:", objects)
	for _, obj := range objects {
		if !strings.Contains(obj, "btnAddextension") {
			continue
		}
		size, err := client.GetObjectSize(windowName, obj)
		fmt.Println("  label value:", obj, size, err)
		props, err := client.GetObjectProperties(windowName, obj)
		fmt.Println("  props:", props, err)
	}
}

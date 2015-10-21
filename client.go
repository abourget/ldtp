package ldtp

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kolo/xmlrpc"
)

// Client is the interface to the LDTP server, on which you can make
// all future calls.
type Client struct {
	rpcClient *xmlrpc.Client
}

// New creates a Client, provided with a "host:port" string
func New(hostPort string) *Client {
	rpcClient, err := xmlrpc.NewClient(fmt.Sprintf("http://%s/RPC2", hostPort), nil)
	if err != nil {
		log.Fatalln("Couldn't create XML-RPC client:", err)
	}

	return &Client{
		rpcClient: rpcClient,
	}
}

// GetWindowList retrieves the list of windows, frames and panels.
func (c *Client) GetWindowList() (res []string, err error) {
	err = c.rpcClient.Call("getwindowlist", nil, &res)
	return
}

// GetAppList retrieves the list of accessibility application windows open.
func (c *Client) GetAppList() (res []string, err error) {
	err = c.rpcClient.Call("getapplist", nil, &res)
	return

}

// GetObjectList retrieves the list of available objects on a window.
func (c *Client) GetObjectList(windowName string) (res []string, err error) {
	err = c.rpcClient.Call("getobjectlist", []interface{}{windowName}, &res)
	return
}

// CaptureScreen takes a full screen capture (calls imagecapture).
func (c *Client) CaptureScreen(filename string) error {
	return c.doImageCapture(filename, "", nil)
}

// CaptureSized takes a full screen capture (calls imagecapture).
func (c *Client) CaptureSized(filename string, size Size) error {
	return c.doImageCapture(filename, "", &size)
}

// CaptureWindow takes a screen capture of a single window (calls imagecapture).
func (c *Client) CaptureWindow(filename, windowName string) error {
	return c.doImageCapture(filename, windowName, nil)
}

func (c *Client) doImageCapture(filename, windowName string, size *Size) error {
	var res string
	params := []interface{}{windowName}
	if size != nil {
		params = append(params, size.X, size.Y, size.Width, size.Height)
	}
	err := c.rpcClient.Call("imagecapture", params, &res)
	if err != nil {
		return err
	}

	imageContent, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, imageContent, 0644)
}

// Click does a click, press, select, check on buttons, combo boxes, etc..
func (c *Client) Click(windowName, objectName string) error {
	return c.rpcClient.Call("click", []interface{}{windowName, objectName}, nil)
}

// MouseMove moves the mouse in the middle of a certain object.
func (c *Client) MouseMove(windowName, objectName string) error {
	return c.rpcClient.Call("mousemove", []interface{}{windowName, objectName}, nil)
}

// MouseLeftClick clicks the left mouse button in the middle of the specified object.
func (c *Client) MouseLeftClick(windowName, objectName string) error {
	return c.rpcClient.Call("mouseleftclick", []interface{}{windowName, objectName}, nil)
}

// MouseRightClick clicks the right mouse button in the middle of the specified object.
func (c *Client) MouseRightClick(windowName, objectName string) error {
	return c.rpcClient.Call("mouserightclick", []interface{}{windowName, objectName}, nil)
}

// GenerateMouseEvent does a mouse action at a given screen position.
//
// The eventType can be: "abs" to jump to an absolute position, "rel" to move relatively to previous position, "b1p" to press left button, "b1r" to release left button, "b1c" to click left button, "b1d" to double-click left button. A similar pattern goes with "b2p", "b2r", "b2c" and "b2d" for the middle button, and b3[prcd] for the right button.
func (c *Client) GenerateMouseEvent(x, y int, eventType string) error {
	return c.rpcClient.Call("generatemouseevent", []interface{}{x, y, eventType}, nil)
}

// GetAllStates retrieves the states for a given object (ex:
// focusable, showing, visible, selectable, enabled, sensitive,
// horizontal, vertical, etc...)
func (c *Client) GetAllStates(windowName, objectName string) (out map[string]bool, err error) {
	var res []string
	err = c.rpcClient.Call("getallstates", []interface{}{windowName, objectName}, &res)
	if err != nil {
		return nil, err
	}

	out = make(map[string]bool)
	for _, state := range res {
		out[state] = true
	}
	return
}

// GetObjectInfo retrieves infos on objects for a given Window name.
func (c *Client) GetObjectInfo(windowName, objectName string) (res []string, err error) {
	err = c.rpcClient.Call("getobjectinfo", []interface{}{windowName, objectName}, &res)
	return
}

// GetObjectProperty retrieves a property of an object. Must be a
// property that exists as returned by `GetObjectInfo`.
func (c *Client) GetObjectProperty(windowName, objectName, propertyName string) (value interface{}, err error) {
	err = c.rpcClient.Call("getobjectproperty", []interface{}{windowName, objectName, propertyName}, &value)
	return
}

// GetTextValue retrieves the text displayed for an object. Specify -1 for `startOffset` and `endOffset` for default values.
func (c *Client) GetTextValue(windowName, objectName string, startOffset, endOffset int) (value string, err error) {
	params := []interface{}{windowName, objectName}
	if endOffset != -1 && startOffset == -1 {
		return "", fmt.Errorf("cannot specify endOffset if startOffset isn't specified")
	}
	if startOffset != -1 {
		params = append(params, startOffset)
		if endOffset != -1 {
			params = append(params, endOffset)
		}
	}
	err = c.rpcClient.Call("gettextvalue", params, &value)
	return
}

// GetChild gets the list of child objects available in the window, by
// matching componentName or role, or both (if non-empty string).
func (c *Client) GetChild(windowName, componentName, role string) (res []string, err error) {
	err = c.rpcClient.Call("getchild", []interface{}{windowName, componentName, role}, &res)
	return
}

// GetObjectProperties is a shorthand to retrieve all the
// properties. If you care about speed, use GetObjectProperty instead.
func (c *Client) GetObjectProperties(windowName, objectName string) (map[string]interface{}, error) {
	props, err := c.GetObjectInfo(windowName, objectName)
	if err != nil {
		return nil, err
	}

	out := make(map[string]interface{})
	for _, prop := range props {
		value, err := c.GetObjectProperty(windowName, objectName, prop)
		if err != nil {
			return nil, err
		}
		out[prop] = value
	}

	return out, nil
}

func (c *Client) GetWindowSize(windowName string) (*Size, error) {
	var res []int
	err := c.rpcClient.Call("getwindowsize", []interface{}{windowName}, &res)
	if err != nil {
		return nil, err
	}

	if len(res) != 4 {
		return nil, fmt.Errorf("didn't receive 4 arguments in return")
	}
	return &Size{res[0], res[1], res[2], res[3]}, nil
}

func (c *Client) GetObjectSize(windowName, objectName string) (*Size, error) {
	var res []int
	err := c.rpcClient.Call("getobjectsize", []interface{}{windowName, objectName}, &res)
	if err != nil {
		return nil, err
	}

	if len(res) != 4 {
		return nil, fmt.Errorf("didn't receive 4 arguments in return")
	}
	return &Size{res[0], res[1], res[2], res[3]}, nil
}

// GUIExist checks whether a window or component exists.
func (c *Client) GUIExist(windowName string) (out bool, err error) {
	var res int
	err = c.rpcClient.Call("guiexist", []interface{}{windowName}, &res)
	out = res == 1
	return
}

// GUITimeout sets the global timeout for windows, in seconds.
func (c *Client) GUITimeout(guiTimeout int) (out bool, err error) {
	var res int
	err = c.rpcClient.Call("guitimeout", []interface{}{guiTimeout}, &res)
	out = res == 1
	return
}

// ObjTimeout sets the global timeout for objects, in seconds.
func (c *Client) ObjTimeout(objTimeout int) (out bool, err error) {
	var res int
	err = c.rpcClient.Call("objtimeout", []interface{}{objTimeout}, &res)
	out = res == 1
	return
}

// GUIExistObject returns whether the objectName exists. This replaces
// the `ldtp` `objectexist` call.
func (c *Client) GUIExistObject(windowName, objectName string) (out bool, err error) {
	var res int
	err = c.rpcClient.Call("guiexist", []interface{}{windowName, objectName}, &res)
	out = res == 1
	return
}

type Size struct {
	X, Y, Width, Height int
}

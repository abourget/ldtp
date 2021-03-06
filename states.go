package ldtp

type LDTPState string

var StateIconified = LDTPState("iconified")
var StateInvalid = LDTPState("invalid")
var StatePressed = LDTPState("pressed")
var StateExpandable = LDTPState("expandable")
var StateVisible = LDTPState("visible")
var StateLastDefined = LDTPState("last_defined")
var StateBusy = LDTPState("busy")
var StateExpanded = LDTPState("expanded")
var StateManagesDescendants = LDTPState("manages_descendants")
var StateIsDefault = LDTPState("is_default")
var StateIndeterminate = LDTPState("indeterminate")
var StateRequired = LDTPState("required")
var StateTransient = LDTPState("transient")
var StateChecked = LDTPState("checked")
var StateSensitive = LDTPState("sensitive")
var StateCollapsed = LDTPState("collapsed")
var StateStale = LDTPState("stale")
var StateOpaque = LDTPState("opaque")
var StateEnabled = LDTPState("enabled")
var StateHasTooltip = LDTPState("has_tooltip")
var StateSupportsAutocompletion = LDTPState("supports_autocompletion")
var StateFocusable = LDTPState("focusable")
var StateSelectable = LDTPState("selectable")
var StateActive = LDTPState("active")
var StateHorizontal = LDTPState("horizontal")
var StateVisited = LDTPState("visited")
var StateInvalidEntry = LDTPState("invalid_entry")
var StateFocused = LDTPState("focused")
var StateModal = LDTPState("modal")
var StateVertical = LDTPState("vertical")
var StateSelected = LDTPState("selected")
var StateShowing = LDTPState("showing")
var StateAnimated = LDTPState("animated")
var StateEditable = LDTPState("editable")
var StateMultiLine = LDTPState("multi_line")
var StateSingleLine = LDTPState("single_line")
var StateSelectableText = LDTPState("selectable_text")
var StateArmed = LDTPState("armed")
var StateDefunct = LDTPState("defunct")
var StateMultiselectable = LDTPState("multiselectable")
var StateResizable = LDTPState("resizable")
var Statetruncated = LDTPState("truncated")

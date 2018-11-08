// +build js,wasm

package dom

import (
	"syscall/js"
)

var (
	Document Doc
	Body     js.Value
)

func init() {
	Document = Doc{NodeBase{js.Global().Get("document")}}
	Body = Document.Get("body")
}

type Node interface {
	JSValue() js.Value

	AppendChild(c Node)
	RemoveChild(c Node) Node
}

type NodeBase struct{ js.Value }

func (n NodeBase) JSValue() js.Value { return n.Value }

func (n NodeBase) AppendChild(c Node)      { n.Call("appendChild", c.JSValue()) }
func (n NodeBase) RemoveChild(c Node) Node { return NodeBase{n.Call("removeChild", c.JSValue())} }
func (n NodeBase) FirstChild() Node        { return NodeBase{n.Get("firstChild")} }
func (n NodeBase) RemoveAllChildren() {
	for c := n.FirstChild(); c.JSValue() != js.Null(); c = n.FirstChild() {
		n.RemoveChild(c)
	}
}

func (n NodeBase) AddEventListener(typ string, fn func(args []js.Value)) js.Callback {
	callBack := js.NewCallback(fn)
	n.Call("addEventListener", typ, callBack)
	return callBack
}

type Doc struct{ NodeBase }

func (d Doc) CreateElement(tag string) Elt { return Elt{NodeBase{Document.Call("createElement", tag)}} }
func (d Doc) GetElementById(id string) Elt { return Elt{NodeBase{Document.Call("getElementById", id)}} }

type Attrs map[string]string

type Elt struct{ NodeBase }

func Element(tag string) Elt {
	return Document.CreateElement(tag)
}

func (e Elt) SetInnerHTML(s string) { e.Set("innerHTML", s) }

func (e Elt) SetAttribute(k, v string) { e.Call("setAttribute", k, v) }

func (e Elt) WithAttribute(k, v string) Elt {
	e.SetAttribute(k, v)
	return e
}

func (e Elt) WithAttributes(attrs Attrs) Elt {
	for k, v := range attrs {
		e.SetAttribute(k, v)
	}
	return e
}

func (e Elt) WithChildren(ns ...Node) Elt {
	for _, n := range ns {
		e.AppendChild(n)
	}
	return e
}

func Text(s string) Elt { return Elt{NodeBase{Document.Call("createTextNode", s)}} }

func Div() Elt { return Element("div") }

type Tbl struct{ Elt }

func Table() Tbl { return Tbl{Element("table")} }

func (t Tbl) WithHeader(ns ...Node) Tbl {
	header := t.Call("createTHead")
	newRow := header.Call("insertRow", -1)
	for _, n := range ns {
		NodeBase{newRow.Call("insertCell", -1)}.AppendChild(n)
	}
	return t
}

func (t Tbl) AppendRow(ns ...Node) Elt {
	newRow := NodeBase{t.Call("insertRow", -1)}
	for _, n := range ns {
		NodeBase{newRow.Call("insertCell", -1)}.AppendChild(n)
	}
	return Elt{newRow}
}

type Inp struct{ Elt }

func Input(typ string) Inp {
	return Inp{Element("input").WithAttribute("type", typ)}
}

func (i Inp) WithPlaceholder(p string) Inp {
	i.Set("placeholder", p)
	return i
}

func (i Inp) WithValue(val string) Inp {
	i.Set("value", val)
	return i
}

func (i Inp) Value() string { return i.Get("value").String() }

func TextInput() Inp { return Input("text") }

func (i Inp) OnKeyUp(fn func(args []js.Value)) js.Callback { return i.AddEventListener("keyup", fn) }

type Btn struct{ Elt }

func Button(s string) Btn {
	btn := Element("button")
	btn.SetInnerHTML(s)
	return Btn{btn}
}

func (b Btn) OnClick(fn func(args []js.Value)) js.Callback { return b.AddEventListener("click", fn) }

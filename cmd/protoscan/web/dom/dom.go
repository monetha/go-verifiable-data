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

func (n NodeBase) AddEventListener(flags js.EventCallbackFlag, typ string, fn func(js.Value)) js.Callback {
	callBack := js.NewEventCallback(flags, fn)
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

func (e Elt) SetClass(c string) { e.SetAttribute("class", c) }

func (e Elt) WithAttribute(k, v string) Elt {
	e.SetAttribute(k, v)
	return e
}

func (e Elt) WithClass(c string) Elt {
	e.SetClass(c)
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

func Label(s string) Elt {
	l := Element("label")
	l.SetInnerHTML(s)
	return l
}

func Div() Elt { return Element("div") }

func Form() Elt { return Element("form") }

type Tbl struct {
	Elt
	thead Elt
	tbody Elt
	tfoot Elt
}

func Table() Tbl {
	thead := Element("thead")
	tbody := Element("tbody")
	tfoot := Element("tfoot")

	return Tbl{
		Elt:   Element("table").WithChildren(thead, tbody, tfoot),
		thead: thead,
		tbody: tbody,
		tfoot: tfoot,
	}
}

func (t Tbl) WithClass(c string) Tbl {
	t.SetClass(c)
	return t
}

func (t Tbl) WithHeader(ns ...Node) Tbl {
	newRow := NodeBase{t.thead.Call("insertRow", -1)}
	for _, n := range ns {
		headerCell := Element("th")
		headerCell.AppendChild(n)
		newRow.AppendChild(headerCell)
	}
	return t
}

func (t Tbl) AppendRow(ns ...Node) Elt {
	newRow := NodeBase{t.tbody.Call("insertRow", -1)}
	for _, n := range ns {
		NodeBase{newRow.Call("insertCell", -1)}.AppendChild(n)
	}
	return Elt{newRow}
}

type Inp struct{ Elt }

func Input(typ string) Inp {
	return Inp{Element("input").WithAttribute("type", typ)}
}

func (i Inp) WithClass(c string) Inp {
	i.SetClass(c)
	return i
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

func (i Inp) OnKeyUp(flags js.EventCallbackFlag, fn func(js.Value)) js.Callback { return i.AddEventListener(flags, "keyup", fn) }

type Btn struct{ Elt }

func Button(s string) Btn {
	btn := Element("button")
	btn.SetInnerHTML(s)
	return Btn{btn}
}

func (b Btn) WithClass(c string) Btn {
	b.SetClass(c)
	return b
}

func (b Btn) OnClick(flags js.EventCallbackFlag, fn func(js.Value)) js.Callback { return b.AddEventListener(flags, "click", fn) }

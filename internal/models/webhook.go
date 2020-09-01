package models

import (
	"github.com/google/go-github/github"
	"strconv"
)

type BaseTypeAndAttachment struct {
	Type string `json:"type"`
	Attachments []BaseContent `json:"attachments"`
}

func NewBaseTypeAndAttachment(
	types string, attachments []BaseContent)   BaseTypeAndAttachment {
	return BaseTypeAndAttachment{Type: types, Attachments: attachments}
}

type BaseContent struct {
	ContentType string `json:"contentType"`
	Content AdaptiveCard `json:"content"`
}

func NewBaseContent(contentType string,  content AdaptiveCard)   BaseContent {
	return BaseContent{ContentType: contentType, Content: content}
}

type AdaptiveCard struct {
	Schema string `json:"$schema"`
	Type string `json:"type"`
	Version string `json:"version"`
	Body []Container `json:"body"`
	Padding string `json:"padding"`
}

func NewAdaptiveCard(schema string,

	types string, version string, body []Container, padding string)   AdaptiveCard {
	return AdaptiveCard{Schema: schema, Type: types, Version: version, Body: body, Padding: padding}
}

type Container struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Padding string `json:"padding"`
	Items []interface{} `json:"items"`
}

func NewContainer(
	types string, id string, padding string, items []interface{})   Container {
	return Container{Type: types, Id: id, Padding: padding, Items: items}
}

type NormalItem struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Text string `json:"text"`
	Wrap bool `json:"wrap"`
	Weight string `json:"weight"`
	Size string `json:"size"`
}

func NewNormalItem(
	types string, id string, text string, wrap bool, weight string,size string,)   NormalItem {
	return NormalItem{Type: types, Id: id, Text: text, Wrap: wrap, Size: size, Weight: weight}
}

type FactSetItem struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Facts []FactsType `json:"facts"`
}

func NewFactSetItem(
types string, id string, facts []FactsType)   FactSetItem {
return FactSetItem{Type: types, Id: id, Facts: facts}
}
type FactsType struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

func NewFactsType(title string, value string)   FactsType {
	return FactsType{Title: title, Value: value}
}

type ActionSetItem struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Actions []ActionType `json:"actions"`
}

func NewActionSetItem(
types string, id string, actions []ActionType)   ActionSetItem {
return ActionSetItem{Type: types, Id: id, Actions: actions}
}

type ActionType struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Style string `json:"style"`
	IsPrimary bool `json:"isPrimary"`
}

func NewActionType(types string, id string, title string, url string, style string, isPrimary bool)   ActionType {
	return ActionType{Type: types, Id: id, Title: title, Url: url, Style: style, IsPrimary: isPrimary}
}

func NewHeadingItem(heading string) interface{}{
	return   NewNormalItem("TextBlock","210d4e03-5d00-4d6f-170c-b155820a26a3",heading,true,"Bolder","ExtraLarge")
}

func NewHeadingContainer(heading string) Container {
	items :=[]interface{}{NewHeadingItem(heading)}
	return   NewContainer("Container","af80d38b-672d-9f01-d7b8-52358cbb5eae","Medium",items)
}

func NewElementContainer(title string, due string,assignees []  *github.User, issuesUrl string) Container{
	var items  []interface{}
	titleBlock:=  NewNormalItem("TextBlock","594ddb64-859c-7728-c029-56a21b543ead",title,true,"Bolder","Medium")
	items = append(items, titleBlock)
	dueBlock:= struct {
		Type string `json:"type"`
		Id string `json:"id"`
		Text string `json:"text"`
		Wrap bool `json:"wrap"`
		Color string `json:"color"`
	}{
		Type: "TextBlock",
		Id: "7d1e4979-976a-721a-f940-367c22d9f887",
		Text: due,
		Wrap: true,
		Color: "Attention",
	}
	items = append(items, dueBlock)
	if len(assignees) == 0{
		items = append(items, NewNormalItem("TextBlock","56f5a066-b0ad-50c7-6a25-c0114c31c540","No Assignees",true,"Bolder","Medium"))
	}else {
		items = append(items, NewNormalItem("TextBlock","56f5a066-b0ad-50c7-6a25-c0114c31c540","Assignees",true,"Bolder","Medium"))
		var facts []FactsType
		for idx,x := range assignees{
			facts = append(facts,   NewFactsType(strconv.Itoa(idx+1),  *x.Login))
		}
		factSet:=  NewFactSetItem("FactSet","344d22c3-e483-4b95-95ff-9731e765fa0e",facts)
		items = append(items, factSet)
	}
	var actions []ActionType
	actions = append(actions,   NewActionType("Action.OpenUrl","2c40ca8f-17a6-b06e-6758-444b2910fb53","View Issue in Github",issuesUrl,"positive",true))
	actionSet:=  NewActionSetItem("ActionSet","88c98c11-2f50-6c69-f608-dc6c4b6f6332",actions)
	items = append(items, actionSet)

	return  NewContainer("Container","af80d38b-672d-9f01-d7b8-52358cbb5eae","Medium",items)
}

func NewPostBody(heading string) BaseTypeAndAttachment {
	var containers []Container
	containers = append(containers, NewHeadingContainer(heading))
	content:=NewAdaptiveCard("http://adaptivecards.io/schemas/adaptive-card.json","AdaptiveCard","1.0",containers,"None")
	baseContent:= NewBaseContent("application/vnd.microsoft.card.adaptive",  content)
	var attachments []BaseContent
	attachments = append(attachments,   baseContent)
	return   NewBaseTypeAndAttachment("message",attachments)
}


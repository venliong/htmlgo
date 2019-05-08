package htmlgo_test

import (
	"context"
	"os"

	. "github.com/theplant/htmlgo"
)

// ### Create a simple div, Text will be escaped by html
func ExampleTag_01simplediv() {
	comp := Div(
		Text("123<h1>"),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <div>123&lt;h1&gt;</div>
}

// ### Create a full html page
func ExampleTag_02fullhtml() {
	comp := HTML(
		Head(
			Meta().Charset("utf8"),
			Title("My test page"),
		),
		Body(
			Img("images/firefox-icon.png").Alt("My test image"),
		),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <!DOCTYPE html>
	//
	// <html>
	// <head>
	// <meta charset='utf8'></meta>
	//
	// <title>My test page</title>
	// </head>
	//
	// <body>
	// <img src='images/firefox-icon.png' alt='My test image'></img>
	// </body>
	// </html>
}

// ### Use RawHTML and Component
func ExampleTag_03rawhtmlandcomponent() {
	userProfile := func(username string, avatarURL string) HTMLComponent {
		return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
			return Div(
				H1(username).Class("profileName"),
				Img(avatarURL).Class("profileImage"),
				RawHTML("<svg>complicated svg</svg>\n"),
			).Class("userProfile").MarshalHTML(ctx)
		})
	}

	comp := Ul(
		Li(
			userProfile("felix<h1>", "http://image.com/img1.png"),
		),
		Li(
			userProfile("john", "http://image.com/img2.png"),
		),
	)
	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <ul>
	// <li>
	// <div class='userProfile'>
	// <h1 class='profileName'>felix&lt;h1&gt;</h1>
	//
	// <img src='http://image.com/img1.png' class='profileImage'></img>
	// <svg>complicated svg</svg>
	// </div>
	// </li>
	//
	// <li>
	// <div class='userProfile'>
	// <h1 class='profileName'>john</h1>
	//
	// <img src='http://image.com/img2.png' class='profileImage'></img>
	// <svg>complicated svg</svg>
	// </div>
	// </li>
	// </ul>
}

type MySelectBuilder struct {
	options  [][]string
	selected string
}

func MySelect() *MySelectBuilder {
	return &MySelectBuilder{}
}

func (b *MySelectBuilder) Options(opts [][]string) (r *MySelectBuilder) {
	b.options = opts
	return b
}

func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
	b.selected = selected
	return b
}

func (b *MySelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	opts := []HTMLComponent{}
	for _, op := range b.options {
		var opt HTMLComponent
		if op[0] == b.selected {
			opt = Option(op[1]).Value(op[0]).Attr("selected", "true")
		} else {
			opt = Option(op[1]).Value(op[0])
		}
		opts = append(opts, opt)
	}
	return Select(opts...).MarshalHTML(ctx)
}

// ### More complicated customized component
func ExampleTag_04newcomponentstyle() {

	/*
		Define MySelect as follows:

		type MySelectBuilder struct {
			options  [][]string
			selected string
		}

		func MySelect() *MySelectBuilder {
			return &MySelectBuilder{}
		}

		func (b *MySelectBuilder) Options(opts [][]string) (r *MySelectBuilder) {
			b.options = opts
			return b
		}

		func (b *MySelectBuilder) Selected(selected string) (r *MySelectBuilder) {
			b.selected = selected
			return b
		}

		func (b *MySelectBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
			opts := []HTMLComponent{}
			for _, op := range b.options {
				var opt HTMLComponent
				if op[0] == b.selected {
					opt = Option(op[1]).Value(op[0]).Attr("selected", "true")
				} else {
					opt = Option(op[1]).Value(op[0])
				}
				opts = append(opts, opt)
			}
			return Select(opts...).MarshalHTML(ctx)
		}
	*/

	comp := MySelect().Options([][]string{
		{"1", "label 1"},
		{"2", "label 2"},
		{"3", "label 3"},
	}).Selected("2")

	Fprint(os.Stdout, comp, context.TODO())
	//Output:
	// <select>
	// <option value='1'>label 1</option>
	//
	// <option value='2' selected='true'>label 2</option>
	//
	// <option value='3'>label 3</option>
	// </select>
}
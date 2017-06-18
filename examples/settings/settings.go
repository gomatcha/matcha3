package settings

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/overcyn/matcha/comm"
	"github.com/overcyn/matcha/env"
	"github.com/overcyn/matcha/layout/constraint"
	"github.com/overcyn/matcha/layout/table"
	"github.com/overcyn/matcha/paint"
	"github.com/overcyn/matcha/text"
	"github.com/overcyn/matcha/touch"
	"github.com/overcyn/matcha/view"
	"github.com/overcyn/matcha/view/basicview"
	"github.com/overcyn/matcha/view/imageview"
	"github.com/overcyn/matcha/view/resimageview"
	"github.com/overcyn/matcha/view/scrollview"
	"github.com/overcyn/matcha/view/stackscreen"
	"github.com/overcyn/matcha/view/switchview"
	"github.com/overcyn/matcha/view/textview"
	"github.com/overcyn/matchabridge"
	"golang.org/x/image/colornames"
)

func init() {
	matchabridge.RegisterFunc("github.com/overcyn/matcha/examples/settings New", func() *view.Root {
		return view.NewRoot(NewApp())
	})
}

type App struct {
	comm.Storer
	store          *comm.AsyncStore
	stackScreen    *stackscreen.Screen
	wifiController *WifiController
}

func NewApp() *App {
	st := &comm.AsyncStore{}
	app := &App{Storer: st, store: st}

	rootScreen := view.ScreenFunc(func(ctx *view.Context) view.View {
		return NewRootView(ctx, "", app)
	})

	app.stackScreen = stackscreen.NewScreen()
	app.store.Set("stackScreen", app.stackScreen)
	app.stackScreen.SetChildren(rootScreen)

	app.wifiController = NewWifiStore()
	app.store.Set("wifi", app.wifiController)

	return app
}

func (app *App) View(ctx *view.Context) view.View {
	return app.StackScreen().View(ctx)
}

func (app *App) StackScreen() *stackscreen.Screen {
	return app.stackScreen
}

func (app *App) WifiController() *WifiController {
	return app.wifiController
}

type RootView struct {
	*view.Embed
	app *App
}

func NewRootView(ctx *view.Context, key string, app *App) *RootView {
	if v, ok := ctx.Prev(key).(*RootView); ok {
		return v
	}
	v := &RootView{Embed: view.NewEmbed(ctx.NewId(key)), app: app}
	v.Subscribe(app.WifiController())
	return v
}

func (v *RootView) Build(ctx *view.Context) *view.Model {
	v.app.Lock()
	defer v.app.Unlock()

	l := &table.Layout{}
	{
		ctx := ctx.WithPrefix("1")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer)

		switchView := switchview.New(ctx, "switch")
		cell1 := NewBasicCell(ctx, "airplane")
		cell1.Title = "Airplane Mode"
		cell1.Icon = env.MustLoad("Airplane")
		cell1.AccessoryView = switchView
		cell1.HasIcon = true
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, "wifi")
		cell2.Title = "Wi-Fi"
		if v.app.WifiController().Enabled() {
			cell2.Subtitle = v.app.WifiController().CurrentNetworkSSID()
		} else {
			cell2.Subtitle = ""
		}
		cell2.HasIcon = true
		cell2.Icon = env.MustLoad("Wifi")
		cell2.Chevron = true
		cell2.OnTap = func() {
			v.app.Lock()
			defer v.app.Unlock()
			v.app.StackScreen().Push(view.ScreenFunc(func(ctx *view.Context) view.View {
				return NewWifiView(ctx, "", v.app, v.app.wifiController)
			}))
		}
		group = append(group, cell2)

		cell3 := NewBasicCell(ctx, "bluetooth")
		cell3.HasIcon = true
		cell3.Icon = env.MustLoad("Bluetooth")
		cell3.Title = "Bluetooth"
		cell3.Subtitle = "On"
		cell3.Chevron = true
		cell3.OnTap = func() {
			v.app.Lock()
			defer v.app.Unlock()
			v.app.StackScreen().Push(view.ScreenFunc(func(ctx *view.Context) view.View {
				return NewBluetoothView(ctx, "", v.app)
			}))
		}
		group = append(group, cell3)

		cell4 := NewBasicCell(ctx, "cellular")
		cell4.HasIcon = true
		cell4.Icon = env.MustLoad("Cellular")
		cell4.Title = "Cellular"
		cell4.Chevron = true
		cell4.OnTap = func() {
			v.app.Lock()
			defer v.app.Unlock()
			v.app.StackScreen().Push(view.ScreenFunc(func(ctx *view.Context) view.View {
				return NewCellularView(ctx, "", v.app)
			}))
		}
		group = append(group, cell4)

		cell5 := NewBasicCell(ctx, "hotspot")
		cell5.HasIcon = true
		cell5.Icon = env.MustLoad("Hotspot")
		cell5.Title = "Personal Hotspot"
		cell5.Subtitle = "Off"
		cell5.Chevron = true
		group = append(group, cell5)

		cell6 := NewBasicCell(ctx, "carrier")
		cell6.HasIcon = true
		cell6.Icon = env.MustLoad("Carrier")
		cell6.Title = "Carrier"
		cell6.Subtitle = "T-Mobile"
		cell6.Chevron = true
		group = append(group, cell6)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i)
		}
	}
	{
		ctx := ctx.WithPrefix("2")
		group := []view.View{}

		spacer := NewSpacer(ctx, "spacer")
		l.Add(spacer)

		cell1 := NewBasicCell(ctx, "notifications")
		cell1.HasIcon = true
		cell1.Icon = env.MustLoad("Notifications")
		cell1.Title = "Notifications"
		cell1.Chevron = true
		group = append(group, cell1)

		cell2 := NewBasicCell(ctx, "controlcenter")
		cell2.HasIcon = true
		cell2.Icon = env.MustLoad("ControlCenter")
		cell2.Title = "Control Center"
		cell2.Chevron = true
		group = append(group, cell2)

		cell3 := NewBasicCell(ctx, "donotdisturb")
		cell3.HasIcon = true
		cell3.Icon = env.MustLoad("DoNotDisturb")
		cell3.Title = "Do Not Disturb"
		cell3.Chevron = true
		group = append(group, cell3)

		for _, i := range AddSeparators(ctx, group) {
			l.Add(i)
		}
	}

	scrollChild := basicview.New(ctx, "scrollChild")
	scrollChild.Layouter = l
	scrollChild.Children = l.Views()

	scrollView := scrollview.New(ctx, "scrollView")
	scrollView.ContentView = scrollChild
	// scrollView.OnScroll = func(position float64) {
	// 	// v.Update()
	// }

	return &view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

func (v *RootView) StackBar(ctx *view.Context) *stackscreen.StackBar {
	return &stackscreen.StackBar{Title: "Settings"}
}

var (
	cellColor            = color.Gray{255}
	cellColorHighlighted = color.Gray{217}
	chevronColor         = color.RGBA{199, 199, 204, 255}
	separatorColor       = color.RGBA{203, 202, 207, 255}
	backgroundColor      = color.RGBA{239, 239, 244, 255}
	subtitleColor        = color.Gray{142}
	titleColor           = color.Gray{0}
	spacerTitleColor     = color.Gray{102}
)

func AddSeparators(ctx *view.Context, vs []view.View) []view.View {
	ctx.WithPrefix("sep")
	newViews := []view.View{}

	top := NewSeparator(ctx, "top")
	newViews = append(newViews, top)

	for idx, i := range vs {
		newViews = append(newViews, i)

		if idx != len(vs)-1 { // Don't add short separator after last view
			sep := NewSeparator(ctx, strconv.Itoa(idx))
			sep.LeftPadding = 60
			newViews = append(newViews, sep)
		}
	}

	bot := NewSeparator(ctx, "bottom")
	newViews = append(newViews, bot)
	return newViews
}

type Separator struct {
	*view.Embed
	LeftPadding float64
}

func NewSeparator(ctx *view.Context, key string) *Separator {
	if v, ok := ctx.Prev(key).(*Separator); ok {
		return v
	}
	return &Separator{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *Separator) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(0.5))
		s.WidthEqual(l.MaxGuide().Width())
	})

	chl := basicview.New(ctx, "child")
	chl.Painter = &paint.Style{BackgroundColor: separatorColor}
	l.Add(chl, func(s *constraint.Solver) {
		s.HeightEqual(l.Height())
		s.LeftEqual(l.Left().Add(v.LeftPadding))
		s.RightEqual(l.Right())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: cellColor},
	}
}

type Spacer struct {
	*view.Embed
	Height float64
}

func NewSpacer(ctx *view.Context, key string) *Spacer {
	if v, ok := ctx.Prev(key).(*Spacer); ok {
		return v
	}
	return &Spacer{
		Embed:  view.NewEmbed(ctx.NewId(key)),
		Height: 35,
	}
}

func (v *Spacer) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(v.Height))
		s.WidthEqual(l.MaxGuide().Width())
	})

	return &view.Model{
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

type SpacerHeader struct {
	*view.Embed
	Height float64
	Title  string
}

func NewSpacerHeader(ctx *view.Context, key string) *SpacerHeader {
	if v, ok := ctx.Prev(key).(*SpacerHeader); ok {
		return v
	}
	return &SpacerHeader{
		Embed:  view.NewEmbed(ctx.NewId(key)),
		Height: 50,
	}
}

func (v *SpacerHeader) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(v.Height))
		s.WidthEqual(l.MaxGuide().Width())
	})

	titleView := textview.New(ctx, "title")
	titleView.String = strings.ToTitle(v.Title)
	titleView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   13,
	})
	titleView.Style.SetTextColor(spacerTitleColor)
	// titleView.Painter = &paint.Style{BackgroundColor: colornames.Red}

	titleGuide := l.Add(titleView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(15))
		s.RightEqual(l.Right().Add(-15))
		s.BottomEqual(l.Bottom().Add(-10))
		s.TopGreater(l.Top())
	})
	_ = titleGuide

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

type SpacerDescription struct {
	*view.Embed
	Description string
}

func NewSpacerDescription(ctx *view.Context, key string) *SpacerDescription {
	if v, ok := ctx.Prev(key).(*SpacerDescription); ok {
		return v
	}
	return &SpacerDescription{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *SpacerDescription) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(0))
		s.WidthEqual(l.MaxGuide().Width())
	})

	titleView := textview.New(ctx, "title")
	titleView.String = v.Description
	titleView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   13,
	})
	titleView.Style.SetTextColor(spacerTitleColor)

	titleGuide := l.Add(titleView, func(s *constraint.Solver) {
		s.LeftEqual(l.Left().Add(15))
		s.RightEqual(l.Right().Add(-15))
		s.TopGreater(l.Top().Add(15))
		s.HeightGreater(constraint.Const(0))
	})

	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(titleGuide.Height().Add(30))
		s.WidthEqual(l.MaxGuide().Width())
	})

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
	}
}

type BasicCell struct {
	*view.Embed
	HasIcon       bool
	Icon          *env.Resource
	Title         string
	Subtitle      string
	AccessoryView view.View
	Chevron       bool
	OnTap         func()
	highlighted   bool
}

func NewBasicCell(ctx *view.Context, key string) *BasicCell {
	if v, ok := ctx.Prev(key).(*BasicCell); ok {
		return v
	}
	return &BasicCell{Embed: view.NewEmbed(ctx.NewId(key))}
}

func (v *BasicCell) Build(ctx *view.Context) *view.Model {
	l := constraint.New()
	l.Solve(func(s *constraint.Solver) {
		s.HeightEqual(constraint.Const(44))
		s.WidthEqual(l.MaxGuide().Width())
	})

	leftAnchor := l.Left()
	if v.HasIcon {
		iconView := resimageview.New(ctx, "icon")
		iconView.Resource = v.Icon
		iconView.ResizeMode = imageview.ResizeModeFill
		pIconView := view.WithPainter(iconView, &paint.Style{BackgroundColor: colornames.Lightgray, CornerRadius: 5})

		iconGuide := l.Add(pIconView, func(s *constraint.Solver) {
			s.WidthEqual(constraint.Const(30))
			s.HeightEqual(constraint.Const(30))
			s.LeftEqual(l.Left().Add(15))
			s.CenterYEqual(l.CenterY())
		})
		leftAnchor = iconGuide.Right()
	}

	rightAnchor := l.Right()
	if v.Chevron {
		chevronView := resimageview.New(ctx, "chevron")
		chevronView.Resource = env.MustLoad("TableArrow")
		chevronView.ResizeMode = imageview.ResizeModeCenter
		chevronView.Tint = chevronColor

		chevronGuide := l.Add(chevronView, func(s *constraint.Solver) {
			s.RightEqual(rightAnchor.Add(-15))
			s.LeftGreater(leftAnchor)
			s.CenterYEqual(l.CenterY())
			s.TopGreater(l.Top())
			s.BottomLess(l.Bottom())
		})
		rightAnchor = chevronGuide.Left()
	}

	if v.AccessoryView != nil {
		accessoryGuide := l.Add(v.AccessoryView, func(s *constraint.Solver) {
			s.RightEqual(rightAnchor.Add(-10))
			s.LeftGreater(leftAnchor)
			s.CenterYEqual(l.CenterY())
		})
		rightAnchor = accessoryGuide.Left()
	}

	if len(v.Subtitle) > 0 {
		subtitleView := textview.New(ctx, "subtitle")
		subtitleView.String = v.Subtitle
		subtitleView.Style.SetFont(text.Font{
			Family: "Helvetica Neue",
			Size:   14,
		})
		subtitleView.Style.SetTextColor(subtitleColor)

		subtitleGuide := l.Add(subtitleView, func(s *constraint.Solver) {
			s.RightEqual(rightAnchor.Add(-10))
			s.LeftGreater(leftAnchor)
			s.CenterYEqual(l.CenterY())
		})
		rightAnchor = subtitleGuide.Left()
	}

	titleView := textview.New(ctx, "title")
	titleView.String = v.Title
	titleView.Style.SetFont(text.Font{
		Family: "Helvetica Neue",
		Size:   14,
	})
	titleView.Style.SetTextColor(titleColor)

	titleGuide := l.Add(titleView, func(s *constraint.Solver) {
		s.LeftEqual(leftAnchor.Add(15))
		s.RightLess(rightAnchor.Add(-10))
		s.CenterYEqual(l.CenterY())
	})
	_ = titleGuide

	values := map[interface{}]interface{}{}
	if v.OnTap != nil {
		tap := &touch.ButtonRecognizer{
			OnTouch: func(e *touch.ButtonEvent) {
				switch e.Kind {
				case touch.EventKindPossible:
					v.highlighted = e.Inside
				case touch.EventKindFailed:
					v.highlighted = false
				case touch.EventKindRecognized:
					v.highlighted = false
					v.OnTap()
				}
				v.Update()
			},
		}
		values[touch.Key] = []touch.Recognizer{tap}
	}

	var color color.Color
	if v.highlighted {
		color = cellColorHighlighted
	} else {
		color = cellColor
	}

	return &view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: color},
		Values:   values,
	}
}

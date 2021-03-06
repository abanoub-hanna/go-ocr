package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/otiai10/gosseract"
)

func main() {
	gtk.Init(nil)

	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("Image to Text OCR Software")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, _ := gtk.GridNew()
	win.Add(grid)

	imgview, _ := gtk.ImageNew()
	pbuf, _ := gdk.PixbufNewFromFileAtScale("./img/default.png", 400, 400, true)
	imgview.SetFromPixbuf(pbuf)
	grid.Attach(imgview, 0, 0, 1, 1)

	textentry, _ := gtk.EntryNew()
	grid.AttachNextTo(textentry, imgview, gtk.POS_BOTTOM, 1, 1)

	btnChoose, _ := gtk.ButtonNewWithLabel("Choose An Image")
	grid.AttachNextTo(btnChoose, textentry, gtk.POS_BOTTOM, 1, 1)

	btnChoose.Connect("clicked", func() {
		dlg, _ := gtk.FileChooserDialogNewWith2Buttons(
			"choose an image", nil, gtk.FILE_CHOOSER_ACTION_OPEN,
			"Open", gtk.RESPONSE_OK, "Cancel", gtk.RESPONSE_CANCEL,
		)
		dlg.SetDefaultResponse(gtk.RESPONSE_OK)

		filter, _ := gtk.FileFilterNew()
		filter.SetName("images")
		filter.AddMimeType("image/png")
		filter.AddMimeType("image/jpeg")
		filter.AddPattern("*.png")
		filter.AddPattern("*.jpg")
		filter.AddPattern("*.jpeg")
		dlg.SetFilter(filter)

		response := dlg.Run()

		if response == gtk.RESPONSE_OK {
			filename := dlg.GetFilename()
			// imgview.SetFromFile(filename)
			pixbuf, _ := gdk.PixbufNewFromFileAtScale(filename, 400, 400, true)
			imgview.SetFromPixbuf(pixbuf)
			extracted := ocr(filename)
			textentry.SetText(extracted)
		}

		dlg.Destroy()
	})

	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}

func ocr(imgpath string) string {
	client := gosseract.NewClient()
	defer client.Close()

	client.Languages = []string{"eng", "ara"}

	client.SetImage(imgpath)

	//boundingBox, _ := client.GetBoundingBoxes(PageIteratorLevel.RIL_SYMBOL)

	text, _ := client.Text()

	return text
}

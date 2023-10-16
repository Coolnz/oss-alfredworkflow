



```shell

var=$(/usr/local/bin/python3 oss.py)

# echo -n $(($var != false?1:0))

if [ "$var" != "err" ]
then
	echo -n $var
else 
	echo -n 0
fi

```



---



- [atotto/clipboard: clipboard for golang](https://github.com/atotto/clipboard) 只支持读写字符串
- [golang-design/clipboard: 📋 cross-platform clipboard package that supports accessing text and image in Go (macOS/Linux/Windows/Android/iOS)](https://github.com/golang-design/clipboard) 支持读取图片，但是读取图片时报错




---


[add support for NSPasteboard by flexzuu · Pull Request #25 · progrium/macdriver](https://github.com/progrium/macdriver/pull/25)




macdriver 的 `NSPasteboard_GeneralPasteboard`

拿不到图片数据



```go

func getPasteImgFile() string {

	// gp := cocoa.NSPasteboard_GeneralPasteboard()
	//
	// types := gp.Types()
	// fmt.Print(types)
	// fmt.Print(cocoa.NSPasteboardTypePNG)
	// if slices.Contains(types, cocoa.NSPasteboardTypePNG) {
	// 	fmt.Print("is png")
	// }

	gp := cocoa.NSPasteboard_GeneralPasteboard()
	pbType := gp.AvailableTypeFromArray([]cocoa.NSPasteboardType{cocoa.NSPasteboardTypePNG, cocoa.NSPasteboardTypeTIFF})
	if pbType == "" {
		// check which types are in the pasteboard right now
		log.Printf("no matching type found. only found: %s \n", gp.Types())
		panic("no matching type found. only found: %s \n")
	}
	// fetch data
	data := gp.DataForType(pbType)
	var bs []byte
	// turn into go bytes
	bs = data.Bytes()
	fileName := fmt.Sprintf("%d.png", time.Now().Unix())
	filePath := filepath.Join("/tmp", fileName)

	data.WriteToFile_atomically_(core.NSStringRef(), false)
	// f, err := ioutil.TempFile("", "out-img")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// _, err = f.Write(bs)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("png written to file: %s\n", f.Name())

	err := ioutil.WriteFile("~/Downloads/1.png", bs, 0666)
	if err != nil {
		panic(err)
	}
	// gp.DataForType()
	// data := gp.Send("dataForType:", core.NSString_FromString(string(cocoa.NSPasteboardTypePNG)))
	// objcBytes := data.Get("bytes")
	// length := data.Get("length")
	// bs := C.GoBytes(unsafe.Pointer(objcBytes.Pointer()), C.int(length.Uint()))
	// err := ioutil.WriteFile("~/Documents/1.png", bs, 0666)
	// if err != nil {
	// 	panic(err)
	// }
}

```



[Alfred Gallery • Workflows • Rename Action](https://alfred.app/workflows/vitor/rename-action/)



[Question: How to check if workflow is a rerun? · Issue #44 · deanishe/awgo](https://github.com/deanishe/awgo/issues/44)

The {var:path} adds the file path set in the Workflow Environment Variables
The {query} adds the name of the file selected in the yellow List Filter object
The {var:task} inserts the content; taken from the very first object where I typed the task I wanted to add.



[Copy to Clipboard Output - Alfred Help and Support](https://www.alfredapp.com/help/workflows/outputs/copy-to-clipboard/)

[Using Clipboard History Items In Workflows and Snippets - Alfred Help and Support](https://www.alfredapp.com/help/features/clipboard/accessing-clipboard-history/)
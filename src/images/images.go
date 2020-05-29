package images

import ( 
    "image/png"
    "github.com/nfnt/resize"
    "Response"
    "fmt"
    "os"
    "image"
    "image/draw"
    "os/exec"
    "runtime"
)

func DrawOnTerm(name string) {

	var cmd *exec.Cmd

	if runtime.GOOS == "linux" {
        cmd = exec.Command("feh", name)
    } else {
        cmd = exec.Command("sh", "catimg.sh", name)
    }
    stdout, err := cmd.Output()

    if err != nil {
        Response.Print(fmt.Sprintf("%s\n", err))
        return
    }

    fmt.Println(string(stdout))
}

func Resize(name string, ratio float64) {

	size := uint(ratio * 1000)

	file, err := os.Open(name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return
	}

	img, err := png.Decode(file)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return
	}
	file.Close()

	m := resize.Resize(size, 0, img, resize.Lanczos3)

	out, err := os.Create(name)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return
	}
	defer out.Close()
	png.Encode(out, m)
}

func Append(final_name, name1, name2 string) {

	imgFile1, err := os.Open(name1)
	imgFile2, err := os.Open(name2)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}

	img1, _, err := image.Decode(imgFile1)
	img2, _, err := image.Decode(imgFile2)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}

	sp2 := image.Point{img1.Bounds().Dx(), 0}

	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
	rgba := image.NewRGBA(r)

	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)

	out, err := os.Create(final_name)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}
	png.Encode(out, rgba)
	os.Remove(name1)
	os.Remove(name2)
}

func AppendRow(final_name, name1, name2 string) {

	imgFile1, err := os.Open(name1)
	imgFile2, err := os.Open(name2)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}

	img1, _, err := image.Decode(imgFile1)
	img2, _, err := image.Decode(imgFile2)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}

	asd := img1.Bounds()
	sp2 := image.Point{0, asd.Max.Y}

	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
	rgba := image.NewRGBA(r)

	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)

	out, err := os.Create(final_name)
	if err != nil {
	    Response.Print(fmt.Sprintf("%s\n", err))
	    return
	}
	png.Encode(out, rgba)
	os.Remove(name1)
	os.Remove(name2)
}
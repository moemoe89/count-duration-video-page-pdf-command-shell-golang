package main

import (
	"log"
	"math"
	"os/exec"
	"strings"
	"strconv"
)


func GetDuration(location string)(int,error){

	out, err := exec.Command("sh","-c","ffmpeg -i "+location+" 2>&1 | grep Duration | cut -d ' ' -f 4 | sed s/,//").Output()
	if err != nil {
		return 0,err
	}

	outString := string(out[:])
	split := strings.Split(outString, ":")

	var seconds int

	hour := strings.Replace(split[0], "\n", "", -1)
	hourFloat, err := strconv.ParseFloat(hour,64)
	if err != nil {
		return 0,err
	}

	hourInt := int(math.Ceil(hourFloat))

	minute := strings.Replace(split[1], "\n", "", -1)
	minuteFloat, err := strconv.ParseFloat(minute,64)
	if err != nil {
		return 0,err
	}

	minuteInt := int(math.Ceil(minuteFloat))

	second := strings.Replace(split[2], "\n", "", -1)
	secondFloat, err := strconv.ParseFloat(second,64)
	if err != nil {
		return 0,err
	}

	secondInt := int(math.Ceil(secondFloat))

	seconds = (hourInt*3600) + (minuteInt*60) + (secondInt)

	return seconds,nil

}

func GetPage(location string)(int,error){

	out, err := exec.Command("sh","-c","pdfinfo "+location+" | grep ^Pages:").Output()
	if err != nil {
		return 0,err
	}

	outString := string(out[:])
	outString = strings.Replace(outString, "Pages:          ", "", -1)
	outString = strings.Replace(outString, "\n", "", -1)

	pages,err := strconv.Atoi(outString)
	if err != nil {
		return 0,err
	}

	return pages,nil

}

func AddBackSlashToWhiteSpace(filename string)(string){

	split := strings.Split(filename, " ")

	j := len(split) - 1

	var newFilename string
	for k,i := range split {
		if j != k {
			newFilename += i + `\ `
		} else {
			newFilename += i
		}
	}

	return newFilename

}

func main(){

	videoName := AddBackSlashToWhiteSpace("video sample.mp4")

	seconds, err := GetDuration("/Users/bismo/go/src/system-command/"+videoName)
	if err != nil {
		log.Print(err)
	}

	log.Print("Video duration (seconds) : ",seconds)

	pdfName := AddBackSlashToWhiteSpace("pdf sample.pdf")

	pages, err := GetPage("/Users/bismo/go/src/system-command/"+ pdfName)
	if err != nil {
		log.Print(err)
	}

	log.Print("Total pages : ",pages)

}